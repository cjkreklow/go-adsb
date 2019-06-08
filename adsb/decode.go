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
	"encoding/binary"
	"encoding/hex"
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

// decode DF4 altitude reply
func (m *Message) decode4() error {
	m.FltStat = FS(int(m.raw[0]) & 0x07) // bits 6-8
	m.setICAOFromAP()

	a, err := alt(binary.BigEndian.Uint16(m.raw[2:4]))
	if err != nil {
		return err
	}

	m.Alt = a

	return nil
}

// decode DF11 all-call reply
func (m *Message) decode11() error {
	m.Cap = CA(int(m.raw[0]) & 0x07) // bits 6-8
	m.ICAO = hex.EncodeToString(m.raw[1:4])
	return nil
}

// utility function to set ICAO when XORed into an AP field
func (m *Message) setICAOFromAP() {
	b := binary.BigEndian.Uint32(m.raw[len(m.raw)-4:len(m.raw)]) & 0x00FFFFFF
	m.ICAO = fmt.Sprintf("%06x", b^m.parity)
}
