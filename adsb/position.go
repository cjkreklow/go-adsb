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
	"errors"
	"math"
)

// CPR is an extended squitter compact position report
type CPR struct {
	Nb  uint8  // number of encoded bits (17, 19, 14 or 12)
	T   uint8  // time bit
	F   uint8  // format bit
	Lat uint32 // encoded latitude
	Lon uint32 // encoded longitude
}

// DecodeLocal decodes an encoded position to a global latitude and
// longitude by comparing the position to a known reference point.
// Argument and return value is in the format [latitude, longitude].
func (c *CPR) DecodeLocal(rp []float64) ([]float64, error) {
	if len(rp) != 2 {
		return nil, errors.New("must provide [lat, lon] as argument")
	}
	if rp[0] > 90 || rp[0] < -90 {
		return nil, errors.New("latitude out of range (-90 to 90)")
	}
	if rp[1] > 190 || rp[1] < -180 {
		return nil, errors.New("longitude out of range (-180 to 180)")
	}

	latr := rp[0]
	lonr := rp[1]
	latc := float64(c.Lat) / 131072
	lonc := float64(c.Lon) / 131072

	dlat := 360.0 / float64(60-c.F)

	j := math.Floor(latr/dlat) +
		math.Floor((mod(latr, dlat)/dlat)-latc+0.5)

	coord := make([]float64, 2)

	coord[0] = dlat * (j + latc)

	var dlon float64

	nl := float64(cprNL(coord[0]) - c.F)

	if nl == 0 {
		dlon = 360.0
	} else {
		dlon = 360.0 / nl
	}

	m := math.Floor(lonr/dlon) +
		math.Floor((mod(lonr, dlon)/dlon)-lonc+0.5)

	coord[1] = dlon * (m + lonc)

	return coord, nil
}

// DecodeGlobalPosition decodes an encoded position to a globally
// unabmiguous latitude and longitude by combining two CPR messages.
// The two messages must have different formats (CPR.F) and must have
// a time difference of less than 10 seconds (3 NM distance). The
// return value is in the format [latitude, longitude].
func DecodeGlobalPosition(c1 *CPR, c2 *CPR) ([]float64, error) {
	if c1 == nil || c2 == nil {
		return nil, errors.New("incomplete arguments")
	}
	if c1.Nb != c2.Nb {
		return nil, errors.New("bit encoding must be equal")
	}
	if c1.F == c2.F {
		return nil, errors.New("format must be different")
	}

	var t0 bool // set t0 to true if the even format is the later message
	var lat0, lon0, lat1, lon1 float64

	if c1.F == 0 {
		t0 = false
		lat0 = float64(c1.Lat) / 131072
		lon0 = float64(c1.Lon) / 131072
		lat1 = float64(c2.Lat) / 131072
		lon1 = float64(c2.Lon) / 131072
	} else {
		t0 = true
		lat0 = float64(c2.Lat) / 131072
		lon0 = float64(c2.Lon) / 131072
		lat1 = float64(c1.Lat) / 131072
		lon1 = float64(c1.Lon) / 131072
	}

	dlat0 := 360.0 / 60.0
	dlat1 := 360.0 / 59.0

	// 2**17 = 131072
	j := math.Floor(((59 * lat0) - (60 * lat1)) + 0.5)

	rlat0 := dlat0 * (mod(j, 60) + lat0)
	if rlat0 >= 270 {
		rlat0 -= 360
	}

	rlat1 := dlat1 * (mod(j, 59) + lat1)
	if rlat1 >= 270 {
		rlat1 -= 360
	}

	if cprNL(rlat0) != cprNL(rlat1) {
		return nil, errors.New("positions cross latitude boundary")
	}

	var nl, ni, dlon, lonc float64
	coord := make([]float64, 2)

	if t0 {
		coord[0] = rlat0
		nl = float64(cprNL(rlat0))
		if nl <= 1 {
			ni = 1
		} else {
			ni = nl
		}
		dlon = 360.0 / ni
		lonc = lon0
	} else {
		coord[0] = rlat1
		nl = float64(cprNL(rlat1))
		if nl-1 <= 1 {
			ni = 1
		} else {
			ni = nl - 1
		}
		dlon = 360.0 / ni
		lonc = lon1
	}

	m := math.Floor(((lon0 * (nl - 1)) - (lon1 * nl)) + 0.5)

	coord[1] = dlon * (mod(m, ni) + lonc)
	if coord[1] >= 180 {
		coord[1] -= 360
	}

	return coord, nil
}

