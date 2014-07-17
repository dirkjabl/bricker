// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package generator

import (
	"testing"
)

func TestNew(t *testing.T) {
	a := New()
	var id1 uint32 = a.Get()
	var id2 uint32 = a.Get()
	if id1 == id2 {
		t.Fatalf("Error TestNew: To different ids are equal (%d = %d).", id1, id2)
	}
	if id1 != (id2 - 1) {
		t.Fatalf("Error TestNew: Two ids should increment (%d = %d - 1).", id1, id2)
	}
	b := New()
	c := New()
	if b.Get() != c.Get() {
		t.Fatalf("Error TestNew: To generators should start with the same id.")
	}
}

func TestGet(t *testing.T) {
	a := New()
	var id1 uint32 = a.Get()
	var id2 uint32 = a.Get()
	if id1 != (id2 - 1) {
		t.Fatalf("Error TestGet: Two ids should increment (%d = %d - 1).", id1, id2)
	}
}

func TestGetUint8(t *testing.T) {
	//	t.Parallel()
	a := New()
	var id1 uint8 = a.Get8()
	var id2 uint8 = a.Get8()
	if id1 != (id2 - 1) {
		t.Fatalf("Error TestGet8: Two ids should increment (%d = %d - 1).", id1, id2)
	}
	for i := id2; i < 255; i++ {
		if i != id2 {
			t.Fatalf("Error TestGet8: Two ids should be equal (%d = %d).", i, id2)
		}
		id2 = a.Get8()
	}
	if id2 != 255 {
		t.Fatalf("Error TestGet8: Two ids should be equal (%d = %d).", id2, 255)
	}
	id2 = a.Get8()
	if id2 != 0 {
		t.Fatalf("Error TestGet8: Two ids should be equal (%d = %d).", id2, 0)
	}
	id2 = a.Get8()
	if id2 != 1 {
		t.Fatalf("Error TestGet8: Two ids should be equal (%d = %d).", id2, 1)
	}
}
