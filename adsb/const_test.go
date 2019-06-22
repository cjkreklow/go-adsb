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

// TestConst tests string formatting of constant values
func TestConst(t *testing.T) {
	t.Run("DF", testDF)
	t.Run("CA", testCA)
	t.Run("FS", testFS)
	t.Run("VS", testVS)
	t.Run("TC", testTC)
	t.Run("SS", testSS)
}

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

func testVS(t *testing.T) {
	for val, out := range map[VS]string{
		VS0: "VS: Airborne",
		VS1: "VS: On Ground",
		99:  "VS: Unknown value 99",
	} {
		var result = fmt.Sprintf("VS: %s", val)
		if result != out {
			t.Errorf("%s: expected %s | received %s\n", val, out, result)
		}
	}
}

func testTC(t *testing.T) {
	for val, out := range map[TC]string{
		TC0:  "TC: No position information",
		TC1:  "TC: Identification (Category Set D)",
		TC2:  "TC: Identification (Category Set C)",
		TC3:  "TC: Identification (Category Set B)",
		TC4:  "TC: Identification (Category Set A)",
		TC5:  "TC: Surface position, 7.5 meter",
		TC6:  "TC: Surface position, 25 meter",
		TC7:  "TC: Surface position, 0.1 NM",
		TC8:  "TC: Surface position",
		TC9:  "TC: Airborne position, 7.5 meter, barometric altitude",
		TC10: "TC: Airborne position, 25 meter, barometric altitude",
		TC11: "TC: Airborne position, 0.1 NM, barometric altitude",
		TC12: "TC: Airborne position, 0.2 NM, barometric altitude",
		TC13: "TC: Airborne position, 0.5 NM, barometric altitude",
		TC14: "TC: Airborne position, 1.0 NM, barometric altitude",
		TC15: "TC: Airborne position, 2.0 NM, barometric altitude",
		TC16: "TC: Airborne position, 10 NM, barometric altitude",
		TC17: "TC: Airborne position, 20 NM, barometric altitude",
		TC18: "TC: Airborne position, barometric altitude",
		TC19: "TC: Airborne velocity",
		TC20: "TC: Airborne position, 7.5 meter, GNSS height",
		TC21: "TC: Airborne position, 25 meter, GNSS height",
		TC22: "TC: Airborne position, GNSS height",
		TC28: "TC: Emergency priority status",
		TC31: "TC: Operational status",
		99:   "TC: Unknown value 99",
	} {
		var result = fmt.Sprintf("TC: %s", val)
		if result != out {
			t.Errorf("%s: expected %s | received %s\n", val, out, result)
		}
	}
}

func testSS(t *testing.T) {
	for val, out := range map[SS]string{
		SS0: "SS: No condition information",
		SS1: "SS: Permanent alert (emergency)",
		SS2: "SS: Temporary alert (ident change)",
		SS3: "SS: SPI",
		99:  "SS: Unknown value 99",
	} {
		var result = fmt.Sprintf("SS: %s", val)
		if result != out {
			t.Errorf("%s: expected %s | received %s\n", val, out, result)
		}
	}
}
