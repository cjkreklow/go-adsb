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
)

type fieldError string

func (e fieldError) Error() string {
	return string(e)
}

const (
	errNotLoaded    fieldError = "data not loaded"
	errNotAvailable fieldError = "field not available"
)

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

// AA returns the Address Announced field.
func (r RawMessage) AA() ([]byte, error) {
	df, err := r.DF()
	if err != nil {
		return nil, err
	}
	switch df[0] {
	case 11, 17, 18:
		return r.bytes(9, 32), nil
	default:
		return nil, errNotAvailable
	}
}

// AC returns the Altitude Code field.
func (r RawMessage) AC() ([]byte, error) {
	df, err := r.DF()
	if err != nil {
		return nil, err
	}
	switch df[0] {
	case 0, 4, 16, 20:
		return r.bytes(20, 32), nil
	default:
		return nil, errNotAvailable
	}
}

// AF returns the Application Field.
func (r RawMessage) AF() ([]byte, error) {
	df, err := r.DF()
	if err != nil {
		return nil, err
	}
	switch df[0] {
	case 19:
		return r.bytes(6, 8), nil
	default:
		return nil, errNotAvailable
	}
}

// AP returns the Address / Parity field.
func (r RawMessage) AP() ([]byte, error) {
	df, err := r.DF()
	if err != nil {
		return nil, err
	}
	switch df[0] {
	case 0, 4, 5:
		return r.bytes(33, 56), nil
	case 16, 20, 21, 24:
		return r.bytes(89, 112), nil
	default:
		return nil, errNotAvailable
	}
}

// CA returns the Capability field.
func (r RawMessage) CA() ([]byte, error) {
	df, err := r.DF()
	if err != nil {
		return nil, err
	}
	switch df[0] {
	case 11, 17:
		return r.bytes(6, 8), nil
	default:
		return nil, errNotAvailable
	}
}

// CC returns the Cross-link Capability field.
func (r RawMessage) CC() ([]byte, error) {
	df, err := r.DF()
	if err != nil {
		return nil, err
	}
	switch df[0] {
	case 0:
		return r.bytes(7, 7), nil
	default:
		return nil, errNotAvailable
	}
}

// CF returns the Control Field.
func (r RawMessage) CF() ([]byte, error) {
	df, err := r.DF()
	if err != nil {
		return nil, err
	}
	switch df[0] {
	case 18:
		return r.bytes(6, 8), nil
	default:
		return nil, errNotAvailable
	}
}

// DF returns the Downlink Format field.
func (r RawMessage) DF() ([]byte, error) {
	if len(r) == 0 {
		return nil, errNotLoaded
	}
	b := r.bytes(1, 5)
	if b[0] > 24 {
		b[0] = 24
	}
	return b, nil
}

// DP returns the Data Parity field.
func (r RawMessage) DP() ([]byte, error) {
	df, err := r.DF()
	if err != nil {
		return nil, err
	}
	switch df[0] {
	case 20, 21:
		return r.bytes(89, 112), nil
	default:
		return nil, errNotAvailable
	}
}

// DR returns the Downlink Request field.
func (r RawMessage) DR() ([]byte, error) {
	df, err := r.DF()
	if err != nil {
		return nil, err
	}
	switch df[0] {
	case 4, 5, 20, 21:
		return r.bytes(9, 13), nil
	default:
		return nil, errNotAvailable
	}
}

// FS returns the Flight Status field.
func (r RawMessage) FS() ([]byte, error) {
	df, err := r.DF()
	if err != nil {
		return nil, err
	}
	switch df[0] {
	case 4, 5, 20, 21:
		return r.bytes(6, 8), nil
	default:
		return nil, errNotAvailable
	}
}

// ID returns the Identity field.
func (r RawMessage) ID() ([]byte, error) {
	df, err := r.DF()
	if err != nil {
		return nil, err
	}
	switch df[0] {
	case 5, 21:
		return r.bytes(20, 32), nil
	default:
		return nil, errNotAvailable
	}
}

