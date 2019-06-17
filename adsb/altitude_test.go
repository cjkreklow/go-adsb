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

package adsb

import (
	"testing"
)

func TestAlt13Errors(t *testing.T) {
	_, err := decodeAlt13(0xF000)
	if err == nil {
		t.Error("received nil, expected invalid data length")
	} else if err.Error() != "invalid data length" {
		t.Errorf("received %s, expected invalid data length", err)
	}
	a, err := decodeAlt13(0x0000)
	if err != nil {
		t.Errorf("received %s, expected nil", err)
	}
	if a != 0 {
		t.Errorf("received %d, expected 0", a)
	}
	_, err = decodeAlt13(0x0FFF)
	if err == nil {
		t.Error("received nil, expected metric altitude not supported")
	} else if err.Error() != "metric altitude not supported" {
		t.Errorf("received %s, expected metric altitude not supported", err)
	}
	_, err = decodeAlt13(0x1FAF)
	if err == nil {
		t.Error("received nil, expected invalid altitude value")
	} else if err.Error() != "invalid altitude value" {
		t.Errorf("received %s, expected invalid altitude value", err)
	}
}

func TestAlt12Errors(t *testing.T) {
	_, err := decodeAlt12(0xF000)
	if err == nil {
		t.Error("received nil, expected invalid data length")
	} else if err.Error() != "invalid data length" {
		t.Errorf("received %s, expected invalid data length", err)
	}
	a, err := decodeAlt12(0x0000)
	if err != nil {
		t.Errorf("received %s, expected nil", err)
	}
	if a != 0 {
		t.Errorf("received %d, expected 0", a)
	}
}
