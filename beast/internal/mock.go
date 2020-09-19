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

//nolint:golint,goerr113 // internal test package
package internal

import (
	"bytes"
	"errors"
)

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

func (r *MockReader) Buffered() int {
	return r.Buf.Len()
}

func (r *MockReader) Discard(n int) (int, error) {
	if n > r.DiscardCount {
		if r.DiscardErr != nil {
			return 0, r.DiscardErr
		}

		return 0, errors.New("unexpected error in Discard")
	}

	r.DiscardCount -= n

	return n, nil
}

func (r *MockReader) Peek(n int) ([]byte, error) {
	if n > r.PeekCount {
		if r.PeekErr != nil {
			return nil, r.PeekErr
		}

		return nil, errors.New("unexpected error in Peek")
	}

	r.PeekCount -= n

	return r.Buf.Next(n), nil
}

func (r *MockReader) ReadByte() (byte, error) {
	if r.ReadCount == 0 {
		if r.ReadErr != nil {
			return 0, r.ReadErr
		}

		return 0, errors.New("unexpected error in ReadByte")
	}

	r.ReadCount--

	return r.Buf.ReadByte()
}

func (r *MockReader) UnreadByte() error {
	if r.UnreadCount == 0 {
		if r.UnreadErr != nil {
			return r.UnreadErr
		}

		return errors.New("unexpected error in UnreadByte")
	}

	r.UnreadCount--

	return nil
}
