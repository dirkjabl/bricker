// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lcd20x4

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
	misc "github.com/dirkjabl/bricker/util/miscellaneous"
)

func SetConfig(id string, uid uint32, cursor *Cursor, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SetConfig"),
		Fid:        function_set_config,
		Uid:        uid,
		Data:       NewCursorRaw(cursor),
		Handler:    handler,
		WithPacket: true}.CreateDevice()
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
	return device.Generator{
		Id:         device.FallbackId(id, "GetConfig"),
		Fid:        function_get_config,
		Uid:        uid,
		Result:     &Cursor{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
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
	rc := new(CursorRaw)
	err := p.Payload.Decode(rc)
	if err == nil && rc != nil {
		c.FromCursorRaw(rc)
	}
	return err
}

// String fullfill the stringer interface.
func (c *Cursor) String() string {
	txt := "Cursor "
	if c != nil {
		txt += fmt.Sprintf("[Show: %t, Blinking: %t]", c.Show, c.Blinking)
	} else {
		txt += "[nil]"
	}
	return txt
}

// FromCursorRaw converts a CursorRaw type into a Cursor type.
func (c *Cursor) FromCursorRaw(cr *CursorRaw) {
	if c == nil || cr == nil {
		return
	}
	c.Show = misc.Uint8ToBool(cr.Show)
	c.Blinking = misc.Uint8ToBool(cr.Blinking)
}

// CursorRaw is the real de/encoding type for a cursor.
type CursorRaw struct {
	Show     uint8
	Blinking uint8
}

// NewFromCursor creates a CursorRaw object from a Cursor.
func NewCursorRaw(c *Cursor) *CursorRaw {
	if c == nil {
		return nil
	}
	cr := new(CursorRaw)
	cr.Show = misc.BoolToUint8(c.Show)
	cr.Blinking = misc.BoolToUint8(c.Blinking)
	return cr
}
