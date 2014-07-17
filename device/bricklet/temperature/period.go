// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package temperature

import (
	"bricker"
	"bricker/device"
	"bricker/net/packet"
	"bricker/subscription"
	"bricker/util/hash"
)

// SetTemperatureCallbackPeriod creates the subscriber to set the callback period.
// Default value is 0. A value of 0 deactivates the peridicall callbacks.
func SetTemperatureCallbackPeriod(id string, uid uint32, pe *device.Period, handler func(device.Resulter, error)) *device.Device {
	fid := function_set_temperature_callback_period
	stcp := device.New(device.FallbackId(id, "SetTemperatureCallbackPeriod"))
	p := packet.NewSimpleHeaderPayload(uid, fid, true, pe)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	stcp.SetSubscription(sub)
	stcp.SetResult(&device.EmptyResult{})
	stcp.SetHandler(handler)
	return stcp
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
	fid := function_get_temperature_callback_period
	gtcp := device.New(device.FallbackId(id, "GetTemperatureCallbackPeriod"))
	p := packet.NewSimpleHeaderOnly(uid, fid, true)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	gtcp.SetSubscription(sub)
	gtcp.SetResult(&device.Period{})
	gtcp.SetHandler(handler)
	return gtcp
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
	fid := callback_temperature
	tp := device.New(device.FallbackId(id, "TemperaturePeriod"))
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, nil, true)
	tp.SetSubscription(sub)
	tp.SetResult(&Temperature{})
	tp.SetHandler(handler)
	return tp
}
