// Copyright 2020 Collin Kreklow
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

package beast_test

import (
	"bytes"
	"encoding"
	"encoding/hex"
	"errors"
	"testing"
	"time"

	"kreklow.us/go/go-adsb/beast"
)

func TestInterfaces(t *testing.T) {
	t.Run("BinaryUnmarshaler", testBinaryUnmarshaler)
	t.Run("BinaryMarshaler", testBinaryMarshaler)
}

func testBinaryUnmarshaler(t *testing.T) {
	var f interface{} = new(beast.Frame)

	_, ok := f.(encoding.BinaryUnmarshaler)
	if !ok {
		t.Error("Frame does not support BinaryUnmarshaler")
	}
}

func testBinaryMarshaler(t *testing.T) {
	var f interface{} = new(beast.Frame)

	_, ok := f.(encoding.BinaryMarshaler)
	if !ok {
		t.Error("Frame does not support BinaryMarshaler")
	}
}

func TestUnmarshalError(t *testing.T) {
	t.Run("NoFormat", testUnmarshalShort)
	t.Run("BadLength1", testUnmarshalBadLength1)
	t.Run("BadLength2", testUnmarshalBadLength2)
	t.Run("BadLength3", testUnmarshalBadLength3)
	t.Run("BadEscape", testUnmarshalBadEscape)
	t.Run("BadType", testUnmarshalBadType)
	t.Run("NoType", testUnmarshalNoType)
}

func testUnmarshalShort(t *testing.T) {
	testUnmarshalError(t, "1affffff", "received truncated data")
}

func testUnmarshalBadLength1(t *testing.T) {
	testUnmarshalError(t, "1a31ffffffffffffff1a1a", "expected 11 bytes, received 10")
}

func testUnmarshalBadLength2(t *testing.T) {
	testUnmarshalError(t, "1a32ffff1a1affff1a1affff", "expected 16 bytes, received 10")
}

func testUnmarshalBadLength3(t *testing.T) {
	testUnmarshalError(t, "1a33ffffffffffffffffffffffffffffffffffffffffffffff", "expected 23 bytes, received 25")
}

func testUnmarshalBadEscape(t *testing.T) {
	testUnmarshalError(t, "1a31ffffffffffffffffffffffffff1a", "expected 11 bytes, received 16")
}

func testUnmarshalBadType(t *testing.T) {
	testUnmarshalError(t, "1a3affffffffffffffffff", "invalid data format: 1a3a")
}

func testUnmarshalNoType(t *testing.T) {
	testUnmarshalError(t, "ff00ff00ff00ff00ff00", "invalid data format: ff00")
}

func testUnmarshalError(t *testing.T, msg string, e string) {
	f := new(beast.Frame)

	b, err := hex.DecodeString(msg)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = f.UnmarshalBinary(b)
	if err == nil {
		t.Errorf("expected %s, received nil", e)
	} else if err.Error() != e {
		t.Errorf("expected %s, received %s", e, err.Error())
	}
}

func TestNoDataErrors(t *testing.T) {
	t.Run("Marshal", testNoDataMarshal)
	t.Run("ADSB", testNoDataADSB)
	t.Run("ModeAC", testNoDataModeAC)
	t.Run("Timestamp", testNoDataTimestamp)
	t.Run("Signal", testNoDataSignal)
}

func testNoDataMarshal(t *testing.T) {
	f := new(beast.Frame)

	b, err := f.MarshalBinary()
	if err == nil {
		t.Error("expected error, received nil")
	} else if !errors.Is(err, beast.ErrNoData) {
		t.Fatal("unexpected error:", err)
	}

	if b != nil {
		t.Errorf("expected nil, received %x", b)
	}
}

func testNoDataADSB(t *testing.T) {
	f := new(beast.Frame)

	b, err := f.MarshalADSB()
	if err == nil {
		t.Error("expected error, received nil")
	} else if !errors.Is(err, beast.ErrNoData) {
		t.Fatal("unexpected error:", err)
	}

	if b != nil {
		t.Errorf("expected nil, received %x", b)
	}
}

func testNoDataModeAC(t *testing.T) {
	f := new(beast.Frame)

	b, err := f.ModeAC()
	if err == nil {
		t.Error("expected error, received nil")
	} else if !errors.Is(err, beast.ErrNoData) {
		t.Fatal("unexpected error:", err)
	}

	if b != nil {
		t.Errorf("expected nil, received %x", b)
	}
}

