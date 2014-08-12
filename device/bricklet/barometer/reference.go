// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package barometer

import (
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
)

/*
SetReferenceAirPressure creates the subscriber to set the reference air pressure.
Setting the reference to the current air pressure results in a calculated altitude of 0cm.
Passing 0 is a shortcut for passing the current air pressure as reference.

Well known reference values are the Q codes QNH and QFE used in aviation.

The default value is 1013.25mbar.
*/
func SetReferenceAirPressure(id string, uid uint32, a *AirPressure, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, ""),
		Fid:        function_set_reference_air_pressure,
		Uid:        uid,
		Data:       a,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SetReferenceAirPressureFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetReferenceAirPressureFuture(brick *bricker.Bricker, connectorname string, uid uint32, a *AirPressure) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetReferenceAirPressure("setreferenceairpressurefuture"+device.GenId(), uid, a,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetReferenceAirPressure creates the subscriber to get reference air pressure.
func GetReferenceAirPressure(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetReferenceAirPressure"),
		Fid:        function_get_reference_air_pressure,
		Uid:        uid,
		Result:     &AirPressure{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetReferenceAirPressureFuture is a future pattern version for a synchronized all of the subscriber.
// If an error occur, the result is nil.
func GetReferenceAirPressureFuture(brick *bricker.Bricker, connectorname string, uid uint32) *AirPressure {
	future := make(chan *AirPressure)
	defer close(future)
	sub := GetReferenceAirPressure("getreferenceairpressure"+device.GenId(), uid,
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
