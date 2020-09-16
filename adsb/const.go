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

import (
	"fmt"
)

// DF is the downlink format of the received Mode S message.
type DF int

// Downlink Format values.
const (
	DF0  DF = 0  // Short air-air surveillance (ACAS)
	DF4  DF = 4  // Surveillance altitude reply
	DF5  DF = 5  // Surveillance identify reply
	DF11 DF = 11 // All-call reply
	DF16 DF = 16 // Long air-air surveillance (ACAS)
	DF17 DF = 17 // Extended squitter
	DF18 DF = 18 // Extended squitter / non-transponder
	DF19 DF = 19 // Military extended squitter
	DF20 DF = 20 // Comm-B altitude reply
	DF21 DF = 21 // Comm-B identify reply
	DF22 DF = 22 // Reserved for military use
	DF24 DF = 24 // Comm-D (ELM)
)

var mDF = map[DF]string{
	DF0:  "Short air-air surveillance (ACAS)",
	DF4:  "Surveillance altitude reply",
	DF5:  "Surveillance identify reply",
	DF11: "All-call reply",
	DF16: "Long air-air surveillance (ACAS)",
	DF17: "Extended squitter",
	DF18: "Extended squitter / non-transponder",
	DF19: "Military extended squitter",
	DF20: "Comm-B altitude reply",
	DF21: "Comm-B identify reply",
	DF22: "Reserved for military use",
	DF24: "Comm-D (ELM)",
}

// String representation of DF.
func (c DF) String() string {
	if str, ok := mDF[c]; ok {
		return str
	}

	return fmt.Sprintf("Unknown value %d", c)
}

// CA is the capability.
type CA int

// Capability values.
const (
	CA0 CA = 0 // Level 1
	CA4 CA = 4 // Level 2 - On ground
	CA5 CA = 5 // Level 2 - Airborne
	CA6 CA = 6 // Level 2
	CA7 CA = 7 // CA=7
)

var mCA = map[CA]string{
	CA0: "Level 1",
	CA4: "Level 2 - On ground",
	CA5: "Level 2 - Airborne",
	CA6: "Level 2",
	CA7: "CA=7",
}

// String representation of CA.
func (c CA) String() string {
	if str, ok := mCA[c]; ok {
		return str
	}

	return fmt.Sprintf("Unknown value %d", c)
}

// CF is the control field.
type CF int

// Control Field values.
const (
	CF0 CF = 0 // ADS-B message, non-transponder device with ICAO address
	CF1 CF = 1 // ADS-B message, anonymous, ground vehicle or obstruction address
	CF2 CF = 2 // Fine format TIS-B message
	CF3 CF = 3 // Coarse format TIS-B message
	CF4 CF = 4 // TIS-B or ADS-R management message
	CF5 CF = 5 // Fine format TIS-B message, non-ICAO address
	CF6 CF = 6 // ADS-B rebroadcast message
)

var mCF = map[CF]string{
	CF0: "ADS-B message, non-transponder device with ICAO address",
	CF1: "ADS-B message, anonymous, ground vehicle or obstruction address",
	CF2: "Fine format TIS-B message",
	CF3: "Coarse format TIS-B message",
	CF4: "TIS-B or ADS-R management message",
	CF5: "Fine format TIS-B message, non-ICAO address",
	CF6: "ADS-B rebroadcast message",
}

// String representation of CF.
func (c CF) String() string {
	if str, ok := mCF[c]; ok {
		return str
	}

	return fmt.Sprintf("Unknown value %d", c)
}

// FS is the flight status.
type FS int

// Flight Status values.
const (
	FS0 FS = 0 // No alert, no SPI, airborne
	FS1 FS = 1 // No alert, no SPI, on ground
	FS2 FS = 2 // Alert, no SPI, airborne
	FS3 FS = 3 // Alert, no SPI, on ground
	FS4 FS = 4 // Alert, SPI
	FS5 FS = 5 // No alert, SPI
)

