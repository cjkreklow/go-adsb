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
)

// DF is the downlink format of the received Mode S message
type DF int

// Downlink Format values
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

// String representation of DF
func (c DF) String() string {
	switch c {
	case DF0:
		return "Short air-air surveillance (ACAS)"
	case DF4:
		return "Surveillance altitude reply"
	case DF5:
		return "Surveillance identify reply"
	case DF11:
		return "All-call reply"
	case DF16:
		return "Long air-air surveillance (ACAS)"
	case DF17:
		return "Extended squitter"
	case DF18:
		return "Extended squitter / non-transponder"
	case DF19:
		return "Military extended squitter"
	case DF20:
		return "Comm-B altitude reply"
	case DF21:
		return "Comm-B identify reply"
	case DF22:
		return "Reserved for military use"
	case DF24:
		return "Comm-D (ELM)"
	default:
		return fmt.Sprintf("Unknown value %d", c)
	}
}

// CA is the capability
type CA int

// Capability values
const (
	CA0 CA = 0 // Level 1
	CA4 CA = 4 // Level 2 - On Ground
	CA5 CA = 5 // Level 2 - Airborne
	CA6 CA = 6 // Level 2
	CA7 CA = 7 // CA=7
)

// String representation of CA
func (c CA) String() string {
	switch c {
	case CA0:
		return "Level 1"
	case CA4:
		return "Level 2 - On Ground"
	case CA5:
		return "Level 2 - Airborne"
	case CA6:
		return "Level 2"
	case CA7:
		return "CA=7"
	default:
		return fmt.Sprintf("Unknown value %d", c)
	}
}

// FS is the flight status
type FS int

// Flight Status values
const (
	FS0 FS = 0 // No Alert, No SPI, Airborne
	FS1 FS = 1 // No Alert, No SPI, On Ground
	FS2 FS = 2 // Alert, No SPI, Airborne
	FS3 FS = 3 // Alert, No SPI, On Ground
	FS4 FS = 4 // Alert, SPI
	FS5 FS = 5 // No Alert, SPI
)

// String representation of FS
func (c FS) String() string {
	switch c {
	case FS0:
		return "No Alert, No SPI, Airborne"
	case FS1:
		return "No Alert, No SPI, On Ground"
	case FS2:
		return "Alert, No SPI, Airborne"
	case FS3:
		return "Alert, No SPI, On Ground"
	case FS4:
		return "Alert, SPI"
	case FS5:
		return "No Alert, SPI"
	default:
		return fmt.Sprintf("Unknown value %d", c)
	}
}

// VS is the vertical status
type VS int

// Vertical Status values
const (
	VS0 VS = 0 // Airborne
	VS1 VS = 1 // On Ground
)

// String representation of VS
func (c VS) String() string {
	switch c {
	case VS0:
		return "Airborne"
	case VS1:
		return "On Ground"
	default:
		return fmt.Sprintf("Unknown value %d", c)
	}
}

// TC is the extended squitter type
type TC int

// Extended squitter type values
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

// String representation of TC
func (c TC) String() string {
	switch c {
	case TC0:
		return "No position information"
	case TC1:
		return "Identification (Category Set D)"
	case TC2:
		return "Identification (Category Set C)"
	case TC3:
		return "Identification (Category Set B)"
	case TC4:
		return "Identification (Category Set A)"
	case TC5:
		return "Surface position, 7.5 meter"
	case TC6:
		return "Surface position, 25 meter"
	case TC7:
		return "Surface position, 0.1 NM"
	case TC8:
		return "Surface position"
	case TC9:
		return "Airborne position, 7.5 meter, barometric altitude"
	case TC10:
		return "Airborne position, 25 meter, barometric altitude"
	case TC11:
		return "Airborne position, 0.1 NM, barometric altitude"
	case TC12:
		return "Airborne position, 0.2 NM, barometric altitude"
	case TC13:
		return "Airborne position, 0.5 NM, barometric altitude"
	case TC14:
		return "Airborne position, 1.0 NM, barometric altitude"
	case TC15:
		return "Airborne position, 2.0 NM, barometric altitude"
	case TC16:
		return "Airborne position, 10 NM, barometric altitude"
	case TC17:
		return "Airborne position, 20 NM, barometric altitude"
	case TC18:
		return "Airborne position, barometric altitude"
	case TC19:
		return "Airborne velocity"
	case TC20:
		return "Airborne position, 7.5 meter, GNSS height"
	case TC21:
		return "Airborne position, 25 meter, GNSS height"
	case TC22:
		return "Airborne position, GNSS height"
	case TC28:
		return "Emergency priority status"
	case TC31:
		return "Operational status"
	default:
		return fmt.Sprintf("Unknown value %d", c)
	}
}

