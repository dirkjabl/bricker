// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package device

import (
	"github.com/dirkjabl/bricker/net/packet"
	"github.com/dirkjabl/bricker/subscription"
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
		t.Fatalf("Error: TestNew faild, Identifer mismatch (%s != \"\").", d.Id())
	}
	d = nil
	if d.Id() != "" {
		t.Fatalf("Error: TestNew faild, Identifer mismatch (%s != \"\").", d.Id())
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

func TestResulter(t *testing.T) {
	d := New("Test")
	a := &EmptyResult{}
	if d.Result() != nil {
		t.Fatalf("Error: TestResulter gets a resulter without setting one.")
	}
	d.SetResult(a)
	b := d.Result()
	if b.String() != "EmptyResult []" {
		t.Fatalf("Error: TestResulter gets a resulter but is not the empty resulter (%s).", b)
	}
	d = nil
	d.SetResult(b)
	if d.Result() != nil {
		t.Fatalf("Error: TestResulter empty device should not give a resulter. (%s)", b)
	}
}
