// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io16

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
	misc "github.com/dirkjabl/bricker/util/miscellaneous"
)

/*
SetPort creates a subscriber to set a value mask (bitmap 8bit) to a port.

This function does nothing for pins that are configured as input.
Pull-up resistors can be switched on with SetPortConfiguration.
*/
func SetPort(id string, uid uint32, pv *PortValue, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SetPort"),
		Fid:        function_set_port,
		Uid:        uid,
		Data:       pv,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SetPortFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetPortFuture(brick *bricker.Bricker, connectorname string, uid uint32, pv *PortValue) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetPort("setportinterruptfuture"+device.GenId(), uid, pv,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetPort creates a subscriber to get the value bitmask (8bit) for a port.
// Returns a bitmask of the values that are currently measured on the specified port.
func GetPort(id string, uid uint32, po *Port, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetPort"),
		Fid:        function_get_port,
		Uid:        uid,
		Result:     &Value{},
		Data:       po,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetPortFuture is a future pattern version for a synchronized all of the subscriber.
// If an error occur, the result is nil.
func GetPortFuture(brick *bricker.Bricker, connectorname string, uid uint32, po *Port) *Value {
	future := make(chan *Value)
	defer close(future)
	sub := GetPort("getportfuture"+device.GenId(), uid, po,
		func(r device.Resulter, err error) {
			var v *Value = nil
			if err == nil {
				if value, ok := r.(*Value); ok {
					v = value
				}
			}
			future <- v
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return nil
	}
	return <-future
}

/*
PortValue is for setting a value mask for a specific port.
Ports could only be 'a' or 'b'.
The bitmask has a size of 8bit.
A value of 1 in the bitmask means 'hight' and a 0 is 'low'.

*/
type PortValue struct {
	Port      byte
	ValueMask uint8
}

// Value is the type for the output bitmap mask (8bit).
type Value struct {
	Mask uint8
}

// FromPacket creates a Value from a packet.
func (v *Value) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(v, p); err != nil {
		return err
	}
	return p.Payload.Decode(v)
}

// String fullfill the stringer interface.
func (v *Value) String() string {
	txt := "Value "
	if v == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("[Mask: %d (%s)]", v.Mask, misc.MaskToString(v.Mask, 8, true))
	}
	return txt
}

// Copy creates a copy of the content.
func (v *Value) Copy() device.Resulter {
	if v == nil {
		return nil
	}
	return &Value{Mask: v.Mask}
}
