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

package adsb_test

import (
	"bytes"
	"encoding"
	"encoding/hex"
	"testing"

	"kreklow.us/go/go-adsb/adsb"
)

func TestRawUnmarshalErrors(t *testing.T) {
	if t.Run("Interface", testRawUnmarshalInterface) {
		t.Run("NilPointer", testRawUnmarshalNil)
		t.Run("ShortMsg", testRawUnmarshalShort)
	}
}

func testRawUnmarshalInterface(t *testing.T) {
	var i interface{} = new(adsb.RawMessage)
	if _, ok := i.(encoding.BinaryUnmarshaler); !ok {
		t.Fatal("RawMessage does not implement encoding.BinaryUnmarshaler")
	}
}

func testRawUnmarshalNil(t *testing.T) {
	errm := "can't unmarshal to nil pointer"
	var m *adsb.RawMessage
	err := m.UnmarshalBinary([]byte{0xf0, 0x0f})
	if err == nil {
		t.Fatal("expected error, received nil")
	} else if err.Error() != errm {
		t.Fatalf("expected: %s ; received: %s", errm, err)
	}
}

func testRawUnmarshalShort(t *testing.T) {
	errm := "incorrect data length"
	m := new(adsb.RawMessage)
	err := m.UnmarshalBinary([]byte{0xf0, 0x0f})
	if err == nil {
		t.Fatal("expected error, received nil")
	} else if err.Error() != errm {
		t.Fatalf("expected: %s ; received: %s", errm, err)
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

	var r adsb.RawMessage = b
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

	var r adsb.RawMessage = b
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

	var r adsb.RawMessage = b
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

	var r adsb.RawMessage = b
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
	t.Run("Reverse64", testRawBits64Rev)
	t.Run("Big64", testRawBits64Big)
	t.Run("Good64", testRawBits64Good)
	t.Run("Reverse32", testRawBits32Rev)
	t.Run("Big32", testRawBits32Big)
	t.Run("Good32", testRawBits32Good)
	t.Run("Reverse16", testRawBits16Rev)
	t.Run("Big16", testRawBits16Big)
	t.Run("Good16", testRawBits16Good)
	t.Run("Reverse8", testRawBits8Rev)
	t.Run("Big8", testRawBits8Big)
	t.Run("Good8", testRawBits8Good)
}

func testRawBitsNeg(t *testing.T) {
	b, err := hex.DecodeString("00aabbccddeeff")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r adsb.RawMessage = b
	defer func() {
		p := recover()
		if p != "bit must be greater than 0" {
			t.Error("unexpected panic:", p)
		}
	}()
	bits := r.Bits64(-10, 20)
	if bits != 0 {
		t.Error("received unexpected value:", bits)
	}
}

func testRawBitsZero(t *testing.T) {
	b, err := hex.DecodeString("00aabbccddeeff")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r adsb.RawMessage = b
	defer func() {
		p := recover()
		if p != "bit must be greater than 0" {
			t.Error("unexpected panic:", p)
		}
	}()
	bits := r.Bits64(0, 20)
	if bits != 0 {
		t.Error("received unexpected value:", bits)
	}
}

func testRawBitsLarge(t *testing.T) {
	b, err := hex.DecodeString("00aabbccddeeff")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r adsb.RawMessage = b
	defer func() {
		p := recover()
		if p != "bit must be within message length" {
			t.Error("unexpected panic:", p)
		}
	}()
	bits := r.Bits64(20, 80)
	if bits != 0 {
		t.Error("received unexpected value:", bits)
	}
}

func testRawBits64Rev(t *testing.T) {
	b, err := hex.DecodeString("00aabbccddeeff")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r adsb.RawMessage = b
	defer func() {
		p := recover()
		if p != "upper bound must be greater than lower bound" {
			t.Error("unexpected panic:", p)
		}
	}()
	bits := r.Bits64(20, 20)
	if bits != 0 {
		t.Error("received unexpected value:", bits)
	}
}

func testRawBits64Big(t *testing.T) {
	b, err := hex.DecodeString("00aabbccddeeff")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r adsb.RawMessage = b
	defer func() {
		p := recover()
		if p != "maximum of 64 bits exceeded" {
			t.Error("unexpected panic:", p)
		}
	}()
	bits := r.Bits64(1, 70)
	if bits != 0 {
		t.Error("received unexpected value:", bits)
	}
}

func testRawBits64Good(t *testing.T) {
	b, err := hex.DecodeString("00aabbccddeeff")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r adsb.RawMessage = b
	defer func() {
		p := recover()
		if p != nil {
			t.Error("unexpected panic:", p)
		}
	}()
	bits := r.Bits64(20, 30)
	if bits != 0x06F3 {
		t.Errorf("received unexpected value: %x", bits)
	}
}

func testRawBits32Rev(t *testing.T) {
	b, err := hex.DecodeString("00aabbccddeeff")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r adsb.RawMessage = b
	defer func() {
		p := recover()
		if p != "upper bound must be greater than lower bound" {
			t.Error("unexpected panic:", p)
		}
	}()
	bits := r.Bits32(20, 10)
	if bits != 0 {
		t.Error("received unexpected value:", bits)
	}
}

func testRawBits32Big(t *testing.T) {
	b, err := hex.DecodeString("00aabbccddeeff")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r adsb.RawMessage = b
	defer func() {
		p := recover()
		if p != "maximum of 32 bits exceeded" {
			t.Error("unexpected panic:", p)
		}
	}()
	bits := r.Bits32(1, 70)
	if bits != 0 {
		t.Error("received unexpected value:", bits)
	}
}

func testRawBits32Good(t *testing.T) {
	b, err := hex.DecodeString("00aabbccddeeff")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r adsb.RawMessage = b
	defer func() {
		p := recover()
		if p != nil {
			t.Error("unexpected panic:", p)
		}
	}()
	bits := r.Bits32(25, 40)
	if bits != 0xCCDD {
		t.Errorf("received unexpected value: %x", bits)
	}
}

func testRawBits16Rev(t *testing.T) {
	b, err := hex.DecodeString("00aabbccddeeff")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r adsb.RawMessage = b
	defer func() {
		p := recover()
		if p != "upper bound must be greater than lower bound" {
			t.Error("unexpected panic:", p)
		}
	}()
	bits := r.Bits16(20, 20)
	if bits != 0 {
		t.Error("received unexpected value:", bits)
	}
}

func testRawBits16Big(t *testing.T) {
	b, err := hex.DecodeString("00aabbccddeeff")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r adsb.RawMessage = b
	defer func() {
		p := recover()
		if p != "maximum of 16 bits exceeded" {
			t.Error("unexpected panic:", p)
		}
	}()
	bits := r.Bits16(1, 70)
	if bits != 0 {
		t.Error("received unexpected value:", bits)
	}
}

func testRawBits16Good(t *testing.T) {
	b, err := hex.DecodeString("00aabbccddeeff")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r adsb.RawMessage = b
	defer func() {
		p := recover()
		if p != nil {
			t.Error("unexpected panic:", p)
		}
	}()
	bits := r.Bits16(13, 20)
	if bits != 0xAB {
		t.Errorf("received unexpected value: %x", bits)
	}
}

func testRawBits8Rev(t *testing.T) {
	b, err := hex.DecodeString("00aabbccddeeff00aabbccddeeff")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r adsb.RawMessage = b
	defer func() {
		p := recover()
		if p != "upper bound must be greater than lower bound" {
			t.Error("unexpected panic:", p)
		}
	}()
	bits := r.Bits8(100, 99)
	if bits != 0 {
		t.Error("received unexpected value:", bits)
	}
}

func testRawBits8Big(t *testing.T) {
	b, err := hex.DecodeString("00aabbccddeeff")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r adsb.RawMessage = b
	defer func() {
		p := recover()
		if p != "maximum of 8 bits exceeded" {
			t.Error("unexpected panic:", p)
		}
	}()
	bits := r.Bits8(1, 70)
	if bits != 0 {
		t.Error("received unexpected value:", bits)
	}
}

func testRawBits8Good(t *testing.T) {
	b, err := hex.DecodeString("00aabbccddeeff")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r adsb.RawMessage = b
	defer func() {
		p := recover()
		if p != nil {
			t.Error("unexpected panic:", p)
		}
	}()
	bits := r.Bits8(10, 15)
	if bits != 0x15 {
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
	results := map[string][]byte{
		"AC": []byte{0x03, 0xbb},
		"AP": []byte{0x45, 0x1e, 0x00},
		"CC": []byte{0x01},
		"DF": []byte{0x00},
		"RI": []byte{0x03},
		"SL": []byte{0x05},
		"VS": []byte{0x00},
	}

	testRaw(t, "02a183bb451e00", results)
}

func testRawDF4(t *testing.T) {
	results := map[string][]byte{
		"AC": []byte{0x0d, 0xb8},
		"AP": []byte{0x67, 0x65, 0x2d},
		"DF": []byte{0x04},
		"DR": []byte{0x00},
		"FS": []byte{0x00},
		"UM": []byte{0x00},
	}

	testRaw(t, "20000db867652d", results)
}

func testRawDF5(t *testing.T) {
	results := map[string][]byte{
		"AP": []byte{0x3a, 0x57, 0xd0},
		"DF": []byte{0x05},
		"DR": []byte{0x17},
		"ID": []byte{0x00, 0x67},
		"FS": []byte{0x02},
		"UM": []byte{0x00},
	}

	testRaw(t, "2ab800673a57d0", results)
}

func testRawDF11(t *testing.T) {
	results := map[string][]byte{
		"AA": []byte{0xaa, 0x23, 0x4a},
		"CA": []byte{0x05},
		"DF": []byte{0x0b},
		"PI": []byte{0x91, 0x28, 0x89},
	}

	testRaw(t, "5daa234a912889", results)
}

func testRawDF16(t *testing.T) {
	results := map[string][]byte{
		"AC": []byte{0x15, 0x30},
		"AP": []byte{0x09, 0xc8, 0x6e},
		"DF": []byte{0x10},
		"MV": []byte{0x58, 0xab, 0x01, 0x60, 0xa0, 0x9b, 0xe8},
		"RI": []byte{0x03},
		"SL": []byte{0x07},
		"VS": []byte{0x00},
	}

	testRaw(t, "80e1953058ab0160a09be809c86e", results)
}

func testRawDF17(t *testing.T) {
	results := map[string][]byte{
		"AA": []byte{0xa2, 0xf1, 0x11},
		"CA": []byte{0x05},
		"DF": []byte{0x11},
		"ME": []byte{0x58, 0x1f, 0xb4, 0x84, 0x2d, 0x1f, 0x59},
		"PI": []byte{0xee, 0xa2, 0xb7},
	}

	testRaw(t, "8da2f111581fb4842d1f59eea2b7", results)
}

func testRawDF18(t *testing.T) {
	results := map[string][]byte{
		"AA": []byte{0xa1, 0xce, 0x0e},
		"CF": []byte{0x02},
		"DF": []byte{0x12},
		"ME": []byte{0x90, 0xb9, 0x73, 0xc2, 0x6a, 0x38, 0x0f},
		"PI": []byte{0x56, 0x25, 0x4c},
	}

	testRaw(t, "92a1ce0e90b973c26a380f56254c", results)
}

func testRawDF19(t *testing.T) {
	results := map[string][]byte{
		"AF": []byte{0x02},
		"DF": []byte{0x13},
	}

	testRaw(t, "9aa1ce0e90b973c26a380f56254c", results)
}

func testRawDF20(t *testing.T) {
	results := map[string][]byte{
		"AC": []byte{0x14, 0x97},
		"AP": []byte{0x5b, 0x75, 0x7a},
		"DF": []byte{0x14},
		"DP": []byte{0x5b, 0x75, 0x7a},
		"DR": []byte{0x00},
		"FS": []byte{0x00},
		"MB": []byte{0x10, 0x03, 0x0a, 0x80, 0xe5, 0x00, 0x00},
		"UM": []byte{0x00},
	}

	testRaw(t, "a000149710030a80e500005b757a", results)
	//testRaw(t, "a52e1487466f35e5af8b4db41af0", results)
}

func testRawDF21(t *testing.T) {
	results := map[string][]byte{
		"AP": []byte{0xba, 0x4f, 0x91},
		"DF": []byte{0x15},
		"DP": []byte{0xba, 0x4f, 0x91},
		"DR": []byte{0x1d},
		"ID": []byte{0x18, 0x60},
		"FS": []byte{0x00},
		"MB": []byte{0x15, 0xa6, 0x8e, 0x5b, 0xae, 0xdb, 0x2a},
		"UM": []byte{0x0b},
	}

	testRaw(t, "a8e9786015a68e5baedb2aba4f91", results)
}

func testRawDF24(t *testing.T) {
	results := map[string][]byte{
		"AP": []byte{0x6d, 0xb1, 0xa1},
		"DF": []byte{0x18},
		"KE": []byte{0x00},
		"MD": []byte{0x25, 0x54, 0x48, 0xac, 0x2a, 0x74, 0xd0, 0x03, 0x54, 0x7a},
		"ND": []byte{0x02},
	}

	testRaw(t, "c2255448ac2a74d003547a6db1a1", results)
}

func testRaw(t *testing.T, m string, results map[string][]byte) {
	msg, err := hex.DecodeString(m)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	rm := new(adsb.RawMessage)
	err = rm.UnmarshalBinary(msg)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	funcs := map[string]func() ([]byte, error){
		"AA": rm.AA, "AC": rm.AC, "AF": rm.AF, "AP": rm.AP,
		"CA": rm.CA, "CC": rm.CC, "CF": rm.CF,
		"DF": rm.DF, "DP": rm.DP, "DR": rm.DR,
		"FS": rm.FS, "ID": rm.ID, "KE": rm.KE,
		"MB": rm.MB, "MD": rm.MD, "ME": rm.ME, "MV": rm.MV,
		"ND": rm.ND, "PI": rm.PI, "RI": rm.RI,
		"SL": rm.SL, "UM": rm.UM, "VS": rm.VS,
	}

	for n, f := range funcs {
		if r, ok := results[n]; ok {
			b, err := f()
			if err != nil {
				t.Errorf("%s  unexpected error: %v", n, err)
			}
			if !bytes.Equal(r, b) {
				t.Errorf("%s  expected: %x  received: %x", n, r, b)
			}
		} else {
			b, err := f()
			if err == nil {
				t.Errorf("%s  expected: error  received: %v", n, err)
			} else if err.Error() != "field not available" {
				t.Errorf("%s  expected: field not available  received: %v", n, err)
			}
			if b != nil {
				t.Errorf("%s  expected: nil  received: %v", n, b)
			}
		}
	}
}

func TestRawFieldErrors(t *testing.T) {
	t.Run("NotLoaded", testRawFieldsNotLoaded)
	//t.Run("NotAvaliable", testRawFieldsNotAvailable)
}

func testRawFieldsNotLoaded(t *testing.T) {
	rm := new(adsb.RawMessage)
	fields := map[string]func() ([]byte, error){
		"AA": rm.AA, "AC": rm.AC, "AF": rm.AF, "AP": rm.AP,
		"CA": rm.CA, "CC": rm.CC, "CF": rm.CF,
		"DF": rm.DF, "DP": rm.DP, "DR": rm.DR,
		"FS": rm.FS, "ID": rm.ID, "KE": rm.KE,
		"MB": rm.MB, "MD": rm.MD, "ME": rm.ME, "MV": rm.MV,
		"ND": rm.ND, "PI": rm.PI, "RI": rm.RI,
		"SL": rm.SL, "UM": rm.UM, "VS": rm.VS,
	}

	for n, f := range fields {
		b, err := f()
		if err == nil {
			t.Errorf("%s  expected: error  received: nil", n)
		} else if err.Error() != "data not loaded" {
			t.Errorf("%s  expected: data not loaded  received: %v", n, err)
		}
		if b != nil {
			t.Errorf("%s  expected: nil  received: %v", n, b)
		}
	}
}