// mod implements the MOD function as defined in the ADS-B
// specifications, specifically adding 360 degrees to a negative value
// prior to caclulating the remainder.
func mod(a float64, b float64) float64 {
	return a - (b * math.Floor(a/b))
}

// cprNL implements the longitude zone lookup table
func cprNL(x float64) uint8 {
	/* Lookup table computed with the following code:

	tbl := make(map[int]float64)

	for nl := 59; nl > 1; nl-- {
		a := 1 - math.Cos(math.Pi/30)
		b := 1 - math.Cos(2*math.Pi/float64(nl))
		c := math.Sqrt(a / b)

		tbl[nl] = (180 / math.Pi) * math.Acos(c)

		fmt.Printf("case x < %s: return %d\n", big.NewFloat(tbl[nl]).String(), nl)
	}
	*/

	switch x = math.Abs(x); {
	case x < 10.4704713:
		return 59
	case x < 14.82817437:
		return 58
	case x < 18.18626357:
		return 57
	case x < 21.02939493:
		return 56
	case x < 23.54504487:
		return 55
	case x < 25.82924707:
		return 54
	case x < 27.9389871:
		return 53
	case x < 29.91135686:
		return 52
	case x < 31.77209708:
		return 51
	case x < 33.53993436:
		return 50
	case x < 35.22899598:
		return 49
	case x < 36.85025108:
		return 48
	case x < 38.41241892:
		return 47
	case x < 39.92256684:
		return 46
	case x < 41.38651832:
		return 45
	case x < 42.80914012:
		return 44
	case x < 44.19454951:
		return 43
	case x < 45.54626723:
		return 42
	case x < 46.86733252:
		return 41
	case x < 48.16039128:
		return 40
	case x < 49.42776439:
		return 39
	case x < 50.67150166:
		return 38
	case x < 51.89342469:
		return 37
	case x < 53.09516153:
		return 36
	case x < 54.27817472:
		return 35
	case x < 55.44378444:
		return 34
	case x < 56.59318756:
		return 33
	case x < 57.72747354:
		return 32
	case x < 58.84763776:
		return 31
	case x < 59.95459277:
		return 30
	case x < 61.04917774:
		return 29
	case x < 62.13216659:
		return 28
	case x < 63.20427479:
		return 27
	case x < 64.26616523:
		return 26
	case x < 65.3184531:
		return 25
	case x < 66.36171008:
		return 24
	case x < 67.39646774:
		return 23
	case x < 68.42322022:
		return 22
	case x < 69.44242631:
		return 21
	case x < 70.45451075:
		return 20
	case x < 71.45986473:
		return 19
	case x < 72.45884545:
		return 18
	case x < 73.45177442:
		return 17
	case x < 74.43893416:
		return 16
	case x < 75.42056257:
		return 15
	case x < 76.39684391:
		return 14
	case x < 77.36789461:
		return 13
	case x < 78.33374083:
		return 12
	case x < 79.29428225:
		return 11
	case x < 80.24923213:
		return 10
	case x < 81.19801349:
		return 9
	case x < 82.13956981:
		return 8
	case x < 83.07199445:
		return 7
	case x < 83.99173563:
		return 6
	case x < 84.89166191:
		return 5
	case x < 85.75541621:
		return 4
	case x < 86.53536998:
		return 3
	case x < 87:
		return 2
	default:
		return 1
	}
}
