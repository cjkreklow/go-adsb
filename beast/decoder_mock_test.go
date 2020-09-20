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

//nolint:testpackage,goerr113 // internal test functions
package beast

import (
	"bytes"
	"errors"
	"testing"

	"kreklow.us/go/go-adsb/beast/internal"
)

var _ decoderReader = new(internal.MockReader)
var _ decoderBuffer = new(internal.MockBuffer)

func TestMockDecodeErrors(t *testing.T) {
	t.Run("SeekPeekError", testSeekPeekErr)
	t.Run("SeekDiscardError", testSeekDiscardErr)
	t.Run("AfterSeekPeekError", testAfterSeekPeekErr)
	t.Run("AfterSeekDiscardError", testAfterSeekDiscardErr)
	t.Run("TypeWriteError", testTypeWriteErr)
	t.Run("ReadWriteError", testReadWriteErr)
	t.Run("ReadUnreadError", testReadUnreadErr)
	t.Run("EscapeWriteError", testEscapeWriteErr)
	t.Run("EscapeDiscardError", testEscapeDiscardErr)
}

func testSeekPeekErr(t *testing.T) {
	mr := &internal.MockReader{
		Buf:       bytes.NewBuffer([]byte{0xff, 0xff, 0xff}),
		PeekCount: 2,
		PeekErr:   errors.New("seek peek error"),
	}

	mb := new(bytes.Buffer)

	testMockDecoder(t, mr, mb, mr.PeekErr)
}

func testSeekDiscardErr(t *testing.T) {
	mr := &internal.MockReader{
		Buf:          bytes.NewBuffer([]byte{0xff, 0xff, 0xff, 0xff, 0x1a, 0x31}),
		PeekCount:    6,
		DiscardCount: 0,
		DiscardErr:   errors.New("seek discard error"),
	}

	mb := new(bytes.Buffer)

	testMockDecoder(t, mr, mb, mr.DiscardErr)
}

func testAfterSeekPeekErr(t *testing.T) {
	mr := &internal.MockReader{
		Buf:          bytes.NewBuffer([]byte{0xff, 0xff, 0xff, 0xff, 0x1a, 0x31}),
		PeekCount:    6,
		DiscardCount: 2,
		PeekErr:      errors.New("after seek peek error"),
	}

	mb := new(bytes.Buffer)

	testMockDecoder(t, mr, mb, mr.PeekErr)
}

func testAfterSeekDiscardErr(t *testing.T) {
	mr := &internal.MockReader{
		Buf:          bytes.NewBuffer([]byte{0xff, 0xff, 0xff, 0xff, 0x1a, 0x31, 0x1a, 0x31}),
		PeekCount:    10,
		DiscardCount: 2,
		DiscardErr:   errors.New("after seek discard error"),
	}

	mb := new(bytes.Buffer)

	testMockDecoder(t, mr, mb, mr.DiscardErr)
}

func testTypeWriteErr(t *testing.T) {
	mr := &internal.MockReader{
		Buf:          bytes.NewBuffer([]byte{0x1a, 0x31, 0x1a}),
		PeekCount:    2,
		DiscardCount: 2,
		ReadCount:    1,
	}

	mb := &internal.MockBuffer{
		Buf:      new(bytes.Buffer),
		WriteErr: errors.New("type write error"),
	}

	testMockDecoder(t, mr, mb, mb.WriteErr)
}

func testReadWriteErr(t *testing.T) {
	mr := &internal.MockReader{
		Buf:          bytes.NewBuffer([]byte{0x1a, 0x31, 0xff}),
		PeekCount:    2,
		DiscardCount: 2,
		ReadCount:    1,
	}

	mb := &internal.MockBuffer{
		Buf:          new(bytes.Buffer),
		WriteCount:   2,
		WriteByteErr: errors.New("read write error"),
	}

	testMockDecoder(t, mr, mb, mb.WriteByteErr)
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

	mb := new(bytes.Buffer)

	testMockDecoder(t, mr, mb, mr.UnreadErr)
}

func testEscapeWriteErr(t *testing.T) {
	mr := &internal.MockReader{
		Buf:          bytes.NewBuffer([]byte{0x1a, 0x31, 0x1a, 0x1a, 0x1a}),
		PeekCount:    4,
		DiscardCount: 2,
		ReadCount:    1,
		UnreadCount:  1,
	}

	mb := &internal.MockBuffer{
		Buf:          new(bytes.Buffer),
		WriteCount:   2,
		WriteByteErr: errors.New("escape write error"),
	}

	testMockDecoder(t, mr, mb, mb.WriteByteErr)
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

	mb := new(bytes.Buffer)

	testMockDecoder(t, mr, mb, mr.DiscardErr)
}

func testMockDecoder(t *testing.T, mr decoderReader, mb decoderBuffer, e error) {
	f := new(Frame)
	d := new(Decoder)

	d.r = mr
	d.buf = mb

	err := d.Decode(f)
	if err != nil && !errors.Is(err, e) {
		t.Error("unexpected error:", err)
	} else if err == nil {
		t.Fatal("expected error, received nil")
	}
}
