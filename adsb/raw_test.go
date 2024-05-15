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
	"encoding"
	"encoding/hex"
	"errors"
	"fmt"
	"testing"

	"kreklow.us/go/go-adsb/adsb"
)

func TestRawUnmarshalInterface(t *testing.T) {
	var i interface{} = new(adsb.RawMessage)
	if _, ok := i.(encoding.BinaryUnmarshaler); !ok {
		t.Fatal("RawMessage does not implement encoding.BinaryUnmarshaler")
	}
}

func TestRawUnmarshalErrors(t *testing.T) {
	t.Run("NoData", testRawUnmarshalNoData)
	t.Run("Short0", testRawUnmarshalShort0)
	t.Run("Long4", testRawUnmarshalLong4)
	t.Run("Short17", testRawUnmarshalShort17)
	t.Run("Unknown12", testRawUnmarshalUnk12)
}

func testRawUnmarshalNoData(t *testing.T) {
	testRawUnmarshalErr(t, "", "no data loaded")
}

func testRawUnmarshalShort0(t *testing.T) {
	testRawUnmarshalErr(t, "00ff", "incorrect data length: 16 bits with format 0")
}

func testRawUnmarshalLong4(t *testing.T) {
	testRawUnmarshalErr(t, "20ffffffffffffff", "incorrect data length: 64 bits with format 4")
}

func testRawUnmarshalShort17(t *testing.T) {
	testRawUnmarshalErr(t, "8800ff00", "incorrect data length: 32 bits with format 17")
}

func testRawUnmarshalUnk12(t *testing.T) {
	testRawUnmarshalErr(t, "6000ff", "unknown downlink format: 24 bits with format 12")
}

func testRawUnmarshalErr(t *testing.T, m string, e string) {
	t.Helper()

	b, err := hex.DecodeString(m)
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	rm := new(adsb.RawMessage)

	err = rm.UnmarshalBinary(b)
	if err == nil {
		t.Fatal("expected error, received nil")
	} else if err.Error() != e {
		t.Fatalf("expected %s, received %s", e, err)
	}
}

func TestRawBit(t *testing.T) {
	t.Run("Negative", testRawBitNeg)
	t.Run("Zero", testRawBitZero)
	t.Run("Large", testRawBitLarge)
	t.Run("Good", testRawBitGood)
}

func testRawBitNeg(t *testing.T) {
	b, err := hex.DecodeString("00aabbccddeeff")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	r := new(adsb.RawMessage)

	err = r.UnmarshalBinary(b)
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	defer func() {
		p := recover()
		if p != "bit must be greater than 0" {
			t.Error("unexpected panic:", p)
		}
	}()

	bit := r.Bit(-10)
	if bit != 0 {
		t.Error("received unexpected value:", bit)
	}
}

func testRawBitZero(t *testing.T) {
	b, err := hex.DecodeString("00aabbccddeeff")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	r := new(adsb.RawMessage)

	err = r.UnmarshalBinary(b)
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	defer func() {
		p := recover()
		if p != "bit must be greater than 0" {
			t.Error("unexpected panic:", p)
		}
	}()

	bit := r.Bit(0)
	if bit != 0 {
		t.Error("received unexpected value:", bit)
	}
}

func testRawBitLarge(t *testing.T) {
	b, err := hex.DecodeString("00aabbccddeeff")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	r := new(adsb.RawMessage)

	err = r.UnmarshalBinary(b)
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	defer func() {
		p := recover()
		if p != "bit must be within message length" {
			t.Error("unexpected panic:", p)
		}
	}()

	bit := r.Bit(99)
	if bit != 0 {
		t.Error("received unexpected value:", bit)
	}
}

func testRawBitGood(t *testing.T) {
	b, err := hex.DecodeString("00aabbccddeeff")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	r := new(adsb.RawMessage)

	err = r.UnmarshalBinary(b)
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	defer func() {
		p := recover()
		if p != nil {
			t.Error("unexpected panic:", p)
		}
	}()

	bit := r.Bit(17)
	if bit != 1 {
		t.Error("received unexpected value:", bit)
	}
}

