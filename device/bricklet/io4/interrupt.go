// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io4

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
	misc "github.com/dirkjabl/bricker/util/miscellaneous"
)

// SetInterrupt creates the subscriber to set the interrupt bitmask.
func SetInterrupt(id string, uid uint32, i *Interrupt, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SetInterrupt"),
		Fid:        function_set_interrupt,
		Uid:        uid,
		Data:       i,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SetInterruptFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetInterruptFuture(brick *bricker.Bricker, connectorname string, uid uint32, i *Interrupt) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetInterrupt("setinterruptfuture"+device.GenId(), uid, i,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetInterrupt creates the subscriber to get the interrupt bitmask.
func GetInterrupt(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetInterrupt"),
		Fid:        function_get_interrupt,
		Uid:        uid,
		Result:     &Interrupt{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetInterruptFuture is a future pattern version for a synchronized all of the subscriber.
// If an error occur, the result is nil.
func GetInterruptFuture(brick *bricker.Bricker, connectorname string, uid uint32) *Interrupt {
	future := make(chan *Interrupt)
	defer close(future)
	sub := GetInterrupt("getinterruptfuture"+device.GenId(), uid,
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
	return device.Generator{
		Id:         device.FallbackId(id, "InterruptTrigger"),
		Fid:        callback_interrupt,
		Uid:        uid,
		Result:     &Interrupts{},
		Handler:    handler,
		IsCallback: true,
		WithPacket: false}.CreateDevice()
}

// Interrupt bitmask type.
// Interrupts are triggered on changes of the voltage level of the pin,
// i.e. changes from high to low and low to high.
type Interrupt struct {
	Mask uint8 // bitmask 4bit
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
			i.Mask, misc.MaskToString(i.Mask, 4, true))
	}
	return txt
}

// Copy creates a copy of the content.
func (i *Interrupt) Copy() device.Resulter {
	if i == nil {
		return nil
	}
	return &Interrupt{Mask: i.Mask}
}

// Interrupts is the result type of the interrupt callback.
type Interrupts struct {
	InterruptMask uint8 // bitmap 4bit
	ValueMask     uint8 // bitmap 4nit
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
		txt += fmt.Sprintf("[Interrupt Mask: %d (%s), Value Mask: %d (%s)]",
			i.InterruptMask, misc.MaskToString(i.InterruptMask, 4, true),
			i.ValueMask, misc.MaskToString(i.ValueMask, 4, true))
	}
	return txt
}

// Copy creates a copy of the content.
func (i *Interrupts) Copy() device.Resulter {
	if i == nil {
		return nil
	}
	return &Interrupts{
		InterruptMask: i.InterruptMask,
		ValueMask:     i.ValueMask}
}
