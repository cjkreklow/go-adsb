// Copyright 2019 Collin Kreklow
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

// Package beast provides objects and methods for decoding and
// managing raw Mode-S Beast format frames.
//
package beast

import (
	"bufio"
	"encoding"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"kreklow.us/go/go-adsb/adsb"
)

// Frame is a Mode-S Beast format message
type Frame struct {
	raw []byte

	Format    uint8 // 1: Mode-AC, 2: Short Mode-S, 3: Long Mode-S, 4: Status
	Timestamp uint64
	Signal    uint8

	Msg *adsb.Message
}

// UnmarshalBinary accepts an unescaped Beast frame and populates the
// fields of the Frame object.
func (f *Frame) UnmarshalBinary(data []byte) error {
	if data[0] != 0x1a {
		return errors.New("format identifier not found")
	}
	switch data[1] {
	case 0x32:
		if len(data) != 16 {
			return fmt.Errorf("expected 16 bytes, received %d", len(data))
		}
		f.raw = data
		f.Format = 2
	case 0x33:
		if len(data) != 23 {
			return fmt.Errorf("expected 23 bytes, received %d", len(data))
		}
		f.raw = data
		f.Format = 3
	case 0x31, 0x34:
		return errors.New("format not supported")
	default:
		return errors.New("invalid format identifier")
	}
	f.Timestamp = binary.BigEndian.Uint64(append([]byte{0, 0}, f.raw[2:8]...))
	f.Signal = f.raw[8]

	f.Msg = new(adsb.Message)
	err := f.Msg.Decode(f.raw[9:])
	if err != nil {
		return err
	}

	return nil
}

// Bytes returns a copy of the raw data used to create the Frame.
func (f *Frame) Bytes() []byte {
	b := make([]byte, len(f.raw))
	copy(b, f.raw)
	return b
}

// Decoder reads and decodes a Beast stream
type Decoder struct {
	r *bufio.Reader
}

// NewDecoder returns a Decoder which reads from r.
func NewDecoder(r io.Reader) *Decoder {
	d := new(Decoder)
	d.r = bufio.NewReader(r)
	return d
}

// Decode reads the next Beast frame from the input source and stores it
// in f. The UnmarshalBinary method of f must support unescaped Beast
// format data.
func (d *Decoder) Decode(f encoding.BinaryUnmarshaler) error {
	m := make([]byte, 0, 64)

	b, err := d.r.ReadByte()
	if err != nil {
		return err
	}
	if b != 0x1a {
		return errors.New("data stream corrupt")
	}

	m = append(m, b)

	b, err = d.r.ReadByte()
	if err != nil {
		return err
	}

	var l int

	switch b {
	case 0x31:
		l = 11
	case 0x32:
		l = 16
	case 0x33:
		l = 23
	default:
		return errors.New("unsupported frame type")
	}

	m = append(m, b)

	for j := len(m); j < l; j = len(m) {
		b, err = d.r.ReadByte()
		if err != nil {
			return err
		}

		if b == 0x1a {
			nb, err := d.r.Peek(1)
			if err != nil {
				return err
			}

			if nb[0] == 0x1a {
				_, err = d.r.Discard(1)
				if err != nil {
					return err
				}
			} else {
				return errors.New("frame truncated")
			}
		}
		m = append(m, b)
	}

	err = f.UnmarshalBinary(m)
	if err != nil {
		return err
	}

	return nil
}
