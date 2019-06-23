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

// Message is an ADS-B message
type Message struct {
	raw    *RawMessage
	parity uint32
	df     DF
	icao   string
}

// NewMessage wraps a RawMessage and returns the new Message.
func NewMessage(r *RawMessage) (*Message, error) {
	m := new(Message)
	m.raw = r

	err := m.processRaw()
	if err != nil {
		return nil, err
	}
	return m, nil
}

// UnmarshalBinary implements the BinaryUnmarshaler interface, storing
// the supplied data in the Message.
func (m *Message) UnmarshalBinary(data []byte) error {
	if m == nil {
		return errors.New("can't unmarshal to nil pointer")
	}

	r := new(RawMessage)
	err := r.UnmarshalBinary(data)
	if err != nil {
		return err
	}

	m.raw = r

	return m.processRaw()
}

func (m *Message) processRaw() error {
	m.df = DF(m.raw.Bits8(1, 5))
	m.setParity()

	switch m.df {
	case 0, 4, 5, 16, 20, 21:
		m.setICAOFromAP()
	case 11, 17, 18:
		m.setICAOFromAA()
	default:
		return fmt.Errorf("unsupported format: %d - %s", int(m.df), m.df)
	}

	return nil
}

// DF returns the downlink format.
func (m Message) DF() DF {
	return m.df
}

// CA returns the capability, or -1 if the format does not support the
// field.
func (m Message) CA() CA {
	switch m.df {
	case 11, 17:
		return CA(m.raw.Bits8(6, 8))
	default:
		return CA(-1)
	}
}

// FS returns the flight status, or -1 if the format does not support
// the field.
func (m Message) FS() FS {
	switch m.df {
	case 4, 5, 20, 21:
		return FS(m.raw.Bits8(6, 8))
	default:
		return FS(-1)
	}
}

// VS returns the vertical status, or -1 if the format does not support
// the field.
func (m Message) VS() VS {
	switch m.df {
	case 0, 16:
		return VS(m.raw.Bit(6))
	default:
		return VS(-1)
	}
}

// TC returns the extended squitter format type code, or -1 if the
// format does not support the field.
func (m Message) TC() TC {
	switch m.df {
	case 17, 18:
		return TC(m.raw.Bits8(33, 37))
	default:
		return TC(-1)
	}
}

// AcCat returns the extended squitter aircraft emitter category, or -1
// if the format does not support the field.
func (m Message) AcCat() AcCat {
	switch m.df {
	case 17, 18:
		switch m.TC() {
		case 1:
			return AcCat(fmt.Sprintf("D%d", m.raw.Bits8(38, 40)))
		case 2:
			return AcCat(fmt.Sprintf("C%d", m.raw.Bits8(38, 40)))
		case 3:
			return AcCat(fmt.Sprintf("B%d", m.raw.Bits8(38, 40)))
		case 4:
			return AcCat(fmt.Sprintf("A%d", m.raw.Bits8(38, 40)))
		default:
			return ""
		}
	default:
		return ""
	}
}

// SS returns the extended squitter surveillance status, or -1 if the
// format does not support the field.
func (m Message) SS() SS {
	switch m.df {
	case 17, 18:
		if m.TC() < 9 || m.TC() > 18 {
			return SS(-1)
		}
		return SS(m.raw.Bits8(38, 39))
	default:
		return SS(-1)
	}
}

// ICAO returns the ICAO address.
func (m Message) ICAO() string {
	return m.icao
}

// Alt returns the altitude. Returns error if the altitude cannot be
// obtained.
func (m Message) Alt() (int64, error) {
	switch m.df {
	case 0, 4, 16, 20:
		return decodeAlt13(m.raw.Bits16(20, 32))
	case 17, 18:
		if m.TC() < 9 || m.TC() > 18 {
			return 0, errors.New("altitude not available")
		}
		return decodeAlt12(m.raw.Bits16(41, 52))
	default:
		return 0, errors.New("altitude not available")
	}
}

