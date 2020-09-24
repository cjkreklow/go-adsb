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

package adsbtype

import (
	"fmt"
)

// TYPE is the extended squitter type.
type TYPE uint64

// Extended squitter type values.
const (
	TYPE0  TYPE = 0  // No position information
	TYPE1  TYPE = 1  // Identification (Category Set D)
	TYPE2  TYPE = 2  // Identification (Category Set C)
	TYPE3  TYPE = 3  // Identification (Category Set B)
	TYPE4  TYPE = 4  // Identification (Category Set A)
	TYPE5  TYPE = 5  // Surface position, 7.5 meter
	TYPE6  TYPE = 6  // Surface position, 25 meter
	TYPE7  TYPE = 7  // Surface position, 0.1 NM
	TYPE8  TYPE = 8  // Surface position
	TYPE9  TYPE = 9  // Airborne position, 7.5 meter, barometric altitude
	TYPE10 TYPE = 10 // Airborne position, 25 meter, barometric altitude
	TYPE11 TYPE = 11 // Airborne position, 0.1 NM, barometric altitude
	TYPE12 TYPE = 12 // Airborne position, 0.2 NM, barometric altitude
	TYPE13 TYPE = 13 // Airborne position, 0.5 NM, barometric altitude
	TYPE14 TYPE = 14 // Airborne position, 1.0 NM, barometric altitude
	TYPE15 TYPE = 15 // Airborne position, 2.0 NM, barometric altitude
	TYPE16 TYPE = 16 // Airborne position, 10 NM, barometric altitude
	TYPE17 TYPE = 17 // Airborne position, 20 NM, barometric altitude
	TYPE18 TYPE = 18 // Airborne position, barometric altitude
	TYPE19 TYPE = 19 // Airborne velocity
	TYPE20 TYPE = 20 // Airborne position, 7.5 meter, GNSS height
	TYPE21 TYPE = 21 // Airborne position, 25 meter, GNSS height
	TYPE22 TYPE = 22 // Airborne position, GNSS height
	TYPE28 TYPE = 28 // Emergency priority status
	TYPE31 TYPE = 31 // Operational status
)

var mTYPE = map[TYPE]string{
	TYPE0:  "No position information",
	TYPE1:  "Identification (Category Set D)",
	TYPE2:  "Identification (Category Set C)",
	TYPE3:  "Identification (Category Set B)",
	TYPE4:  "Identification (Category Set A)",
	TYPE5:  "Surface position, 7.5 meter",
	TYPE6:  "Surface position, 25 meter",
	TYPE7:  "Surface position, 0.1 NM",
	TYPE8:  "Surface position",
	TYPE9:  "Airborne position, 7.5 meter, barometric altitude",
	TYPE10: "Airborne position, 25 meter, barometric altitude",
	TYPE11: "Airborne position, 0.1 NM, barometric altitude",
	TYPE12: "Airborne position, 0.2 NM, barometric altitude",
	TYPE13: "Airborne position, 0.5 NM, barometric altitude",
	TYPE14: "Airborne position, 1.0 NM, barometric altitude",
	TYPE15: "Airborne position, 2.0 NM, barometric altitude",
	TYPE16: "Airborne position, 10 NM, barometric altitude",
	TYPE17: "Airborne position, 20 NM, barometric altitude",
	TYPE18: "Airborne position, barometric altitude",
	TYPE19: "Airborne velocity",
	TYPE20: "Airborne position, 7.5 meter, GNSS height",
	TYPE21: "Airborne position, 25 meter, GNSS height",
	TYPE22: "Airborne position, GNSS height",
	TYPE28: "Emergency priority status",
	TYPE31: "Operational status",
}

// String representation of TYPE.
func (c TYPE) String() string {
	if str, ok := mTYPE[c]; ok {
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
