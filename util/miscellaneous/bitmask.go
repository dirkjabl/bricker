// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package miscellaneous

import (
	"strings"
)

/*
MaskToString converts for better printing a bitmask to a string.

  mask - bitmask to print (maximum 8 bit)
  len - length of the used bits inside the bitmap (e.g. 4 bits at IO-4 Bricklet)
  fillup - fillup the bitmask up to 8 bit (maximum)
*/
func MaskToString(mask, len uint8, fillup bool) string {
	ol := int(len)
	tm := ""
	for len > 0 {
		if (mask & 0x01) == 0x01 {
			tm = "1" + tm
		} else {
			tm = "0" + tm
		}
		len--
		mask >>= 1
	}
	if fillup && ol < 8 {
		tm = strings.Repeat("0", 8-ol) + tm
	}
	return tm
}
