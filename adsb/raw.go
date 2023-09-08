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

package adsb

import (
	"bytes"
)

// RawMessage is a raw binary ADS-B message with helper methods for
// retrieving message fields and arbitrary bit sequences.
type RawMessage struct {
	data bytes.Buffer
}

// UnmarshalBinary implements the BinaryUnmarshaler interface for
// storing ADS-B data.
func (r *RawMessage) UnmarshalBinary(data []byte) error {
	r.data.Reset()
	r.data.Write(data)

	df, err := r.DF()
	if err != nil {
		return err
	}

	switch df {
	case 0, 4, 5, 11:
		if len(data) != 7 {
			return newErrorf(nil, "incorrect data length: %d bits with format %d", len(data)*8, df)
		}
	case 16, 17, 18, 19, 20, 21, 24:
		if len(data) != 14 {
			return newErrorf(nil, "incorrect data length: %d bits with format %d", len(data)*8, df)
		}
	default:
		return newErrorf(nil, "unknown downlink format: %d bits with format %d", len(data)*8, df)
	}

	return nil
}

// AA returns the Address Announced field.
func (r *RawMessage) AA() (uint64, error) {
	df, err := r.DF()
	if err != nil {
		return 0, err
	}

	switch df {
	case 11, 17, 18:
		return r.Bits(9, 32), nil
	default:
		return 0, newErrorf(ErrNotAvailable, "error retrieving %s from %d",
			"AA", df)
	}
}

// AC returns the Altitude Code field.
func (r *RawMessage) AC() (uint64, error) {
	df, err := r.DF()
	if err != nil {
		return 0, err
	}

	switch df {
	case 0, 4, 16, 20:
		return r.Bits(20, 32), nil
	default:
		return 0, newErrorf(ErrNotAvailable, "error retrieving %s from %d",
			"AC", df)
	}
}

// AF returns the Application Field.
func (r *RawMessage) AF() (uint64, error) {
	df, err := r.DF()
	if err != nil {
		return 0, err
	}

	switch df {
	case 19:
		return r.Bits(6, 8), nil
	default:
		return 0, newErrorf(ErrNotAvailable, "error retrieving %s from %d",
			"AF", df)
	}
}

// AP returns the Address / Parity field.
func (r *RawMessage) AP() (uint64, error) {
	df, err := r.DF()
	if err != nil {
		return 0, err
	}

	switch df {
	case 0, 4, 5:
		return r.Bits(33, 56), nil
	case 16, 20, 21, 24:
		return r.Bits(89, 112), nil
	default:
		return 0, newErrorf(ErrNotAvailable, "error retrieving %s from %d",
			"AP", df)
	}
}

// CA returns the Capability field.
func (r *RawMessage) CA() (uint64, error) {
	df, err := r.DF()
	if err != nil {
		return 0, err
	}

	switch df {
	case 11, 17:
		return r.Bits(6, 8), nil
	default:
		return 0, newErrorf(ErrNotAvailable, "error retrieving %s from %d",
			"CA", df)
	}
}

// CC returns the Cross-link Capability field.
func (r *RawMessage) CC() (uint64, error) {
	df, err := r.DF()
	if err != nil {
		return 0, err
	}

	switch df {
	case 0:
		return r.Bits(7, 7), nil
	default:
		return 0, newErrorf(ErrNotAvailable, "error retrieving %s from %d",
			"CC", df)
	}
}

// CF returns the Control Field.
func (r *RawMessage) CF() (uint64, error) {
	df, err := r.DF()
	if err != nil {
		return 0, err
	}

	switch df {
	case 18:
		return r.Bits(6, 8), nil
	default:
		return 0, newErrorf(ErrNotAvailable, "error retrieving %s from %d",
			"CF", df)
	}
}

// DF returns the Downlink Format field.
func (r *RawMessage) DF() (uint64, error) {
	if r.data.Len() == 0 {
		return 0, newError(nil, "no data loaded")
	}

	b := r.Bits(1, 5)
	if b > 24 {
		b = 24
	}

	return b, nil
}

