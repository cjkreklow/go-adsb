// Copyright 2024 Collin Kreklow
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

package adsbtype

import (
	"fmt"
)

// ATS is the altitude type subfield.
type ATS uint64

// Altitude Type Subfield values.
const (
	ATS0 ATS = 0 // Barometric altitude
	ATS1 ATS = 1 // Navigation-derived altitude
)

var mATS = map[ATS]string{
	ATS0: "Barometric altitude",
	ATS1: "Navigation-derived altitude",
}

// String representation of ATS.
func (c ATS) String() string {
	if str, ok := mATS[c]; ok {
		return str
	}

	return fmt.Sprintf("Unknown value %d", c)
}

// BDS is the Comm-B data selector.
type BDS uint64

// Comm-B Data Selector values.
const (
	BDS02 BDS = 0x02 // Linked Comm-B, segment 2
	BDS03 BDS = 0x03 // Linked Comm-B, segment 3
	BDS04 BDS = 0x04 // Linked Comm-B, segment 4
	BDS05 BDS = 0x05 // Extended squitter airborne position
	BDS06 BDS = 0x06 // Extended squitter surface position
	BDS07 BDS = 0x07 // Extended squitter status
	BDS08 BDS = 0x08 // Extended squitter identification and category
	BDS09 BDS = 0x09 // Extended squitter airborne velocity
	BDS0A BDS = 0x0A // Extended squitter event-driven information
	BDS0B BDS = 0x0B // Air / air information 1 (aircraft state)
	BDS0C BDS = 0x0C // Air / air information 2 (aircraft intent)

	BDS10 BDS = 0x10 // Data link capability report

	BDS17 BDS = 0x17 // Common usage GICB capability report

	BDS20 BDS = 0x20 // Aircraft identification
	BDS21 BDS = 0x21 // Aircraft and airline registration markings
	BDS22 BDS = 0x22 // Antenna positions

	BDS25 BDS = 0x25 // Aircraft type

	BDS30 BDS = 0x30 // ACAS active resolution advisory

	BDS40 BDS = 0x40 // Selected vertical intention
	BDS41 BDS = 0x41 // Next waypoint identifier
	BDS42 BDS = 0x42 // Next waypoint position
	BDS43 BDS = 0x43 // Next waypoint information
	BDS44 BDS = 0x44 // Meteorological routine air report
	BDS45 BDS = 0x45 // Meteorological hazard report

	BDS48 BDS = 0x48 // VHF channel report

	BDS50 BDS = 0x50 // Track and turn report
	BDS51 BDS = 0x51 // Position report coarse
	BDS52 BDS = 0x52 // Position report fine
	BDS53 BDS = 0x53 // Air-referenced state vector
	BDS54 BDS = 0x54 // Waypoint 1
	BDS55 BDS = 0x55 // Waypoint 2
	BDS56 BDS = 0x56 // Waypoint 3

	BDS5F BDS = 0x5F // Quasi-static parameter monitoring
	BDS60 BDS = 0x60 // Heading and speed report
	BDS61 BDS = 0x61 // Extended squitter emergency / priority status

	BDS65 BDS = 0x65 // Extended squitter aircraft operational status

	BDSE3 BDS = 0xE3 // Transponder type / part number
	BDSE4 BDS = 0xE4 // Transponder software revision number
	BDSE5 BDS = 0xE5 // ACAS unit part number
	BDSE6 BDS = 0xE6 // ACAS unit software revision number
	BDSE7 BDS = 0xE7 // Transponder status and diagnostics

	BDSEA BDS = 0xEA // Vendor specific status and diagnostics

	BDSF1 BDS = 0xF1 // Military applications (F1)
	BDSF2 BDS = 0xF2 // Military applications (F2)
)

var mBDS = map[BDS]string{
	BDS02: "Linked Comm-B, segment 2",
	BDS03: "Linked Comm-B, segment 3",
	BDS04: "Linked Comm-B, segment 4",
	BDS05: "Extended squitter airborne position",
	BDS06: "Extended squitter surface position",
	BDS07: "Extended squitter status",
	BDS08: "Extended squitter identification and category",
	BDS09: "Extended squitter airborne velocity",
	BDS0A: "Extended squitter event-driven information",
	BDS0B: "Air / air information 1 (aircraft state)",
	BDS0C: "Air / air information 2 (aircraft intent)",
	BDS10: "Data link capability report",
	BDS17: "Common usage GICB capability report",
	BDS20: "Aircraft identification",
	BDS21: "Aircraft and airline registration markings",
	BDS22: "Antenna positions",
	BDS25: "Aircraft type",
	BDS30: "ACAS active resolution advisory",
	BDS40: "Selected vertical intention",
	BDS41: "Next waypoint identifier",
	BDS42: "Next waypoint position",
	BDS43: "Next waypoint information",
	BDS44: "Meteorological routine air report",
	BDS45: "Meteorological hazard report",
	BDS48: "VHF channel report",
	BDS50: "Track and turn report",
	BDS51: "Position report coarse",
	BDS52: "Position report fine",
	BDS53: "Air-referenced state vector",
	BDS54: "Waypoint 1",
	BDS55: "Waypoint 2",
	BDS56: "Waypoint 3",
	BDS5F: "Quasi-static parameter monitoring",
	BDS60: "Heading and speed report",
	BDS61: "Extended squitter emergency / priority status",
	BDS65: "Extended squitter aircraft operational status",
	BDSE3: "Transponder type / part number",
	BDSE4: "Transponder software revision number",
	BDSE5: "ACAS unit part number",
	BDSE6: "ACAS unit software revision number",
	BDSE7: "Transponder status and diagnostics",
	BDSEA: "Vendor specific status and diagnostics",
	BDSF1: "Military applications (F1)",
	BDSF2: "Military applications (F2)",
}

// String representation of BDS.
func (c BDS) String() string {
	if str, ok := mBDS[c]; ok {
		return str
	}

	return fmt.Sprintf("Unknown value %02x", uint64(c))
}

// SSS is the surveillance status subfield.
type SSS uint64

// Surveillance Status Subfield values.
const (
	SSS0 SSS = 0 // No condition information
	SSS1 SSS = 1 // Permanent alert (emergency)
	SSS2 SSS = 2 // Temporary alert (ident change)
	SSS3 SSS = 3 // SPI
)

var mSSS = map[SSS]string{
	SSS0: "No condition information",
	SSS1: "Permanent alert (emergency)",
	SSS2: "Temporary alert (ident change)",
	SSS3: "SPI",
}

// String representation of SSS.
func (c SSS) String() string {
	if str, ok := mSSS[c]; ok {
		return str
	}

	return fmt.Sprintf("Unknown value %d", c)
}

// TRS is the transmission rate subfield.
type TRS uint64

// Transmission Rate Subfield values.
const (
	TRS0 TRS = 0 // No capability
	TRS1 TRS = 1 // High surface squitter rate
	TRS2 TRS = 2 // Low surface squitter rate
)

var mTRS = map[TRS]string{
	TRS0: "No capability",
	TRS1: "High surface squitter rate",
	TRS2: "Low surface squitter rate",
}

// String representation of TRS.
func (c TRS) String() string {
	if str, ok := mTRS[c]; ok {
		return str
	}

	return fmt.Sprintf("Unknown value %d", c)
}
