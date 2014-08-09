// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package device

import (
	//	"github.com/dirkjabl/bricker/net/packet"
	//	"github.com/dirkjabl/bricker/subscription"
	//"fmt"
	"github.com/dirkjabl/bricker/util/hash"
	"testing"
)

type Testdata struct {
	a uint8
	b uint8
}

func TestCreateDevice(t *testing.T) {
	base := Generator{
		Id: "Test",
	} // WithPacket and IsCallback should be false, Resulter, Data and Handler not set
	d := base.CreateDevice()
	if d.Id() != "Test" {
		t.Fatalf("Error: TestCreateDevice failed, Identifer mismatch (%s != Test).", d.Id())
	}
	s0 := d.Subscription()

	if s0.Uid != 0 || s0.FunctionID != 0 || s0.Request != nil || s0.Choosen != hash.ChoosenFunctionIDUid {
		t.Fatalf("Error: TestCreateDevice failed, subscription content not zero (%v).", s0)
	}
	base.WithPacket = true
	base.Fid = uint8(1)
	d = base.CreateDevice()
	s1 := d.Subscription()
	if s1.FunctionID != 1 {
		t.Fatalf("Error: TestCreateDevice failed, Funtion-ID of subscription mismatch (%d != %d).", s1.FunctionID, uint8(1))
	}
	if s1.CompareHash(s0.Hash()) {
		t.Fatalf("Error: TestCreateDevice failed, Subscriptions are equal but should not (%v != %v).", s0, s1)
	}
	if s1.Request == nil || s1.Request.Payload != nil {
		t.Fatalf("Error: TestCreateDevice failed, subscription has no request packet (%v).", s1)
	}
	if d.result == nil {
		t.Fatalf("Error: TestCreateDevice failed, the create device object should have a Resulter (%v).", d)
	}
	if er, erok := d.result.(*EmptyResult); !erok {
		t.Fatalf("Error: TestCreateDevice failed, the device object should have a EmptyResult resulter (%v)", er)
	}
	base.Fid = uint8(2)
	base.Data = &Testdata{a: 1, b: 2}
	base.Result = &Period{}
	d = base.CreateDevice()
	s2 := d.Subscription()
	if s2.Request == nil || s2.Request.Payload == nil {
		t.Fatalf("Error: TestCreateDevice failed, subscription has no request packet or the payload is empty (%v).", s2)
	}
	if pr, prok := d.result.(*Period); !prok {
		t.Fatalf("Error: TestCreateDevice failed, the device object should have a Period resulter (%v)", pr)
	}
}
