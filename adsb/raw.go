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

import "errors"

// RawMessage is a raw binary ADS-B message with helper methods for
// unmarshaling and retrieving arbitrary bit sequences.
type RawMessage []byte

// UnmarshalBinary implements the BinaryUnmarshaler interface, storing
// the supplied data in the RawMessage.
func (r *RawMessage) UnmarshalBinary(data []byte) error {
	if r == nil {
		return errors.New("can't unmarshal to nil pointer")
	}
	if len(data) != 7 && len(data) != 14 {
		return errors.New("incorrect data length")
	}
	*r = append((*r)[0:0], data...)
	return nil
}

// Bit returns the n-th bit of the RawMessage, where the first bit is
// numbered 1. Bit will panic if n is zero or beyond the end of the
// message.
func (r RawMessage) Bit(n int) uint8 {
	if n <= 0 {
		panic("bit must be greater than 0")
	}
	if n > len(r)*8 {
		panic("bit must be within message length")
	}

	n--

	return (r[n/8] >> (7 - uint(n%8))) & 0x01
}

// Bits64 returns bits n through z of the RawMessage, where the first
// bit is numbered 1. Bits64 will panic if n is not less than z, if n is
// zero, if z is beyond the end of the message, or if the result is
// greater than 64 bits.
func (r RawMessage) Bits64(n int, z int) uint64 {
	if n >= z {
		panic("upper bound must be greater than lower bound")
	}
	if (z - n) > 64 {
		panic("maximum of 64 bits exceeded")
	}

	var b uint64

	for i := n; i <= z; i++ {
		b <<= 1
		b |= uint64(r.Bit(i))
	}

	return b
}

// Bits32 returns bits n through z of the RawMessage, where the first
// bit is numbered 1. Bits32 will panic if n is not less than z, if n is
// zero, if z is beyond the end of the message, or if the result is
// greater than 32 bits.
func (r RawMessage) Bits32(n int, z int) uint32 {
	if n >= z {
		panic("upper bound must be greater than lower bound")
	}
	if (z - n) > 32 {
		panic("maximum of 32 bits exceeded")
	}

	var b uint32

	for i := n; i <= z; i++ {
		b <<= 1
		b |= uint32(r.Bit(i))
	}

	return b
}

// Bits16 returns bits n through z of the RawMessage, where the first
// bit is numbered 1. Bits16 will panic if n is not less than z, if n is
// zero, if z is beyond the end of the message, or if the result is
// greater than 16 bits.
func (r RawMessage) Bits16(n int, z int) uint16 {
	if n >= z {
		panic("upper bound must be greater than lower bound")
	}
	if (z - n) > 16 {
		panic("maximum of 16 bits exceeded")
	}

	var b uint16

	for i := n; i <= z; i++ {
		b <<= 1
		b |= uint16(r.Bit(i))
	}

	return b
}

// Bits8 returns bits n through z of the RawMessage, where the first bit
// is numbered 1. Bits8 will panic if n is not less than z, if n is
// zero, if z is beyond the end of the message, or if the result is
// greater than 8 bits.
func (r RawMessage) Bits8(n int, z int) uint8 {
	if n >= z {
		panic("upper bound must be greater than lower bound")
	}
	if (z - n) > 8 {
		panic("maximum of 8 bits exceeded")
	}

	var b uint8

	for i := n; i <= z; i++ {
		b <<= 1
		b |= r.Bit(i)
	}

	return b
}
