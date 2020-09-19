// Copyright 2020 Collin Kreklow
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
	"io"
)

// Decoder reads and decodes a Beast stream.
type Decoder struct {
	r   *bufio.Reader
	buf bytes.Buffer
}

// NewDecoder returns a Decoder which reads from r.
func NewDecoder(r io.Reader) *Decoder {
	d := new(Decoder)
	d.r = bufio.NewReader(r)

	return d
}

// Decode reads the next Beast frame from the input source and stores it
// in f. The UnmarshalBinary method of f must support unescaped Beast
// format data. The data passed to f remains valid only until the next
// call to Decode().
func (d *Decoder) Decode(f encoding.BinaryUnmarshaler) error {
	defer d.buf.Reset()

	b, err := d.r.ReadByte()
	if err != nil {
		return newError(err, "error reading stream")
	}

	if b != 0x1a {
		return newError(nil, "data stream corrupt")
	}

	d.buf.WriteByte(b)

	t, err := d.r.Peek(1)
	if err != nil {
		return newError(err, "error reading stream")
	}

	var msglen int

	switch t[0] {
	case 0x31:
		msglen = 11
	case 0x32:
		msglen = 16
	case 0x33:
		msglen = 23
	default:
		return newErrorf(nil, "unsupported frame type: %x", t[0])
	}

	err = d.readMsg(msglen)
	if err != nil {
		return err
	}

	err = f.UnmarshalBinary(d.buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func (d *Decoder) readMsg(l int) error {
	for j := d.buf.Len(); j < l; j = d.buf.Len() {
		b, err := d.r.ReadByte()
		if err != nil {
			return newError(err, "error reading stream")
		}

		if b == 0x1a { // TODO:refactor complexity
			nb, err := d.r.Peek(1)
			if err != nil {
				return newError(err, "error reading stream")
			}

			if nb[0] == 0x1a {
				_, err = d.r.Discard(1)
				if err != nil {
					return newError(err, "error reading stream")
				}
			} else {
				return newError(nil, "frame truncated")
			}
		}

		d.buf.WriteByte(b)
	}

	return nil
}
