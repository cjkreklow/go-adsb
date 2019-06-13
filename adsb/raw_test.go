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

// TestRawBit tests RawBytes.Bit
func TestRawBit(t *testing.T) {
	t.Run("Negative n", testBitNeg)
	t.Run("Zero n", testBitZero)
	t.Run("Large n", testBitLarge)
	t.Run("Good n", testBitGood)
}

// TestRawBits tests RawBytes.Bits*
func TestRawBits(t *testing.T) {
	t.Run("Negative n", testBitsNeg)
	t.Run("Zero n", testBitsZero)
	t.Run("Large z", testBitsLarge)
	t.Run("Reverse 64", testBits64Rev)
	t.Run("Big 64", testBits64Big)
	t.Run("Good 64", testBits64Good)
	t.Run("Reverse 32", testBits32Rev)
	t.Run("Big 32", testBits32Big)
	t.Run("Good 32", testBits32Good)
	t.Run("Reverse 16", testBits16Rev)
	t.Run("Big 16", testBits16Big)
	t.Run("Good 16", testBits16Good)
	t.Run("Reverse 8", testBits8Rev)
	t.Run("Big 8", testBits8Big)
	t.Run("Good 8", testBits8Good)
}

// test negative bit address
func testBitNeg(t *testing.T) {
	msg, err := hex.DecodeString("ff00ff00")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r RawBytes = msg
	defer func() {
		p := recover()
		if p != "bit must be greater than 0" {
			t.Error("unexpected panic:", p)
		}
	}()
	b := r.Bit(-10)
	if b != 0 {
		t.Error("received unexpected value:", b)
	}
}

// test zero bit address
func testBitZero(t *testing.T) {
	msg, err := hex.DecodeString("ff00ff00")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r RawBytes = msg
	defer func() {
		p := recover()
		if p != "bit must be greater than 0" {
			t.Error("unexpected panic:", p)
		}
	}()
	b := r.Bit(0)
	if b != 0 {
		t.Error("received unexpected value:", b)
	}
}

// test large bit address
func testBitLarge(t *testing.T) {
	msg, err := hex.DecodeString("ff00ff00")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r RawBytes = msg
	defer func() {
		p := recover()
		if p != "bit must be within message length" {
			t.Error("unexpected panic:", p)
		}
	}()
	b := r.Bit(99)
	if b != 0 {
		t.Error("received unexpected value:", b)
	}
}

// test good bit address
func testBitGood(t *testing.T) {
	msg, err := hex.DecodeString("ff00ff00")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r RawBytes = msg
	defer func() {
		p := recover()
		if p != nil {
			t.Error("unexpected panic:", p)
		}
	}()
	b := r.Bit(18)
	if b != 1 {
		t.Error("received unexpected value:", b)
	}
}

// test negative bit address
func testBitsNeg(t *testing.T) {
	msg, err := hex.DecodeString("ff00ff00")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r RawBytes = msg
	defer func() {
		p := recover()
		if p != "bit must be greater than 0" {
			t.Error("unexpected panic:", p)
		}
	}()
	b := r.Bits64(-10, 20)
	if b != 0 {
		t.Error("received unexpected value:", b)
	}
}

// test zero bit address
func testBitsZero(t *testing.T) {
	msg, err := hex.DecodeString("ff00ff00")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r RawBytes = msg
	defer func() {
		p := recover()
		if p != "bit must be greater than 0" {
			t.Error("unexpected panic:", p)
		}
	}()
	b := r.Bits64(0, 20)
	if b != 0 {
		t.Error("received unexpected value:", b)
	}
}

// test large bit address
func testBitsLarge(t *testing.T) {
	msg, err := hex.DecodeString("ff00ff00")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r RawBytes = msg
	defer func() {
		p := recover()
		if p != "bit must be within message length" {
			t.Error("unexpected panic:", p)
		}
	}()
	b := r.Bits64(20, 80)
	if b != 0 {
		t.Error("received unexpected value:", b)
	}
}

// test reverse bit address
func testBits64Rev(t *testing.T) {
	msg, err := hex.DecodeString("ff00ff00")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r RawBytes = msg
	defer func() {
		p := recover()
		if p != "upper bound must be greater than lower bound" {
			t.Error("unexpected panic:", p)
		}
	}()
	b := r.Bits64(20, 20)
	if b != 0 {
		t.Error("received unexpected value:", b)
	}
}

