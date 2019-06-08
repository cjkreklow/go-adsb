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
