// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io16

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
	"github.com/dirkjabl/bricker/subscription"
	"github.com/dirkjabl/bricker/util/hash"
	misc "github.com/dirkjabl/bricker/util/miscellaneous"
)

// SetPortInterrupt creates the subscriber to set the interrupt bitmask for a port.
func SetPortInterrupt(id string, uid uint32, pi *PortInterrupt, handler func(device.Resulter, error)) *device.Device {
	fid := function_set_port_interrupt
	spi := device.New(device.FallbackId(id, "SetPortInterrupt"))
	p := packet.NewSimpleHeaderPayload(uid, fid, true, pi)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	spi.SetSubscription(sub)
	spi.SetResult(&device.EmptyResult{})
	spi.SetHandler(handler)
	return spi
}

// SetPortInterruptFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetPortInterruptFuture(brick *bricker.Bricker, connectorname string, uid uint32, pi *PortInterrupt) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetPortInterrupt("setportinterruptfuture"+device.GenId(), uid, pi,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetPortInterrupt creates the subscriber to get the interrupt bitmask for a port.
func GetPortInterrupt(id string, uid uint32, po *Port, handler func(device.Resulter, error)) *device.Device {
	fid := function_get_port_interrupt
	gpi := device.New(device.FallbackId(id, "GetPortInterrupt"))
	p := packet.NewSimpleHeaderPayload(uid, fid, true, po)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	gpi.SetSubscription(sub)
	gpi.SetResult(&Interrupt{})
	gpi.SetHandler(handler)
	return gpi
}

// GetPortInterruptFuture is a future pattern version for a synchronized all of the subscriber.
// If an error occur, the result is nil.
func GetPortInterruptFuture(brick *bricker.Bricker, connectorname string, uid uint32, po *Port) *Interrupt {
	future := make(chan *Interrupt)
	defer close(future)
	sub := GetPortInterrupt("getinterruptfuture"+device.GenId(), uid, po,
		func(r device.Resulter, err error) {
			var v *Interrupt = nil
			if err == nil {
				if value, ok := r.(*Interrupt); ok {
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

// InterruptTrigger creates a subscriber for the interrupt callback.
// This callback is triggered whenever a change of the voltage level is detected
// on pins where the interrupt was activated with SetInterrupt.
func InterruptTrigger(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := callback_interrupt
	it := device.New(device.FallbackId(id, "InterruptTrigger"))
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, nil, true)
	it.SetSubscription(sub)
	it.SetResult(&Interrupts{})
	it.SetHandler(handler)
	return it
}

// PortInterrupt is a combined type.
type PortInterrupt struct {
	Port          byte
	InterruptMask uint8
}

/*
 Interrupt bitmask type.
 Interrupts are triggered on changes of the voltage level of the pin,
 i.e. changes from high to low and low to high.
*/
type Interrupt struct {
	Mask uint8 // bitmask 8bit
}

// FromPacket creates a Interrupt from a packet.
func (i *Interrupt) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(i, p); err != nil {
		return err
	}
	return p.Payload.Decode(i)
}

// String fullfill the stringer interface.
func (i *Interrupt) String() string {
	txt := "Interrupt "
	if i == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("[Mask: %d (%s)]",
			i.Mask, misc.MaskToString(i.Mask, 8, false))
	}
	return txt
}

// Interrupts is the result type of the interrupt callback.
type Interrupts struct {
	Port          byte  // port a or b
	InterruptMask uint8 // bitmap 8bit
	ValueMask     uint8 // bitmap 8bit
}

// FromPacket creates a Interrupts object from a packet.
func (i *Interrupts) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(i, p); err != nil {
		return err
	}
	return p.Payload.Decode(i)
}

// String fullfill the stringer interface.
func (i *Interrupts) String() string {
	txt := "Interrupts "
	if i == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("[Port: %s, Interrupt Mask: %d (%s), Value Mask: %d (%s)]",
			i.Port,
			i.InterruptMask, misc.MaskToString(i.InterruptMask, 8, false),
			i.ValueMask, misc.MaskToString(i.ValueMask, 8, false))
	}
	return txt
}
