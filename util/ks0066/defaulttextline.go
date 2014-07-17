// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ks0066

import (
	"github.com/dirkjabl/bricker/device/bricklet/lcd20x4"
	"unicode/utf8"
)

// NewDefaultTextLine converts a unicode string to the ks0066 lcd byte string.
func NewDefaultTextLine(line uint8, txt string) *lcd20x4.DefaultTextLine {
	var i int
	ltl := &lcd20x4.DefaultTextLine{Line: line}
	// convert string to bytes
	text := []byte(txt)
	for len(text) > 0 && i < 20 {
		r, size := utf8.DecodeRune(text)
		text = text[size:]
		ltl.Text[i] = ToByte(r)
		i++
	}
	return ltl
}
