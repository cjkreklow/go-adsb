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

package adsb_test

import (
	"fmt"
	"testing"

	"kreklow.us/go/go-adsb/adsb"
)

// TestConst tests string formatting of constant values
func TestConst(t *testing.T) {
	t.Run("DF", testDF)
	t.Run("CA", testCA)
	t.Run("CF", testCF)
	t.Run("FS", testFS)
	t.Run("VS", testVS)
	t.Run("TC", testTC)
	t.Run("SS", testSS)
	t.Run("AcCat", testAcCat)
}

func testDF(t *testing.T) {
	for val, out := range map[adsb.DF]string{
		adsb.DF0:  "DF: Short air-air surveillance (ACAS)",
		adsb.DF4:  "DF: Surveillance altitude reply",
		adsb.DF5:  "DF: Surveillance identify reply",
		adsb.DF11: "DF: All-call reply",
		adsb.DF16: "DF: Long air-air surveillance (ACAS)",
		adsb.DF17: "DF: Extended squitter",
		adsb.DF18: "DF: Extended squitter / non-transponder",
		adsb.DF19: "DF: Military extended squitter",
		adsb.DF20: "DF: Comm-B altitude reply",
		adsb.DF21: "DF: Comm-B identify reply",
		adsb.DF22: "DF: Reserved for military use",
		adsb.DF24: "DF: Comm-D (ELM)",
		99:        "DF: Unknown value 99",
	} {
		var result = fmt.Sprintf("DF: %s", val)
		if result != out {
			t.Errorf("%s: expected %s | received %s\n", val, out, result)
		}
	}
}

func testCA(t *testing.T) {
	for val, out := range map[adsb.CA]string{
		adsb.CA0: "CA: Level 1",
		adsb.CA4: "CA: Level 2 - On ground",
		adsb.CA5: "CA: Level 2 - Airborne",
		adsb.CA6: "CA: Level 2",
		adsb.CA7: "CA: CA=7",
		99:       "CA: Unknown value 99",
	} {
		var result = fmt.Sprintf("CA: %s", val)
		if result != out {
			t.Errorf("%s: expected %s | received %s\n", val, out, result)
		}
	}
}

func testCF(t *testing.T) {
	for val, out := range map[adsb.CF]string{
		adsb.CF0: "CF: ADS-B message, non-transponder device with ICAO address",
		adsb.CF1: "CF: ADS-B message, anonymous, ground vehicle or obstruction address",
		adsb.CF2: "CF: Fine format TIS-B message",
		adsb.CF3: "CF: Coarse format TIS-B message",
		adsb.CF4: "CF: TIS-B or ADS-R management message",
		adsb.CF5: "CF: Fine format TIS-B message, non-ICAO address",
		adsb.CF6: "CF: ADS-B rebroadcast message",
		99:       "CF: Unknown value 99",
	} {
		var result = fmt.Sprintf("CF: %s", val)
		if result != out {
			t.Errorf("%s: expected %s | received %s\n", val, out, result)
		}
	}
}
func testFS(t *testing.T) {
	for val, out := range map[adsb.FS]string{
		adsb.FS0: "FS: No alert, no SPI, airborne",
		adsb.FS1: "FS: No alert, no SPI, on ground",
		adsb.FS2: "FS: Alert, no SPI, airborne",
		adsb.FS3: "FS: Alert, no SPI, on ground",
		adsb.FS4: "FS: Alert, SPI",
		adsb.FS5: "FS: No alert, SPI",
		99:       "FS: Unknown value 99",
	} {
		var result = fmt.Sprintf("FS: %s", val)
		if result != out {
			t.Errorf("%s: expected %s | received %s\n", val, out, result)
		}
	}
}

func testVS(t *testing.T) {
	for val, out := range map[adsb.VS]string{
		adsb.VS0: "VS: Airborne",
		adsb.VS1: "VS: On ground",
		99:       "VS: Unknown value 99",
	} {
		var result = fmt.Sprintf("VS: %s", val)
		if result != out {
			t.Errorf("%s: expected %s | received %s\n", val, out, result)
		}
	}
}

