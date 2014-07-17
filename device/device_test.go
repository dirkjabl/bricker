// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package device

import (
	"bricker/net/packet"
	"bricker/subscription"
	"testing"
)

func TestNew(t *testing.T) {
	d := New("Test")
	if d.Id() != "Test" {
		t.Fatalf("Error: TestNew faild, Identifer mismatch (%s != Test).", d.Id())
	}
}

func TestId(t *testing.T) {
	d := New("")
	if d.Id() != "" {
		t.Fatalf("Error: TestNew faild, Identifer mismatch (%s != ).", d.Id())
	}
}

func TestSubscription(t *testing.T) {
	d := New("Test")
	s0 := subscription.NewFid(1, packet.NewSimpleHeaderOnly(0, 1, true), true)
	d.SetSubscription(s0)
	s1 := d.Subscription()
	if s0.FunctionID != s1.FunctionID {
		t.Fatalf("Error: TestSubscription seted subscription mismatch.")
	}
}
