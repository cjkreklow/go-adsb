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
	"testing"
)

type testCase struct {
	Msg  string
	DF   int
	CA   int
	FS   int
	TC   int
	ICAO string
	Alt  int64
	Sqk  string
	Call string
}

// TestDecode runs the test cases for message decoding
func TestDecode(t *testing.T) {
	t.Run("DF4", testDF4A)
	t.Run("DF4 Gillham", testDF4B)
	t.Run("DF5", testDF5)
	t.Run("DF11", testDF11)
	t.Run("DF17 Altitude", testDF17Alt)
	t.Run("DF20", testDF20)
	t.Run("DF21", testDF21)
}

// test DF4 with 25ft altitude report
func testDF4A(t *testing.T) {
	tc := &testCase{
		Msg:  "20001910bc45e9",
		DF:   4,
		CA:   -1,
		FS:   0,
		TC:   -1,
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
		Msg:  "2000042210fc86",
		DF:   4,
		CA:   -1,
		FS:   0,
		TC:   -1,
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
		Msg:  "28001b0601970d",
		DF:   5,
		CA:   -1,
		FS:   0,
		TC:   -1,
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
		Msg:  "5dac22c54b7a07",
		DF:   11,
		CA:   5,
		FS:   -1,
		TC:   -1,
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
		Msg:  "a0000f9820057273df8d20e2cf30",
		DF:   20,
		CA:   -1,
		FS:   0,
		TC:   -1,
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
		Msg:  "ac19b29573482f6963663636022b",
		DF:   21,
		CA:   -1,
		FS:   4,
		TC:   -1,
		ICAO: "a97db4",
		Alt:  0,
		Sqk:  "6017",
		Call: "",
	}

	testDecode(t, tc)
}

// test DF17 extended squitter altitude
func testDF17Alt(t *testing.T) {
	tc := &testCase{
		Msg:  "8da9450d60bde138e8638c939134",
		DF:   17,
		CA:   5,
		FS:   -1,
		TC:   12,
		ICAO: "a9450d",
		Alt:  36950,
		Sqk:  "",
		Call: "",
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
	if msg.TC != TC(tc.TC) {
		t.Errorf("TC: received %v, expected %v", int(msg.TC), tc.TC)
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
}
