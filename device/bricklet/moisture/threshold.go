// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package moisture

import (
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
)

// SetMoistureCallbackThreshold creates the subscriber to set the callback thresold.
// Default value is ('x', 0, 0).
func SetMoistureCallbackThreshold(id string, uid uint32, t *device.Threshold16, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SetMoistureCallbackThreshold"),
		Fid:        function_set_moisture_callback_threshold,
		Uid:        uid,
		Data:       t,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SetTemperatureCallbackThresholdFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetMoistureCallbackThresholdFuture(brick *bricker.Bricker, connectorname string, uid uint32, t *device.Threshold16) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetMoistureCallbackThreshold("setmoisturecallbackthresholdfuture"+device.GenId(), uid, t,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetMoistureCallbackThreshold creates the subscriber to get the callback thresold.
func GetMoistureCallbackThreshold(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetMoistureCallbackThreshold"),
		Fid:        function_get_moisture_callback_threshold,
		Uid:        uid,
		Result:     &device.Threshold16{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetMoistureCallbackThresholdFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func GetMoistureCallbackThresholdFuture(brick *bricker.Bricker, connectorname string, uid uint32) *device.Threshold16 {
	future := make(chan *device.Threshold16)
	defer close(future)
	sub := GetMoistureCallbackThreshold("getmoisturecallbackthresholdfuture"+device.GenId(), uid,
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

// MoistureReached creates a subscriber for the theshold triggered temperature callback.
func MoistureReached(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "MoistureReached"),
		Fid:        callback_moisture_reached,
		Uid:        uid,
		Result:     &Moisture{},
		Handler:    handler,
		IsCallback: true,
		WithPacket: false}.CreateDevice()
}
