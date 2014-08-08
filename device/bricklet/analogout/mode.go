// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package analogout

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
)

/*
SetMode creates a subscriber to set the modes of the analog value.
The default value is 0.

  0: Normal Mode (Analog value as set by set_voltage is applied)
  1: 1k Ohm resistor to ground
  2: 100k Ohm resistor to ground
  3: 500k Ohm resistor to ground
*/
func SetMode(id string, uid uint32, m *Mode, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SetMode"),
		Fid:        function_set_mode,
		Uid:        uid,
		Data:       m,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SetModeFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetModeFuture(brick bricker.Bricker, connectorname string, uid uint32, m *Mode) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetMode("setmodefuture"+device.GenId(), uid, m,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetMode creates a subscriber to get the measurement mode value.
func GetMode(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetMode"),
		Fid:        function_get_mode,
		Uid:        uid,
		Result:     &Mode{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetModeFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func GetModeFuture(brick bricker.Bricker, connectorname string, uid uint32) *Mode {
	future := make(chan *Mode)
	defer close(future)
	sub := GetMode("getmodefuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *Mode = nil
			if err == nil {
				if value, ok := r.(*Mode); ok {
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

// Constants for the modes.
const (
	ModeNormal             = 0
	Mode1kResistorGround   = 1
	Mode100kResistorGround = 2
	Mode500kResistorGround = 3
)

// Mode result type
type Mode struct {
	Value uint8 // range identifer
}

// FromPacket creates a Range from a packet.
func (m *Mode) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(m, p); err != nil {
		return err
	}
	return p.Payload.Decode(m)
}

// Name converts the mode identifer value to a readable string.
func (m *Mode) Name() string {
	switch m.Value {
	case ModeNormal:
		return "Normal Mode"
	case Mode1kResistorGround:
		return "1k Ohm resistor to ground"
	case Mode100kResistorGround:
		return "100k Ohm resistor to ground"
	case Mode500kResistorGround:
		return "500k Ohm resistor to ground"
	default:
		return "Unknown"
	}
}

// String fullfill the stringer interface.
func (m *Mode) String() string {
	txt := "Mode "
	if m == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("[Value: %s (%d)]", m.Name(), m.Value)
	}
	return txt
}
