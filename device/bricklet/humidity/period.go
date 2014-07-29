// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package humidity

import (
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
	"github.com/dirkjabl/bricker/subscription"
	"github.com/dirkjabl/bricker/util/hash"
)

// SetHumidityCallbackPeriod creates the subscriber to set the callback period.
// Default value is 0. A value of 0 deactivates the periodical callbacks.
// HumidityPeriod is only triggered if the voltage has changed since the last triggering.
func SetHumidityCallbackPeriod(id string, uid uint32, pe *device.Period, handler func(device.Resulter, error)) *device.Device {
	fid := function_set_humidity_callback_period
	shcp := device.New(device.FallbackId(id, "SetHumidityCallbackPeriod"))
	p := packet.NewSimpleHeaderPayload(uid, fid, true, pe)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	shcp.SetSubscription(sub)
	shcp.SetResult(&device.EmptyResult{})
	shcp.SetHandler(handler)
	return shcp
}

// SetHumidityCallbackPeriodFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetHumidityCallbackPeriodFuture(brick *bricker.Bricker, connectorname string, uid uint32, pe *device.Period) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetHumidityCallbackPeriod("sethumiditycallbackperiodfuture"+device.GenId(), uid, pe,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetHumidityCallbackPeriod creates a subsctiber to get the callback period value.
func GetHumidityCallbackPeriod(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := function_get_humidity_callback_period
	ghcp := device.New(device.FallbackId(id, "GetHumidityCallbackPeriod"))
	p := packet.NewSimpleHeaderOnly(uid, fid, true)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	ghcp.SetSubscription(sub)
	ghcp.SetResult(&device.Period{})
	ghcp.SetHandler(handler)
	return ghcp
}

// GetHumidityCallbackPeriodFuture is a future pattern version for a synchronized call of the subsctiber.
// If an error occur, the result is nil.
func GetHumidityCallbackPeriodFuture(brick *bricker.Bricker, connectorname string, uid uint32) *device.Period {
	future := make(chan *device.Period)
	defer close(future)
	sub := GetHumidityCallbackPeriod("gethumiditycallbackperiodfuture"+device.GenId(), uid,
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

// SetAnalogValueCallbackPeriod creates the subscriber to set the callback period.
// Default value is 0. A value of 0 deactivates the periodical callbacks.
// AnalogValuePeriod is only triggered if the voltage has changed since the last triggering.
func SetAnalogValueCallbackPeriod(id string, uid uint32, pe *device.Period, handler func(device.Resulter, error)) *device.Device {
	fid := function_set_analog_value_callback_period
	savcp := device.New(device.FallbackId(id, "SetAnalogValueCallbackPeriod"))
	p := packet.NewSimpleHeaderPayload(uid, fid, true, pe)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	savcp.SetSubscription(sub)
	savcp.SetResult(&device.EmptyResult{})
	savcp.SetHandler(handler)
	return savcp
}

// SetAnalogValueCallbackPeriodFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetAnalogValueCallbackPeriodFuture(brick *bricker.Bricker, connectorname string, uid uint32, pe *device.Period) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetAnalogValueCallbackPeriod("setanalogvaluecallbackperiodfuture"+device.GenId(), uid, pe,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetAnalogValueCallbackPeriod creates a subsctiber to get the callback period value.
func GetAnalogValueCallbackPeriod(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := function_get_analog_value_callback_period
	gavcp := device.New(device.FallbackId(id, "GetAnalogValueCallbackPeriod"))
	p := packet.NewSimpleHeaderOnly(uid, fid, true)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	gavcp.SetSubscription(sub)
	gavcp.SetResult(&device.Period{})
	gavcp.SetHandler(handler)
	return gavcp
}

// GetAnalogValueCallbackPeriodFuture is a future pattern version for a synchronized call of the subsctiber.
// If an error occur, the result is nil.
func GetAnalogValueCallbackPeriodFuture(brick *bricker.Bricker, connectorname string, uid uint32) *device.Period {
	future := make(chan *device.Period)
	defer close(future)
	sub := GetAnalogValueCallbackPeriod("getanalogvaluecallbackperiodfuture"+device.GenId(), uid,
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

// HumidityPeriod creates a subscriber for the periodical humidity callback.
// Is only triggered if the voltage changed, since last triggering.
func HumidityPeriod(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := callback_humidity
	hp := device.New(device.FallbackId(id, "HumidityPeriod"))
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, nil, true)
	hp.SetSubscription(sub)
	hp.SetResult(&Humidity{})
	hp.SetHandler(handler)
	return hp
}

// AnalogValuePeriod creates a subscriber for the periodical analog value callback.
// Is only triggered if the value changed, since last triggering.
func AnalogValuePeriod(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := callback_analog_value
	avp := device.New(device.FallbackId(id, "AnalogValuePeriod"))
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, nil, true)
	avp.SetSubscription(sub)
	avp.SetResult(&AnalogValue{})
	avp.SetHandler(handler)
	return avp
}
