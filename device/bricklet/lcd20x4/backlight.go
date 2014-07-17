// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lcd20x4

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
	"github.com/dirkjabl/bricker/subscription"
	"github.com/dirkjabl/bricker/util/hash"
)

func BacklightOn(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := function_backlight_on
	blon := device.New(device.FallbackId(id, "BacklightOn"))
	p := packet.NewSimpleHeaderOnly(uid, fid, true)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	blon.SetSubscription(sub)
	blon.SetResult(&device.EmptyResult{})
	blon.SetHandler(handler)
	return blon
}

func BacklightOnFuture(brick *bricker.Bricker, connectorname string, uid uint32) bool {
	future := make(chan bool)
	sub := BacklightOn("backlightonfuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	b := <-future
	close(future)
	return b
}

func BacklightOff(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := function_backlight_off
	bloff := device.New(device.FallbackId(id, "BacklightOff"))
	p := packet.NewSimpleHeaderOnly(uid, fid, true)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	bloff.SetSubscription(sub)
	bloff.SetResult(&device.EmptyResult{})
	bloff.SetHandler(handler)
	return bloff
}

func BacklightOffFuture(brick *bricker.Bricker, connectorname string, uid uint32) bool {
	future := make(chan bool)
	sub := BacklightOff("backlightofffuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	b := <-future
	close(future)
	return b
}

func IsBacklightOn(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := function_is_backlight_on
	iblo := device.New(device.FallbackId(id, "IsBacklightOn"))
	p := packet.NewSimpleHeaderOnly(uid, fid, true)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	iblo.SetSubscription(sub)
	iblo.SetResult(&Backlight{})
	iblo.SetHandler(handler)
	return iblo
}

func IsBacklightOnFuture(brick *bricker.Bricker, connectorname string, uid uint32) *Backlight {
	future := make(chan *Backlight)
	sub := IsBacklightOn("isbacklightonfuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *Backlight = nil
			if err == nil {
				if value, ok := r.(*Backlight); ok {
					v = value
				}
			}
			future <- v
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return nil
	}
	v := <-future
	close(future)
	return v
}

func IsBacklightOnFutureSimple(brick *bricker.Bricker, connectorname string, uid uint32) bool {
	bl := IsBacklightOnFuture(brick, connectorname, uid)
	if bl != nil && bl.IsOn {
		return true
	}
	return false
}

// Backlight is a type for the return of the IsBacklightOn subscriber.
type Backlight struct {
	IsOn bool // is the backlight on
}

// FromPacket converts the packet payload to the Backlight type.
func (bl *Backlight) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(bl, p); err != nil {
		return err
	}
	blr := NewBacklightRawFromBacklight(bl)
	err := p.Payload.Decode(blr)
	if err == nil {
		bl.FromBacklightRaw(blr)
	}
	return err
}

// String fullfill the stringer interface.
func (bl *Backlight) String() string {
	txt := "Backlight ["
	if bl != nil {
		txt += fmt.Sprintf("IsOn: %t", bl.IsOn)
	} else {
		txt += "nil"
	}
	return txt + "]"
}

func (bl *Backlight) FromBacklightRaw(br *BacklightRaw) {
	if bl == nil || br == nil {
		return
	}
	bl.IsOn = (br.IsOn & 0x01) == 0x01
}

// BacklightRaw is a type for raw coding of the backlight.
type BacklightRaw struct {
	IsOn uint8
}

func NewBacklightRawFromBacklight(b *Backlight) *BacklightRaw {
	br := new(BacklightRaw)
	br.FromBacklight(b)
	return br
}

func (br *BacklightRaw) FromBacklight(b *Backlight) {
	if b == nil {
		return
	}
	if b.IsOn {
		br.IsOn = 0x01
	} else {
		br.IsOn = 0x00
	}
}
