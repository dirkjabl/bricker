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

// SetValue create the subscriber to set the output value with a bitmask (high or low).
// A 1 in the bitmask means high and a 0 in the bitmask means low.
// Only 4 bit are supported.
//
// This function does nothing for pins that are configured as input.
// Pull-up resistors can be switched on with SetConfiguration.
func SetValue(id string, uid uint32, v *Value, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SetValue"),
		Fid:        function_set_value,
		Uid:        uid,
		Data:       v,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SetValueFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetValueFuture(brick *bricker.Bricker, connectorname string, uid uint32, v *Value) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetValue("setvaluefuture"+device.GenId(), uid, v,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetValue creates the subscriber to get the output value.
func GetValue(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetValue"),
		Fid:        function_get_value,
		Uid:        uid,
		Result:     &Value{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetValueFuture is a future pattern version for a synchronized all of the subscriber.
// If an error occur, the result is nil.
func GetValueFuture(brick *bricker.Bricker, connectorname string, uid uint32) *Value {
	future := make(chan *Value)
	defer close(future)
	sub := GetValue("getvaluefuture"+device.GenId(), uid,
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

// SetSelectedValues creates a subscriber for setting values per bitmap (4bit).
// This function does nothing for pins that are configured as input.
// Pull-up resistors can be switched on with set_configuration.
func SetSelectedValues(id string, uid uint32, v *Values, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SetSelectedValues"),
		Fid:        function_set_selected_values,
		Uid:        uid,
		Data:       v,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SetSelectedValuesFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetSelectedValuesFuture(brick *bricker.Bricker, connectorname string, uid uint32, v *Values) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetSelectedValues("setselectedvalues"+device.GenId(), uid, v,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// Value is the type for the output bitmap mask (4bit).
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
		txt += fmt.Sprintf("[Mask: %d (%s)]", v.Mask, misc.MaskToString(v.Mask, 4, true))
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
