// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Coverters for custom characters from the LCD 20x4 Bricklet.

For a simple handling of custom characters.
The characters have 5x8 pixel.

For a simple approach, a array of 8 strings could convert to a lcd custom character.
Every element of the array of strings should not have more than 5 character (runes) inside.
A SPACE(" ") oder POINT(".") represent a zero, every other character in the string is a one.

Example to generate a "H" custom character:
  lines := [8]string{
    "O...O",
    "O...O",
    "O...O",
    "OOOOO",
    "O...O",
    "O...O",
    "O...O",
    "....." }
  cc := ConvertStringToCharacter(lines)
*/
package lcdcharacter

import (
	"github.com/dirkjabl/bricker/device/bricklet/lcd20x4"
	"unicode/utf8"
)

// ConvertStringToCharacter converts the strings to the custom character representation.
func ConvertStringToCharacter(lines [8]string) *lcd20x4.Character {
	cc := new(lcd20x4.Character)
	for i, line := range lines {
		cc[i] = convertstringline(line)
	}
	return cc
}

// Internal function: convertstringline converts a single line to the resulting uint8 value.
func convertstringline(line string) uint8 {
	var i int
	var c uint8
	var e uint8 = 16
	text := []byte(line)
	for len(text) > 0 && i < 5 { // only 5 characters possible
		r, size := utf8.DecodeRune(text)
		text = text[size:]
		if !(r == ' ' || r == '.') { // not by spaces and points
			c += e
		}
		e = e >> 1
		i++
	}
	return uint8(c)
}
