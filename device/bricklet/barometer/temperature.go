// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package barometer

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
)

// GetChipTemperature create the subscriber to get the chip temperature value.
func GetChipTemperature(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetChipTemperature"),
		Fid:        function_get_chip_temperature,
		Uid:        uid,
		Result:     &Temperature{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetAveragingFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func GetChipTemperatureFuture(brick bricker.Bricker, connectorname string, uid uint32) *Temperature {
	future := make(chan *Temperature)
	defer close(future)
	sub := GetChipTemperature("getchiptemperaturefuture"+device.GenId(), uid,
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

// Temperature type with a value 100/°C in a range between -4000 to 8500.
type Temperature struct {
	Value int16
}

// FromPacket creates from a packet a Temperature.
func (t *Temperature) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(t, p); err != nil {
		return err
	}
	return p.Payload.Decode(t)
}

// String fullfill the stringer interface.
func (t *Temperature) String() string {
	txt := "Temperature "
	if t == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("[Value: %d, Temperature: &02.02f °C]", t.Value, t.Float64())
	}
	return txt
}

// Copy creates a copy of the content.
func (t *Temperature) Copy() device.Resulter {
	if t == nil {
		return nil
	}
	return &Temperature{Value: t.Value}
}

// Float64 converts the temperature value to a float.
func (t *Temperature) Float64() float64 {
	f := float64(t.Value) / 100.00
	return f
}

// Float32 converts the temperature value to a float.
func (t *Temperature) Float32() float32 {
	f := float32(t.Value) / 100.00
	return f
}
