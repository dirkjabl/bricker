// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Collection of subscriber for the Moisture Bricklet.
package moisture

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
)

const (
	function_get_moisture_value              = uint8(1)
	function_set_moving_average              = uint8(10)
	function_get_moving_average              = uint8(11)
	function_set_moisture_callback_period    = uint8(2)
	function_get_moisture_callback_period    = uint8(3)
	function_set_moisture_callback_threshold = uint8(4)
	function_get_moisture_callback_threshold = uint8(5)
	function_set_debounce_period             = uint8(6)
	function_get_debounce_period             = uint8(7)
	callback_moisture                        = uint8(8)
	callback_moisture_reached                = uint8(9)
)

// GetMoistureValue creates a subscriber to get the moisture value.
func GetMoistureValue(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetMoistureValue"),
		Fid:        function_get_moisture_value,
		Uid:        uid,
		Result:     &Moisture{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetMoistureValueFuture is a future pattern version for a synchronized calll of the subscriber.
// If an error occur, the result is nil.
func GetMoistureValueFuture(brick *bricker.Bricker, connectorname string, uid uint32) *Moisture {
	future := make(chan *Moisture)
	defer close(future)
	sub := GetMoistureValue("getmoisturevaluefuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *Moisture = nil
			if err == nil {
				if value, ok := r.(*Moisture); ok {
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

// Moisture is the type of the moisture value.
// The value has a range of 0 to 4095.
// A small value means low moisture, a big value corresponds to much moisture.
type Moisture struct {
	Value uint16
}

// FromPacket create from a packet a moisture value.
func (m *Moisture) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(m, p); err != nil {
		return err
	}
	return p.Payload.Decode(m)
}

// String fullfill the stringer interface.
func (m *Moisture) String() string {
	txt := "Moisture "
	if m == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("[Value: %d]", m.Value)
	}
	return txt
}
