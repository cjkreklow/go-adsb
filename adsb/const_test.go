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
	"fmt"
	"testing"
)

func testDF(t *testing.T) {
	for val, out := range map[DF]string{
		DF0:  "DF: Short air-air surveillance (ACAS)",
		DF4:  "DF: Surveillance altitude reply",
		DF5:  "DF: Surveillance identify reply",
		DF11: "DF: All-call reply",
		DF16: "DF: Long air-air surveillance (ACAS)",
		DF17: "DF: Extended squitter",
		DF18: "DF: Extended squitter / non-transponder",
		DF19: "DF: Military extended squitter",
		DF20: "DF: Comm-B altitude reply",
		DF21: "DF: Comm-B identify reply",
		DF22: "DF: Reserved for military use",
		DF24: "DF: Comm-D (ELM)",
		99:   "DF: Unknown value 99",
	} {
		var result = fmt.Sprintf("DF: %s", val)
		if result != out {
			t.Errorf("%s: expected %s | received %s\n", val, out, result)
		}
	}
}

func testCA(t *testing.T) {
	for val, out := range map[CA]string{
		CA0: "CA: Level 1",
		CA4: "CA: Level 2 - On Ground",
		CA5: "CA: Level 2 - Airborne",
		CA6: "CA: Level 2",
		CA7: "CA: CA=7",
		99:  "CA: Unknown value 99",
	} {
		var result = fmt.Sprintf("CA: %s", val)
		if result != out {
			t.Errorf("%s: expected %s | received %s\n", val, out, result)
		}
	}
}

func testFS(t *testing.T) {
	for val, out := range map[FS]string{
		FS0: "FS: No Alert, No SPI, Airborne",
		FS1: "FS: No Alert, No SPI, On Ground",
		FS2: "FS: Alert, No SPI, Airborne",
		FS3: "FS: Alert, No SPI, On Ground",
		FS4: "FS: Alert, SPI",
		FS5: "FS: No Alert, SPI",
		99:  "FS: Unknown value 99",
	} {
		var result = fmt.Sprintf("FS: %s", val)
		if result != out {
			t.Errorf("%s: expected %s | received %s\n", val, out, result)
		}
	}
}
