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
	"encoding/hex"
	"errors"
	"io"
	"testing"

	"kreklow.us/go/go-adsb/beast"
)

type testCase struct {
	Msg  string
	ADSB string

	Format    uint8
	Signal    uint8
	Timestamp int64
}

func TestDecode(t *testing.T) {
	t.Run("Type1", testDecode1)
	t.Run("Type2", testDecode2)
	t.Run("Type3", testDecode3)
}

func testDecode1(t *testing.T) {
	tc := &testCase{
		Msg:       "1a3116f933baf325c4cdab",
		ADSB:      "cdab",
		Format:    1,
		Timestamp: 2104964213144500,
		Signal:    196,
	}
	testDecoder(t, tc)
}

func testDecode2(t *testing.T) {
	tc := &testCase{
		Msg:       "1a3216f933baf325c45da99adad95ff6",
		ADSB:      "5da99adad95ff6",
		Format:    2,
		Timestamp: 2104964213144500,
		Signal:    196,
	}
	testDecoder(t, tc)
}

func testDecode3(t *testing.T) {
	tc := &testCase{
		Msg:       "1a3316f933bbc63ec68da99ada58b98446e703357e2417",
		ADSB:      "8da99ada58b98446e703357e2417",
		Format:    3,
		Timestamp: 2104964217648000,
		Signal:    198,
	}
	testDecoder(t, tc)
}

func testDecoder(t *testing.T, tc *testCase) {
	b, err := hex.DecodeString(tc.Msg)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	r := bytes.NewReader(b)
	d := beast.NewDecoder(r)
	f := new(beast.Frame)

	err = d.Decode(f)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	ts, err := f.Timestamp()
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if tc.Timestamp != ts.Nanoseconds() {
		t.Errorf("Timestamp: expected %d, received %d", tc.Timestamp, ts.Nanoseconds())
	}

	sig, err := f.Signal()
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if tc.Signal != sig {
		t.Errorf("Signal: expected %d, received %d", tc.Signal, sig)
	}

	adsb, err := f.MarshalADSB()
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if tc.ADSB != hex.EncodeToString(adsb) {
		t.Errorf("ADSB: expected %s, received %s", tc.ADSB, hex.EncodeToString(adsb))
	}
}

func TestDecodeError(t *testing.T) {
	t.Run("Null", testDecodeNull)
	t.Run("Short1", testDecodeShort1)
	t.Run("Short2", testDecodeShort2)
	t.Run("Short3", testDecodeShort3)
	t.Run("ShortUnescape", testDecodeShortUnescape)
	t.Run("BadStart", testDecodeBadStart)
	t.Run("NoData", testDecodeNoData)
	t.Run("Truncate", testDecodeTruncated)
	t.Run("Unsupported", testDecodeUnsupported)
	t.Run("Corrupt", testDecodeCorrupt)
	t.Run("Long", testDecodeLong)
}

func testDecodeNull(t *testing.T) {
	testDecoderError(t, "", "error reading stream: EOF", io.EOF)
}

func testDecodeShort1(t *testing.T) {
	testDecoderError(t, "1a", "error reading stream: EOF", io.EOF)
}

func testDecodeShort2(t *testing.T) {
	testDecoderError(t, "1a31", "error unmarshalling data: expected 11 bytes, received 2", nil)
}

func testDecodeShort3(t *testing.T) {
	testDecoderError(t, "1a331a", "error unmarshalling data: expected 23 bytes, received 2", nil)
}

func testDecodeShortUnescape(t *testing.T) {
	testDecoderError(t, "1a331a1a", "error unmarshalling data: expected 23 bytes, received 3", nil)
}

func testDecodeBadStart(t *testing.T) {
	testDecoderError(t, "ff00111a32ffff", "error unmarshalling data: expected 16 bytes, received 4", nil)
}

func testDecodeNoData(t *testing.T) {
	testDecoderError(t, "ff00", "no frame data found", nil)
}

func testDecodeTruncated(t *testing.T) {
	testDecoderError(t, "1a32ffff1a33ff", "error unmarshalling data: expected 16 bytes, received 4", nil)
}

func testDecodeUnsupported(t *testing.T) {
	testDecoderError(t, "1affffff", "no frame data found", nil)
}

func testDecodeCorrupt(t *testing.T) {
	testDecoderError(t, "1a32ff1aff", "data stream corrupt", nil)
}

func testDecodeLong(t *testing.T) {
	testDecoderError(t, "1a32ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", "data stream corrupt", nil)
}

func testDecoderError(t *testing.T, msg string, str string, we error) {
	b, err := hex.DecodeString(msg)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	r := bytes.NewReader(b)
	d := beast.NewDecoder(r)
	f := new(beast.Frame)

	err = d.Decode(f)
	if err == nil {
		t.Errorf("expected %s, received nil", str)

		return
	}

	if err.Error() != str {
		t.Errorf("expected %s, received %s", str, err.Error())
	}

	if we != nil && !errors.Is(err, we) {
		t.Errorf("expected type %T, received type %T", we, err)
	}
}
