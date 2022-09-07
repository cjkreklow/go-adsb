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

// Package internal contains mock objects for testing.
package internal

import (
	"bytes"
	"errors"
)

// MockReader implements decoderReader.
type MockReader struct {
	Buf *bytes.Buffer

	DiscardCount int
	PeekCount    int
	ReadCount    int
	UnreadCount  int

	DiscardErr error
	PeekErr    error
	ReadErr    error
	UnreadErr  error
}

// Buffered returns the number of bytes in Buf.
func (r *MockReader) Buffered() int {
	return r.Buf.Len()
}

// Discard the specified number of bytes from Buf.
func (r *MockReader) Discard(n int) (int, error) {
	if n > r.DiscardCount {
		if r.DiscardErr != nil {
			return 0, r.DiscardErr
		}

		return 0, errors.New("unexpected error in Discard") //nolint:goerr113 // no error to wrap
	}

	r.DiscardCount -= n

	return n, nil
}

// Peek returns the specified number of bytes from Buf without
// discarding.
func (r *MockReader) Peek(n int) ([]byte, error) {
	if n > r.PeekCount {
		if r.PeekErr != nil {
			return nil, r.PeekErr
		}

		return nil, errors.New("unexpected error in Peek") //nolint:goerr113 // no error to wrap
	}

	r.PeekCount -= n

	return r.Buf.Next(n), nil
}

// ReadByte returns the next available byte from Buf.
func (r *MockReader) ReadByte() (byte, error) {
	if r.ReadCount == 0 {
		if r.ReadErr != nil {
			return 0, r.ReadErr
		}

		return 0, errors.New("unexpected error in ReadByte") //nolint:goerr113 // no error to wrap
	}

	r.ReadCount--

	return r.Buf.ReadByte() //nolint:wrapcheck // unnecessary wrapping
}

// UnreadByte places the last read byte back into Buf.
func (r *MockReader) UnreadByte() error {
	if r.UnreadCount == 0 {
		if r.UnreadErr != nil {
			return r.UnreadErr
		}

		return errors.New("unexpected error in UnreadByte") //nolint:goerr113 // no error to wrap
	}

	r.UnreadCount--

	return nil
}

// MockFrame implements BinaryUnmarshaler.
type MockFrame struct {
	Buf bytes.Buffer
}

// UnmarshalBinary stores data into Buf.
func (f *MockFrame) UnmarshalBinary(data []byte) error {
	f.Buf.Write(data)

	return nil
}
