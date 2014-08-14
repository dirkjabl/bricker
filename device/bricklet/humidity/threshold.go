// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package humidity

import (
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
)

// SetHumidityCallbackThreshold creates the subscriber to set the callback thresold.
// Default value is ('x', 0, 0).
func SetHumidityCallbackThreshold(id string, uid uint32, t *device.Threshold16, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SetHumidityCallbackThreshold"),
		Fid:        function_set_humidity_callback_threshold,
		Uid:        uid,
		Data:       t,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SetHumidityCallbackThresholdFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetHumidityCallbackThresholdFuture(brick *bricker.Bricker, connectorname string, uid uint32, t *device.Threshold16) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetHumidityCallbackThreshold("sethumiditycallbackthresholdfuture"+device.GenId(), uid, t,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetHumidityCallbackThreshold creates the subscriber to get the callback thresold.
func GetHumidityCallbackThreshold(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetHumidityCallbackThreshold"),
		Fid:        function_get_humidity_callback_threshold,
		Uid:        uid,
		Result:     &device.Threshold16{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetHumidityCallbackThresholdFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func GetHumidityCallbackThresholdFuture(brick *bricker.Bricker, connectorname string, uid uint32) *device.Threshold16 {
	future := make(chan *device.Threshold16)
	defer close(future)
	sub := GetHumidityCallbackThreshold("gethumiditycallbackthresholdfuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *device.Threshold16 = nil
			if err == nil {
				if value, ok := r.(*device.Threshold16); ok {
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

// SetAnalogValueCallbackThreshold creates the subscriber to set the callback thresold.
// Default value is ('x', 0, 0).
func SetAnalogValueCallbackThreshold(id string, uid uint32, t *device.Threshold16, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SetAnalogValueCallbackThreshold"),
		Fid:        function_set_analog_value_callback_threshold,
		Uid:        uid,
		Data:       t,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SetAnalogValueCallbackThresholdFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetAnalogValueCallbackThresholdFuture(brick *bricker.Bricker, connectorname string, uid uint32, t *device.Threshold16) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetAnalogValueCallbackThreshold("setanalogvaluecallbackthresholdfuture"+device.GenId(), uid, t,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetAnalogValueCallbackThreshold creates the subscriber to get the callback thresold.
func GetAnalogValueCallbackThreshold(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetAnalogValueCallbackThreshold"),
		Fid:        function_get_analog_value_callback_threshold,
		Uid:        uid,
		Result:     &device.Threshold16{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetAnalogValueCallbackThresholdFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func GetAnalogValueCallbackThresholdFuture(brick *bricker.Bricker, connectorname string, uid uint32) *device.Threshold16 {
	future := make(chan *device.Threshold16)
	defer close(future)
	sub := GetAnalogValueCallbackThreshold("getanalogvaluecallbackthresholdfuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *device.Threshold16 = nil
			if err == nil {
				if value, ok := r.(*device.Threshold16); ok {
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

// HumidityReached creates a subscriber for the theshold triggered voltage callback.
func HumidityReached(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "HumidityReached"),
		Fid:        callback_humidity_reached,
		Uid:        uid,
		Result:     &Humidity{},
		Handler:    handler,
		IsCallback: true,
		WithPacket: false}.CreateDevice()
}

// AnalogValueReached creates a subscriber for the theshold triggered voltage callback.
func AnalogValueReached(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "AnalogValueReached"),
		Fid:        callback_analog_value_reached,
		Uid:        uid,
		Result:     &AnalogValue{},
		Handler:    handler,
		IsCallback: true,
		WithPacket: false}.CreateDevice()
}
