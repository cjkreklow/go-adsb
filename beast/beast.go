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

// Package beast provides objects and methods for decoding and managing
// raw Mode-S Beast format frames.
package beast

import (
	"fmt"
)

// beastError is the error type for the beast library.
type beastError struct {
	msg  string // error message string from this library
	werr error  // wrapped error from downstream function
}

// Error returns the string value of an error.
func (e beastError) Error() string {
	if e.werr == nil {
		return e.msg
	}

	return e.msg + ": " + e.werr.Error()
}

// Unwrap returns an underlying error if applicable.
func (e beastError) Unwrap() error {
	return e.werr
}

// newError returns a new beastError.
func newError(w error, m string) beastError {
	return beastError{
		msg:  m,
		werr: w,
	}
}

// newErrorf returns a new beastError with a Printf-style message.
func newErrorf(w error, m string, v ...interface{}) beastError { //nolint:unparam // consistent with newError
	return beastError{
		msg:  fmt.Sprintf(m, v...),
		werr: w,
	}
}

// ErrNoData is returned when the frame does not contain the data
// necessary to return the requested information.
var ErrNoData beastError = newError(nil, "data not available")
