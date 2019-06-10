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

package adsb

import (
	"errors"
)

// decodeAlt13 converts a 13-bit altitude code field to an integer
// altitude value in feet.  The highest three bits of the uint16
// argument passed are ignored.
func decodeAlt13(a uint16) (int64, error) {
	if a&0xE000 != 0 { // data is not properly aligned
		return 0, errors.New("invalid data length")
	}
	if a == 0 { // altitude is 0 or invalid
		return 0, nil
	}
	if a&0x40 != 0 { // M bit designates feet vs meters
		return 0, errors.New("metric altitude not supported")
	}

	if a&0x10 == 0 { // Q bit designates 100 ft vs 25 ft increments
		// Gillham encoding

		// trailing 3 bits is 100 ft increments
		h := grayDecode(uint64(((a >> 10) & 0x04) |
			((a >> 9) & 0x02) | ((a >> 8) & 0x01))) // C1(20) C2(22) C4(24)
		if h == 5 || h == 6 {
			return 0, errors.New("invalid altitude value")
		}
		if h == 7 {
			h = 5
		}

		// first 8 bits is 500 ft increments
		f := grayDecode(uint64(((a << 5) & 0x80) | ((a << 6) & 0x40) | // D2(30) D4(32)
			((a >> 6) & 0x20) | ((a >> 5) & 0x10) | ((a >> 4) & 0x08) | // A1(21) A2(23) A4(25)
			((a >> 3) & 0x04) | ((a >> 2) & 0x02) | ((a >> 1) & 0x01))) // B1(27) B2(29) B4(31)
		if f%2 == 1 {
			h = 6 - h
		}
		return int64((f*500)+(h*100)) - 1300, nil
	}

	// must be an 11 bit altitude
	return (int64((a&0x0F)|((a&0x20)>>1)|((a&0x1F80)>>2)) * 25) - 1000, nil
}

// decodeAlt12 converts a 12-bit extended squitter altitude field to an
// integer altitude value in feet.  The highest four bits of the uint16
// argument passed are ignored.
func decodeAlt12(a uint16) (int64, error) {
	if a&0xF000 != 0 { // data is not properly aligned
		return 0, errors.New("invalid data length")
	}
	if a == 0 { // altitude is 0 or invalid
		return 0, nil
	}

	// insert a ZERO M-bit
	a = ((a & 0x0FC0) << 1) | (a & 0x3F)

	return decodeAlt13(a)
}

// grayDecode converts a value in "reflected binary code" aka "Gray
// code" to the standard decimal value
func grayDecode(b uint64) uint64 {
	for z := uint(32); z >= 1; z /= 2 {
		b = b ^ (b >> z)
	}
	return b
}
