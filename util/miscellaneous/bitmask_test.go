// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package miscellaneous

import (
	"testing"
)

func TestMaskToString(t *testing.T) {
	test2b := map[uint8]string{
		0x0: "00",
		0x1: "01",
		0x2: "10",
		0x3: "11"}
	test4b := map[uint8]string{
		0x00: "0000",
		0x01: "0001",
		0x03: "0011",
		0x05: "0101",
		0x0a: "1010",
		0x0f: "1111",
		0xf0: "0000",
		0xff: "1111"}
	test8b := map[uint8]string{
		0x00: "00000000",
		0x55: "01010101",
		0x81: "10000001",
		0xaa: "10101010",
		0xff: "11111111"}
	testfu1 := map[uint8]string{
		0x0: "00000000",
		0x1: "00000001"}
	testfu4 := map[uint8]string{
		0x00: "00000000",
		0x0f: "00001111",
		0x0a: "00001010",
		0xf0: "00000000"}
	tests := []struct {
		tm *map[uint8]string
		l  uint8
		fu bool
	}{{tm: &test2b, l: 2, fu: false},
		{tm: &test4b, l: 4, fu: false},
		{tm: &test8b, l: 8, fu: false},
		{tm: &testfu1, l: 1, fu: true},
		{tm: &testfu4, l: 4, fu: true}}
	for _, ts := range tests {
		for k, v := range *ts.tm {
			r := MaskToString(k, ts.l, ts.fu)
			if v != r {
				t.Fatalf("Error TestMaskToString: Not the right bitmask string generated (0x%02x -> %s != %s).",
					k, v, r)
			}
		}
	}
}
