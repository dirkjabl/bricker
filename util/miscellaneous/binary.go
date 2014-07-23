// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package miscellaneous

// Helping methods for the binary en/decoding

// Converts from a boolean to a uint8 for encoding.
func BoolToUint8(b bool) uint8 {
	if b {
		return 0x01
	}
	return 0x00
}

// Converts from a uint8 to a bool for decoding.
// true is only if the first bit is set (0x01).
func Uint8ToBool(u uint8) bool {
	return (u & 0x01) == 0x01
}
