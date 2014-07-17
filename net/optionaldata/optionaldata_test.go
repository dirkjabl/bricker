// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package optionaldata

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNewOptionalData(t *testing.T) {
	o := New([]byte("123")) // 3 bytes
	if (o == nil || o == &OptionalData{}) {
		t.Fatalf("Error TestOptionalData: OptionalData empty, should not.")
	}
	n := New(nil) // 0 bytes, empty OptionalData
	if o == n {
		t.Fatalf("Error TestNewOptionalData: OptionalData %v != %v.", o, n)
	}
	if len(*o) != 3 {
		t.Fatalf("Error TestNewOptionalData: OptionalData has wrong len %d != 3.", len(*o))
	}
}

func TestBytesOptionalData(t *testing.T) {
	a := []byte("12345")
	p := New(a) // 5 bytes
	b := p.Bytes()
	if len(a) != len(b) {
		t.Fatalf("Error TestBytesOptionalData: get not same len for byte slices %v != %v. ", len(a), len(b))
	}
	if bytes.Compare(a, b) != 0 {
		t.Fatalf("Error TestBytesOptionalData: get not same byte slices %v != %v. \n", a, b)
	}
}

func TestStringOptionalData(t *testing.T) {
	a := "1234567890"
	b := []byte(a)
	c := ""
	o := New(b) // 10 bytes
	s := o.String()
	if s == "" {
		t.Fatalf("Error TestStringOptionalData: String is empty.")
	}
	for _, v := range b {
		c += fmt.Sprintf("0x%02x ", v)
	}
	c = "Optional Data: [ " + c + "]"
	if s != c {
		t.Fatalf("Error TestStringOptionalData: Strings mismatch: %s != %s. \n", s, c)
	}
}

// t.Fatalf
