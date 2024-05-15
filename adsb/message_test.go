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

package adsb_test

import (
	"bytes"
	"encoding/hex"
	"errors"
	"math/big"
	"testing"

	"kreklow.us/go/go-adsb/adsb"
)

func TestMessageErrors(t *testing.T) {
	t.Run("EmptyRaw", testMsgEmptyRaw)
	t.Run("Unsupported", testMsgUnsupported)
	t.Run("Unknown", testMsgUnknown)
	t.Run("ICAOError0", testMsgICAOErr0)
	t.Run("ICAOError19", testMsgICAOErr19)
	t.Run("AltError", testMsgAltErr)
	t.Run("CallError", testMsgCallErr)
	t.Run("SqkError", testMsgSqkErr)
	t.Run("CPRError", testMsgCPRErr)
}

func testMsgEmptyRaw(t *testing.T) {
	rm := new(adsb.RawMessage)

	_, err := adsb.NewMessage(rm)
	if err == nil {
		t.Fatal("received nil, expected error")
	}

	if err.Error() != "no data loaded" { //nolint:goconst // test output
		t.Error("received unexpected error", err)
	}
}

func testMsgUnsupported(t *testing.T) {
	raw, err := hex.DecodeString("980000000000ff000000000000ff")
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	rm := new(adsb.RawMessage)

	err = rm.UnmarshalBinary(raw)
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	_, err = adsb.NewMessage(rm)
	if err == nil {
		t.Fatal("received nil, expected error")
	}

	if err.Error() != "downlink format 19: format unsupported" {
		t.Error("received unexpected error", err)
	}

	if !errors.Is(err, adsb.ErrUnsupported) {
		t.Error("expected error type ErrUnsupported not received")
	}
}

func testMsgUnknown(t *testing.T) {
	raw, err := hex.DecodeString("600000000000ff000000000000ff")
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	m := new(adsb.Message)

	err = m.UnmarshalBinary(raw)
	if err == nil {
		t.Fatal("received nil, expected error")
	}

	if err.Error() != "unknown downlink format: 112 bits with format 12" {
		t.Error("received unexpected error", err)
	}
}

func testMsgICAOErr0(t *testing.T) {
	raw, err := hex.DecodeString("00000000000000")
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	m := new(adsb.Message)

	err = m.UnmarshalBinary(raw)
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	rm := m.Raw()

	err = rm.UnmarshalBinary([]byte{})
	if err == nil {
		t.Fatal("received nil, expected error")
	}

	if err.Error() != "no data loaded" {
		t.Error("received unexpected error", err)
	}

	icao, err := m.ICAO()
	if err == nil {
		t.Fatal("received nil, expected error")
	}

	if err.Error() != "no data loaded" {
		t.Error("received unexpected error", err)
	}

	if icao != 0 {
		t.Errorf("expected 0, received %x", icao)
	}
}

func testMsgICAOErr19(t *testing.T) {
	raw, err := hex.DecodeString("00000000000000")
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	m := new(adsb.Message)

	err = m.UnmarshalBinary(raw)
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	raw, err = hex.DecodeString("9800000000000000000000000000")
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	rm := m.Raw()

	err = rm.UnmarshalBinary(raw)
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	icao, err := m.ICAO()
	if err == nil {
		t.Fatal("received nil, expected error")
	}

	if err.Error() != "error retrieving AP from 19: field not available" {
		t.Error("received unexpected error", err)
	}

	if icao != 0 {
		t.Errorf("expected 0, received %x", icao)
	}
}

func testMsgAltErr(t *testing.T) {
	raw, err := hex.DecodeString("00000000000000")
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	m := new(adsb.Message)

	err = m.UnmarshalBinary(raw)
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	rm := m.Raw()

	err = rm.UnmarshalBinary([]byte{})
	if err == nil {
		t.Fatal("received nil, expected error")
	}

	if err.Error() != "no data loaded" {
		t.Error("received unexpected error", err)
	}

	alt, err := m.Alt()
	if err == nil {
		t.Fatal("received nil, expected error")
	}

	if err.Error() != "error retrieving altitude: no data loaded" {
		t.Error("received unexpected error", err)
	}

	if alt != 0 {
		t.Errorf("expected 0, received %x", alt)
	}
}

func testMsgCallErr(t *testing.T) {
	raw, err := hex.DecodeString("00000000000000")
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	m := new(adsb.Message)

	err = m.UnmarshalBinary(raw)
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	rm := m.Raw()

	err = rm.UnmarshalBinary([]byte{})
	if err == nil {
		t.Fatal("received nil, expected error")
	}

	if err.Error() != "no data loaded" {
		t.Error("received unexpected error", err)
	}

	call, err := m.Call()
	if err == nil {
		t.Fatal("received nil, expected error")
	}

	if err.Error() != "error retrieving callsign: no data loaded" {
		t.Error("received unexpected error", err)
	}

	if call != "" {
		t.Errorf("expected nil, received %s", call)
	}
}

