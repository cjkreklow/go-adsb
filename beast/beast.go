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
	"bytes"
	"fmt"
	"time"
)

// Frame is a Mode-S Beast format message. A Frame is safe to reuse by
// calling UnmarshalBinary with new data.
type Frame struct {
	data bytes.Buffer
}

// UnmarshalBinary stores data from a single Beast message into the
// Frame.
//
// The data must have the escape codes removed from any escaped 0x1a
// values. Decoder.Decode() removes these escape codes by default.
func (f *Frame) UnmarshalBinary(data []byte) error {
	f.data.Reset()

	if data[0] != 0x1a {
		return newError(nil, "format identifier not found")
	}

	switch data[1] {
	case 0x31:
		if len(data) != 11 {
			return newErrorf(nil, "expected 11 bytes, received %d", len(data))
		}
	case 0x32:
		if len(data) != 16 {
			return newErrorf(nil, "expected 16 bytes, received %d", len(data))
		}
	case 0x33:
		if len(data) != 23 {
			return newErrorf(nil, "expected 23 bytes, received %d", len(data))
		}
	case 0x34:
		return newErrorf(nil, "format not supported: %x", data[1])
	default:
		return newErrorf(nil, "invalid format identifier: %x", data[1])
	}

	f.data.Write(data)

	return nil
}

// Bytes returns the full frame data. The returned slice remains valid
// until the next call to UnmarshalBinary. Modifying the returned slice
// directly may impact future Frame method calls.
func (f Frame) Bytes() []byte {
	return f.data.Bytes()
}

// MarshalADSB returns ADS-B or Mode-AC data in the Frame. The returned
// slice remains valid until the next call to UnmarshalBinary.
// Modifying the returned slice directly may impact future Frame method
// calls.
func (f Frame) MarshalADSB() ([]byte, error) {
	if f.data.Len() < 10 {
		return nil, ErrNoData
	}

	return f.data.Bytes()[9:], nil
}

// Timestamp returns the MLAT timestamp as a time.Duration.
func (f Frame) Timestamp() (time.Duration, error) {
	if f.data.Len() < 8 {
		return time.Duration(0), ErrNoData
	}

	d := f.data.Bytes()
	ts := int64(d[7]) | int64(d[6])<<8 | int64(d[5])<<16 |
		int64(d[4])<<24 | int64(d[3])<<32 | int64(d[2])<<40

	return time.Duration(ts * 1000 / 12).Round(time.Microsecond / 2), nil
}

// Signal returns the signal level.
func (f Frame) Signal() (uint8, error) {
	if f.data.Len() < 9 {
		return 0, ErrNoData
	}

	return f.data.Bytes()[8], nil
}

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
