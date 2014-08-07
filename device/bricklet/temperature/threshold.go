// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package temperature

import (
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
)

// SetTemperatureCallbackThreshold creates the subscriber to set the callback thresold.
// Default value is ('x', 0, 0).
func SetTemperatureCallbackThreshold(id string, uid uint32, t *device.Threshold, handler func(device.Resulter, error)) *device.Device {
	return device.NewHeaderPayloadEmptyResult(device.FallbackId(id, "SetTemperatureCallbackThreshold"),
		uid, function_set_temperature_callback_threshold, false, t, handler)
}

// SetTemperatureCallbackThresholdFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetTemperatureCallbackThresholdFuture(brick *bricker.Bricker, connectorname string, uid uint32, t *device.Threshold) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetTemperatureCallbackThreshold("settemperaturecallbackthresholdfuture"+device.GenId(), uid, t,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetTemperatureCallbackThreshold creates the subscriber to get the callback thresold.
func GetTemperatureCallbackThreshold(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.NewHeaderOnlyWithResult(device.FallbackId(id, "GetTemperatureCallbackThreshold"),
		uid, function_get_temperature_callback_threshold, false, &device.Threshold{}, handler)
}

// GetTemperatureCallbackThresholdFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func GetTemperatureCallbackThresholdFuture(brick *bricker.Bricker, connectorname string, uid uint32) *device.Threshold {
	future := make(chan *device.Threshold)
	defer close(future)
	sub := GetTemperatureCallbackThreshold("gettemperaturecallbackthresholdfuture"+device.GenId(), uid,
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

// TemperatureReached creates a subscriber for the theshold triggered temperature callback.
func TemperatureReached(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.NewHeaderNoPacketWithResult(device.FallbackId(id, "TemperatureReached"),
		uid, callback_temperature_reached, true, &Temperature{}, handler)
}