// Call returns the callsign, or an empty string if the callsign is
// unknown or unavailable.
func (m Message) Call() string {
	switch m.df {
	case 17, 18:
		if m.TC() < 1 || m.TC() > 4 {
			return ""
		}
	case 20, 21:
		if m.raw.Bits8(33, 40) != 0x20 {
			return ""
		}
	default:
		return ""
	}

	b := m.raw.Bits64(41, 88)
	c := []byte("?ABCDEFGHIJKLMNOPQRSTUVWXYZ????? ???????????????0123456789??????")

	call := make([]byte, 8)

	var i uint
	for i = 0; i < 8; i++ {
		call[i] = c[(b>>(42-(i*6)))&0x3F]
	}

	return string(bytes.TrimRight(call, " "))
}

// Sqk returns the squawk code, or an empty string if the squawk code is
// unknown or unavailable.
func (m Message) Sqk() string {
	switch m.df {
	case 5, 21:
	default:
		return ""
	}

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

	return strconv.FormatUint(id, 8)
}

// CPR returns the compact position report. Returns error if the
// altitude cannot be obtained.
func (m Message) CPR() (*CPR, error) {
	switch m.df {
	case 17, 18:
		if m.TC() < 9 || m.TC() > 18 {
			return nil, errors.New("position not available")
		}
	default:
		return nil, errors.New("position not available")
	}

	c := new(CPR)
	c.Nb = 17
	c.T = m.raw.Bit(53)
	c.F = m.raw.Bit(54)
	c.Lat = m.raw.Bits32(55, 71)
	c.Lon = m.raw.Bits32(72, 88)

	return c, nil
}

// set the ICAO address from the AA field
func (m *Message) setICAOFromAA() {
	m.icao = strconv.FormatUint(m.raw.Bits64(9, 32), 16)
}

// set the ICAO address from the AP field by XORing the calculated
// parity
func (m *Message) setICAOFromAP() {
	var b uint32
	if int(m.df) < 16 {
		b = m.raw.Bits32(33, 56)
	} else {
		b = m.raw.Bits32(89, 112)
	}
	m.icao = strconv.FormatUint(uint64(b^m.parity), 16)
}

// calculate the message parity
func (m *Message) setParity() {
	pt := []uint32{
		0x3935ea, 0x1c9af5, 0xf1b77e, 0x78dbbf,
		0xc397db, 0x9e31e9, 0xb0e2f0, 0x587178,
		0x2c38bc, 0x161c5e, 0x0b0e2f, 0xfa7d13,
		0x82c48d, 0xbe9842, 0x5f4c21, 0xd05c14,
		0x682e0a, 0x341705, 0xe5f186, 0x72f8c3,
		0xc68665, 0x9cb936, 0x4e5c9b, 0xd8d449,
		0x939020, 0x49c810, 0x24e408, 0x127204,
		0x093902, 0x049c81, 0xfdb444, 0x7eda22,
		0x3f6d11, 0xe04c8c, 0x702646, 0x381323,
		0xe3f395, 0x8e03ce, 0x4701e7, 0xdc7af7,
		0x91c77f, 0xb719bb, 0xa476d9, 0xadc168,
		0x56e0b4, 0x2b705a, 0x15b82d, 0xf52612,
		0x7a9309, 0xc2b380, 0x6159c0, 0x30ace0,
		0x185670, 0x0c2b38, 0x06159c, 0x030ace,
		0x018567, 0xff38b7, 0x80665f, 0xbfc92b,
		0xa01e91, 0xaff54c, 0x57faa6, 0x2bfd53,
		0xea04ad, 0x8af852, 0x457c29, 0xdd4410,
		0x6ea208, 0x375104, 0x1ba882, 0x0dd441,
		0xf91024, 0x7c8812, 0x3e4409, 0xe0d800,
		0x706c00, 0x383600, 0x1c1b00, 0x0e0d80,
		0x0706c0, 0x038360, 0x01c1b0, 0x00e0d8,
		0x00706c, 0x003836, 0x001c1b, 0xfff409,
		0x000000, 0x000000, 0x000000, 0x000000,
		0x000000, 0x000000, 0x000000, 0x000000,
		0x000000, 0x000000, 0x000000, 0x000000,
		0x000000, 0x000000, 0x000000, 0x000000,
		0x000000, 0x000000, 0x000000, 0x000000,
		0x000000, 0x000000, 0x000000, 0x000000,
	}

	var l int
	var o int
	if m.df < 16 {
		l = 32
		o = 56
	} else {
		l = 88
		o = 0
	}

	for i := 1; i <= l; i++ {
		if m.raw.Bit(i) != 0 {
			m.parity ^= pt[i+o-1]
		}
	}
}
