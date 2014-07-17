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

// GetAnalogValue creates A subscriber to return the raw 12-bit analog value.
// It is only useful, if you need the full resolution of the analog-to-digital converter.
// Please use normaly GetVoltage.
func GetAnalogValue(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := function_get_analog_value
	gav := device.New(device.FallbackId(id, "GetAnalogValue"))
	p := packet.NewSimpleHeaderOnly(uid, fid, true)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	gav.SetSubscription(sub)
	gav.SetResult(&AnalogValue{})
	gav.SetHandler(handler)
	return gav
}

// GetAnalogValueFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func GetAnalogValueFuture(brick bricker.Bricker, connectorname string, uid uint32) *AnalogValue {
	future := make(chan *AnalogValue)
	defer close(future)
	sub := GetVoltage("getanalogvaluefuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *AnalogValue = nil
			if err == nil {
				if value, ok := r.(*AnalogValue); ok {
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

/*
AnalogValue is a type for the 12-bit analog-to-digial converter value.
It can have values between 0 and 4095. This is the raw unfiltered analog value.
Please see the original documentation
http://www.tinkerforge.com/en/doc/Software/Bricklets/AnalogIn_Bricklet_TCPIP.html#advanced-functions
for more information.
*/
type AnalogValue struct {
	Value uint16
}

// FromPacket creates from a packet a analog value.
func (av *AnalogValue) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(av, p); err != nil {
		return err
	}
	return p.Payload.Decode(av)
}

// String fullfill the stringer interface.
func (av *AnalogValue) String() string {
	txt := "AnalogValue "
	if av == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("[Value: %d]", av.Value)
	}
	return txt
}
