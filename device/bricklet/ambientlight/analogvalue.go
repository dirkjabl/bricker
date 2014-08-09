// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ambientlight

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
)

/*
GetAnalogValue creates A subscriber to return the raw 12-bit analog value (0 up to 4095).
It is only useful, if you need the full resolution of the analog-to-digital converter.
Please use normaly GetIlluminance.
*/
func GetAnalogValue(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetAnalogValue"),
		Fid:        function_get_analog_value,
		Uid:        uid,
		Result:     &AnalogValue{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetAnalogValueFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func GetAnalogValueFuture(brick bricker.Bricker, connectorname string, uid uint32) *AnalogValue {
	future := make(chan *AnalogValue)
	defer close(future)
	sub := GetAnalogValue("getanalogvaluefuture"+device.GenId(), uid,
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
http://www.tinkerforge.com/en/doc/Software/Bricklets/AmbientLight_Bricklet_TCPIP.html#BrickletAmbientLight.get_analog_value
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
