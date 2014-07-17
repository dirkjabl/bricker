// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package subscription

import (
	"bricker/util/hash"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	a := New(hash.ChoosenFunctionID, 1, 2, nil, true)
	b := New(hash.ChoosenUid, 1, 2, nil, false)
	if a.String() == b.String() {
		t.Fatalf("Error TestNew: different subscriptions are equal (%s == %s).",
			a.String(), b.String())
	}
}

func TestNewFid(t *testing.T) {
	a := New(hash.ChoosenFunctionID, 0, 1, nil, true)
	b := NewFid(1, nil, true)
	if a.String() != b.String() {
		t.Fatalf("Error TestNewFid: subscriptions are not equal (%s != %s).",
			a.String(), b.String())
	}
}

func TestHash(t *testing.T) {
	a := New(hash.ChoosenFunctionID, 1, 2, nil, true)
	b := New(hash.ChoosenUid, 1, 2, nil, false)
	h := a.Hash()
	if h.Equal(b.Hash()) {
		t.Fatalf("Error TestHash: hashes are equal but should different (%s = %s).",
			a.Hash(), b.Hash())
	}
}

func TestCompareHash(t *testing.T) {
	a := New(hash.ChoosenFunctionID, 1, 2, nil, true)
	b := New(hash.ChoosenUid, 1, 2, nil, false)
	h := b.Hash()
	if a.CompareHash(h) {
		t.Fatalf("Error TestCompareHash: different subscriptions are equal (%s = %s).",
			a.String(), b.String())
	}
}

func TestSubscriptionString(t *testing.T) {
	a := New(hash.ChoosenFunctionID, 1, 2, nil, true)
	if strings.Index(a.String(), "Request-Packet:") != 60 {
		t.Fatalf("Error TestSubscriptionString: String result not correct (%s - %d).",
			a.String(), strings.Index(a.String(), "Request-Packet:"))
	}
}
