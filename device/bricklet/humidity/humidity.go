// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Collection of subscriber for the Humidity Bricklet.
package humidity

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
)

const (
	function_get_humidity                        = uint8(1)
	function_get_analog_value                    = uint8(2)
	function_set_humidity_callback_period        = uint8(3)
	function_get_humidity_callback_period        = uint8(4)
	function_set_analog_value_callback_period    = uint8(5)
	function_get_analog_value_callback_period    = uint8(6)
	function_set_humidity_callback_threshold     = uint8(7)
	function_get_humidity_callback_threshold     = uint8(8)
	function_set_analog_value_callback_threshold = uint8(9)
	function_get_analog_value_callback_threshold = uint8(10)
	function_set_debounce_period                 = uint8(11)
	function_get_debounce_period                 = uint8(12)
	callback_humidity                            = uint8(13)
	callback_analog_value                        = uint8(14)
	callback_humidity_reached                    = uint8(15)
	callback_analog_value_reached                = uint8(16)
)

// GetHumidity creates a subsriber to read out the humidity sensor.
// Use the callbacks to get periodical the value.
func GetHumidity(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetHumidity"),
		Fid:        function_get_humidity,
		Uid:        uid,
		Result:     &Humidity{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetHumidityFuture is a future pattern version for a synchronized calll of the subscriber.
// If an error occur, the result is nil.
func GetHumidityFuture(brick *bricker.Bricker, connectorname string, uid uint32) *Humidity {
	future := make(chan *Humidity)
	defer close(future)
	sub := GetHumidity("gethumidityfuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *Humidity = nil
			if err == nil {
				if value, ok := r.(*Humidity); ok {
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
 Humidity type.
The value has a range of 0 to 1000 and is given in %RH/10 (Relative Humidity),
i.e. a value of 421 means that a humidity of 42.1 %RH is measured.
*/
type Humidity struct {
	Value uint16
}

// FromPacket create from a packet a Humidity.
func (h *Humidity) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(h, p); err != nil {
		return err
	}
	return p.Payload.Decode(h)
}

// String fullfill the stringer interface.
func (h *Humidity) String() string {
	txt := "Humidity "
	if h == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("[Value: %d, Humidity: %02.02f %RH]", h.Value, h.Float64())
	}
	return txt
}

// Float64 converts the humidity value from int16 to float64.
func (h *Humidity) Float64() float64 {
	f := float64(h.Value) / 10.00
	return f
}

// Float32 converts the humidity value from int16 to float32.
func (h *Humidity) Float32() float32 {
	f := float32(h.Value) / 10.00
	return f
}