func testMsgSqkErr(t *testing.T) {
	raw, err := hex.DecodeString("00000000000000")
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	m := new(adsb.Message)

	err = m.UnmarshalBinary(raw)
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	rm := m.Raw()

	err = rm.UnmarshalBinary([]byte{})
	if err == nil {
		t.Fatal("received nil, expected error")
	}

	if err.Error() != "no data loaded" {
		t.Error("received unexpected error", err)
	}

	sqk, err := m.Sqk()
	if err == nil {
		t.Fatal("received nil, expected error")
	}

	if err.Error() != "error retrieving squawk: no data loaded" {
		t.Error("received unexpected error", err)
	}

	if !bytes.Equal(sqk, []byte{}) {
		t.Errorf("expected nil, received %x", sqk)
	}
}

func testMsgCPRErr(t *testing.T) {
	raw, err := hex.DecodeString("00000000000000")
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	m := new(adsb.Message)

	err = m.UnmarshalBinary(raw)
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	rm := m.Raw()

	err = rm.UnmarshalBinary([]byte{})
	if err == nil {
		t.Fatal("received nil, expected error")
	}

	if err.Error() != "no data loaded" {
		t.Error("received unexpected error", err)
	}

	cpr, err := m.CPR()
	if err == nil {
		t.Fatal("received nil, expected error")
	}

	if err.Error() != "error retrieving position: no data loaded" {
		t.Error("received unexpected error", err)
	}

	if cpr != nil {
		t.Error("received unexpected data")
	}
}

type testCase struct {
	Msg string

	DF int
	CA int
	FS int
	VS int

	TC  int
	SS  int
	Cat string

	CPR       bool
	LocalPos  bool
	GlobalPos bool
	RefPt     []float64
	Msg2      string

	Lat float64
	Lon float64

	ICAO uint64
	Sqk  []byte
	Call string
	Alt  int64

	AltError string
}

// TestDecode runs the test cases for message decoding.
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
	t.Run("DF24", testDF24)
}

// TestDecodeErrors runs test cases for message decoding errors.
func TestDecodeErrors(t *testing.T) {
	t.Run("MetricAltitude", testAltErrMetric)
	t.Run("InvalidAltitude", testAltErrInvalid)
}

// test DF0 air-to-air surveillance.
func testDF0(t *testing.T) {
	tc := &testCase{
		Msg: "02e19718e70f6c",

		DF: 0,
		CA: -1,
		FS: -1,
		VS: 0,

		TC:  -1,
		SS:  -1,
		Cat: "",

		CPR: false,

		ICAO: 0xabd94d,
		Alt:  36000,
		Sqk:  []byte{},
		Call: "",
	}

	testDecode(t, tc)
}

// test DF4 with 25ft altitude report.
func testDF4A(t *testing.T) {
	tc := &testCase{
		Msg: "20001910bc45e9",

		DF: 4,
		CA: -1,
		FS: 0,
		VS: -1,

		TC:  -1,
		SS:  -1,
		Cat: "",

		CPR: false,

		ICAO: 0xa27aee,
		Alt:  39000,
		Sqk:  []byte{},
		Call: "",
	}

	testDecode(t, tc)
}

// test DF4 with a Gillham-encoded altitude.
func testDF4B(t *testing.T) {
	tc := &testCase{
		Msg: "2000102a10fc86",

		DF: 4,
		CA: -1,
		FS: 0,
		VS: -1,

		TC:  -1,
		SS:  -1,
		Cat: "",

		CPR: false,

		ICAO: 0x71ef1e,
		Alt:  1300,
		Sqk:  []byte{},
		Call: "",
	}

	testDecode(t, tc)
}

// test DF5 identity reply.
func testDF5(t *testing.T) {
	tc := &testCase{
		Msg: "28001b0601970d",

		DF: 5,
		CA: -1,
		FS: 0,
		VS: -1,

		TC:  -1,
		SS:  -1,
		Cat: "",

		CPR: false,

		ICAO: 0xa3696e,
		Alt:  0,
		Sqk:  []byte{3, 4, 5, 2},
		Call: "",
	}

	testDecode(t, tc)
}

