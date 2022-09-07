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

package adsb

// decodeAC decodes the Altitude Code field to an altitude in feet.
func decodeAC(a uint64) (int64, error) {
	if a == 0 || a&0xffffffffffffe000 != 0 {
		return 0, newError(nil, "invalid altitude data")
	}

	if a&0b0000001000000 != 0 { // M bit designates feet vs meters
		return 0, newError(nil, "metric altitude not supported")
	}

	if a&0b0000000010000 == 0 { // Q bit designates 100 ft vs 25 ft increments
		// Gillham encoding
		// 100 ft increments
		h := grayDecode(((a & 0b1000000000000) >> 10) | // C1(20)
			((a & 0b0010000000000) >> 9) | // C2(22)
			((a & 0b0000100000000) >> 8)) // C4(24)

		if h == 0 || h == 5 || h == 6 || h > 7 {
			return 0, newError(nil, "invalid altitude value")
		}

		if h == 7 {
			h = 5
		}

		// 500 ft increments
		f := grayDecode(((a & 0b0000000000100) << 5) | // D2(30)
			((a & 0b0000000000001) << 6) | // D4(32)
			((a & 0b0100000000000) >> 6) | // A1(21)
			((a & 0b0001000000000) >> 5) | // A2(23)
			((a & 0b0000010000000) >> 4) | // A4(25)
			((a & 0b0000000100000) >> 3) | // B1(27)
			((a & 0b0000000001000) >> 2) | // B2(29)
			((a & 0b0000000000010) >> 1)) // B4(31)

		// reverse 100s if 500s is even
		if f%2 == 1 {
			h = 6 - h
		}

		return int64((f*500)+(h*100)) - 1300, nil
	}

	// must be an 11 bit altitude
	a = ((a & 0b1111110000000) >> 2) |
		((a & 0b0000000100000) >> 1) |
		(a & 0b0000000001111)

	return int64(a*25) - 1000, nil
}

// decodeESAlt decodes the extended squitter Altitude field to an
// altitude feet.
func decodeESAlt(a uint64) (int64, error) {
	if a == 0 || a&0xfffffffffffff000 != 0 {
		return 0, newError(nil, "invalid altitude data")
	}

	// insert M bit
	a = ((a & 0b111111000000) << 1) | (a & 0b000000111111)

	return decodeAC(a)
}

// grayDecode converts a value in "reflected binary code" aka "Gray
// code" to a decimal value.
func grayDecode(b uint64) uint64 {
	for z := 32; z >= 1; z /= 2 {
		b ^= (b >> z)
	}

	return b
}
