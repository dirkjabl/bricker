// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package miscellaneous

import (
	"testing"
)

func TestBoolToUint8(t *testing.T) {
	a := true
	if BoolToUint8(a) != 0x01 {
		t.Fatalf("Error TestBoolToUint8: Value should be 0x01 but is not (Ox%02x != %t).",
			BoolToUint8(a), a)
	}
	a = false
	if BoolToUint8(a) != 0x00 {
		t.Fatalf("Error TestBoolToUint8: Value should be 0x00 but is not (0x%02x != %t).",
			BoolToUint8(a), a)
	}
}

func TestUint8ToBool(t *testing.T) {
	a := uint8(0x01)
	if Uint8ToBool(a) != true {
		t.Fatalf("Error TestUint8ToBool: Want a true value, but did not get (%t != 0x%02x).",
			Uint8ToBool(a), a)
	}
	a = uint8(0x07)
	if Uint8ToBool(a) != true {
		t.Fatalf("Error TestUint8ToBool: Want a true value, but did not get (%t != 0x%02x).",
			Uint8ToBool(a), a)
	}
	a = uint8(0x00)
	if Uint8ToBool(a) != false {
		t.Fatalf("Error TestUint8ToBool: Want a false value, but did not get (%t != 0x%02x).",
			Uint8ToBool(a), a)
	}
	a = uint8(0x7e)
	if Uint8ToBool(a) != false {
		t.Fatalf("Error TestUint8ToBool: Want a false value, but did not get (%t != 0x%02x).",
			Uint8ToBool(a), a)
	}
}

func Test(t *testing.T) {
	a := true
	b := BoolToUint8(a)
	c := Uint8ToBool(b)
	if c != a {
		t.Fatalf("Error Test: Converted values should be equal (%t != 0x%02x).",
			c, b)
	}
	a = false
	b = BoolToUint8(a)
	c = Uint8ToBool(b)
	if c != a {
		t.Fatalf("Error Test: Converted values should be equal (%t != 0x%02x).",
			c, b)
	}
}
