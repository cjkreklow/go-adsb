// Copyright 2024 Collin Kreklow
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS
// BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
// ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package beast

import (
	"bufio"
	"bytes"
	"encoding"
	"errors"
	"io"
)

// decoderReader allows mocking the bufio.Reader in Decoder.
type decoderReader interface {
	Buffered() int
	Discard(n int) (int, error)
	Peek(n int) ([]byte, error)
	ReadByte() (byte, error)
	UnreadByte() error
}

// Decoder reads a Beast stream and stores individual frames. It must be
// created with NewDecoder().
type Decoder struct {
	// Setting StripEscape to true will remove the extra escape
	// character from 0x1a values in the data passed to the
	// BinaryUnmarshaler by Decode.
	StripEscape bool

	r   decoderReader
	buf bytes.Buffer
}

// NewDecoder returns a Decoder which reads from r.
func NewDecoder(r io.Reader) *Decoder {
	d := new(Decoder)
	d.r = bufio.NewReader(r)

	return d
}

// Decode reads the next Beast frame from the input source and stores it
// in f. The data passed to f remains valid only until the next call to
// Decode().
func (d *Decoder) Decode(f encoding.BinaryUnmarshaler) error {
	// make sure the stream is at the beginning of a frame
	t, err := d.r.Peek(2)
	if err != nil {
		return readError(err)
	}

	if !(t[0] == 0x1a &&
		(t[1] == 0x31 || t[1] == 0x32 || t[1] == 0x33 || t[1] == 0x34)) {
		err = d.seekNext()
		if err != nil {
			return err
		}

		t, err = d.r.Peek(2)
		if err != nil {
			return readError(err)
		}
	}

	d.buf.Reset()

	// store the frame type escape sequence
	d.buf.Write(t)

	_, err = d.r.Discard(2)
	if err != nil {
		return readError(err)
	}

	// read the remainder of the message
	err = d.readMsg()
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	err = f.UnmarshalBinary(d.buf.Bytes())
	if err != nil {
		return newError(err, "error unmarshalling data")
	}

	return nil
}

// seekNext attempts to seek the input buffer to the next frame start
// sequence.
func (d *Decoder) seekNext() error {
	ct := 100 // don't read more than 100 bytes
	if d.r.Buffered() < ct {
		ct = d.r.Buffered()
	}

	b, err := d.r.Peek(ct)
	if err != nil {
		return readError(err)
	}

	var n int

	for _, t := range []byte{0x31, 0x32, 0x33, 0x34} {
		nx := bytes.Index(b, []byte{0x1a, t})
		if n == 0 && nx > 0 || nx > 0 && nx < n {
			n = nx
		}
	}

	if n == 0 {
		return newError(nil, "no frame data found")
	}

	_, err = d.r.Discard(n)
	if err != nil {
		return readError(err)
	}

	return nil
}

// readMsg writes frame data to the output buffer.
func (d *Decoder) readMsg() error {
	for i := 0; i < 100; i++ { // don't read more than 100 bytes
		b, err := d.r.ReadByte()
		if err != nil {
			return readError(err)
		}

		// not an escape byte
		if b != 0x1a {
			d.buf.WriteByte(b)

			continue
		}

		// return byte to the read buffer
		err = d.r.UnreadByte()
		if err != nil {
			return readError(err)
		}

		nb, err := d.r.Peek(2)
		if err != nil {
			return readError(err)
		}

		switch nb[1] {
		// next frame, message is complete
		case 0x31, 0x32, 0x33, 0x34:
			return nil
		// escaped 0x1a, write and continue
		case 0x1a:
			if d.StripEscape {
				d.buf.WriteByte(nb[1])
			} else {
				d.buf.Write(nb)
			}

			_, err = d.r.Discard(2)
			if err != nil {
				return readError(err)
			}

			continue
		}

		// unrecognized escape code
		break
	}

	return newError(nil, "data stream corrupt")
}

// readError returns a read error.
func readError(w error) beastError {
	return beastError{
		msg:  "error reading stream",
		werr: w,
	}
}