func TestRawBits(t *testing.T) {
	t.Run("Negative", testRawBitsNeg)
	t.Run("Zero", testRawBitsZero)
	t.Run("Large", testRawBitsLarge)
	t.Run("Reverse", testRawBitsRev)
	t.Run("Big", testRawBitsBig)
	t.Run("Good", testRawBitsGood)
}

func testRawBitsNeg(t *testing.T) {
	b, err := hex.DecodeString("00aabbccddeeff")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	r := new(adsb.RawMessage)

	err = r.UnmarshalBinary(b)
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	defer func() {
		p := recover()
		if p != "lower bound must be greater than 0" {
			t.Error("unexpected panic:", p)
		}
	}()

	bits := r.Bits(-10, 20)
	if bits != 0 {
		t.Error("received unexpected value:", bits)
	}
}

func testRawBitsZero(t *testing.T) {
	b, err := hex.DecodeString("00aabbccddeeff")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	r := new(adsb.RawMessage)

	err = r.UnmarshalBinary(b)
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	defer func() {
		p := recover()
		if p != "lower bound must be greater than 0" {
			t.Error("unexpected panic:", p)
		}
	}()

	bits := r.Bits(0, 20)
	if bits != 0 {
		t.Error("received unexpected value:", bits)
	}
}

func testRawBitsLarge(t *testing.T) {
	b, err := hex.DecodeString("00aabbccddeeff")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	r := new(adsb.RawMessage)

	err = r.UnmarshalBinary(b)
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	defer func() {
		p := recover()
		if p != "upper bound must be within message length" {
			t.Error("unexpected panic:", p)
		}
	}()

	bits := r.Bits(20, 80)
	if bits != 0 {
		t.Error("received unexpected value:", bits)
	}
}

func testRawBitsRev(t *testing.T) {
	b, err := hex.DecodeString("00aabbccddeeff")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	r := new(adsb.RawMessage)

	err = r.UnmarshalBinary(b)
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	defer func() {
		p := recover()
		if p != "upper bound must be greater than lower bound" {
			t.Error("unexpected panic:", p)
		}
	}()

	bits := r.Bits(21, 20)
	if bits != 0 {
		t.Error("received unexpected value:", bits)
	}
}

func testRawBitsBig(t *testing.T) {
	b, err := hex.DecodeString("88aabbccddeeff00aabbccddeeff")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	r := new(adsb.RawMessage)

	err = r.UnmarshalBinary(b)
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	defer func() {
		p := recover()
		if p != "maximum of 64 bits exceeded" {
			t.Error("unexpected panic:", p)
		}
	}()

	bits := r.Bits(1, 70)
	if bits != 0 {
		t.Error("received unexpected value:", bits)
	}
}

func testRawBitsGood(t *testing.T) {
	b, err := hex.DecodeString("00aabbccddeeff")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	r := new(adsb.RawMessage)

	err = r.UnmarshalBinary(b)
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	defer func() {
		p := recover()
		if p != nil {
			t.Error("unexpected panic:", p)
		}
	}()

	bits := r.Bits(20, 30)
	if bits != 0x06F3 {
		t.Errorf("received unexpected value: %x", bits)
	}
}

func TestRawDecode(t *testing.T) {
	t.Run("DF0", testRawDF0)
	t.Run("DF4", testRawDF4)
	t.Run("DF5", testRawDF5)
	t.Run("DF11", testRawDF11)
	t.Run("DF16", testRawDF16)
	t.Run("DF17", testRawDF17)
	t.Run("DF18", testRawDF18)
	t.Run("DF19", testRawDF19)
	t.Run("DF20", testRawDF20)
	t.Run("DF21", testRawDF21)
	t.Run("DF24", testRawDF24)
}

func testRawDF0(t *testing.T) {
	results := map[string]uint64{
		"AC": 0x03bb,
		"AP": 0x451e00,
		"CC": 0x01,
		"DF": 0x00,
		"RI": 0x03,
		"SL": 0x05,
		"VS": 0x00,
	}

	testRaw(t, "02a183bb451e00", results)
}

func testRawDF4(t *testing.T) {
	results := map[string]uint64{
		"AC": 0x0db8,
		"AP": 0x67652d,
		"DF": 0x04,
		"DR": 0x00,
		"FS": 0x00,
		"UM": 0x00,
	}

	testRaw(t, "20000db867652d", results)
}

