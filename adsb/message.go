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
	"errors"
	"fmt"
)

// Message is an ADS-B message
type Message struct {
	raw    []uint8 // raw message
	parity uint32  // parity

	DF DF // downlink format
	CA CA // capability
	FS FS // flight status

	TC TC // extended squitter type

	ICAO string // ICAO transponder address
	Alt  int64  // altitude
	Sqk  string // transponder (squawk) code
	Call string // callsign
}

// Decode takes a []byte containing a raw 56- or 112-bit ADS-B message
// and populates the Message struct.
func (m *Message) Decode(msg []byte) error {
	var err error

	if len(msg) != 7 && len(msg) != 14 {
		return errors.New("invalid message length")
	}

	m.raw = msg
	m.setParity()
	m.DF = DF(m.raw[0] >> 3) // bits 1-5
	m.CA = -1
	m.FS = -1
	m.TC = -1

	switch m.DF {
	case DF4:
		err = m.decodeAltMsg()
	case DF5:
		err = m.decodeIdentMsg()
	case DF11:
		err = m.decode11()
	case DF17:
		err = m.decode17()
	case DF20:
		err = m.decodeAltMsg()
	case DF21:
		err = m.decodeIdentMsg()
	default:
		return fmt.Errorf("unsupported format: %v", int(m.DF))
	}
	if err != nil {
		return err
	}

	return nil
}
