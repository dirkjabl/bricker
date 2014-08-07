// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package temperature

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
)

const (
	I2CModeFast = 0 // 400kHz, default
	I2CModeSlow = 1 // 100kHz
)

// I2C mode type.
type I2CMode struct {
	Value uint8
}

// SetI2CMode creates the subscriber to set the I2C mode.
func SetI2CMode(id string, uid uint32, m *I2CMode, handler func(device.Resulter, error)) *device.Device {
	return device.NewHeaderPayloadEmptyResult(device.FallbackId(id, "SetI2CMode"), uid,
		function_set_i2c_mode, false, m, handler)
}

// SetI2CModeFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetI2CModeFuture(brick *bricker.Bricker, connectorname string, uid uint32, m *I2CMode) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetI2CMode("seti2cmodefuture"+device.GenId(), uid, m,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetI2CMode creates the subscriber to get the I2C mode.
func GetI2CMode(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.NewHeaderOnlyWithResult(device.FallbackId(id, "GetI2CMode"), uid,
		function_get_i2c_mode, false, &I2CMode{}, handler)
}

// GetI2CModeFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func GetI2CModeFuture(brick *bricker.Bricker, connectorname string, uid uint32) *I2CMode {
	future := make(chan *I2CMode)
	defer close(future)
	sub := GetI2CMode("geti2cmodefuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *I2CMode = nil
			if err == nil {
				if value, ok := r.(*I2CMode); ok {
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

// FromPacket converts from packet to Mode.
func (m *I2CMode) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(m, p); err != nil {
		return err
	}
	return p.Payload.Decode(m)
}

// Name convert the Mode to a readable string.
func (m *I2CMode) Name() string {
	switch m.Value {
	case I2CModeFast:
		return "Fast (400kHz, Standard)"
	case I2CModeSlow:
		return "Slow (100kHz)"
	default:
		return "Unknown"
	}
}

// String fullfill the stringer interface.
func (m *I2CMode) String() string {
	return fmt.Sprintf("I2C Mode [Value: %d Name: %s]", m.Value, m.Name())
}
