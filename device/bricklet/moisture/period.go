// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package moisture

import (
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
)

// SetMoistureCallbackPeriod creates the subscriber to set the callback period.
// Default value is 0. A value of 0 deactivates the peridicall callbacks.
func SetMoistureCallbackPeriod(id string, uid uint32, pe *device.Period, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SetMoistureCallbackPeriod"),
		Fid:        function_set_moisture_callback_period,
		Uid:        uid,
		Data:       pe,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SetMoistureCallbackPeriodFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetMoistureCallbackPeriodFuture(brick *bricker.Bricker, connectorname string, uid uint32, pe *device.Period) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetMoistureCallbackPeriod("setmoisturecallbackperiodfuture"+device.GenId(), uid, pe,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetMoistureCallbackPeriod creates a subsctiber to get the callback period value.
func GetMoistureCallbackPeriod(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetMoistureCallbackPeriod"),
		Fid:        function_get_moisture_callback_period,
		Uid:        uid,
		Result:     &device.Period{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetMoistureCallbackPeriodFuture is a future pattern version for a synchronized call of the subsctiber.
// If an error occur, the result is nil.
func GetMoistureCallbackPeriodFuture(brick *bricker.Bricker, connectorname string, uid uint32) *device.Period {
	future := make(chan *device.Period)
	defer close(future)
	sub := GetMoistureCallbackPeriod("getmoisturecallbackperiodfuture"+device.GenId(), uid,
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

// MoisturePeriod creates a subscriber for the periodical moisture callback.
func MoisturePeriod(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "MoisturePeriod"),
		Fid:        callback_moisture,
		Uid:        uid,
		Result:     &Moisture{},
		Handler:    handler,
		IsCallback: true,
		WithPacket: false}.CreateDevice()
}