// test DF11 all call reply.
func testDF11(t *testing.T) {
	tc := &testCase{
		Msg: "5dac22c54b7a07",

		DF: 11,
		CA: 5,
		FS: -1,
		VS: -1,

		TC:  -1,
		SS:  -1,
		Cat: "",

		CPR: false,

		ICAO: 0xac22c5,
		Alt:  0,
		Sqk:  []byte{},
		Call: "",
	}

	testDecode(t, tc)
}

// test DF20 Comm-B altitude reply.
func testDF20(t *testing.T) {
	tc := &testCase{
		Msg: "a0000f9820057273df8d20e2cf30",

		DF: 20,
		CA: -1,
		FS: 0,
		VS: -1,

		TC:  -1,
		SS:  -1,
		Cat: "",

		CPR: false,

		ICAO: 0xa52333,
		Alt:  24000,
		Sqk:  []byte{},
		Call: "AWI3784",
	}

	testDecode(t, tc)
}

// test DF21 Comm-B identity reply.
func testDF21(t *testing.T) {
	tc := &testCase{
		Msg: "ac19b29573482f6963663636022b",

		DF: 21,
		CA: -1,
		FS: 4,
		VS: -1,

		TC:  -1,
		SS:  -1,
		Cat: "",

		CPR: false,

		ICAO: 0xa97db4,
		Alt:  0,
		Sqk:  []byte{6, 0, 1, 7},
		Call: "",
	}

	testDecode(t, tc)
}

func testDF24(t *testing.T) {
	tc := &testCase{
		Msg: "c4576da66a68295e7d22ed5dd112",

		DF: 24,
		CA: -1,
		FS: -1,
		VS: -1,

		TC:  -1,
		SS:  -1,
		Cat: "",

		CPR: false,

		ICAO: 0xab4531,
		Alt:  0,
		Sqk:  []byte{},
		Call: "",
	}

	testDecode(t, tc)
}

// test DF17 extended squitter position, local decode.
func testDF17PosLocal(t *testing.T) {
	tc := &testCase{
		Msg: "8da9450d60bde138e8638c939134",

		DF: 17,
		CA: 5,
		FS: -1,
		VS: -1,

		TC:  12,
		SS:  0,
		Cat: "",

		CPR:      true,
		LocalPos: true,
		RefPt:    []float64{43.14, -89.33},

		Lat: 43.83300781,
		Lon: -90.46484375,

		ICAO: 0xa9450d,
		Alt:  36950,
		Sqk:  []byte{},
		Call: "",
	}

	testDecode(t, tc)
}

// test DF17 extended squitter position, global decode.
func testDF17PosGlobal(t *testing.T) {
	tc := &testCase{
		Msg: "8da8028758ab0028de078689d437",

		DF: 17,
		CA: 5,
		FS: -1,
		VS: -1,

		TC:  11,
		SS:  0,
		Cat: "",

		CPR:       true,
		GlobalPos: true,
		Msg2:      "8da8028758ab07b0b8876e81eb25",

		Lat: 42.23945229,
		Lon: -89.87851165,

		ICAO: 0xa80287,
		Alt:  33000,
		Sqk:  []byte{},
		Call: "",
	}

	testDecode(t, tc)
}

// test DF17 extended squitter position, global decode, reversed.
func testDF17PosGlobalRev(t *testing.T) {
	tc := &testCase{
		Msg: "8dab9448589ff40a4e62a6c8b7a6",

		DF: 17,
		CA: 5,
		FS: -1,
		VS: -1,

		TC:  11,
		SS:  0,
		Cat: "",

		CPR:       true,
		GlobalPos: true,
		Msg2:      "8dab9448589ff083bbe2387219c5",

		Lat: 42.77183532,
		Lon: -90.47590775,

		ICAO: 0xab9448,
		Alt:  30975,
		Sqk:  []byte{},
		Call: "",
	}

	testDecode(t, tc)
}

// test DF17 extended squitter identity.
func testDF17Ident(t *testing.T) {
	tc := &testCase{
		Msg: "8dacf84e23101332cf3ca037ef13",

		DF: 17,
		CA: 5,
		FS: -1,
		VS: -1,

		TC:  4,
		SS:  -1,
		Cat: "A3",

		CPR: false,

		ICAO: 0xacf84e,
		Alt:  0,
		Sqk:  []byte{},
		Call: "DAL2332",
	}

	testDecode(t, tc)
}

func testDecode(t *testing.T, tc *testCase) {
	t.Helper()

	b, err := hex.DecodeString(tc.Msg)
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	msg := new(adsb.Message)

	err = msg.UnmarshalBinary(b)
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	testICAO(t, tc, msg)
	testSqk(t, tc, msg)
	testCall(t, tc, msg)
	testAlt(t, tc, msg)
	testCPR(t, tc, msg)
}

