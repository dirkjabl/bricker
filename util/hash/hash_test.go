// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hash

import (
	"testing"
)

func TestNew(t *testing.T) {
	a := New(ChoosenFunctionID|ChoosenUid, 1, 2)
	b := New(ChoosenFunctionID, 1, 2)
	c := New(ChoosenUid, 1, 2)
	if a.String() == b.String() {
		t.Fatalf("Error TestNew: To different hash are equal (%s = %s).", a.String(), b.String())
	}
	if b.String() == c.String() {
		t.Fatalf("Error TestNew: To different hash are equal (%s = %s).", b.String(), c.String())
	}
	if a.String() == c.String() {
		t.Fatalf("Error TestNew: To different hash are equal (%s = %s).", a.String(), c.String())
	}
	c = New(ChoosenFunctionID, 1, 2)
	if c.String() != b.String() {
		t.Fatalf("Error TestNew: To same hash are not equal (%s != %s).", b.String(), c.String())
	}
	c = New(ChoosenNothing, 1, 2)
	if b.String() == c.String() {
		t.Fatalf("Error TestNew: To different hash are equal (%s = %s).", b.String(), c.String())
	}
	c = New(ChoosenFunctionIDUid, 1, 2)
	if c.String() != a.String() {
		t.Fatalf("Error TestNew: To same hash are not equal (%s != %s).", a.String(), c.String())
	}
}

func TestEqual(t *testing.T) {
	a := New(ChoosenFunctionID|ChoosenUid, 1, 2)
	b := New(ChoosenFunctionID, 1, 2)
	c := New(ChoosenFunctionID|ChoosenUid, 1, 2)
	if a.Equal(b) {
		t.Fatalf("Error TestEqual: To different hash are equal (%s = %s).", a.String(), b.String())
	}
	if !a.Equal(c) {
		t.Fatalf("Error TestEqual: To equal hash are not equal (%s = %s).", b.String(), c.String())
	}
	c = New(ChoosenFunctionIDUid, 1, 2)
	if !a.Equal(c) {
		t.Fatalf("Error TestEqual: To equal hash are not equal (%s = %s).", a.String(), c.String())
	}
}

func TestString(t *testing.T) {
	a := New(ChoosenFunctionID|ChoosenUid, 1, 2)
	b := New(ChoosenFunctionID, 1, 2)
	if b.String() != "Hash [ 48 166 236 22 46 87 55 77 24 214 250 90 32 100 93 22 ]" {
		t.Fatalf("Error TestString: String not correct (%s).", b.String())
	}
	if a.String() == "Hash [ 48 166 236 22 46 87 55 77 24 214 250 90 32 100 93 22 ]" {
		t.Fatalf("Error TestString: Strings should not be equal (%s).", a.String())
	}
}

func TestAll(t *testing.T) {
	all := All()
	for _, c := range all {
		switch c {
		case ChoosenFunctionID:
		case ChoosenUid:
		case ChoosenFunctionIDUid:
		case ChoosenNothing:
		default:
			t.Fatalf("Error TestAll: Unknown chooser (%d)", c)
		}
	}
}