var mFS = map[FS]string{
	FS0: "No alert, no SPI, airborne",
	FS1: "No alert, no SPI, on ground",
	FS2: "Alert, no SPI, airborne",
	FS3: "Alert, no SPI, on ground",
	FS4: "Alert, SPI",
	FS5: "No alert, SPI",
}

// String representation of FS.
func (c FS) String() string {
	if str, ok := mFS[c]; ok {
		return str
	}

	return fmt.Sprintf("Unknown value %d", c)
}

// VS is the vertical status.
type VS int

// Vertical Status values.
const (
	VS0 VS = 0 // Airborne
	VS1 VS = 1 // On ground
)

var mVS = map[VS]string{
	VS0: "Airborne",
	VS1: "On ground",
}

// String representation of VS.
func (c VS) String() string {
	if str, ok := mVS[c]; ok {
		return str
	}

	return fmt.Sprintf("Unknown value %d", c)
}

// TC is the extended squitter type.
type TC int

// Extended squitter type values.
const (
	TC0  TC = 0  // No position information
	TC1  TC = 1  // Identification (Category Set D)
	TC2  TC = 2  // Identification (Category Set C)
	TC3  TC = 3  // Identification (Category Set B)
	TC4  TC = 4  // Identification (Category Set A)
	TC5  TC = 5  // Surface position, 7.5 meter
	TC6  TC = 6  // Surface position, 25 meter
	TC7  TC = 7  // Surface position, 0.1 NM
	TC8  TC = 8  // Surface position
	TC9  TC = 9  // Airborne position, 7.5 meter, barometric altitude
	TC10 TC = 10 // Airborne position, 25 meter, barometric altitude
	TC11 TC = 11 // Airborne position, 0.1 NM, barometric altitude
	TC12 TC = 12 // Airborne position, 0.2 NM, barometric altitude
	TC13 TC = 13 // Airborne position, 0.5 NM, barometric altitude
	TC14 TC = 14 // Airborne position, 1.0 NM, barometric altitude
	TC15 TC = 15 // Airborne position, 2.0 NM, barometric altitude
	TC16 TC = 16 // Airborne position, 10 NM, barometric altitude
	TC17 TC = 17 // Airborne position, 20 NM, barometric altitude
	TC18 TC = 18 // Airborne position, barometric altitude
	TC19 TC = 19 // Airborne velocity
	TC20 TC = 20 // Airborne position, 7.5 meter, GNSS height
	TC21 TC = 21 // Airborne position, 25 meter, GNSS height
	TC22 TC = 22 // Airborne position, GNSS height
	TC28 TC = 28 // Emergency priority status
	TC31 TC = 31 // Operational status
)

var mTC = map[TC]string{
	TC0:  "No position information",
	TC1:  "Identification (Category Set D)",
	TC2:  "Identification (Category Set C)",
	TC3:  "Identification (Category Set B)",
	TC4:  "Identification (Category Set A)",
	TC5:  "Surface position, 7.5 meter",
	TC6:  "Surface position, 25 meter",
	TC7:  "Surface position, 0.1 NM",
	TC8:  "Surface position",
	TC9:  "Airborne position, 7.5 meter, barometric altitude",
	TC10: "Airborne position, 25 meter, barometric altitude",
	TC11: "Airborne position, 0.1 NM, barometric altitude",
	TC12: "Airborne position, 0.2 NM, barometric altitude",
	TC13: "Airborne position, 0.5 NM, barometric altitude",
	TC14: "Airborne position, 1.0 NM, barometric altitude",
	TC15: "Airborne position, 2.0 NM, barometric altitude",
	TC16: "Airborne position, 10 NM, barometric altitude",
	TC17: "Airborne position, 20 NM, barometric altitude",
	TC18: "Airborne position, barometric altitude",
	TC19: "Airborne velocity",
	TC20: "Airborne position, 7.5 meter, GNSS height",
	TC21: "Airborne position, 25 meter, GNSS height",
	TC22: "Airborne position, GNSS height",
	TC28: "Emergency priority status",
	TC31: "Operational status",
}

