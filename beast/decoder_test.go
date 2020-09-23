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
	"kreklow.us/go/go-adsb/beast/internal"
)

func TestDecode(t *testing.T) {
	t.Run("Strip", testDecodeStrip)
	t.Run("NoStrip", testDecodeNoStrip)
}

func testDecodeNoStrip(t *testing.T) {
	i := "1a331a1af933bbc63ec68f1a1a9ada58b98446e703357e241a1a"

	testDecoder(t, i, "")
}

func testDecodeStrip(t *testing.T) {
	i := "1a331a1af933bbc63ec68f1a1a9ada58b98446e703357e241a1a"
	o := "1a331af933bbc63ec68f1a9ada58b98446e703357e241a"

	testDecoder(t, i, o)
}

func testDecoder(t *testing.T, in string, out string) {
	ib, err := hex.DecodeString(in)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	r := bytes.NewReader(ib)
	d := beast.NewDecoder(r)
	f := new(internal.MockFrame)

	ob := ib

	if out != "" {
		d.StripEscape = true

		ob, err = hex.DecodeString(out)
		if err != nil {
			t.Fatal("unexpected error:", err)
		}
	}

	err = d.Decode(f)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if !bytes.Equal(ob, f.Buf.Bytes()) {
		t.Errorf("expected %x, received %x", ob, f.Buf.Bytes())
	}
}

func TestDecodeError(t *testing.T) {
	t.Run("Null", testDecodeNull)
	t.Run("Short1", testDecodeShort1)
	t.Run("Short2", testDecodeShort2)
	t.Run("Short3", testDecodeShort3)
	t.Run("NoData", testDecodeNoData)
	t.Run("Truncate", testDecodeTruncated)
	t.Run("Unsupported", testDecodeUnsupported)
	t.Run("Corrupt", testDecodeCorrupt)
}

func testDecodeNull(t *testing.T) {
	testDecoderError(t, "", "error reading stream: EOF", io.EOF)
}

func testDecodeShort1(t *testing.T) {
	testDecoderError(t, "1a", "error reading stream: EOF", io.EOF)
}

func testDecodeShort2(t *testing.T) {
	testDecoderError(t, "1a31", "error unmarshalling data: received truncated data", nil)
}

func testDecodeShort3(t *testing.T) {
	testDecoderError(t, "1a331a1aff00ff00ff00ff00ff00", "error unmarshalling data: expected 23 bytes, received 13", nil)
}

func testDecodeNoData(t *testing.T) {
	testDecoderError(t, "ff00ff00ff00ff00ff00ff00ff00", "no frame data found", nil)
}

func testDecodeTruncated(t *testing.T) {
	testDecoderError(t, "1a32ffffffffffffffffffff1a33ff", "error unmarshalling data: expected 16 bytes, received 12", nil)
}

func testDecodeUnsupported(t *testing.T) {
	testDecoderError(t, "1affffff", "no frame data found", nil)
}

func testDecodeCorrupt(t *testing.T) {
	testDecoderError(t, "1a32ff1aff", "data stream corrupt", nil)
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
