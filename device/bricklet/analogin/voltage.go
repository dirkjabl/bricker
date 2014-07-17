// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package analogin

import (
	"bricker"
	"bricker/device"
	"bricker/net/packet"
	"bricker/subscription"
	"bricker/util/hash"
	"fmt"
)

// GetVoltage creates A subscriber to return the actual voltage (mV).
func GetVoltage(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := function_get_voltage
	gv := device.New(device.FallbackId(id, "GetVoltage"))
	p := packet.NewSimpleHeaderOnly(uid, fid, true)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	gv.SetSubscription(sub)
	gv.SetResult(&Voltage{})
	gv.SetHandler(handler)
	return gv
}

// GetVoltageFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func GetVoltageFuture(brick bricker.Bricker, connectorname string, uid uint32) *Voltage {
	future := make(chan *Voltage)
	defer close(future)
	sub := GetVoltage("getvoltagefuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *Voltage = nil
			if err == nil {
				if value, ok := r.(*Voltage); ok {
					v = value
				}
			}
			future <- v
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return nil
	}
	v := <-future
	return v
}

// Voltage result type
type Voltage struct {
	Value uint16 // mV
}

// FromPacket creates from a packet a Voltage.
func (v *Voltage) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(v, p); err != nil {
		return err
	}
	return p.Payload.Decode(v)
}

// String fullfill the stringer interface.
func (v *Voltage) String() string {
	txt := "Voltage "
	if v == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("[Value: %d mV]", v.Value)
	}
	return txt
}
