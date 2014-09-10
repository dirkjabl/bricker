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

// GetAltitude creates the subscriber to get the altitude value.
func GetAltitude(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetAltitude"),
		Fid:        function_get_altitude,
		Uid:        uid,
		Result:     &Altitude{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetAltitudeFuture is a future pattern version for a synchronized all of the subscriber.
// If an error occur, the result is nil.
func GetAltitudeFuture(brick *bricker.Bricker, connectorname string, uid uint32) *Altitude {
	future := make(chan *Altitude)
	defer close(future)
	sub := GetAltitude("getaltitudefuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *Altitude = nil
			if err == nil {
				if value, ok := r.(*Altitude); ok {
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
Altitude is a type for the altitude value.

The value is given in cm and is calculated based on the difference
between the current air pressure and the reference air pressure
that can be set with SetReferenceAirPressure.
*/
type Altitude struct {
	Value int32
}

// FromPacket converts from packet into Altitude.
func (a *Altitude) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(a, p); err != nil {
		return err
	}
	return p.Payload.Decode(a)
}

// String fullfill the stringer interface.
func (a *Altitude) String() string {
	txt := "Altitude "
	if a == nil {
		txt += "[nil]"
	} else {
		fmt.Sprintf("[Value: %d cm]", a.Value)
	}
	return txt
}

// Copy creates a copy of the content.
func (a *Altitude) Copy() device.Resulter {
	if a == nil {
		return nil
	}
	return &Altitude{Value: a.Value}
}
