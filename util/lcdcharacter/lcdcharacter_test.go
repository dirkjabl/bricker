// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lcdcharacter

import (
	"testing"
)

func TestConvertStringToCharacter(t *testing.T) {
	wanted := [8]uint8{17, 17, 17, 31, 17, 17, 17, 0}
	lines := [8]string{
		"O...O",
		"O...O",
		"O...O",
		"OOOOO",
		"O...O",
		"O...O",
		"O...O",
		"....."}
	cc := ConvertStringToCharacter(lines)
	for i, v := range wanted {
		if cc[i] != v {
			t.Fatalf("Error TestConvertStringToCharacter: Not the expected value (0x%x != 0x%x) in line %d.",
				cc[i], v, i)
		}
	}
}
