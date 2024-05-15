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

// CA is the capability.
type CA uint64

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

// CC is the Cross-link Capability.
type CC uint64

// Cross-link Capability values.
const (
	CC0 CC = 0 // Not supported
	CC1 CC = 1 // Supported
)

var mCC = map[CC]string{
	CC0: "Not supported",
	CC1: "Supported",
}

// String representation of CC.
func (c CC) String() string {
	if str, ok := mCC[c]; ok {
		return str
	}

	return fmt.Sprintf("Unknown value %d", c)
}

// CF is the control field.
type CF uint64

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

// DF is the downlink format.
type DF uint64

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

// DR is the downlink request.
type DR uint64

// Downlink Request values.
const (
	DR0  DR = 0  // No request
	DR1  DR = 1  // Send Comm-B message
	DR2  DR = 2  // ACAS message available
	DR3  DR = 3  // Comm-B message and ACAS message available
	DR4  DR = 4  // Comm-B broadcast message 1 available
	DR5  DR = 5  // Comm-B broadcast message 2 available
	DR6  DR = 6  // Comm-B broadcast message 1 and ACAS message available
	DR7  DR = 7  // Comm-B broadcast message 2 and ACAS message available
	DR16 DR = 16 // 1 segment ELM message available
	DR17 DR = 17 // 2 segment ELM message available
	DR18 DR = 18 // 3 segment ELM message available
	DR19 DR = 19 // 4 segment ELM message available
	DR20 DR = 20 // 5 segment ELM message available
	DR21 DR = 21 // 6 segment ELM message available
	DR22 DR = 22 // 7 segment ELM message available
	DR23 DR = 23 // 8 segment ELM message available
	DR24 DR = 24 // 9 segment ELM message available
	DR25 DR = 25 // 10 segment ELM message available
	DR26 DR = 26 // 11 segment ELM message available
	DR27 DR = 27 // 12 segment ELM message available
	DR28 DR = 28 // 13 segment ELM message available
	DR29 DR = 29 // 14 segment ELM message available
	DR30 DR = 30 // 15 segment ELM message available
	DR31 DR = 31 // 15 segment ELM message available
)

var mDR = map[DR]string{
	DR0:  "No request",
	DR1:  "Comm-B message available",
	DR2:  "ACAS message available",
	DR3:  "Comm-B message and ACAS message available",
	DR4:  "Comm-B broadcast message 1 available",
	DR5:  "Comm-B broadcast message 2 available",
	DR6:  "Comm-B broadcast message 1 and ACAS message available",
	DR7:  "Comm-B broadcast message 2 and ACAS message available",
	DR16: "1 segment ELM message available",
	DR17: "2 segment ELM message available",
	DR18: "3 segment ELM message available",
	DR19: "4 segment ELM message available",
	DR20: "5 segment ELM message available",
	DR21: "6 segment ELM message available",
	DR22: "7 segment ELM message available",
	DR23: "8 segment ELM message available",
	DR24: "9 segment ELM message available",
	DR25: "10 segment ELM message available",
	DR26: "11 segment ELM message available",
	DR27: "12 segment ELM message available",
	DR28: "13 segment ELM message available",
	DR29: "14 segment ELM message available",
	DR30: "15 segment ELM message available",
	DR31: "15 segment ELM message available",
}

// String representation of DR.
func (c DR) String() string {
	if str, ok := mDR[c]; ok {
		return str
	}

	return fmt.Sprintf("Unknown value %d", c)
}

// FS is the flight status.
type FS uint64

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

// RI is the reply information.
type RI uint64

// Reply Information values.
const (
	RI0  RI = 0  // No ACAS
	RI2  RI = 2  // ACAS with resolution inhibited
	RI3  RI = 3  // ACAS with vertical-only resolution
	RI4  RI = 4  // ACAS with vertical and horizontal resolution
	RI8  RI = 8  // No maximum airspeed available
	RI9  RI = 9  // Maximum airspeed < 75 kt
	RI10 RI = 10 // Maximum airspeed 75 - 150 kt
	RI11 RI = 11 // Maximum airspeed 150 - 300 kt
	RI12 RI = 12 // Maximum airspeed 300 - 600 kt
	RI13 RI = 13 // Maximum airspeed 600 - 1200 kt
	RI14 RI = 14 // Maximum airspeed > 1200 kt
)

var mRI = map[RI]string{
	RI0:  "No ACAS",
	RI2:  "ACAS with resolution inhibited",
	RI3:  "ACAS with vertical-only resolution",
	RI4:  "ACAS with vertical and horizontal resolution",
	RI8:  "No maximum airspeed available",
	RI9:  "Maximum airspeed < 75 kt",
	RI10: "Maximum airspeed 75 - 150 kt",
	RI11: "Maximum airspeed 150 - 300 kt",
	RI12: "Maximum airspeed 300 - 600 kt",
	RI13: "Maximum airspeed 600 - 1200 kt",
	RI14: "Maximum airspeed > 1200 kt",
}

// String representation of RI.
func (c RI) String() string {
	if str, ok := mRI[c]; ok {
		return str
	}

	return fmt.Sprintf("Unknown value %d", c)
}

// SL is the sensitivity level report.
type SL uint64

// Sensitivity Level values.
const (
	SL0 SL = 0 // ACAS inoperative
	SL1 SL = 1 // ACAS sensitivity level 1
	SL2 SL = 2 // ACAS sensitivity level 2
	SL3 SL = 3 // ACAS sensitivity level 3
	SL4 SL = 4 // ACAS sensitivity level 4
	SL5 SL = 5 // ACAS sensitivity level 5
	SL6 SL = 6 // ACAS sensitivity level 6
	SL7 SL = 7 // ACAS sensitivity level 7
)

var mSL = map[SL]string{
	SL0: "ACAS inoperative",
	SL1: "ACAS sensitivity level 1",
	SL2: "ACAS sensitivity level 2",
	SL3: "ACAS sensitivity level 3",
	SL4: "ACAS sensitivity level 4",
	SL5: "ACAS sensitivity level 5",
	SL6: "ACAS sensitivity level 6",
	SL7: "ACAS sensitivity level 7",
}

// String representation of SL.
func (c SL) String() string {
	if str, ok := mSL[c]; ok {
		return str
	}

	return fmt.Sprintf("Unknown value %d", c)
}

// VS is the vertical status.
type VS uint64

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
