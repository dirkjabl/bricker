// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package virtual

import (
	"errors"
	"github.com/dirkjabl/bricker/event"
	"github.com/dirkjabl/bricker/net/packet"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	v := New()
	if v == nil {
		t.Fatal("Error TestNew: New do not create a object (nil pointer).")
	}
	v.Send(nil)
	if v.fallback == nil {
		t.Fatal("Error TestNew: New creates an object without fallback generator.")
	}
}

func TestAttachFallbackGenerator(t *testing.T) {
	v := New()
	v.AttachFallbackGenerator(nilEventGenerator)
	a := nilEventGenerator(nil)
	b := v.fallback(nil)
	if !(a.Packet == nil && b.Packet == nil) {
		t.Fatalf("Error TestAttachFallbackGenerator: Fallback generator, one packet not nil (%v != %v).",
			a.Packet, b.Packet)
	}
	v.AttachFallbackGenerator(copyEventGenerator)
	b = v.fallback(nil)
	if b != nil {
		t.Fatalf("Error TestAttachFallbackGenerator: Fallback generator, event not nil (%v).", b)
	}
	a = sampleEventGenerator(nil)
	b = v.fallback(a)
	if b.Packet.Head.Uid != a.Packet.Head.Uid ||
		b.Packet.Head.FunctionID != a.Packet.Head.FunctionID ||
		b.Packet.Head.Sequence() != (a.Packet.Head.Sequence()+1) {
		t.Fatalf("Error TestAttachFallbackGenerator: Fallback generator, Packets does not match (%v != %v).",
			a.Packet, b.Packet)
	}
}

func TestFallback(t *testing.T) {
	e := Fallback(nilEventGenerator(nil))
	if e != nil {
		t.Fatalf("Error TestFallback: Get not nil back (%v).", e)
	}
	e = Fallback(nil) // nil should ok, too
	if e != nil {
		t.Fatalf("Error TestFallback: Get not nil back by nil call (%v).", e)
	}
}

func nilEventGenerator(e *event.Event) *event.Event {
	return event.New(errors.New("Error"), time.Now(), nil)
}

func copyEventGenerator(e *event.Event) *event.Event {
	var ev *event.Event = nil
	if e != nil {
		ev = e.Copy()
		ev.Packet.Head.SetSequence(e.Packet.Head.Sequence() + 1)
	}
	return ev
}

func sampleEventGenerator(e *event.Event) *event.Event {
	var ev *event.Event
	if e != nil {
		ev = e.Copy()
		ev.Packet.Head.SetSequence(ev.Packet.Head.Sequence() + 1)
	} else {
		ev = event.New(nil, time.Now(), packet.NewSimpleHeaderOnly(123, 1, true))
	}
	return ev
}
