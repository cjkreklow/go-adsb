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
	"bytes"
	"encoding"
	"time"
)

var _ encoding.BinaryUnmarshaler = new(Frame)
var _ encoding.BinaryMarshaler = new(Frame)

// Frame is a Mode-S Beast format message. A Frame is safe to reuse by
// calling UnmarshalBinary with new data.
type Frame struct {
	data bytes.Buffer
}

// UnmarshalBinary stores a Beast message.
func (f *Frame) UnmarshalBinary(data []byte) error {
	f.data.Reset()

	if len(data) < 9 {
		return newError(nil, "received truncated data")
	}

	if !(data[0] == 0x1a &&
		(data[1] == 0x31 || data[1] == 0x32 || data[1] == 0x33 || data[1] == 0x34)) {
		return newErrorf(nil, "invalid data format: %04x", data[0:2])
	}

	for i := 0; i < len(data); i++ {
		if i > 0 && data[i-1] == 0x1a && data[i] == 0x1a {
			continue
		}

		f.data.WriteByte(data[i])
	}

	switch data[1] {
	case 0x31:
		if f.data.Len() != 11 {
			return newErrorf(nil, "expected 11 bytes, received %d", f.data.Len())
		}
	case 0x32:
		if f.data.Len() != 16 {
			return newErrorf(nil, "expected 16 bytes, received %d", f.data.Len())
		}
	case 0x33:
		if f.data.Len() != 23 {
			return newErrorf(nil, "expected 23 bytes, received %d", f.data.Len())
		}
	}

	return nil
}

// MarshalBinary returns a Beast message.
func (f *Frame) MarshalBinary() ([]byte, error) {
	if f.data.Len() < 9 {
		return nil, ErrNoData
	}

	ob := bytes.NewBuffer(make([]byte, 0, 25))

	for i, b := range f.data.Bytes() {
		if i > 0 && b == 0x1a {
			ob.WriteByte(0x1a)
		}

		ob.WriteByte(b)
	}

	return ob.Bytes(), nil
}

// Bytes returns the stored frame data. Escaped 0x1a values in the data
// are stripped before storing. To obtain the frame in the escaped wire
// format, use MarshalBinary.
//
// The returned slice remains valid until the next call to
// UnmarshalBinary. Modifying the returned slice directly may impact
// future Frame method calls.
func (f Frame) Bytes() []byte {
	return f.data.Bytes()
}

// MarshalADSB returns the ADS-B or Mode-AC data in the Frame.
//
// The returned slice remains valid until the next call to
// UnmarshalBinary. Modifying the returned slice directly may impact
// future Frame method calls.
func (f Frame) MarshalADSB() ([]byte, error) {
	if f.data.Len() < 10 {
		return nil, ErrNoData
	}

	return f.data.Bytes()[9:], nil
}

// Timestamp returns the MLAT timestamp as a time.Duration.
func (f Frame) Timestamp() (time.Duration, error) {
	if f.data.Len() < 8 {
		return time.Duration(0), ErrNoData
	}

	d := f.data.Bytes()
	ts := int64(d[2])<<40 | int64(d[3])<<32 | int64(d[4])<<24 |
		int64(d[5])<<16 | int64(d[6])<<8 | int64(d[7])

	return time.Duration(ts * 1000 / 12).Round(time.Microsecond / 2), nil
}

// Signal returns the signal level.
func (f Frame) Signal() (uint8, error) {
	if f.data.Len() < 9 {
		return 0, ErrNoData
	}

	return f.data.Bytes()[8], nil
}
