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
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

// Decode takes a []byte containing a raw 56- or 112-bit ADS-B message
// and populates the Message struct.
func (m *Message) Decode(msg []byte) error {
	var err error

	if len(msg) != 7 && len(msg) != 14 {
		return errors.New("invalid message length")
	}

	m.raw = msg
	m.setParity()
	m.DF = DF(m.raw.Bits8(1, 5))
	m.CA = -1
	m.FS = -1
	m.VS = -1
	m.TC = -1
	m.SS = -1

	switch m.DF {
	case DF0, DF16, DF4, DF20:
		err = m.decodeAltMsg()
	case DF5, DF21:
		err = m.decodeIdentMsg()
	case DF11:
		err = m.decode11()
	case DF17:
		err = m.decode17()
	default:
		err = fmt.Errorf("unsupported format: %v", int(m.DF))
	}
	if err != nil {
		return err
	}

	return nil
}

// decode DF11 all call reply
func (m *Message) decode11() error {
	m.CA = CA(m.raw.Bits8(6, 8))
	m.ICAO = strconv.FormatUint(m.raw.Bits64(9, 32), 16)

	return nil
}

// decode DF0/DF16 air-to-air and DF4/DF20 altitude reply
func (m *Message) decodeAltMsg() error {
	switch m.DF {
	case DF0, DF16:
		m.VS = VS(m.raw.Bit(6))
	case DF4, DF20:
		m.FS = FS(m.raw.Bits8(6, 8))
	}

	m.setICAOFromAP()

	a, err := decodeAlt13(m.raw.Bits16(20, 32))
	if err != nil {
		return err
	}

	m.Alt = a

	if m.DF == DF20 {
		err := m.decodeCommB()
		if err != nil {
			return err
		}
	}

	return nil
}

// decode DF5 and DF21 identity reply
func (m *Message) decodeIdentMsg() error {
	m.FS = FS(m.raw.Bits8(6, 8))
	m.setICAOFromAP()

	f := [][]int{
		{25, 23, 21},
		{31, 29, 27},
		{24, 22, 20},
		{32, 30, 28},
	}

	var id uint64

	for _, v := range f {
		for _, x := range v {
			id <<= 1
			id |= uint64(m.raw.Bit(x))
		}
	}

	m.Sqk = strconv.FormatUint(id, 8)

	if m.DF == DF21 {
		err := m.decodeCommB()
		if err != nil {
			return err
		}
	}

	return nil
}

// decode DF17 extended squitter message
func (m *Message) decode17() error {
	m.CA = CA(m.raw.Bits8(6, 8))
	m.ICAO = strconv.FormatUint(m.raw.Bits64(9, 32), 16)

	m.TC = TC(m.raw.Bits8(33, 37))

	if m.TC >= 1 && m.TC <= 4 {
		m.Cat = m.raw.Bits8(38, 40)
		m.setCall(m.raw.Bits64(41, 88))
	}

	if m.TC >= 9 && m.TC <= 18 {
		m.SS = SS(m.raw.Bits8(38, 39))

		a, err := decodeAlt12(m.raw.Bits16(41, 52))
		if err != nil {
			return err
		}

		m.Alt = a

		m.CPR = new(CPR)
		m.CPR.Nb = 17
		m.CPR.T = m.raw.Bit(53)
		m.CPR.F = m.raw.Bit(54)
		m.CPR.Lat = m.raw.Bits32(55, 71)
		m.CPR.Lon = m.raw.Bits32(72, 88)
	}

	return nil
}

// utility function to set ICAO when XORed into an AP field
func (m *Message) setICAOFromAP() {
	b := m.raw.Bits32((len(m.raw)*8)-23, len(m.raw)*8)
	m.ICAO = fmt.Sprintf("%06x", b^m.parity)
}

// utility function to set Call from a data field
func (m *Message) setCall(b uint64) {
	c := []byte("?ABCDEFGHIJKLMNOPQRSTUVWXYZ????? ???????????????0123456789??????")

	call := make([]byte, 8)

	var i uint
	for i = 0; i < 8; i++ {
		call[i] = c[(b>>(42-(i*6)))&0x3F]
	}

	m.Call = string(bytes.TrimRight(call, " "))
}
