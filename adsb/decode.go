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
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
)

/*
56 bits
           1          2          3          4           5
 12345678 90123456 78901234 56789012 34567890 12345678 90123456
0        1        2        3        4        5        6        7

 12345678
 84218421
*/

// decode DF11 all call reply
func (m *Message) decode11() error {
	m.CA = CA(int(m.raw[0]) & 0x07) // bits 6-8
	m.ICAO = hex.EncodeToString(m.raw[1:4])

	return nil
}

// decode DF4 and DF20 altitude reply
func (m *Message) decodeAltMsg() error {
	m.FS = FS(int(m.raw[0]) & 0x07) // bits 6-8
	m.setICAOFromAP()

	a, err := decodeAlt13(binary.BigEndian.Uint16(m.raw[2:4]) & 0x1FFF)
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
	m.FS = FS(int(m.raw[0]) & 0x07) // bits 6-8
	m.setICAOFromAP()

	i := binary.BigEndian.Uint16(m.raw[2:4]) & 0x1FFF

	oct := make([]uint8, 4)

	oct[0] = uint8(((i & 0x80) >> 5) | ((i & 0x0200) >> 8) | ((i & 0x0800) >> 11))   // A4 A2 A1
	oct[1] = uint8(((i & 0x02) << 1) | ((i & 0x08) >> 2) | ((i & 0x20) >> 5))        // B4 B2 B1
	oct[2] = uint8(((i & 0x0100) >> 6) | ((i & 0x0400) >> 9) | ((i & 0x1000) >> 12)) // C4 C2 C1
	oct[3] = uint8(((i & 0x01) << 2) | ((i & 0x04) >> 1) | ((i & 0x10) >> 4))        // A4 A2 A1

	m.Sqk = fmt.Sprintf("%o%o%o%o", oct[0], oct[1], oct[2], oct[3])

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
	m.CA = CA(int(m.raw[0]) & 0x07) // bits 6-8
	m.ICAO = hex.EncodeToString(m.raw[1:4])

	m.TC = TC(m.raw[4] >> 3)

	if m.TC >= 1 && m.TC <= 4 {
		err := m.setCall(m.raw[5:11])
		if err != nil {
			return err
		}
	}

	if m.TC >= 9 && m.TC <= 18 {
		a, err := decodeAlt12(binary.BigEndian.Uint16(m.raw[5:7]) >> 4)
		if err != nil {
			return err
		}

		m.Alt = a
	}

	return nil
}

// utility function to set ICAO when XORed into an AP field
func (m *Message) setICAOFromAP() {
	b := binary.BigEndian.Uint32(m.raw[len(m.raw)-4:len(m.raw)]) & 0x00FFFFFF
	m.ICAO = fmt.Sprintf("%06x", b^m.parity)
}

// utility function to set Call from a data field
func (m *Message) setCall(b []byte) error {
	if len(b) != 6 {
		return errors.New("incorrect data length")
	}

	c := []byte("?ABCDEFGHIJKLMNOPQRSTUVWXYZ????? ???????????????0123456789??????")

	a := binary.BigEndian.Uint64(append([]byte{0x00, 0x00}, b...))

	call := make([]byte, 8)

	var i uint
	for i = 0; i < 8; i++ {
		call[i] = c[(a>>(42-(i*6)))&0x3F]
	}

	m.Call = string(bytes.TrimRight(call, " "))

	return nil
}
