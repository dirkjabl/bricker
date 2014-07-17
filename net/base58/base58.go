// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Implementation of the base58 encoding used by Tinkerforge.
*/
package base58

const (
	Base58Alphabet = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
)

var Base58AlphabetMap = map[string]uint64{
	"1": 0, "2": 1, "3": 2, "4": 3, "5": 4, "6": 5, "7": 6, "8": 7, "9": 8, "a": 9,
	"b": 10, "c": 11, "d": 12, "e": 13, "f": 14, "g": 15, "h": 16, "i": 17, "j": 18, "k": 19,
	"m": 20, "n": 21, "o": 22, "p": 23, "q": 24, "r": 25, "s": 26, "t": 27, "u": 28, "v": 29,
	"w": 30, "x": 31, "y": 32, "z": 33, "A": 34, "B": 35, "C": 36, "D": 37, "E": 38, "F": 39,
	"G": 40, "H": 41, "J": 42, "K": 43, "L": 44, "M": 45, "N": 46, "P": 47, "Q": 48, "R": 49,
	"S": 50, "T": 51, "U": 52, "V": 53, "W": 54, "X": 55, "Y": 56, "Z": 57}

// Decode converts a 8 byte character string (old c style) base58 representation to a uint64 number.
func Decode(src [8]byte) uint64 {
	var (
		i           int
		value, base uint64
	)
	value = 0
	base = 1
	for i = 7; i >= 0; i-- {
		if src[i] == 0x0 {
			continue
		}
		value += Base58AlphabetMap[string(src[i])] * base
		base *= 58
	}
	return value
}

// Encode converts a uint64 number to a 8 byte base58 representation (c style string).
func Encode(src uint64) [8]byte {
	var (
		rdest []byte = make([]byte, 0, 8)
		dest  [8]byte
		i, l  int = 0, 0
	)
	for src >= 58 {
		rdest = append(rdest, byte(Base58Alphabet[src%58]))
		src /= 58
		l++
	}
	rdest = append(rdest, byte(Base58Alphabet[src]))
	l++
	if l > 8 {
		l = 8
	}
	for i = 0; i < l; i++ {
		dest[i] = rdest[l-1-i]
	}

	return dest
}

// Converts a uint64 Number to a uint32, but it is not a cast.
// It is special for the used user identifer numbers by bricks.
func Convert32(uid uint64) uint32 {
	var newuid uint32

	if uid > 0xFFFFFFFF {
		var value1, value2 uint32
		value1 = uint32(uid & 0xFFFFFFFF)
		value2 = uint32(uid>>32) & 0xFFFFFFFF
		newuid = value1 & 0x00000FFF
		newuid |= (value1 & 0x0F000000) >> 12
		newuid |= (value2 & 0x0000003F) << 16
		newuid |= (value2 & 0x000F0000) << 6
		newuid |= (value2 & 0x3F000000) << 2
	} else {
		newuid = uint32(uid & 0xFFFFFFFF)
	}

	return newuid
}
