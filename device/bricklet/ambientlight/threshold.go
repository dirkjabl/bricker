// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ambientlight

import (
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
)

// SetIlluminanceCallbackThreshold creates the subscriber to set the callback thresold.
// Default value is ('x', 0, 0).
func SetIlluminanceCallbackThreshold(id string, uid uint32, t *device.Threshold, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SetIlluminanceCallbackThreshold"),
		Fid:        function_set_illuminance_callback_threshold,
		Uid:        uid,
		Data:       t,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SetIlluminanceCallbackThresholdFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetIlluminanceCallbackThresholdFuture(brick *bricker.Bricker, connectorname string, uid uint32, t *device.Threshold) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetIlluminanceCallbackThreshold("setilluminancecallbackthresholdfuture"+device.GenId(), uid, t,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetIlluminanceCallbackThreshold creates the subscriber to get the callback thresold.
func GetIlluminanceCallbackThreshold(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetIlluminanceCallbackThreshold"),
		Fid:        function_get_illuminance_callback_threshold,
		Uid:        uid,
		Result:     &device.Threshold{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetIlluminanceCallbackThresholdFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func GetIlluminanceCallbackThresholdFuture(brick *bricker.Bricker, connectorname string, uid uint32) *device.Threshold {
	future := make(chan *device.Threshold)
	defer close(future)
	sub := GetIlluminanceCallbackThreshold("getilluminancecallbackthresholdfuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *device.Threshold = nil
			if err == nil {
				if value, ok := r.(*device.Threshold); ok {
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
func SetAnalogValueCallbackThreshold(id string, uid uint32, t *device.Threshold, handler func(device.Resulter, error)) *device.Device {
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
func SetAnalogValueCallbackThresholdFuture(brick *bricker.Bricker, connectorname string, uid uint32, t *device.Threshold) bool {
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
		Result:     &device.Threshold{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetAnalogValueCallbackThresholdFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func GetAnalogValueCallbackThresholdFuture(brick *bricker.Bricker, connectorname string, uid uint32) *device.Threshold {
	future := make(chan *device.Threshold)
	defer close(future)
	sub := GetAnalogValueCallbackThreshold("getanalogvaluecallbackthresholdfuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *device.Threshold = nil
			if err == nil {
				if value, ok := r.(*device.Threshold); ok {
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

// IlluminanceReached creates a subscriber for the theshold triggered voltage callback.
func IlluminanceReached(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "IlluminanceReached"),
		Fid:        callback_illuminance_reached,
		Uid:        uid,
		Result:     &Illuminance{},
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
