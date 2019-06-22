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
	"encoding"
	"encoding/hex"
	"fmt"
	"testing"

	"kreklow.us/go/go-adsb/adsb"
)

// Test that RawMessage correctly implements unmarshaling
func TestRawUnmarshal(t *testing.T) {
	if t.Run("Interface", testRawUnmarshalInterface) {
		t.Run("NilPointer", testRawUnmarshalNil)
		t.Run("ShortMsg", testRawUnmarshalShort)
		t.Run("Success", testRawUnmarshalSuccess)
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

func testRawUnmarshalSuccess(t *testing.T) {
	b, err := hex.DecodeString("00aabbccddeeff")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	m := new(adsb.RawMessage)
	err = m.UnmarshalBinary(b)
	if err != nil {
		t.Fatalf("expected: nil; received: %s", err)
	}
	ev := "&[0 170 187 204 221 238 255]"
	rv := fmt.Sprintf("%v", m)
	if ev != rv {
		t.Fatalf("expected: %s; received %s", ev, rv)
	}
}

// Test Bit method
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

// Test Bits* methods
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
