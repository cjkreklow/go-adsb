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
	"bytes"
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
	if msg.Cap != -1 {
		t.Errorf("Cap: received %v, expected %v", int(msg.Cap), -1)
	}
	if msg.FltStat != 0 {
		t.Errorf("FltStat: received %v, expected %v", int(msg.FltStat), 0)
	}
	if msg.Alt != 39000 {
		t.Errorf("Alt: received %v, expected %v", msg.Alt, 39000)
	}
	if msg.Sqk != "" {
		t.Errorf("Sqk: received %s, expected %s", msg.Sqk, "")
	}
	if msg.MsgB != nil {
		t.Errorf("MsgB: received %s, expected nil", msg.MsgB)
	}
}

// TestDF4B tests DF4 altitude reply with a Gillham-encoded altitude
func TestDF4B(t *testing.T) {
	msg := new(Message)
	err := msg.Decode([]byte{0x20, 0x00, 0x04, 0x22, 0x10, 0xfc, 0x86})
	if err != nil {
		t.Fatal("received unexpected error", err)
	}
	if msg.ICAO != "a97172" {
		t.Errorf("ICAO: received %s, expected %s", msg.ICAO, "a97172")
	}
	if msg.Format != DF4 {
		t.Errorf("Format: received %v, expected %v", int(msg.Format), 4)
	}
	if msg.Cap != -1 {
		t.Errorf("Cap: received %v, expected %v", int(msg.Cap), -1)
	}
	if msg.FltStat != 0 {
		t.Errorf("FltStat: received %v, expected %v", int(msg.FltStat), 0)
	}
	if msg.Alt != 2000 {
		t.Errorf("Alt: received %v, expected %v", msg.Alt, 2000)
	}
	if msg.Sqk != "" {
		t.Errorf("Sqk: received %s, expected %s", msg.Sqk, "")
	}
	if msg.MsgB != nil {
		t.Errorf("MsgB: received %s, expected nil", msg.MsgB)
	}
}

// TestDF5 tests DF5 identity reply
func TestDF5(t *testing.T) {
	msg := new(Message)
	err := msg.Decode([]byte{0x28, 0x00, 0x1b, 0x06, 0x01, 0x97, 0x0d})
	if err != nil {
		t.Fatal("received unexpected error", err)
	}
	if msg.ICAO != "a3696e" {
		t.Errorf("ICAO: received %s, expected %s", msg.ICAO, "a3696e")
	}
	if msg.Format != DF5 {
		t.Errorf("Format: received %v, expected %v", int(msg.Format), 5)
	}
	if msg.Cap != -1 {
		t.Errorf("Cap: received %v, expected %v", int(msg.Cap), -1)
	}
	if msg.FltStat != 0 {
		t.Errorf("FltStat: received %v, expected %v", int(msg.FltStat), 0)
	}
	if msg.Alt != 0 {
		t.Errorf("Alt: received %v, expected %v", msg.Alt, 0)
	}
	if msg.Sqk != "3452" {
		t.Errorf("Sqk: received %s, expected %s", msg.Sqk, "3452")
	}
	if msg.MsgB != nil {
		t.Errorf("MsgB: received %s, expected nil", msg.MsgB)
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
	if msg.FltStat != -1 {
		t.Errorf("FltStat: received %v, expected %v", int(msg.FltStat), -1)
	}
	if msg.Alt != 0 {
		t.Errorf("Alt: received %v, expected %v", msg.Alt, 0)
	}
	if msg.Sqk != "" {
		t.Errorf("Sqk: received %s, expected %s", msg.Sqk, "")
	}
	if msg.MsgB != nil {
		t.Errorf("MsgB: received %s, expected nil", msg.MsgB)
	}
}

// TestDF20 tests DF5 identity reply
func TestDF20(t *testing.T) {
	msg := new(Message)
	err := msg.Decode([]byte{0xa0, 0x00, 0x0f, 0x98, 0x20, 0x05, 0x72,
		0x73, 0xdf, 0x8d, 0x20, 0xe2, 0xcf, 0x30})
	if err != nil {
		t.Fatal("received unexpected error", err)
	}
	if msg.ICAO != "a52333" {
		t.Errorf("ICAO: received %s, expected %s", msg.ICAO, "a52333")
	}
	if msg.Format != DF20 {
		t.Errorf("Format: received %v, expected %v", int(msg.Format), 20)
	}
	if msg.Cap != -1 {
		t.Errorf("Cap: received %v, expected %v", int(msg.Cap), -1)
	}
	if msg.FltStat != 0 {
		t.Errorf("FltStat: received %v, expected %v", int(msg.FltStat), 0)
	}
	if msg.Alt != 24000 {
		t.Errorf("Alt: received %v, expected %v", msg.Alt, 24000)
	}
	if msg.Sqk != "" {
		t.Errorf("Sqk: received %s, expected %s", msg.Sqk, "")
	}
	if !bytes.Equal(msg.MsgB, []byte{0x20, 0x05, 0x72, 0x73, 0xdf, 0x8d, 0x20}) {
		t.Errorf("MsgB: received %s, expected %s", msg.MsgB, "rs^")
	}
}

// TestDF21 tests DF21 Comm-B identity reply
func TestDF21(t *testing.T) {
	msg := new(Message)
	err := msg.Decode([]byte{0xac, 0x19, 0xb2, 0x95, 0x73, 0x48, 0x2f,
		0x69, 0x63, 0x66, 0x36, 0x36, 0x02, 0x2b})
	if err != nil {
		t.Fatal("received unexpected error", err)
	}
	if msg.ICAO != "a97db4" {
		t.Errorf("ICAO: received %s, expected %s", msg.ICAO, "a97db4")
	}
	if msg.Format != DF21 {
		t.Errorf("Format: received %v, expected %v", int(msg.Format), 21)
	}
	if msg.Cap != -1 {
		t.Errorf("Cap: received %v, expected %v", int(msg.Cap), -1)
	}
	if msg.FltStat != 4 {
		t.Errorf("FltStat: received %v, expected %v", int(msg.FltStat), 4)
	}
	if msg.Alt != 0 {
		t.Errorf("Alt: received %v, expected %v", msg.Alt, 0)
	}
	if msg.Sqk != "6017" {
		t.Errorf("Sqk: received %s, expected %s", msg.Sqk, "6017")
	}
	if !bytes.Equal(msg.MsgB, []byte{0x73, 0x48, 0x2f, 0x69, 0x63, 0x66, 0x36}) {
		t.Errorf("MsgB: received %s, expected %s", msg.MsgB, "sH/icf6")
	}
}
