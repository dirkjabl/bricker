// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package payload

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNewPayload(t *testing.T) {
	p := New([]byte("123")) // 3 bytes
	if (p == nil || p == &Payload{}) {
		t.Fatalf("Error TestNew: Payload empty, should not. ")
	}
	n := New(nil) // 0 bytes, empty payload
	if p == n {
		t.Fatalf("Error TestNew: Payload %v != %v. ", p, n)
	}
	if len(*p) != 3 {
		t.Fatalf("Error TestNew: Payload has wrong len %d != 3. ", len(*p))
	}
}

func TestBytesPayload(t *testing.T) {
	a := []byte("12345")
	p := New(a) // 5 bytes
	b := p.Bytes()
	if len(a) != len(b) {
		t.Fatalf("Error TestBytesPayload: get not same len for byte slices %v != %v. ", len(a), len(b))
	}
	if bytes.Compare(a, b) != 0 {
		t.Fatalf("Error TestBytesPayload: get not same byte slices %v != %v. ", a, b)
	}
}

func TestStringPayload(t *testing.T) {
	a := "1234567890"
	b := []byte(a)
	c := ""
	p := New(b) // 10 bytes
	s := p.String()
	if s == "" {
		t.Fatalf("Error TestStringPayload: String is empty. ")
	}
	for _, v := range b {
		c += fmt.Sprintf("0x%02x ", v)
	}
	c = "Payload: [ " + c + "]"
	if s != c {
		t.Fatalf("Error TestStringPayload: Strings mismatch: %s != %s. ", s, c)
	}
}

func TestWritePayload(t *testing.T) {
	a := []byte("1234567890")
	p := New(a)
	c := p.Bytes()
	buf := bytes.NewBuffer(nil)
	err := p.Write(buf)
	if err != nil {
		t.Fatalf("Error: TestWritePayload: Could not write to buffer: %s ", err.Error())
	}
	b := buf.Bytes()
	if len(c) != len(b) {
		t.Fatalf("Error: TestWritePayload: Length of byte slice differ %d != %d ", len(c), len(b))
	}
	if bytes.Compare(c, b) != 0 {
		t.Fatalf("Error TestWritePayload: Get not same byte slices %v != %v. ", c, b)
	}
	p = New(nil)
	buf = bytes.NewBuffer(a)
	err = p.Write(buf)
	if err != nil {
		t.Fatalf("Error: TestWritePayload: Could not write to buffer: %s ", err.Error())
	}
	c = p.Bytes()
	b = buf.Bytes()
	if len(c) == len(b) {
		t.Fatalf("Error: TestWritePayload: Length of byte slice should differ %d != %d ", len(c), len(b))
	}
	if bytes.Compare(c, b) == 0 {
		t.Fatalf("Error TestWritePayload: Get same byte slices %v != %v. ", c, b)
	}
}
