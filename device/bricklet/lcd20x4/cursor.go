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

func SetConfig(id string, uid uint32, cursor *Cursor, handler func(device.Resulter, error)) *device.Device {
	fid := function_set_config
	sc := device.New(device.FallbackId(id, "SetConfig"))
	p := packet.NewSimpleHeaderPayload(uid, fid, true, NewCursorRawFromCursor(cursor))
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	sc.SetSubscription(sub)
	sc.SetResult(&device.EmptyResult{})
	sc.SetHandler(handler)
	return sc
}

func SetConfigFuture(brick *bricker.Bricker, connectorname string, uid uint32, cursor *Cursor) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetConfig("setconfigfuture"+device.GenId(), uid, cursor,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	b := <-future
	return b
}

func GetConfig(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := function_get_config
	gc := device.New(device.FallbackId(id, "GetConfig"))
	p := packet.NewSimpleHeaderOnly(uid, fid, true)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	gc.SetSubscription(sub)
	gc.SetResult(&Cursor{})
	gc.SetHandler(handler)
	return gc
}

func GetConfigFuture(brick *bricker.Bricker, connectorname string, uid uint32) *Cursor {
	future := make(chan *Cursor)
	sub := GetConfig("getconfigfuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *Cursor = nil
			if err == nil {
				if value, ok := r.(*Cursor); ok {
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

// Cursor config type. For setting or getting the cursor state.
type Cursor struct {
	Show     bool // is the cursor shown as line
	Blinking bool // is the cursor blinking
}

// FromPacket converts the packet payload to the Cursor type.
func (c *Cursor) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(c, p); err != nil {
		return err
	}
	rc := NewCursorRawFromCursor(c)
	err := p.Payload.Decode(rc)
	if err == nil && rc != nil {
		c.FromCursorRaw(rc)
	}
	return err
}

// String fullfill the stringer interface.
func (c *Cursor) String() string {
	txt := "Cursor ["
	if c != nil {
		txt += fmt.Sprintf("Show: %t, Blinking: %t", c.Show, c.Blinking)
	} else {
		txt += "nil"
	}
	return txt + "]"
}

// FromCursorRaw converts a CursorRaw type into a Cursor type.
func (c *Cursor) FromCursorRaw(cr *CursorRaw) {
	if c == nil || cr == nil {
		return
	}
	c.Show = (cr.Show & 0x01) == 0x01
	c.Blinking = (cr.Blinking & 0x01) == 0x01
}

// CursorRaw is the real de/encoding type for a cursor.
type CursorRaw struct {
	Show     uint8
	Blinking uint8
}

// NewFromCursor creates a CursorRaw object from a Cursor.
func NewCursorRawFromCursor(c *Cursor) *CursorRaw {
	if c == nil {
		return nil
	}
	cr := new(CursorRaw)
	cr.FromCursor(c)
	return cr
}

// FromCursor converts a Cursor type to a CursorRaw type.
func (cr *CursorRaw) FromCursor(c *Cursor) {
	if cr == nil || c == nil {
		return
	}
	if c.Show {
		cr.Show = 0x01
	} else {
		cr.Show = 0x00
	}
	if c.Blinking {
		cr.Blinking = 0x01
	} else {
		cr.Blinking = 0x00
	}
}
