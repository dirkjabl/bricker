// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package packet

import (
	// "fmt"
	"bricker/net/head"
	"bricker/net/optionaldata"
	"bricker/net/payload"
	"testing"
)

func TestComputeLength(t *testing.T) {
	p := &Packet{nil, nil, nil}
	l := p.ComputeLength()
	if l != 0 {
		t.Fatalf("Error TestComputeLength: Length = %d should 0", l)
	}
	p.Head = &head.Head{}
	l = p.ComputeLength()
	if l != 8 {
		t.Fatalf("Error TestComputeLength: Length = %d should 8", l)
	}
	p.Payload = payload.New([]byte("123")) // 3 bytes
	l = p.ComputeLength()
	if l != (8 + 3) {
		t.Fatalf("Error TestComputeLength: Length = %d should %d", l, (8 + 3))
	}
	p.OptionalData = optionaldata.New([]byte("1234")) // 4 bytes
	l = p.ComputeLength()
	if l != (8 + 3 + 4) {
		t.Fatalf("Error TestComputeLength: Length = %d should %d", l, (8 + 3 + 4))
	}
	p.Head = nil
	l = p.ComputeLength()
	if l != (3 + 4) {
		t.Fatalf("Error TestComputeLength: Length = %d should %d", l, (3 + 4))
	}
	p.Payload = nil
	l = p.ComputeLength()
	if l != 4 {
		t.Fatalf("Error TestComputeLength: Length = %d should %d", l, 4)
	}
	p.OptionalData = nil
	p.Payload = payload.New([]byte("123")) // 3 bytes
	l = p.ComputeLength()
	if l != 3 {
		t.Fatalf("Error TestComputeLength: Length = %d should %d", l, 3)
	}
	p.Payload = nil
	p.OptionalData = optionaldata.New([]byte("1234")) // 4 bytes
	p.Head = &head.Head{}
	l = p.ComputeLength()
	if l != (8 + 4) {
		t.Fatalf("Error TestComputeLength: Length = %d should %d", l, (8 + 4))
	}
}

// t.Fatalf
