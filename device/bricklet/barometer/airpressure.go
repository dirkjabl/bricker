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

// GetAirPressure creates the subscriber to get the air pressure value once.
func GetAirPressure(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetAirPressure"),
		Fid:        function_get_air_pressure,
		Uid:        uid,
		Result:     &AirPressure{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetAirPressureFuture is a future pattern version for a synchronized all of the subscriber.
// If an error occur, the result is nil.
func GetAirPressureFuture(brick *bricker.Bricker, connectorname string, uid uint32) *AirPressure {
	future := make(chan *AirPressure)
	defer close(future)
	sub := GetAirPressure("getairpressurefuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *AirPressure = nil
			if err == nil {
				if value, ok := r.(*AirPressure); ok {
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
AirPressure is the type for the air pressure value.

The value has a range of 10000 to 1200000 and is given in mbar/1000
*/
type AirPressure struct {
	Value int32
}

// FromPacket converts from packet into AirPressure.
func (a *AirPressure) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(a, p); err != nil {
		return err
	}
	return p.Payload.Decode(a)
}

// fullfill the stringer interface.
func (a *AirPressure) String() string {
	txt := "Air Pressure "
	if a == nil {
		txt += "[]"
	} else {
		fmt.Sprintf("[Value: %d, Air Pressure: %04.03f mbar]", a.Value, a.Float64())
	}
	return txt
}

// Float64 converts the int32 value (mbar/1000) into float64 (mbar)
func (a *AirPressure) Float64() float64 {
	f := float64(a.Value) / 1000.0
	return f
}

// Float32 converts the int32 value (mbar/1000) into float32 (mbar)
func (a *AirPressure) Float32() float32 {
	f := float32(a.Value) / 1000.0
	return f
}