func (r *RawMessage) TC() uint64 {
	return r.Bits(33, 37)
}

// DP returns the Data Parity field.
func (r *RawMessage) DP() (uint64, error) {
	df, err := r.DF()
	if err != nil {
		return 0, err
	}

	switch df {
	case 20, 21:
		return r.Bits(89, 112), nil
	default:
		return 0, newErrorf(ErrNotAvailable, "error retrieving %s from %d",
			"DP", df)
	}
}

// DR returns the Downlink Request field.
func (r *RawMessage) DR() (uint64, error) {
	df, err := r.DF()
	if err != nil {
		return 0, err
	}

	switch df {
	case 4, 5, 20, 21:
		return r.Bits(9, 13), nil
	default:
		return 0, newErrorf(ErrNotAvailable, "error retrieving %s from %d",
			"DR", df)
	}
}

// FS returns the Flight Status field.
func (r *RawMessage) FS() (uint64, error) {
	df, err := r.DF()
	if err != nil {
		return 0, err
	}

	switch df {
	case 4, 5, 20, 21:
		return r.Bits(6, 8), nil
	default:
		return 0, newErrorf(ErrNotAvailable, "error retrieving %s from %d",
			"FS", df)
	}
}

// ID returns the Identity field.
func (r *RawMessage) ID() (uint64, error) {
	df, err := r.DF()
	if err != nil {
		return 0, err
	}

	switch df {
	case 5, 21:
		return r.Bits(20, 32), nil
	default:
		return 0, newErrorf(ErrNotAvailable, "error retrieving %s from %d",
			"ID", df)
	}
}

// KE returns the ELM Control field.
func (r *RawMessage) KE() (uint64, error) {
	df, err := r.DF()
	if err != nil {
		return 0, err
	}

	switch df {
	case 24:
		return r.Bits(4, 4), nil
	default:
		return 0, newErrorf(ErrNotAvailable, "error retrieving %s from %d",
			"KE", df)
	}
}

// MB returns the Comm-B Message field.
func (r *RawMessage) MB() (uint64, error) {
	df, err := r.DF()
	if err != nil {
		return 0, err
	}

	switch df {
	case 20, 21:
		return r.Bits(33, 88), nil
	default:
		return 0, newErrorf(ErrNotAvailable, "error retrieving %s from %d",
			"MB", df)
	}
}

// MD returns the Comm-D Message field.
func (r *RawMessage) MD() ([]byte, error) {
	df, err := r.DF()
	if err != nil {
		return nil, err
	}

	switch df {
	case 24:
		return r.bytes(9, 88), nil
	default:
		return nil, newErrorf(ErrNotAvailable, "error retrieving %s from %d",
			"MD", df)
	}
}

// ME returns the Extended Squitter Message field.
func (r *RawMessage) ME() (uint64, error) {
	df, err := r.DF()
	if err != nil {
		return 0, err
	}

	switch df {
	case 17, 18:
		return r.Bits(33, 88), nil
	default:
		return 0, newErrorf(ErrNotAvailable, "error retrieving %s from %d",
			"ME", df)
	}
}

// MV returns the ACAS Message field.
func (r *RawMessage) MV() (uint64, error) {
	df, err := r.DF()
	if err != nil {
		return 0, err
	}

	switch df {
	case 16:
		return r.Bits(33, 88), nil
	default:
		return 0, newErrorf(ErrNotAvailable, "error retrieving %s from %d",
			"MV", df)
	}
}

// ND returns the Number of D-segment field.
func (r *RawMessage) ND() (uint64, error) {
	df, err := r.DF()
	if err != nil {
		return 0, err
	}

	switch df {
	case 24:
		return r.Bits(5, 8), nil
	default:
		return 0, newErrorf(ErrNotAvailable, "error retrieving %s from %d",
			"ND", df)
	}
}

