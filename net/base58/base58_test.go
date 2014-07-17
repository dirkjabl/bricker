// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package base58

import (
	"testing"
)

func convertByte(src []byte) [8]byte {
	var dest [8]byte
	var i int
	for i = 0; i < len(src); i++ {
		dest[i] = src[i]
	}
	return dest
}

var testset map[uint64][8]byte

func init() {
	// CGy => 123456
	// 6DbsDo => 3702534944
	testset = map[uint64][8]byte{
		123456:     convertByte([]byte("CGy")),
		3702534944: convertByte([]byte("6DbsDo")),
	}
}

func TestDecode(t *testing.T) {
	for k := range testset {
		dest := Decode(testset[k])
		if dest != k {
			t.Fatalf("Error Base58: %s, decode to uint64: %d - has to %d ", testset[k], dest, k)
		}
	}
}

func TestEncode(t *testing.T) {
	for k := range testset {
		dest := Encode(k)
		if dest != testset[k] {
			t.Fatalf("Error TestEncode: %d uint64, encode to bytes: %s - has to %s ", k, dest, testset[k])
		}
	}
}

func TestConvert32(t *testing.T) {
	for k := range testset {
		dest := Convert32(k)
		if uint64(dest) != k {
			t.Fatalf("Error TestConvert32: %d uint64 != %d uint32 ", k, dest)
		}
	}
}
