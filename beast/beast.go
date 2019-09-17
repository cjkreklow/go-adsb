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
	"bytes"
	"errors"
	"fmt"
	"time"
)

// Frame is a Mode-S Beast format message. A Frame is safe to reuse by
// calling UnmarshalBinary with new data.
type Frame struct {
	data bytes.Buffer
}

// UnmarshalBinary stores unescaped Beast data into the Frame.
func (f *Frame) UnmarshalBinary(data []byte) error {
	f.data.Reset()

	if data[0] != 0x1a {
		return errors.New("format identifier not found")
	}
	switch data[1] {
	case 0x32:
		if len(data) != 16 {
			return fmt.Errorf("expected 16 bytes, received %d", len(data))
		}
		f.data.Write(data)
	case 0x33:
		if len(data) != 23 {
			return fmt.Errorf("expected 23 bytes, received %d", len(data))
		}
		f.data.Write(data)
	case 0x31, 0x34:
		return errors.New("format not supported")
	default:
		return errors.New("invalid format identifier")
	}

	return nil
}

// Bytes exposes the underlying raw data used to create the Frame. The
// returned slice remains valid until the next call to UnmarshalBinary.
// Modifying the returned slice directly may impact future calls to
// Bytes() or ADSB().
func (f Frame) Bytes() []byte {
	return f.data.Bytes()
}

// ADSB exposes the 56- or 112-bit ADS-B data encoded in the Frame. The
// returned slice remains valid until the next call to UnmarshalBinary.
// Modifying the returned slice directly may impact future calls to
// Bytes() or ADSB().
func (f Frame) ADSB() []byte {
	return f.data.Bytes()[9:]
}

// Timestamp returns the MLAT timestamp as a time.Duration
func (f Frame) Timestamp() time.Duration {
	d := f.data.Bytes()
	ts := int64(d[7]) | int64(d[6])<<8 | int64(d[5])<<16 |
		int64(d[4])<<24 | int64(d[3])<<32 | int64(d[2])<<40

	return time.Duration(ts * 1000 / 12).Round(time.Microsecond / 2)
}

// Signal returns the signal level.
func (f Frame) Signal() uint8 {
	return f.data.Bytes()[8]
}