func testRawDF5(t *testing.T) {
	results := map[string]uint64{
		"AP": 0x3a57d0,
		"DF": 0x05,
		"DR": 0x17,
		"ID": 0x0067,
		"FS": 0x02,
		"UM": 0x00,
	}

	testRaw(t, "2ab800673a57d0", results)
}

func testRawDF11(t *testing.T) {
	results := map[string]uint64{
		"AA": 0xaa234a,
		"CA": 0x05,
		"DF": 0x0b,
		"PI": 0x912889,
	}

	testRaw(t, "5daa234a912889", results)
}

func testRawDF16(t *testing.T) {
	results := map[string]uint64{
		"AC": 0x1530,
		"AP": 0x09c86e,
		"DF": 0x10,
		"MV": 0x58ab0160a09be8,
		"RI": 0x03,
		"SL": 0x07,
		"VS": 0x00,
	}

	testRaw(t, "80e1953058ab0160a09be809c86e", results)
}

func testRawDF17(t *testing.T) {
	results := map[string]uint64{
		"AA":     0xa2f111,
		"CA":     0x05,
		"DF":     0x11,
		"ME":     0x581fb4842d1f59,
		"PI":     0xeea2b7,
		"ESType": 0x0b,
	}

	testRaw(t, "8da2f111581fb4842d1f59eea2b7", results)
}

func testRawDF18(t *testing.T) {
	results := map[string]uint64{
		"AA":     0xa1ce0e,
		"CF":     0x02,
		"DF":     0x12,
		"ME":     0x90b973c26a380f,
		"PI":     0x56254c,
		"ESType": 0x12,
	}

	testRaw(t, "92a1ce0e90b973c26a380f56254c", results)
}

func testRawDF19(t *testing.T) {
	results := map[string]uint64{
		"AF": 0x02,
		"DF": 0x13,
	}

	testRaw(t, "9aa1ce0e90b973c26a380f56254c", results)
}

func testRawDF20(t *testing.T) {
	results := map[string]uint64{
		"AC": 0x1497,
		"AP": 0x5b757a,
		"DF": 0x14,
		"DP": 0x5b757a,
		"DR": 0x00,
		"FS": 0x00,
		"MB": 0x10030a80e50000,
		"UM": 0x00,
	}

	// testRaw(t, "a52e1487466f35e5af8b4db41af0", results)
	testRaw(t, "a000149710030a80e500005b757a", results)
}

func testRawDF21(t *testing.T) {
	results := map[string]uint64{
		"AP": 0xba4f91,
		"DF": 0x15,
		"DP": 0xba4f91,
		"DR": 0x1d,
		"ID": 0x1860,
		"FS": 0x00,
		"MB": 0x15a68e5baedb2a,
		"UM": 0x0b,
	}

	testRaw(t, "a8e9786015a68e5baedb2aba4f91", results)
	testRawMD(t, "a8e9786015a68e5baedb2aba4f91", nil)
}

func testRawDF24(t *testing.T) {
	results := map[string]uint64{
		"AP": 0x6db1a1,
		"DF": 0x18,
		"KE": 0x00,
		"ND": 0x02,
	}
	md := []byte{0x25, 0x54, 0x48, 0xac, 0x2a, 0x74, 0xd0, 0x03, 0x54, 0x7a}

	testRaw(t, "c2255448ac2a74d003547a6db1a1", results)
	testRawMD(t, "c2255448ac2a74d003547a6db1a1", md)
}

