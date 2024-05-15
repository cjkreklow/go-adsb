// Copyright 2024 Collin Kreklow
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

//nolint:testpackage,goerr113 // internal test functions
package beast

import (
	"bytes"
	"errors"
	"testing"

	"kreklow.us/go/go-adsb/beast/internal"
)

var _ decoderReader = new(internal.MockReader)

func TestMockDecodeErrors(t *testing.T) {
	t.Run("SeekPeekError", testSeekPeekErr)
	t.Run("SeekDiscardError", testSeekDiscardErr)
	t.Run("AfterSeekPeekError", testAfterSeekPeekErr)
	t.Run("AfterSeekDiscardError", testAfterSeekDiscardErr)
	t.Run("ReadUnreadError", testReadUnreadErr)
	t.Run("ReadPeekError", testReadPeekErr)
	t.Run("EscapeDiscardError", testEscapeDiscardErr)
}

func testSeekPeekErr(t *testing.T) {
	mr := &internal.MockReader{
		Buf:       bytes.NewBuffer([]byte{0xff, 0xff, 0xff}),
		PeekCount: 2,
		PeekErr:   errors.New("seek peek error"),
	}

	testMockDecoder(t, mr, mr.PeekErr)
}

func testSeekDiscardErr(t *testing.T) {
	mr := &internal.MockReader{
		Buf:          bytes.NewBuffer([]byte{0xff, 0xff, 0xff, 0xff, 0x1a, 0x31}),
		PeekCount:    6,
		DiscardCount: 0,
		DiscardErr:   errors.New("seek discard error"),
	}

	testMockDecoder(t, mr, mr.DiscardErr)
}

func testAfterSeekPeekErr(t *testing.T) {
	mr := &internal.MockReader{
		Buf:          bytes.NewBuffer([]byte{0xff, 0xff, 0xff, 0xff, 0x1a, 0x31}),
		PeekCount:    6,
		DiscardCount: 2,
		PeekErr:      errors.New("after seek peek error"),
	}

	testMockDecoder(t, mr, mr.PeekErr)
}

func testAfterSeekDiscardErr(t *testing.T) {
	mr := &internal.MockReader{
		Buf:          bytes.NewBuffer([]byte{0xff, 0xff, 0xff, 0xff, 0x1a, 0x31, 0x1a, 0x31}),
		PeekCount:    10,
		DiscardCount: 2,
		DiscardErr:   errors.New("after seek discard error"),
	}

	testMockDecoder(t, mr, mr.DiscardErr)
}

func testReadUnreadErr(t *testing.T) {
	mr := &internal.MockReader{
		Buf:          bytes.NewBuffer([]byte{0x1a, 0x31, 0x1a}),
		PeekCount:    2,
		DiscardCount: 2,
		ReadCount:    1,
		UnreadCount:  0,
		UnreadErr:    errors.New("read unread error"),
	}

	testMockDecoder(t, mr, mr.UnreadErr)
}

func testReadPeekErr(t *testing.T) {
	mr := &internal.MockReader{
		Buf:          bytes.NewBuffer([]byte{0x1a, 0x31, 0x1a}),
		PeekCount:    2,
		DiscardCount: 2,
		ReadCount:    1,
		UnreadCount:  1,
		PeekErr:      errors.New("read peek error"),
	}

	testMockDecoder(t, mr, mr.PeekErr)
}

func testEscapeDiscardErr(t *testing.T) {
	mr := &internal.MockReader{
		Buf:          bytes.NewBuffer([]byte{0x1a, 0x31, 0x1a, 0x1a, 0x1a}),
		PeekCount:    4,
		DiscardCount: 2,
		ReadCount:    1,
		UnreadCount:  1,
		DiscardErr:   errors.New("escape discard error"),
	}

	testMockDecoder(t, mr, mr.DiscardErr)
}

func testMockDecoder(t *testing.T, mr decoderReader, e error) {
	t.Helper()

	f := new(Frame)
	d := new(Decoder)

	d.r = mr

	err := d.Decode(f)
	if err != nil && !errors.Is(err, e) {
		t.Error("unexpected error:", err)
	} else if err == nil {
		t.Fatal("expected error, received nil")
	}
}