// test too many bits
func testBits64Big(t *testing.T) {
	msg, err := hex.DecodeString("ff00ff00ff00ff00ff00ff00")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r RawBytes = msg
	defer func() {
		p := recover()
		if p != "maximum of 64 bits exceeded" {
			t.Error("unexpected panic:", p)
		}
	}()
	b := r.Bits64(1, 70)
	if b != 0 {
		t.Error("received unexpected value:", b)
	}
}

// test good request
func testBits64Good(t *testing.T) {
	msg, err := hex.DecodeString("ff00ff00")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r RawBytes = msg
	defer func() {
		p := recover()
		if p != nil {
			t.Error("unexpected panic:", p)
		}
	}()
	b := r.Bits64(20, 30)
	if b != 0x07C0 {
		t.Errorf("received unexpected value: %x", b)
	}
}

// test reverse bit address
func testBits32Rev(t *testing.T) {
	msg, err := hex.DecodeString("ff00ff00")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r RawBytes = msg
	defer func() {
		p := recover()
		if p != "upper bound must be greater than lower bound" {
			t.Error("unexpected panic:", p)
		}
	}()
	b := r.Bits32(20, 20)
	if b != 0 {
		t.Error("received unexpected value:", b)
	}
}

// test too many bits
func testBits32Big(t *testing.T) {
	msg, err := hex.DecodeString("ff00ff00ff00ff00ff00ff00")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r RawBytes = msg
	defer func() {
		p := recover()
		if p != "maximum of 32 bits exceeded" {
			t.Error("unexpected panic:", p)
		}
	}()
	b := r.Bits32(1, 70)
	if b != 0 {
		t.Error("received unexpected value:", b)
	}
}

// test good request
func testBits32Good(t *testing.T) {
	msg, err := hex.DecodeString("ff00ff00")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r RawBytes = msg
	defer func() {
		p := recover()
		if p != nil {
			t.Error("unexpected panic:", p)
		}
	}()
	b := r.Bits32(20, 30)
	if b != 0x07C0 {
		t.Errorf("received unexpected value: %x", b)
	}
}

// test reverse bit address
func testBits16Rev(t *testing.T) {
	msg, err := hex.DecodeString("ff00ff00")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r RawBytes = msg
	defer func() {
		p := recover()
		if p != "upper bound must be greater than lower bound" {
			t.Error("unexpected panic:", p)
		}
	}()
	b := r.Bits16(20, 20)
	if b != 0 {
		t.Error("received unexpected value:", b)
	}
}

// test too many bits
func testBits16Big(t *testing.T) {
	msg, err := hex.DecodeString("ff00ff00ff00ff00ff00ff00")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r RawBytes = msg
	defer func() {
		p := recover()
		if p != "maximum of 16 bits exceeded" {
			t.Error("unexpected panic:", p)
		}
	}()
	b := r.Bits16(1, 70)
	if b != 0 {
		t.Error("received unexpected value:", b)
	}
}

// test good request
func testBits16Good(t *testing.T) {
	msg, err := hex.DecodeString("ff00ff00")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r RawBytes = msg
	defer func() {
		p := recover()
		if p != nil {
			t.Error("unexpected panic:", p)
		}
	}()
	b := r.Bits16(20, 30)
	if b != 0x07C0 {
		t.Errorf("received unexpected value: %x", b)
	}
}

// test reverse bit address
func testBits8Rev(t *testing.T) {
	msg, err := hex.DecodeString("ff00ff00")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r RawBytes = msg
	defer func() {
		p := recover()
		if p != "upper bound must be greater than lower bound" {
			t.Error("unexpected panic:", p)
		}
	}()
	b := r.Bits8(20, 20)
	if b != 0 {
		t.Error("received unexpected value:", b)
	}
}

// test too many bits
func testBits8Big(t *testing.T) {
	msg, err := hex.DecodeString("ff00ff00ff00ff00ff00ff00")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r RawBytes = msg
	defer func() {
		p := recover()
		if p != "maximum of 8 bits exceeded" {
			t.Error("unexpected panic:", p)
		}
	}()
	b := r.Bits8(1, 70)
	if b != 0 {
		t.Error("received unexpected value:", b)
	}
}

// test good request
func testBits8Good(t *testing.T) {
	msg, err := hex.DecodeString("ff00ff00")
	if err != nil {
		t.Fatal("received unexpected error:", err)
	}

	var r RawBytes = msg
	defer func() {
		p := recover()
		if p != nil {
			t.Error("unexpected panic:", p)
		}
	}()
	b := r.Bits8(20, 25)
	if b != 0x3E {
		t.Errorf("received unexpected value: %x", b)
	}
}