// String representation of TC.
func (c TC) String() string {
	if str, ok := mTC[c]; ok {
		return str
	}

	return fmt.Sprintf("Unknown value %d", c)
}

// SS is the extended squitter surveillance status.
type SS int

// Surveillance Status values.
const (
	SS0 SS = 0 // No condition information
	SS1 SS = 1 // Permanent alert (emergency)
	SS2 SS = 2 // Temporary alert (ident change)
	SS3 SS = 3 // SPI
)

var mSS = map[SS]string{
	SS0: "No condition information",
	SS1: "Permanent alert (emergency)",
	SS2: "Temporary alert (ident change)",
	SS3: "SPI",
}

// String representation of SS.
func (c SS) String() string {
	if str, ok := mSS[c]; ok {
		return str
	}

	return fmt.Sprintf("Unknown value %d", c)
}

// AcCat is the extended squitter aircraft emitter category.
type AcCat string

// Extended squitter aircraft emitter category values.
const (
	A0 AcCat = "A0" // No ADS-B emitter category information
	A1 AcCat = "A1" // Light (< 15500 lbs)
	A2 AcCat = "A2" // Small (15500 to 75000 lbs)
	A3 AcCat = "A3" // Large (75000 to 300000 lbs)
	A4 AcCat = "A4" // High vortex large (aircraft such as B-757)
	A5 AcCat = "A5" // Heavy (> 300000 lbs)
	A6 AcCat = "A6" // High performance (> 5g acceleration and 400 kts)
	A7 AcCat = "A7" // Rotorcraft

	B0 AcCat = "B0" // No ADS-B emitter category information
	B1 AcCat = "B1" // Glider / sailplane
	B2 AcCat = "B2" // Lighter-than-air
	B3 AcCat = "B3" // Parachutist / skydiver
	B4 AcCat = "B4" // Ultralight / hang-glider / paraglider
	B5 AcCat = "B5" // Reserved
	B6 AcCat = "B6" // Unmanned aerial vehicle
	B7 AcCat = "B7" // Space / trans-atmospheric vehicle

	C0 AcCat = "C0" // No ADS-B emitter category information
	C1 AcCat = "C1" // Surface vehicle – emergency vehicle
	C2 AcCat = "C2" // Surface vehicle – service vehicle
	C3 AcCat = "C3" // Point obstacle (includes tethered balloons)
	C4 AcCat = "C4" // Cluster obstacle
	C5 AcCat = "C5" // Line obstacle
	C6 AcCat = "C6" // Reserved
	C7 AcCat = "C7" // Reserved
)

var mAcCat = map[AcCat]string{
	A0: "No ADS-B emitter category information",
	A1: "Light (< 15500 lbs)",
	A2: "Small (15500 to 75000 lbs)",
	A3: "Large (75000 to 300000 lbs)",
	A4: "High vortex large (aircraft such as B-757)",
	A5: "Heavy (> 300000 lbs)",
	A6: "High performance (> 5g acceleration and 400 kts)",
	A7: "Rotorcraft",
	B0: "No ADS-B emitter category information",
	B1: "Glider / sailplane",
	B2: "Lighter-than-air",
	B3: "Parachutist / skydiver",
	B4: "Ultralight / hang-glider / paraglider",
	B5: "Reserved",
	B6: "Unmanned aerial vehicle",
	B7: "Space / trans-atmospheric vehicle",
	C0: "No ADS-B emitter category information",
	C1: "Surface vehicle – emergency vehicle",
	C2: "Surface vehicle – service vehicle",
	C3: "Point obstacle (includes tethered balloons)",
	C4: "Cluster obstacle",
	C5: "Line obstacle",
	C6: "Reserved",
	C7: "Reserved",
}

// String representation of AcCat.
func (c AcCat) String() string {
	if str, ok := mAcCat[c]; ok {
		return str
	}

	return fmt.Sprintf("Unknown value %s", string(c))
}
