// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Collection of subscriber for the Temperature Bricklet.
package temperature

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
)

const (
	function_get_temperature                    = uint8(1)
	function_set_i2c_mode                       = uint8(10)
	function_get_i2c_mode                       = uint8(11)
	function_set_temperature_callback_period    = uint8(2)
	function_get_temperature_callback_period    = uint8(3)
	function_set_temperature_callback_threshold = uint8(4)
	function_get_temperature_callback_threshold = uint8(5)
	function_set_debounce_period                = uint8(6)
	function_get_debounce_period                = uint8(7)
	callback_temperature                        = uint8(8)
	callback_temperature_reached                = uint8(9)
)

// GetTemperature creates a subscriber for getting the actual tempreture.
func GetTemperature(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetTemperature"),
		Fid:        function_get_temperature,
		Uid:        uid,
		Result:     &Temperature{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetTemperatureFuture is a future pattern version for a synchronized calll of the subscriber.
// If an error occur, the result is nil.
func GetTemperatureFuture(brick *bricker.Bricker, connectorname string, uid uint32) *Temperature {
	future := make(chan *Temperature)
	defer close(future)
	sub := GetTemperature("gettemperaturefuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *Temperature = nil
			if err == nil {
				if value, ok := r.(*Temperature); ok {
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

// Temperature type for a single temperature.
// From -2500 up to 8500 as °C/100, means a temperature of 1234 is 12.34 °C.
type Temperature struct {
	Value int16
}

// FromPacket create from a packet a Temperature.
func (t *Temperature) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(t, p); err != nil {
		return err
	}
	return p.Payload.Decode(t)
}

// Float64 convert the temperature from int16 to float64.
func (t *Temperature) Float64() float64 {
	f := float64(t.Value) / 100.00
	return f
}

// Float32 convert the temperature from int16 to float32.
func (t *Temperature) Float32() float32 {
	f := float32(t.Value) / 100.0
	return f
}

// String fullfill the stringer interface.
func (t *Temperature) String() string {
	txt := "Temperature "
	if t == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("[Value: %d, Temperature: %02.02f°C]", t.Value, t.Float64())
	}
	return txt
}
