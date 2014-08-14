// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package barometer

import (
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
)

// SetAirPressureCallbackThreshold creates the subscriber to set the callback threshold.
// Default value is ('x', 0, 0).
func SetAirPressureCallbackThreshold(id string, uid uint32, t *device.Threshold32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SetAirPressureCallbackThreshold"),
		Fid:        function_set_air_pressure_callback_threshold,
		Uid:        uid,
		Data:       t,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SetAirPressureCallbackThresholdFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetAirPressureCallbackThresholdFuture(brick *bricker.Bricker, connectorname string, uid uint32, t *device.Threshold32) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetAirPressureCallbackThreshold("setairpressurecallbackthresholdfuture"+device.GenId(), uid, t,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetAirPressureCallbackThreshold creates the subscriber to get the callback threshold.
func GetAirPressureCallbackThreshold(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetAirPressureCallbackThreshold"),
		Fid:        function_get_air_pressure_callback_threshold,
		Uid:        uid,
		Result:     &device.Threshold32{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetAirPressureCallbackThresholdFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func GetAirPressureCallbackThresholdFuture(brick *bricker.Bricker, connectorname string, uid uint32) *device.Threshold32 {
	future := make(chan *device.Threshold32)
	defer close(future)
	sub := GetAirPressureCallbackThreshold("getairpressurecallbackthresholdfuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *device.Threshold32 = nil
			if err == nil {
				if value, ok := r.(*device.Threshold32); ok {
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

// SetAltitudeCallbackThreshold creates the subscriber to set the callback threshold.
// Default value is ('x', 0, 0).
func SetAltitudeCallbackThreshold(id string, uid uint32, t *device.Threshold32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SetAltitudeCallbackThreshold"),
		Fid:        function_set_altitude_callback_threshold,
		Uid:        uid,
		Data:       t,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SetAltitudeCallbackThresholdFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetAltitudeCallbackThresholdFuture(brick *bricker.Bricker, connectorname string, uid uint32, t *device.Threshold32) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetAltitudeCallbackThreshold("setaltitudecallbackthresholdfuture"+device.GenId(), uid, t,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetAltitudeCallbackThreshold creates the subscriber to get the callback threshold.
func GetAltitudeCallbackThreshold(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetAltitudeCallbackThreshold"),
		Fid:        function_get_altitude_callback_threshold,
		Uid:        uid,
		Result:     &device.Threshold32{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetAltitudeCallbackThresholdFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func GetAltitudeCallbackThresholdFuture(brick *bricker.Bricker, connectorname string, uid uint32) *device.Threshold32 {
	future := make(chan *device.Threshold32)
	defer close(future)
	sub := GetAltitudeCallbackThreshold("getaltitudecallbackthresholdfuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *device.Threshold32 = nil
			if err == nil {
				if value, ok := r.(*device.Threshold32); ok {
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

// AirPressureReached creates a subscriber for the theshold triggered air pressure callback.
func AirPressureReached(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "AirPressureReached"),
		Fid:        callback_air_pressure_reached,
		Uid:        uid,
		Result:     &AirPressure{},
		Handler:    handler,
		IsCallback: true,
		WithPacket: false}.CreateDevice()
}

// AltitudeReached creates a subscriber for the theshold triggered altitude callback.
func AltitudeReached(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "AltitudeReached"),
		Fid:        callback_altitude_reached,
		Uid:        uid,
		Result:     &Altitude{},
		Handler:    handler,
		IsCallback: true,
		WithPacket: false}.CreateDevice()
}
