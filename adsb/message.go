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

	ICAO    string // ICAO transponder address
	Format  DF     // downlink format
	Cap     CA     // capability
	FltStat FS     // flight status
	Alt     int64  // altitude
	Sqk     string // transponder (squawk) code
	MsgB    []byte // data link message
}

// Decode takes a []byte containing a raw 56- or 112-bit ADS-B message
// and populates the Message struct.
//
// Decode currently supports DF4 and DF11 format messages.
func (m *Message) Decode(msg []byte) error {
	var err error

	if len(msg) != 7 && len(msg) != 14 {
		return errors.New("invalid message length")
	}

	m.raw = msg
	m.setParity()
	m.Format = DF(m.raw[0] >> 3) // bits 1-5
	m.Cap = -1
	m.FltStat = -1

	switch m.Format {
	case DF4:
		err = m.decodeAlt()
	case DF5:
		err = m.decodeIdent()
	case DF11:
		err = m.decode11()
	case DF20:
		err = m.decodeAlt()
	case DF21:
		err = m.decodeIdent()
	default:
		return fmt.Errorf("unsupported format: %v", int(m.Format))
	}
	if err != nil {
		return err
	}

	return nil
}