func testTC(t *testing.T) {
	for val, out := range map[adsb.TC]string{
		adsb.TC0:  "TC: No position information",
		adsb.TC1:  "TC: Identification (Category Set D)",
		adsb.TC2:  "TC: Identification (Category Set C)",
		adsb.TC3:  "TC: Identification (Category Set B)",
		adsb.TC4:  "TC: Identification (Category Set A)",
		adsb.TC5:  "TC: Surface position, 7.5 meter",
		adsb.TC6:  "TC: Surface position, 25 meter",
		adsb.TC7:  "TC: Surface position, 0.1 NM",
		adsb.TC8:  "TC: Surface position",
		adsb.TC9:  "TC: Airborne position, 7.5 meter, barometric altitude",
		adsb.TC10: "TC: Airborne position, 25 meter, barometric altitude",
		adsb.TC11: "TC: Airborne position, 0.1 NM, barometric altitude",
		adsb.TC12: "TC: Airborne position, 0.2 NM, barometric altitude",
		adsb.TC13: "TC: Airborne position, 0.5 NM, barometric altitude",
		adsb.TC14: "TC: Airborne position, 1.0 NM, barometric altitude",
		adsb.TC15: "TC: Airborne position, 2.0 NM, barometric altitude",
		adsb.TC16: "TC: Airborne position, 10 NM, barometric altitude",
		adsb.TC17: "TC: Airborne position, 20 NM, barometric altitude",
		adsb.TC18: "TC: Airborne position, barometric altitude",
		adsb.TC19: "TC: Airborne velocity",
		adsb.TC20: "TC: Airborne position, 7.5 meter, GNSS height",
		adsb.TC21: "TC: Airborne position, 25 meter, GNSS height",
		adsb.TC22: "TC: Airborne position, GNSS height",
		adsb.TC28: "TC: Emergency priority status",
		adsb.TC31: "TC: Operational status",
		99:        "TC: Unknown value 99",
	} {
		var result = fmt.Sprintf("TC: %s", val)
		if result != out {
			t.Errorf("%s: expected %s | received %s\n", val, out, result)
		}
	}
}

func testSS(t *testing.T) {
	for val, out := range map[adsb.SS]string{
		adsb.SS0: "SS: No condition information",
		adsb.SS1: "SS: Permanent alert (emergency)",
		adsb.SS2: "SS: Temporary alert (ident change)",
		adsb.SS3: "SS: SPI",
		99:       "SS: Unknown value 99",
	} {
		var result = fmt.Sprintf("SS: %s", val)
		if result != out {
			t.Errorf("%s: expected %s | received %s\n", val, out, result)
		}
	}
}

func testAcCat(t *testing.T) {
	for val, out := range map[adsb.AcCat]string{
		adsb.A0: "AcCat: No ADS-B emitter category information",
		adsb.A1: "AcCat: Light (< 15500 lbs)",
		adsb.A2: "AcCat: Small (15500 to 75000 lbs)",
		adsb.A3: "AcCat: Large (75000 to 300000 lbs)",
		adsb.A4: "AcCat: High vortex large (aircraft such as B-757)",
		adsb.A5: "AcCat: Heavy (> 300000 lbs)",
		adsb.A6: "AcCat: High performance (> 5g acceleration and 400 kts)",
		adsb.A7: "AcCat: Rotorcraft",
		adsb.B0: "AcCat: No ADS-B emitter category information",
		adsb.B1: "AcCat: Glider / sailplane",
		adsb.B2: "AcCat: Lighter-than-air",
		adsb.B3: "AcCat: Parachutist / skydiver",
		adsb.B4: "AcCat: Ultralight / hang-glider / paraglider",
		adsb.B5: "AcCat: Reserved",
		adsb.B6: "AcCat: Unmanned aerial vehicle",
		adsb.B7: "AcCat: Space / trans-atmospheric vehicle",
		adsb.C0: "AcCat: No ADS-B emitter category information",
		adsb.C1: "AcCat: Surface vehicle – emergency vehicle",
		adsb.C2: "AcCat: Surface vehicle – service vehicle",
		adsb.C3: "AcCat: Point obstacle (includes tethered balloons)",
		adsb.C4: "AcCat: Cluster obstacle",
		adsb.C5: "AcCat: Line obstacle",
		adsb.C6: "AcCat: Reserved",
		adsb.C7: "AcCat: Reserved",
		"99":    "AcCat: Unknown value 99",
	} {
		var result = fmt.Sprintf("AcCat: %s", val)
		if result != out {
			t.Errorf("%s: expected %s | received %s\n", val, out, result)
		}
	}
}