func testCPR(t *testing.T, tc *testCase, msg *adsb.Message) {
	t.Helper()

	cpr, err := msg.CPR()
	if err != nil {
		if tc.CPR != false || tc.CPR == false && !errors.Is(err, adsb.ErrNotAvailable) {
			t.Fatal("received unexpected error", err)
		}
	}

	if !tc.CPR && cpr != nil {
		t.Error("CPR: unexpected position report populated")
	}

	if tc.CPR && tc.LocalPos {
		if cpr == nil {
			t.Error("CPR: expected but not present")
		} else {
			testCPRLocal(t, tc, cpr)
		}
	}

	if tc.CPR && tc.GlobalPos {
		if cpr == nil {
			t.Error("CPR: expected but not present")
		} else {
			testCPRGlobal(t, tc, cpr)
		}
	}
}

func testCPRLocal(t *testing.T, tc *testCase, cpr *adsb.CPR) {
	t.Helper()

	c, err := cpr.DecodeLocal(tc.RefPt)
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

func testCPRGlobal(t *testing.T, tc *testCase, cpr *adsb.CPR) {
	t.Helper()

	rm := new(adsb.RawMessage)

	m2b, err := (hex.DecodeString(tc.Msg2))
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	err = rm.UnmarshalBinary(m2b)
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	m, err := adsb.NewMessage(rm)
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	cpr2, err := m.CPR()
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	if cpr2 == nil {
		t.Fatal("no position decoded in Msg2")
	}

	c, err := adsb.DecodeGlobalPosition(cpr, cpr2)
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

func testICAO(t *testing.T, tc *testCase, msg *adsb.Message) {
	t.Helper()

	icao, err := msg.ICAO()
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	if icao != tc.ICAO {
		t.Errorf("ICAO: received %06x, expected %06x", icao, tc.ICAO)
	}
}

func testSqk(t *testing.T, tc *testCase, msg *adsb.Message) {
	t.Helper()

	sqk, err := msg.Sqk()
	if err != nil {
		if len(tc.Sqk) > 0 || len(tc.Sqk) == 0 && !errors.Is(err, adsb.ErrNotAvailable) {
			t.Fatal("received unexpected error", err)
		}
	}

	if !bytes.Equal(sqk, tc.Sqk) {
		t.Errorf("Sqk: received %s, expected %s", sqk, tc.Sqk)
	}
}

func testCall(t *testing.T, tc *testCase, msg *adsb.Message) {
	t.Helper()

	call, err := msg.Call()
	if err != nil {
		if tc.Call != "" || tc.Call == "" && !errors.Is(err, adsb.ErrNotAvailable) {
			t.Fatal("received unexpected error", err)
		}
	}

	if call != tc.Call {
		t.Errorf("Call: received %s, expected %s", call, tc.Call)
	}
}

func testAlt(t *testing.T, tc *testCase, msg *adsb.Message) {
	t.Helper()

	a, err := msg.Alt()
	if err != nil {
		if tc.Alt != 0 || tc.Alt == 0 && !errors.Is(err, adsb.ErrNotAvailable) {
			t.Fatal("received unexpected error", err)
		}
	}

	if a != tc.Alt {
		t.Errorf("Alt: received %v, expected %v", a, tc.Alt)
	}
}

// test DF4 with metric altitude.
func testAltErrMetric(t *testing.T) {
	tc := &testCase{
		Msg: "2000046210fc86",

		AltError: "metric altitude not supported",
	}

	testDecodeErr(t, tc)
}

// test DF4 with invalid altitude.
func testAltErrInvalid(t *testing.T) {
	tc := &testCase{
		Msg: "2000002210fc86",

		AltError: "invalid altitude value",
	}

	testDecodeErr(t, tc)
}

func testDecodeErr(t *testing.T, tc *testCase) {
	t.Helper()

	b, err := hex.DecodeString(tc.Msg)
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	msg := new(adsb.Message)

	err = msg.UnmarshalBinary(b)
	if err != nil {
		t.Fatal("received unexpected error", err)
	}

	if tc.AltError != "" {
		testAltError(t, tc, msg)
	}
}

func testAltError(t *testing.T, tc *testCase, msg *adsb.Message) {
	t.Helper()

	a, err := msg.Alt()
	if err == nil {
		t.Error("expected error, received nil")

		return
	}

	if a != 0 {
		t.Errorf("expected 0, received %d", a)
	}

	if tc.AltError != err.Error() {
		t.Errorf("expected %s, received %s", tc.AltError, err)
	}
}
