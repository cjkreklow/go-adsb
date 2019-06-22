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
	"encoding/hex"
	"math/big"
	"testing"
)

type testCase struct {
	Msg string

	DF int
	CA int
	FS int
	VS int

	TC  int
	SS  int
	Cat uint8

	CPR       bool
	LocalPos  bool
	RefPt     []float64
	GlobalPos bool
	Msg2      string

	Lat float64
	Lon float64

	ICAO string
	Alt  int64
	Sqk  string
	Call string
}

// TestDecode runs the test cases for message decoding
func TestDecode(t *testing.T) {
	t.Run("DF0", testDF0)
	t.Run("DF4", testDF4A)
	t.Run("DF4 Gillham", testDF4B)
	t.Run("DF5", testDF5)
	t.Run("DF11", testDF11)
	t.Run("DF17 Position Local", testDF17PosLocal)
	t.Run("DF17 Position Global", testDF17PosGlobal)
	t.Run("DF17 Position Global Reverse", testDF17PosGlobalRev)
	t.Run("DF17 Identity", testDF17Ident)
	t.Run("DF20", testDF20)
	t.Run("DF21", testDF21)
}

// test DF0 air-to-air surveillance
func testDF0(t *testing.T) {
	tc := &testCase{
		Msg: "02e19718e70f6c",

		DF: 0,
		CA: -1,
		FS: -1,
		VS: 0,

		TC:  -1,
		SS:  -1,
		Cat: 0,

		CPR: false,

		ICAO: "abd94d",
		Alt:  36000,
		Sqk:  "",
		Call: "",
	}

	testDecode(t, tc)
}

// test DF4 with 25ft altitude report
func testDF4A(t *testing.T) {
	tc := &testCase{
		Msg: "20001910bc45e9",

		DF: 4,
		CA: -1,
		FS: 0,
		VS: -1,

		TC:  -1,
		SS:  -1,
		Cat: 0,

		CPR: false,

		ICAO: "a27aee",
		Alt:  39000,
		Sqk:  "",
		Call: "",
	}

	testDecode(t, tc)
}

// test DF4 with a Gillham-encoded altitude
func testDF4B(t *testing.T) {
	tc := &testCase{
		Msg: "2000042210fc86",

		DF: 4,
		CA: -1,
		FS: 0,
		VS: -1,

		TC:  -1,
		SS:  -1,
		Cat: 0,

		CPR: false,

		ICAO: "a97172",
		Alt:  2000,
		Sqk:  "",
		Call: "",
	}

	testDecode(t, tc)
}

// test DF5 identity reply
func testDF5(t *testing.T) {
	tc := &testCase{
		Msg: "28001b0601970d",

		DF: 5,
		CA: -1,
		FS: 0,
		VS: -1,

		TC:  -1,
		SS:  -1,
		Cat: 0,

		CPR: false,

		ICAO: "a3696e",
		Alt:  0,
		Sqk:  "3452",
		Call: "",
	}

	testDecode(t, tc)
}

// test DF11 all call reply
func testDF11(t *testing.T) {
	tc := &testCase{
		Msg: "5dac22c54b7a07",

		DF: 11,
		CA: 5,
		FS: -1,
		VS: -1,

		TC:  -1,
		SS:  -1,
		Cat: 0,

		CPR: false,

		ICAO: "ac22c5",
		Alt:  0,
		Sqk:  "",
		Call: "",
	}

	testDecode(t, tc)
}

// test DF20 Comm-B altitude reply
func testDF20(t *testing.T) {
	tc := &testCase{
		Msg: "a0000f9820057273df8d20e2cf30",

		DF: 20,
		CA: -1,
		FS: 0,
		VS: -1,

		TC:  -1,
		SS:  -1,
		Cat: 0,

		CPR: false,

		ICAO: "a52333",
		Alt:  24000,
		Sqk:  "",
		Call: "AWI3784",
	}

	testDecode(t, tc)
}

// test DF21 Comm-B identity reply
func testDF21(t *testing.T) {
	tc := &testCase{
		Msg: "ac19b29573482f6963663636022b",

		DF: 21,
		CA: -1,
		FS: 4,
		VS: -1,

		TC:  -1,
		SS:  -1,
		Cat: 0,

		CPR: false,

		ICAO: "a97db4",
		Alt:  0,
		Sqk:  "6017",
		Call: "",
	}

	testDecode(t, tc)
}

// test DF17 extended squitter position, local decode
func testDF17PosLocal(t *testing.T) {
	tc := &testCase{
		Msg: "8da9450d60bde138e8638c939134",

		DF: 17,
		CA: 5,
		FS: -1,
		VS: -1,

		TC:  12,
		SS:  0,
		Cat: 0,

		CPR:      true,
		LocalPos: true,
		RefPt:    []float64{43.14, -89.33},

		Lat: 43.83300781,
		Lon: -90.46484375,

		ICAO: "a9450d",
		Alt:  36950,
		Sqk:  "",
		Call: "",
	}

	testDecode(t, tc)
}

// test DF17 extended squitter position, global decode
func testDF17PosGlobal(t *testing.T) {
	tc := &testCase{
		Msg: "8da8028758ab0028de078689d437",

		DF: 17,
		CA: 5,
		FS: -1,
		VS: -1,

		TC:  11,
		SS:  0,
		Cat: 0,

		CPR:       true,
		GlobalPos: true,
		Msg2:      "8da8028758ab07b0b8876e81eb25",

		Lat: 42.23945229,
		Lon: -89.87851165,

		ICAO: "a80287",
		Alt:  33000,
		Sqk:  "",
		Call: "",
	}

	testDecode(t, tc)
}

