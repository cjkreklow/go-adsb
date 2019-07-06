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
)

type msgError string

func (e msgError) Error() string {
	return string(e)
}

const (
	errMsgUnsupported msgError = "unsupported message format"
	errMsgNoAltitude  msgError = "altitude not available"
	errMsgNoPosition  msgError = "position not available"
)

// Message is an ADS-B message
type Message struct {
	raw *RawMessage
}

// NewMessage wraps a RawMessage and returns the new Message.
func NewMessage(r *RawMessage) (*Message, error) {
	m := new(Message)
	m.raw = r

	err := m.validateRaw()
	if err != nil {
		return nil, err
	}
	return m, nil
}

// UnmarshalBinary implements the BinaryUnmarshaler interface, storing
// the supplied data in the Message.
//
// If an error is returned, but the data was successfully Unmarshalled
// into an underlying RawMessage, the Raw() method will still return the
// RawMessage for further inspection.
func (m *Message) UnmarshalBinary(data []byte) error {
	if m.raw == nil {
		m.raw = new(RawMessage)
	}
	err := m.raw.UnmarshalBinary(data)
	if err != nil {
		return err
	}

	return m.validateRaw()
}

// Validate that the downlink format is an expected value.
func (m *Message) validateRaw() error {
	switch m.df() {
	case 0, 4, 5, 11, 16, 17, 18, 20, 21:
		return nil
	default:
		return errMsgUnsupported
	}
}

// Simplify access to the DF value by converting errors to an unknown
// format that will be ignored.
func (m Message) df() int {
	df, err := m.raw.DF()
	if err != nil {
		return -1
	}
	return int(df)
}

// Raw returns the underlying RawMessage either explicitly passed via
// NewMessage or implicitly created via UnmarshalBinary. The returned
// RawMessage will be overwritten by a subsequent call to
// UnmarsahalBinary.
func (m Message) Raw() *RawMessage {
	return m.raw
}

// ICAO returns the ICAO address as an unsigned integer. Since the ICAO
// address is often extracted from the parity field, additional
// validation against a list of known addresses may be warranted.
func (m Message) ICAO() uint64 {
	aa, err := m.raw.AA()
	if err != nil {
		ap, err := m.raw.AP()
		if err != nil {
			return 0
		}
		return ap ^ m.raw.Parity()
	}
	return aa
}

// Alt returns the altitude. Returns error if the altitude cannot be
// obtained.
func (m Message) Alt() (int64, error) {
	switch m.df() {
	case 0, 4, 16, 20:
		ac, err := m.raw.AC()
		if err != nil {
			return 0, err
		}
		return decodeAlt13(uint16(ac))
	case 17, 18:
		alt, err := m.raw.ESAltitude()
		if err != nil {
			return 0, err
		}
		return decodeAlt12(uint16(alt))
	default:
		return 0, errMsgNoAltitude
	}
}

// Call returns the callsign, or an empty string if the callsign is
// unknown or unavailable.
func (m Message) Call() string {
	switch m.df() {
	case 17, 18:
		tc, _ := m.raw.ESType()
		if tc < 1 || tc > 4 {
			return ""
		}
	case 20, 21:
		if m.raw.Bits(33, 40) != 0x20 {
			return ""
		}
	default:
		return ""
	}

	bits := m.raw.Bits(41, 88)
	chars := []byte("?ABCDEFGHIJKLMNOPQRSTUVWXYZ????? ???????????????0123456789??????")

	call := make([]byte, 8)

	var i uint
	for i = 0; i < 8; i++ {
		call[i] = chars[(bits>>(42-(i*6)))&0x3F]
	}

	return string(bytes.TrimRight(call, " "))
}

// Sqk returns the squawk code, or an empty slice if the squawk code is
// unknown or unavailable.
func (m Message) Sqk() []byte {
	sqk := make([]byte, 0, 4)

	switch m.df() {
	case 5, 21:
	default:
		return sqk
	}

	f := [][]int{
		{25, 23, 21},
		{31, 29, 27},
		{24, 22, 20},
		{32, 30, 28},
	}

	sqk = sqk[0:4]

	for i, v := range f {
		for _, x := range v {
			sqk[i] <<= 1
			sqk[i] |= m.raw.Bit(x)
		}
	}

	return sqk
}

// CPR returns the compact position report. Returns error if the
// altitude cannot be obtained.
func (m Message) CPR() (*CPR, error) {
	switch m.df() {
	case 17, 18:
		tc, _ := m.raw.ESType()
		if tc < 9 || tc > 18 {
			return nil, errMsgNoPosition
		}
	default:
		return nil, errMsgNoPosition
	}

	c := new(CPR)
	c.Nb = 17
	c.T = m.raw.Bit(53)
	c.F = m.raw.Bit(54)
	c.Lat = uint32(m.raw.Bits(55, 71))
	c.Lon = uint32(m.raw.Bits(72, 88))

	return c, nil
}