// AcCat is the extended squitter aircraft emitter category
type AcCat string

// Extended squitter aircraft emitter category values
const (
	A0 AcCat = "A0" // No ADS-B Emitter Category Information
	A1 AcCat = "A1" // Light (< 15500 lbs)
	A2 AcCat = "A2" // Small (15500 to 75000 lbs)
	A3 AcCat = "A3" // Large (75000 to 300000 lbs)
	A4 AcCat = "A4" // High Vortex Large (aircraft such as B-757)
	A5 AcCat = "A5" // Heavy (> 300000 lbs)
	A6 AcCat = "A6" // High Performance (> 5g acceleration and 400 kts)
	A7 AcCat = "A7" // Rotorcraft

	B0 AcCat = "B0" // No ADS-B Emitter Category Information
	B1 AcCat = "B1" // Glider / sailplane
	B2 AcCat = "B2" // Lighter-than-air
	B3 AcCat = "B3" // Parachutist / Skydiver
	B4 AcCat = "B4" // Ultralight / hang-glider / paraglider
	B5 AcCat = "B5" // Reserved
	B6 AcCat = "B6" // Unmanned Aerial Vehicle
	B7 AcCat = "B7" // Space / Trans-atmospheric vehicle

	C0 AcCat = "C0" // No ADS-B Emitter Category Information
	C1 AcCat = "C1" // Surface Vehicle – Emergency Vehicle
	C2 AcCat = "C2" // Surface Vehicle – Service Vehicle
	C3 AcCat = "C3" // Point Obstacle (includes tethered balloons)
	C4 AcCat = "C4" // Cluster Obstacle
	C5 AcCat = "C5" // Line Obstacle
	C6 AcCat = "C6" // Reserved
	C7 AcCat = "C7" // Reserved
)

// String representation of AcCat
func (c AcCat) String() string {
	switch c {
	case A0:
		return "No ADS-B Emitter Category Information"
	case A1:
		return "Light (< 15500 lbs)"
	case A2:
		return "Small (15500 to 75000 lbs)"
	case A3:
		return "Large (75000 to 300000 lbs)"
	case A4:
		return "High Vortex Large (aircraft such as B-757)"
	case A5:
		return "Heavy (> 300000 lbs)"
	case A6:
		return "High Performance (> 5g acceleration and 400 kts)"
	case A7:
		return "Rotorcraft"
	case B0:
		return "No ADS-B Emitter Category Information"
	case B1:
		return "Glider / sailplane"
	case B2:
		return "Lighter-than-air"
	case B3:
		return "Parachutist / Skydiver"
	case B4:
		return "Ultralight / hang-glider / paraglider"
	case B5:
		return "Reserved"
	case B6:
		return "Unmanned Aerial Vehicle"
	case B7:
		return "Space / Trans-atmospheric vehicle"
	case C0:
		return "No ADS-B Emitter Category Information"
	case C1:
		return "Surface Vehicle – Emergency Vehicle"
	case C2:
		return "Surface Vehicle – Service Vehicle"
	case C3:
		return "Point Obstacle (includes tethered balloons)"
	case C4:
		return "Cluster Obstacle"
	case C5:
		return "Line Obstacle"
	case C6:
		return "Reserved"
	case C7:
		return "Reserved"
	default:
		return fmt.Sprintf("Unknown value %s", string(c))
	}
}

// SS is the extended squitter surveillance status
type SS int

// Surveillance Status values
const (
	SS0 SS = 0 // No condition information
	SS1 SS = 1 // Permanent alert (emergency)
	SS2 SS = 2 // Temporary alert (ident change)
	SS3 SS = 3 // SPI
)

// String representation of SS
func (c SS) String() string {
	switch c {
	case SS0:
		return "No condition information"
	case SS1:
		return "Permanent alert (emergency)"
	case SS2:
		return "Temporary alert (ident change)"
	case SS3:
		return "SPI"
	default:
		return fmt.Sprintf("Unknown value %d", c)
	}
}