func testNoDataTimestamp(t *testing.T) {
	f := new(beast.Frame)

	ts, err := f.Timestamp()
	if err == nil {
		t.Error("expected error, received nil")
	} else if !errors.Is(err, beast.ErrNoData) {
		t.Fatal("unexpected error:", err)
	}

	if ts != time.Duration(0) {
		t.Errorf("expected nil, received %s", ts)
	}
}

func testNoDataSignal(t *testing.T) {
	f := new(beast.Frame)

	sig, err := f.Signal()
	if err == nil {
		t.Error("expected error, received nil")
	} else if !errors.Is(err, beast.ErrNoData) {
		t.Fatal("unexpected error:", err)
	}

	if sig != 0 {
		t.Errorf("expected nil, received %d", sig)
	}
}

func TestUnmarshal(t *testing.T) {
	t.Run("Bytes", testUnmarshalBytes)
	t.Run("Marshal", testUnmarshalMarshal)
	t.Run("ADSB", testUnmarshalADSB)
	t.Run("ModeAC", testUnmarshalModeAC)
	t.Run("Timestamp", testUnmarshalTimestamp)
	t.Run("Signal", testUnmarshalSignal)
}

func testUnmarshalBytes(t *testing.T) {
	msg, err := hex.DecodeString("1a321a1af933baf325c45da99adad95ff6")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	data, err := hex.DecodeString("1a321af933baf325c45da99adad95ff6")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	f := new(beast.Frame)

	err = f.UnmarshalBinary(msg)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if !bytes.Equal(data, f.Bytes()) {
		t.Errorf("expected %x, received %x", data, f.Bytes())
	}
}

func testUnmarshalMarshal(t *testing.T) {
	msg, err := hex.DecodeString("1a321a1af933baf325c45da99adad95ff6")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	f := new(beast.Frame)

	err = f.UnmarshalBinary(msg)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	rm, err := f.MarshalBinary()
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if !bytes.Equal(msg, rm) {
		t.Errorf("expected %x, received %x", msg, rm)
	}
}

func testUnmarshalADSB(t *testing.T) { //nolint:dupl // ADSB/ModeAC tests are similar
	msg, err := hex.DecodeString("1a321a1af933baf325c45da99adad91a1a1a1a")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	data, err := hex.DecodeString("5da99adad91a1a")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	f := new(beast.Frame)

	err = f.UnmarshalBinary(msg)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	adsb, err := f.MarshalADSB()
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if !bytes.Equal(data, adsb) {
		t.Errorf("expected %x, received %x", data, adsb)
	}

	ac, err := f.ModeAC()
	if err != nil && !errors.Is(err, beast.ErrNoData) {
		t.Fatal("unexpected error:", err)
	} else if err == nil {
		t.Errorf("expected %s, received nil", beast.ErrNoData)
	}

	if ac != nil {
		t.Errorf("expected nil, received %x", ac)
	}
}

func testUnmarshalModeAC(t *testing.T) { //nolint:dupl // ADSB/ModeAC tests are similar
	msg, err := hex.DecodeString("1a311a1af933baf325c45047")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	data, err := hex.DecodeString("5047")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	f := new(beast.Frame)

	err = f.UnmarshalBinary(msg)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	ac, err := f.ModeAC()
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if !bytes.Equal(data, ac) {
		t.Errorf("expected %x, received %x", data, ac)
	}

	adsb, err := f.MarshalADSB()
	if err != nil && !errors.Is(err, beast.ErrNoData) {
		t.Fatal("unexpected error:", err)
	} else if err == nil {
		t.Errorf("expected %s, received nil", beast.ErrNoData)
	}

	if adsb != nil {
		t.Errorf("expected nil, received %x", adsb)
	}
}

func testUnmarshalTimestamp(t *testing.T) {
	msg, err := hex.DecodeString("1a321a1af933baf325c45da99adad95ff6")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	ts := time.Duration(2471468089070000)

	f := new(beast.Frame)

	err = f.UnmarshalBinary(msg)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	rts, err := f.Timestamp()
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if ts != rts {
		t.Errorf("expected %s, received %s", ts, rts)
	}
}

func testUnmarshalSignal(t *testing.T) {
	msg, err := hex.DecodeString("1a321a1af933baf325c45da99adad95ff6")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	var sig uint8 = 0xc4

	f := new(beast.Frame)

	err = f.UnmarshalBinary(msg)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	rs, err := f.Signal()
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if sig != rs {
		t.Errorf("expected %d, received %d", sig, rs)
	}
}
