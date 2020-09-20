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
	"encoding/hex"
	"errors"
	"testing"
	"time"

	"kreklow.us/go/go-adsb/beast"
)

func TestUnmarshalError(t *testing.T) {
	t.Run("NoFormat", testUnmarshalNoFormat)
	t.Run("BadLength1", testUnmarshalBadLength1)
	t.Run("BadLength2", testUnmarshalBadLength2)
	t.Run("BadLength3", testUnmarshalBadLength3)
	t.Run("Type4", testUnmarshalType4)
	t.Run("BadType", testUnmarshalBadType)
}

func testUnmarshalNoFormat(t *testing.T) {
	testUnmarshalError(t, "ff0000ff", "format identifier not found")
}

func testUnmarshalBadLength1(t *testing.T) {
	testUnmarshalError(t, "1a31ffff", "expected 11 bytes, received 4")
}

func testUnmarshalBadLength2(t *testing.T) {
	testUnmarshalError(t, "1a32ffff", "expected 16 bytes, received 4")
}

func testUnmarshalBadLength3(t *testing.T) {
	testUnmarshalError(t, "1a33ffff", "expected 23 bytes, received 4")
}

func testUnmarshalType4(t *testing.T) {
	testUnmarshalError(t, "1a34ffff", "format not supported: 34")
}

func testUnmarshalBadType(t *testing.T) {
	testUnmarshalError(t, "1affffff", "invalid format identifier: ff")
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

func TestNoDataError(t *testing.T) {
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

	ts, err := f.Timestamp()
	if err == nil {
		t.Error("expected error, received nil")
	} else if !errors.Is(err, beast.ErrNoData) {
		t.Fatal("unexpected error:", err)
	}

	if ts != time.Duration(0) {
		t.Errorf("expected nil, received %s", ts)
	}

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

func TestBytes(t *testing.T) {
	m := "1a3216f933baf325c45da99adad95ff6"

	b, err := hex.DecodeString(m)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	f := new(beast.Frame)

	err = f.UnmarshalBinary(b)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	rb := f.Bytes()
	rm := hex.EncodeToString(rb)

	if m != rm {
		t.Errorf("received %s, expected %s", rm, m)
	}
}
