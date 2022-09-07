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

package adsbtype_test

import (
	"fmt"
	"testing"

	"kreklow.us/go/go-adsb/adsbtype"
)

// TestConst tests string formatting of constant values.
func TestConst(t *testing.T) {
	for val, out := range map[interface{}]string{
		adsbtype.CA0: "adsbtype.CA: Level 1",
		adsbtype.CC0: "adsbtype.CC: Not supported",
		adsbtype.CF0: "adsbtype.CF: ADS-B message, non-transponder device with ICAO address",
		adsbtype.DF0: "adsbtype.DF: Short air-air surveillance (ACAS)",
		adsbtype.DR0: "adsbtype.DR: No request",
		adsbtype.FS0: "adsbtype.FS: No alert, no SPI, airborne",
		adsbtype.RI0: "adsbtype.RI: No ACAS",
		adsbtype.SL0: "adsbtype.SL: ACAS inoperative",
		adsbtype.VS0: "adsbtype.VS: Airborne",

		adsbtype.ATS0:  "adsbtype.ATS: Barometric altitude",
		adsbtype.BDS02: "adsbtype.BDS: Linked Comm-B, segment 2",
		adsbtype.SSS0:  "adsbtype.SSS: No condition information",
		adsbtype.TRS0:  "adsbtype.TRS: No capability",

		adsbtype.TYPE0: "adsbtype.TYPE: No position information",
	} {
		result := fmt.Sprintf("%T: %s", val, val)
		if result != out {
			t.Errorf("expected %s | received %s\n", out, result)
		}
	}
}

func TestConstUnknown(t *testing.T) {
	for val, out := range map[interface{}]string{
		adsbtype.CA(99): "adsbtype.CA: Unknown value 99",
		adsbtype.CC(99): "adsbtype.CC: Unknown value 99",
		adsbtype.CF(99): "adsbtype.CF: Unknown value 99",
		adsbtype.DF(99): "adsbtype.DF: Unknown value 99",
		adsbtype.DR(99): "adsbtype.DR: Unknown value 99",
		adsbtype.FS(99): "adsbtype.FS: Unknown value 99",
		adsbtype.RI(99): "adsbtype.RI: Unknown value 99",
		adsbtype.SL(99): "adsbtype.SL: Unknown value 99",
		adsbtype.VS(99): "adsbtype.VS: Unknown value 99",

		adsbtype.ATS(99):   "adsbtype.ATS: Unknown value 99",
		adsbtype.BDS(0x99): "adsbtype.BDS: Unknown value 99",
		adsbtype.SSS(99):   "adsbtype.SSS: Unknown value 99",
		adsbtype.TRS(99):   "adsbtype.TRS: Unknown value 99",

		adsbtype.TYPE(99): "adsbtype.TYPE: Unknown value 99",
	} {
		result := fmt.Sprintf("%T: %s", val, val)
		if result != out {
			t.Errorf("expected %s | received %s\n", out, result)
		}
	}
}