// PI returns the Parity / Interrogator Identifier field.
func (r *RawMessage) PI() (uint64, error) {
	df, err := r.DF()
	if err != nil {
		return 0, err
	}

	switch df {
	case 11:
		return r.Bits(33, 56), nil
	case 17, 18:
		return r.Bits(89, 112), nil
	default:
		return 0, newErrorf(ErrNotAvailable, "error retrieving %s from %d",
			"PI", df)
	}
}

// RI returns the Reply Information field.
func (r *RawMessage) RI() (uint64, error) {
	df, err := r.DF()
	if err != nil {
		return 0, err
	}

	switch df {
	case 0, 16:
		return r.Bits(14, 17), nil
	default:
		return 0, newErrorf(ErrNotAvailable, "error retrieving %s from %d",
			"RI", df)
	}
}

// SL returns the Sensitivity Level field.
func (r *RawMessage) SL() (uint64, error) {
	df, err := r.DF()
	if err != nil {
		return 0, err
	}

	switch df {
	case 0, 16:
		return r.Bits(9, 11), nil
	default:
		return 0, newErrorf(ErrNotAvailable, "error retrieving %s from %d",
			"SL", df)
	}
}

// UM returns the Utility Message field.
func (r *RawMessage) UM() (uint64, error) {
	df, err := r.DF()
	if err != nil {
		return 0, err
	}

	switch df {
	case 4, 5, 20, 21:
		return r.Bits(14, 19), nil
	default:
		return 0, newErrorf(ErrNotAvailable, "error retrieving %s from %d",
			"UM", df)
	}
}

// VS returns the Vertical Status field.
func (r *RawMessage) VS() (uint64, error) {
	df, err := r.DF()
	if err != nil {
		return 0, err
	}

	switch df {
	case 0, 16:
		return r.Bits(6, 6), nil
	default:
		return 0, newErrorf(ErrNotAvailable, "error retrieving %s from %d",
			"VS", df)
	}
}

// Bit returns the n-th bit of the RawMessage, where the first bit is
// numbered 1. Bit will panic if n is out of range.
func (r *RawMessage) Bit(n int) uint8 {
	switch {
	case n <= 0:
		panic("bit must be greater than 0")
	case n > r.data.Len()*8:
		panic("bit must be within message length")
	}

	n--

	return (r.data.Bytes()[n/8] >> (7 - (n % 8))) & 0x01
}

// Bits returns bits n through z of the RawMessage, where the first bit
// is numbered 1. Bits will panic if n or z are out of range, or if the
// result is greater than 64 bits.
func (r *RawMessage) Bits(n int, z int) (bits uint64) {
	switch {
	case n <= 0:
		panic("lower bound must be greater than 0")
	case z > r.data.Len()*8:
		panic("upper bound must be within message length")
	case n > z:
		panic("upper bound must be greater than lower bound")
	case (z - n) > 64:
		panic("maximum of 64 bits exceeded")
	}

	n--
	z--

	nshift := n % 8
	zshift := 7 - (z % 8)
	bshift := 0

	for i := z / 8; i >= n/8; i-- {
		var b uint8
		if i == n/8 {
			b = (r.data.Bytes()[i] << nshift) >> nshift
		} else {
			b = r.data.Bytes()[i]
		}

		bits |= (uint64(b) << bshift) >> zshift
		bshift += 8
	}

	return bits
}

var pTbl = []uint64{
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

// Parity returns the calculated parity for the message data.
func (r *RawMessage) Parity() (p uint64) {
	var length, offset int

	switch r.data.Len() {
	case 7:
		length = 32
		offset = 56
	case 14:
		length = 88
		offset = 0
	default:
		return 0
	}

	for i := 1; i <= length; i++ {
		if r.Bit(i) != 0 {
			p ^= pTbl[i+offset-1]
		}
	}

	return p
}

func (r *RawMessage) bytes(n int, z int) []byte {
	bytes := make([]byte, ((z-n)/8)+1)

	var bits uint8

	var j int

	for i := n; i <= z; i++ {
		bits <<= 1
		bits |= r.Bit(i)

		if (z-i)%8 == 0 {
			bytes[j] = bits
			bits = 0
			j++
		}
	}

	return bytes
}
