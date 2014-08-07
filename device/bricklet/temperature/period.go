// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package temperature

import (
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
)

// SetTemperatureCallbackPeriod creates the subscriber to set the callback period.
// Default value is 0. A value of 0 deactivates the peridicall callbacks.
func SetTemperatureCallbackPeriod(id string, uid uint32, pe *device.Period, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SetTemperatureCallbackPeriod"),
		Fid:        function_set_temperature_callback_period,
		Uid:        uid,
		Data:       pe,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SetTemperatureCallbackPeriodFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetTemperatureCallbackPeriodFuture(brick *bricker.Bricker, connectorname string, uid uint32, pe *device.Period) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetTemperatureCallbackPeriod("settemperaturecallbackperiodfuture"+device.GenId(), uid, pe,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetTemperatureCallbackPeriod creates a subsctiber to get the callback period value.
func GetTemperatureCallbackPeriod(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetTemperatureCallbackPeriod"),
		Fid:        function_get_temperature_callback_period,
		Uid:        uid,
		Result:     &device.Period{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetTemperatureCallbackPeriodFuture is a future pattern version for a synchronized call of the subsctiber.
// If an error occur, the result is nil.
func GetTemperatureCallbackPeriodFuture(brick *bricker.Bricker, connectorname string, uid uint32) *device.Period {
	future := make(chan *device.Period)
	defer close(future)
	sub := GetTemperatureCallbackPeriod("gettemperaturecallbackperiodfuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *device.Period = nil
			if err == nil {
				if value, ok := r.(*device.Period); ok {
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

// TemperaturePeriod creates a subscriber for the periodical temperature callback.
func TemperaturePeriod(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "TemperaturePeriod"),
		Fid:        callback_temperature,
		Uid:        uid,
		Result:     &Temperature{},
		Handler:    handler,
		IsCallback: true,
		WithPacket: false}.CreateDevice()
}
