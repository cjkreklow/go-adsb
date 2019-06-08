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
	"testing"
)

// TestConst test string formatting of constant values
func TestConst(t *testing.T) {
	t.Run("DF", testDF)
	t.Run("CA", testCA)
	t.Run("FS", testFS)
}

// TestDF4 tests DF4 altitude reply
func TestDF4(t *testing.T) {
	msg := new(Message)
	err := msg.Decode([]byte{0x20, 0x00, 0x19, 0x10, 0xbc, 0x45, 0xe9})
	if err != nil {
		t.Fatal("received unexpected error", err)
	}
	if msg.ICAO != "a27aee" {
		t.Errorf("ICAO: received %s, expected %s", msg.ICAO, "a27aee")
	}
	if msg.Format != DF4 {
		t.Errorf("Format: received %v, expected %v", int(msg.Format), 4)
	}
	if msg.Cap != 0 {
		t.Errorf("Cap: received %v, expected %v", int(msg.Cap), 0)
	}
	if msg.FltStat != 0 {
		t.Errorf("FltStat: received %v, expected %v", int(msg.FltStat), 0)
	}
	if msg.Alt != 39000 {
		t.Errorf("Alt: received %v, expected %v", msg.Alt, 39000)
	}

}

// TestDF11 tests DF11 all call reply
func TestDF11(t *testing.T) {
	msg := new(Message)
	err := msg.Decode([]byte{0x5d, 0xac, 0x22, 0xc5, 0x4b, 0x7a, 0x07})
	if err != nil {
		t.Fatal("received unexpected error", err)
	}
	if msg.ICAO != "ac22c5" {
		t.Errorf("ICAO: received %s, expected %s", msg.ICAO, "ac22c5")
	}
	if msg.Format != DF11 {
		t.Errorf("Format: received %v, expected %v", int(msg.Format), 4)
	}
	if msg.Cap != CA5 {
		t.Errorf("Cap: received %v, expected %v", int(msg.Cap), 5)
	}
	if msg.FltStat != 0 {
		t.Errorf("FltStat: received %v, expected %v", int(msg.FltStat), 0)
	}
	if msg.Alt != 0 {
		t.Errorf("Alt: received %v, expected %v", msg.Alt, 0)
	}

}
