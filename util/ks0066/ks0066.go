// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package for converting utf8 to ks0066-00 English-Japanese
package ks0066

// Map of runes to bytes
var ks0066ext = map[rune]byte{
	'ä':      0xe1,
	'Ä':      0xe1,
	'ß':      0xe2,
	'³':      0xe3,
	'ö':      0xef,
	'Ö':      0xef,
	'ü':      0xf5,
	'Ü':      0xf5,
	'\u03b1': 0xe0,
	'\u03b2': 0xe2,
	'\u03bc': 0xe4,
	'\u00b5': 0xe4,
	'\u00f1': 0xee,
	'\u2190': 0x7f,
	'\u2192': 0x7e,
	'\u00f7': 0xfd,
	'Ŧ':      0xfa,
	'\u03c0': 0xf7,
	'\u03a3': 0xf6,
	'\u2211': 0xf6,
	'\u2208': 0xe3,
	'\u220a': 0xe3,
	'¤':      0xeb,
	'°':      0xdf,
	'Ω':      0xf4,
	'~':      0xde,
	'·':      0xa5,
	'–':      0xb0,
	'—':      0xb0,
	'¢':      0xec,
	'∙':      0xa5,
	'\u0087': 0xa5,
	'\u221a': 0xe8,
	'\u220e': 0xff, // ∎
	'\u25a0': 0xff, // ■
	'\u220d': 0xae,
	'\u221d': 0xe0,
}

// ToByte converts a rune to a byte.
func ToByte(r rune) byte {
	if r == 65533 {
		return 0
	}
	if r >= ' ' && r <= '}' {
		return byte(r)
	}
	if v, ok := ks0066ext[r]; ok {
		return v
	}
	return byte(' ')
}
