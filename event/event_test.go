// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package event

import (
	"bricker/net/packet"
	"errors"
	"strings"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	err := errors.New("Error")
	ts := time.Now()
	ev := New(err, ts, nil)
	if ev == nil {
		t.Fatalf("Error TestNew: Event nil")
	}
	if !ev.TimeStamp.Equal(ts) {
		t.Fatalf("Error TestNew: Event time stamp not equal (%s = %s)",
			ev.TimeStamp.Format(time.RFC3339), ts.Format(time.RFC3339))
	}
	if err.Error() != ev.Err.Error() {
		t.Fatalf("Error TestNew: Event error not equal (%s != %s).",
			err.Error(), ev.Err.Error())
	}
	if ev.Packet != nil {
		t.Fatalf("Error TestNew: Packet not nil")
	}
	p := packet.NewSimpleHeaderOnly(0, 0, true)
	ev = New(err, ts, p)
	if ev == nil {
		t.Fatalf("Error TestNew: Event nil")
	}
	if !ev.TimeStamp.Equal(ts) {
		t.Fatalf("Error TestNew: Event time stamp not equal (%s = %s)",
			ev.TimeStamp.Format(time.RFC3339), ts.Format(time.RFC3339))
	}
	if err.Error() != ev.Err.Error() {
		t.Fatalf("Error TestNew: Event error not equal (%s = %s)",
			err.Error(), ev.Err.Error())
	}
	if p.String() != ev.Packet.String() {
		t.Fatalf("Error TestNew: Event packet not equal (%s = %s)",
			p.String(), ev.Packet.String())
	}
}

func TestNewSimple(t *testing.T) {
	err := errors.New("Error")
	p := packet.NewSimpleHeaderOnly(0, 0, true)
	ev := NewSimple(err, p)
	if ev == nil {
		t.Fatalf("Error TestNewSimple: Event nil")
	}
	if p.String() != ev.Packet.String() {
		t.Fatalf("Error TestNew: Event packet not equal (%s = %s)",
			p.String(), ev.Packet.String())
	}
}

func TestNewPacket(t *testing.T) {
	err := errors.New("Error")
	ev := NewError(err)
	if ev == nil {
		t.Fatalf("Error TestNewSimple: Event nil")
	}
	if err.Error() != ev.Err.Error() {
		t.Fatalf("Error TestNew: Event error not equal (%s = %s)",
			err.Error(), ev.Err.Error())
	}
	if ev.Packet != nil {
		t.Fatalf("Error TestNewSimple: packet not nil")
	}
}

func TestCopy(t *testing.T) {
	err := errors.New("Error")
	p := packet.NewSimpleHeaderOnly(0, 0, true)
	ev := NewSimple(err, p)
	nev := ev.Copy()
	if nev == nil {
		t.Fatal("Error TestCopy: copy result is nil.")
	}
	if nev.Err.Error() != ev.Err.Error() || p.String() != ev.Packet.String() {
		t.Fatalf("Error TestCopy: copy result differs from source (%v != %v).", ev, nev)
	}
	ev = nil
	nev = ev.Copy()
	if nev != nil {
		t.Fatal("Error TestCopy: copy result is not nil.")
	}
}

func TestString(t *testing.T) {
	err := errors.New("Error")
	ev := NewError(err)
	// TODO: RFC3339Nano differs in one count !!! Should be changed...
	if !(strings.Index(ev.String(), "Error: Error") == 61 || strings.Index(ev.String(), "Error: Error") == 60) {
		t.Fatalf("Error TestString: String result not correct: %s (%d)",
			ev.String(), strings.Index(ev.String(), "Error: Error"))
	}
}
