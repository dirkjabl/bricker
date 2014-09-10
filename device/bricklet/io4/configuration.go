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

// SetConfiguration create the subscriber to set the value and the direction of the specified pin.
// Direction could be 'i' (input) or 'o' (output).
// If the direction is configured as input, the value is either pull-up or default (set as true or false).
func SetConfiguration(id string, uid uint32, c *Configuration, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SetConfiguration"),
		Fid:        function_set_configuration,
		Uid:        uid,
		Data:       NewConfigurationRaw(c),
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SetConfigurationFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetConfigurationFuture(brick *bricker.Bricker, connectorname string, uid uint32, c *Configuration) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetConfiguration("setconfigurationfuture"+device.GenId(), uid, c,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetConfiguration creates the subscriber to get the configuration of all pins.
// Returns a value bitmask and a direction bitmask.
// A 1 in the direction bitmask means input and a 0 in the bitmask means output.
func GetConfiguration(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetConfiguration"),
		Fid:        function_get_configuration,
		Uid:        uid,
		Result:     &Configurations{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetConfigurationFuture is a future pattern version for a synchronized all of the subscriber.
// If an error occur, the result is nil.
func GetConfigurationFuture(brick *bricker.Bricker, connectorname string, uid uint32) *Configurations {
	future := make(chan *Configurations)
	defer close(future)
	sub := GetConfiguration("getconfigurationfuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *Configurations = nil
			if err == nil {
				if value, ok := r.(*Configurations); ok {
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

// Configuration is a type to set the direction and the value of the specified pin(s).
type Configuration struct {
	SelectionMask uint8 // bitmask (4bit)
	Direction     byte  // 'i' - input, 'o' - output
	Value         bool  // true - hight, pull-up; false - low, default
}

// ConfigurationRaw is the raw type of the Configuration (for de/encoding).
type ConfigurationRaw struct {
	SelectionMask uint8 // bitmask (4bit)
	Direction     byte  // 'i' - input, 'o' - output
	Value         uint8 // 0x01 true or 0x00 false
}

// NewConfigurationRaw creates a ConfigurationRaw object from a Configuration.
func NewConfigurationRaw(c *Configuration) *ConfigurationRaw {
	if c == nil {
		return nil
	}
	cr := new(ConfigurationRaw)
	cr.SelectionMask = c.SelectionMask
	cr.Direction = c.Direction
	cr.Value = misc.BoolToUint8(c.Value)
	return cr
}

// Configurations is the return type for the state of the pins.
type Configurations struct {
	DirectionMask uint8
	ValueMask     uint8
}

// FromPacket creates a Configurations from a packet.
func (c *Configurations) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(c, p); err != nil {
		return err
	}
	return p.Payload.Decode(c)
}

// String fullfill the stringer interface.
func (c *Configurations) String() string {
	txt := "Configurations "
	if c == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("[Direction Mask: %d (%s), Value Mask: %d (%s)]",
			c.DirectionMask, misc.MaskToString(c.DirectionMask, 4, true),
			c.ValueMask, misc.MaskToString(c.ValueMask, 4, true))
	}
	return txt
}

// Copy creates a copy of the content.
func (c *Configurations) Copy() device.Resulter {
	if c == nil {
		return nil
	}
	return &Configurations{
		DirectionMask: c.DirectionMask,
		ValueMask:     c.ValueMask}
}
