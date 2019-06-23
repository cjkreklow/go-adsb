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
	"errors"
	"fmt"
	"math/big"
	"time"
)

// Frame is a Mode-S Beast format message
type Frame struct {
	raw []byte

	Format    uint8 // 1: Mode-AC, 2: Short Mode-S, 3: Long Mode-S, 4: Status
	Signal    uint8
	Timestamp time.Duration
}

// UnmarshalBinary stores unescaped Beast data into the Frame.
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

	f.Signal = f.raw[8]
	f.Timestamp = time.Microsecond *
		time.Duration(big.NewInt(0).SetBytes(f.raw[2:8]).Int64()/12)

	return nil
}

// Bytes returns a copy of the raw data used to create the Frame.
func (f Frame) Bytes() []byte {
	b := make([]byte, len(f.raw))
	copy(b, f.raw)
	return b
}

// ADSB returns a copy of the 56- or 112-bit ADS-B data encoded in the
// Frame.
func (f Frame) ADSB() []byte {
	b := make([]byte, len(f.raw)-9)
	copy(b, f.raw[9:])
	return b
}
