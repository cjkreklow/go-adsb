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

// ESType returns the extended squitter type code.
func (r *RawMessage) ESType() (uint64, error) {
	df, err := r.DF()
	if err != nil {
		return 0, err
	}

	switch df {
	case 17:
		return r.esbits(1, 5), nil
	case 18:
		cf, err := r.CF()
		if err != nil {
			return 0, err
		}

		switch cf {
		case 0, 1, 2, 5, 6:
			return r.esbits(1, 5), nil
		default:
			return 0, newErrorf(ErrNotAvailable,
				"error retrieving %s from %d/%d", "ESType", df, cf)
		}
	default:
		return 0, newErrorf(ErrNotAvailable, "error retrieving %s from %d",
			"ESType", df)
	}
}

// ESAltitude returns the extended squitter altitude field.
func (r *RawMessage) ESAltitude() (uint64, error) {
	tc, err := r.ESType()
	if err != nil {
		return 0, err
	}

	switch tc {
	case 0, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18:
		return r.esbits(9, 20), nil
	default:
		return 0, newErrorf(ErrNotAvailable, "error retrieving %s from %d",
			"ESAltitude", tc)
	}
}

// Get bits from the ME field.
func (r *RawMessage) esbits(n int, z int) uint64 {
	return r.Bits(n+32, z+32)
}