func testRaw(t *testing.T, m string, results map[string]uint64) {
	t.Helper()

	msg, err := hex.DecodeString(m)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	rm := new(adsb.RawMessage)

	err = rm.UnmarshalBinary(msg)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	funcs := map[string]func() (uint64, error){
		"AA": rm.AA, "AC": rm.AC, "AF": rm.AF, "AP": rm.AP,
		"CA": rm.CA, "CC": rm.CC, "CF": rm.CF,
		"DF": rm.DF, "DP": rm.DP, "DR": rm.DR,
		"FS": rm.FS, "ID": rm.ID, "KE": rm.KE,
		"MB": rm.MB, "ME": rm.ME, "MV": rm.MV,
		"ND": rm.ND, "PI": rm.PI, "RI": rm.RI,
		"SL": rm.SL, "UM": rm.UM, "VS": rm.VS,
		"ESType": rm.ESType,
	}

	for n, f := range funcs {
		expErr := fmt.Sprintf(
			"error retrieving %s from %d: field not available",
			n, results["DF"])

		if r, ok := results[n]; ok { //nolint:nestif // recursive testing function
			b, err := f()
			if err != nil {
				t.Errorf("%s  unexpected error: %v", n, err)
			}

			if r != b {
				t.Errorf("%s  expected: %x  received: %x", n, r, b)
			}
		} else {
			b, err := f()
			if err == nil {
				t.Errorf("%s  expected: error  received: %v", n, err)
			} else {
				if err.Error() != expErr {
					t.Errorf("%s  expected: '%s'  received: '%v'", n, expErr, err)
				}

				if !errors.Is(err, adsb.ErrNotAvailable) {
					t.Errorf("%s  unexpected error type, not ErrNotAvailable", n)
				}
			}

			if b != 0 {
				t.Errorf("%s  expected: 0  received: %x", n, b)
			}
		}
	}
}

func testRawMD(t *testing.T, m string, r []byte) {
	t.Helper()

	msg, err := hex.DecodeString(m)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	rm := new(adsb.RawMessage)

	err = rm.UnmarshalBinary(msg)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	n := "MD"

	if r != nil { //nolint:nestif // test chain
		b, err := rm.MD()
		if err != nil {
			t.Errorf("%s  unexpected error: %v", n, err)
		}

		if !bytes.Equal(r, b) {
			t.Errorf("%s  expected: %x  received: %x", n, r, b)
		}
	} else {
		df, err := rm.DF()
		if err != nil {
			t.Errorf("%s  unexpected error: %v", n, err)
		}

		expErr := fmt.Sprintf("error retrieving %s from %d: field not available", n, df)

		b, err := rm.MD()
		if err == nil {
			t.Errorf("%s  expected: error  received: %v", n, err)
		} else if err.Error() != expErr {
			t.Errorf("%s  expected: '%s'  received: '%v'", n, expErr, err)
		}

		if b != nil {
			t.Errorf("%s  expected: nil  received: %v", n, b)
		}
	}
}

func TestRawFieldsNotLoaded(t *testing.T) {
	rm := new(adsb.RawMessage)
	fields := map[string]func() (uint64, error){
		"AA": rm.AA, "AC": rm.AC, "AF": rm.AF, "AP": rm.AP,
		"CA": rm.CA, "CC": rm.CC, "CF": rm.CF,
		"DF": rm.DF, "DP": rm.DP, "DR": rm.DR,
		"FS": rm.FS, "ID": rm.ID, "KE": rm.KE,
		"MB": rm.MB, "ME": rm.ME, "MV": rm.MV,
		"ND": rm.ND, "PI": rm.PI, "RI": rm.RI,
		"SL": rm.SL, "UM": rm.UM, "VS": rm.VS,
		"ESType": rm.ESType, "ESAltitude": rm.ESAltitude,
	}

	expErr := "no data loaded"

	for n, f := range fields {
		b, err := f()
		if err == nil {
			t.Errorf("%s  expected: error  received: nil", n)
		} else if err.Error() != expErr {
			t.Errorf("%s  expected: '%s'  received: '%v'", n, expErr, err)
		}

		if b != 0 {
			t.Errorf("%s  expected: 0  received: %v", n, b)
		}
	}

	n := "MD"

	b, err := rm.MD()
	if err == nil {
		t.Errorf("%s  expected: error  received: nil", n)
	} else if err.Error() != expErr {
		t.Errorf("%s  expected: '%s'  received: '%v'", n, expErr, err)
	}

	if b != nil {
		t.Errorf("%s  expected: nil  received: '%v'", n, b)
	}

	n = "Parity"

	p := rm.Parity()
	if p != 0 {
		t.Errorf("%s  expected: 0  received: %x", n, p)
	}
}
