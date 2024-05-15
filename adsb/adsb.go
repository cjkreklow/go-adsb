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

// Package adsb provides objects and methods for decoding and managing
// raw ADS-B messages.
package adsb

import "fmt"

// Public error variables.
var (
	errNotAvailable = newError(nil, "field not available")
	errUnsupported  = newError(nil, "format unsupported")

	// ErrNotAvailable is used to indicate that a field is not part of the
	// specification for the message format received. Each field error wraps
	// ErrNotAvailable, making it accessible by calling
	// errors.Is(err, adsb.ErrNotAvailable).
	ErrNotAvailable = errNotAvailable

	// ErrUnsupported is returned when the Downlink Format of a message
	// is not supported by Message. The error may be wrapped and should be
	// checked with errors.Is().
	ErrUnsupported = errUnsupported
)

// adsbError is the error type for the adsb library.
type adsbError struct {
	msg  string // error message string from this library
	werr error  // wrapped error from downstream function
}

// Error returns the string value of an error.
func (e adsbError) Error() string {
	if e.werr == nil {
		return e.msg
	}

	return e.msg + ": " + e.werr.Error()
}

// Unwrap returns an underlying error if applicable.
func (e adsbError) Unwrap() error {
	return e.werr
}

// newError returns a new adsbError.
func newError(w error, m string) adsbError {
	return adsbError{
		msg:  m,
		werr: w,
	}
}

// newErrorf returns a new adsbError with a Printf-style message.
func newErrorf(w error, m string, v ...interface{}) adsbError {
	return adsbError{
		msg:  fmt.Sprintf(m, v...),
		werr: w,
	}
}
