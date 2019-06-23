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

package beast_test

import (
	"encoding/hex"
	"testing"

	"kreklow.us/go/go-adsb/beast"
)

func TestUnmarshalBadData(t *testing.T) {
	testUnmarshalError(t, "ff0000ff", "format identifier not found")
}

func TestUnmarshalBadLength2(t *testing.T) {
	testUnmarshalError(t, "1a32ffff", "expected 16 bytes, received 4")
}

func TestUnmarshalBadLength3(t *testing.T) {
	testUnmarshalError(t, "1a33ffff", "expected 23 bytes, received 4")
}

func TestUnmarshalType1(t *testing.T) {
	testUnmarshalError(t, "1a31ffff", "format not supported")
}

func TestUnmarshalType4(t *testing.T) {
	testUnmarshalError(t, "1a34ffff", "format not supported")
}

func TestUnmarshalBadType(t *testing.T) {
	testUnmarshalError(t, "1affffff", "invalid format identifier")
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