// test DF17 extended squitter position, global decode, reversed
func testDF17PosGlobalRev(t *testing.T) {
	tc := &testCase{
		Msg: "8dab9448589ff40a4e62a6c8b7a6",

		DF: 17,
		CA: 5,
		FS: -1,
		VS: -1,

		TC:  11,
		SS:  0,
		Cat: 0,

		CPR:       true,
		GlobalPos: true,
		Msg2:      "8dab9448589ff083bbe2387219c5",

		Lat: 42.77183532,
		Lon: -90.47590775,

		ICAO: "ab9448",
		Alt:  30975,
		Sqk:  "",
		Call: "",
	}

	testDecode(t, tc)
}

// test DF17 extended squitter identity
func testDF17Ident(t *testing.T) {
	tc := &testCase{
		Msg: "8dacf84e23101332cf3ca037ef13",

		DF: 17,
		CA: 5,
		FS: -1,
		VS: -1,

		TC:  4,
		SS:  -1,
		Cat: 3,

		CPR: false,

		ICAO: "acf84e",
		Alt:  0,
		Sqk:  "",
		Call: "DAL2332",
	}

	testDecode(t, tc)
}

func testDecode(t *testing.T, tc *testCase) {
	b, err := hex.DecodeString(tc.Msg)
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	msg := new(Message)
	err = msg.Decode(b)
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	if msg.DF != DF(tc.DF) {
		t.Errorf("DF: received %v, expected %v", int(msg.DF), tc.DF)
	}
	if msg.CA != CA(tc.CA) {
		t.Errorf("CA: received %v, expected %v", int(msg.CA), tc.CA)
	}
	if msg.FS != FS(tc.FS) {
		t.Errorf("FS: received %v, expected %v", int(msg.FS), tc.FS)
	}
	if msg.VS != VS(tc.VS) {
		t.Errorf("VS: received %v, expected %v", int(msg.VS), tc.VS)
	}
	if msg.TC != TC(tc.TC) {
		t.Errorf("TC: received %v, expected %v", int(msg.TC), tc.TC)
	}
	if msg.SS != SS(tc.SS) {
		t.Errorf("SS: received %v, expected %v", int(msg.SS), tc.SS)
	}
	if msg.Cat != tc.Cat {
		t.Errorf("Cat: received %v, expected %v", msg.Cat, tc.Cat)
	}
	if msg.ICAO != tc.ICAO {
		t.Errorf("ICAO: received %s, expected %s", msg.ICAO, tc.ICAO)
	}
	if msg.Alt != tc.Alt {
		t.Errorf("Alt: received %v, expected %v", msg.Alt, tc.Alt)
	}
	if msg.Sqk != tc.Sqk {
		t.Errorf("Sqk: received %s, expected %s", msg.Sqk, tc.Sqk)
	}
	if msg.Call != tc.Call {
		t.Errorf("Call: received %s, expected %s", msg.Call, tc.Call)
	}

	if !tc.CPR && msg.CPR != nil {
		t.Error("CPR: unexpected position report populated")
	}

	if tc.CPR && tc.LocalPos {
		if msg.CPR == nil {
			t.Error("CPR: expected but not present")
		} else {
			c, err := msg.CPR.DecodeLocal(tc.RefPt)
			if err != nil {
				t.Error("CPR: local decode error:", err)
			} else {
				eLat := big.NewFloat(tc.Lat)
				eLat.SetPrec(16)
				cLat := big.NewFloat(c[0])
				cLat.SetPrec(16)
				if eLat.Cmp(cLat) != 0 {
					t.Errorf("Lat: received %s, expected %s", cLat.String(), eLat.String())
				}
				eLon := big.NewFloat(tc.Lon)
				eLon.SetPrec(16)
				cLon := big.NewFloat(c[1])
				cLon.SetPrec(16)
				if eLon.Cmp(cLon) != 0 {
					t.Errorf("Lon: received %s, expected %s", cLon.String(), eLon.String())
				}
			}
		}
	}

	if tc.CPR && tc.GlobalPos {
		if msg.CPR == nil {
			t.Error("CPR: expected but not present")
		} else {
			m2 := new(Message)
			m2b, err := (hex.DecodeString(tc.Msg2))
			if err != nil {
				t.Fatal("received unexpected error", err)
			}
			err = m2.Decode(m2b)
			if err != nil {
				t.Fatal("received unexpected error", err)
			}
			if m2.CPR == nil {
				t.Fatal("no position decoded in Msg2")
			}
			c, err := DecodeGlobalPosition(msg.CPR, m2.CPR)
			if err != nil {
				t.Error("CPR: global decode error:", err)
			} else {
				eLat := big.NewFloat(tc.Lat)
				eLat.SetPrec(16)
				cLat := big.NewFloat(c[0])
				cLat.SetPrec(16)
				if eLat.Cmp(cLat) != 0 {
					t.Errorf("Lat: received %s, expected %s", cLat.String(), eLat.String())
				}
				eLon := big.NewFloat(tc.Lon)
				eLon.SetPrec(16)
				cLon := big.NewFloat(c[1])
				cLon.SetPrec(16)
				if eLon.Cmp(cLon) != 0 {
					t.Errorf("Lon: received %s, expected %s", cLon.String(), eLon.String())
				}
			}
		}
	}
}