// KE returns the ELM Control field.
func (r RawMessage) KE() ([]byte, error) {
	df, err := r.DF()
	if err != nil {
		return nil, err
	}
	switch df[0] {
	case 24:
		return r.bytes(4, 4), nil
	default:
		return nil, errNotAvailable
	}
}

// MB returns the Comm-B Message field.
func (r RawMessage) MB() ([]byte, error) {
	df, err := r.DF()
	if err != nil {
		return nil, err
	}
	switch df[0] {
	case 20, 21:
		return r.bytes(33, 88), nil
	default:
		return nil, errNotAvailable
	}
}

// MD returns the Comm-D Message field.
func (r RawMessage) MD() ([]byte, error) {
	df, err := r.DF()
	if err != nil {
		return nil, err
	}
	switch df[0] {
	case 24:
		return r.bytes(9, 88), nil
	default:
		return nil, errNotAvailable
	}
}

// ME returns the Extended Squitter Message field.
func (r RawMessage) ME() ([]byte, error) {
	df, err := r.DF()
	if err != nil {
		return nil, err
	}
	switch df[0] {
	case 17, 18:
		return r.bytes(33, 88), nil
	default:
		return nil, errNotAvailable
	}
}

// MV returns the ACAS Message field.
func (r RawMessage) MV() ([]byte, error) {
	df, err := r.DF()
	if err != nil {
		return nil, err
	}
	switch df[0] {
	case 16:
		return r.bytes(33, 88), nil
	default:
		return nil, errNotAvailable
	}
}

// ND returns the Number of D-segment field.
func (r RawMessage) ND() ([]byte, error) {
	df, err := r.DF()
	if err != nil {
		return nil, err
	}
	switch df[0] {
	case 24:
		return r.bytes(5, 8), nil
	default:
		return nil, errNotAvailable
	}
}

// PI returns the Parity / Interrogator Identifier field.
func (r RawMessage) PI() ([]byte, error) {
	df, err := r.DF()
	if err != nil {
		return nil, err
	}
	switch df[0] {
	case 11:
		return r.bytes(33, 56), nil
	case 17, 18:
		return r.bytes(89, 112), nil
	default:
		return nil, errNotAvailable
	}
}

// RI returns the Reply Information field.
func (r RawMessage) RI() ([]byte, error) {
	df, err := r.DF()
	if err != nil {
		return nil, err
	}
	switch df[0] {
	case 0, 16:
		return r.bytes(14, 17), nil
	default:
		return nil, errNotAvailable
	}
}

// SL returns the Sensitivity Level field.
func (r RawMessage) SL() ([]byte, error) {
	df, err := r.DF()
	if err != nil {
		return nil, err
	}
	switch df[0] {
	case 0, 16:
		return r.bytes(9, 11), nil
	default:
		return nil, errNotAvailable
	}
}

// UM returns the Utility Message field.
func (r RawMessage) UM() ([]byte, error) {
	df, err := r.DF()
	if err != nil {
		return nil, err
	}
	switch df[0] {
	case 4, 5, 20, 21:
		return r.bytes(14, 19), nil
	default:
		return nil, errNotAvailable
	}
}

// VS returns the Vertical Status field.
func (r RawMessage) VS() ([]byte, error) {
	df, err := r.DF()
	if err != nil {
		return nil, err
	}
	switch df[0] {
	case 0, 16:
		return r.bytes(6, 6), nil
	default:
		return nil, errNotAvailable
	}
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

func (r RawMessage) bytes(n int, z int) []byte {
	if z < n {
		panic("upper bound must not be less than lower bound")
	}

	bytes := make([]byte, ((z-n)/8)+1)

	var bits uint8
	var j int

	for i := n; i <= z; i++ {
		bits <<= 1
		bits |= r.Bit(i)
		if (z-i)%8 == 0 {
			bytes[j] = bits
			j++
			bits = 0
		}
	}

	return bytes
}
