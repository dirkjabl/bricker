// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package analogin

import (
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
	"github.com/dirkjabl/bricker/subscription"
	"github.com/dirkjabl/bricker/util/hash"
)

// SetVoltageCallbackPeriod creates the subscriber to set the callback period.
// Default value is 0. A value of 0 deactivates the periodical callbacks.
// VoltagePeriod is only triggered if the voltage has changed since the last triggering.
func SetVoltageCallbackPeriod(id string, uid uint32, pe *device.Period, handler func(device.Resulter, error)) *device.Device {
	fid := function_set_voltage_callback_period
	svcp := device.New(device.FallbackId(id, "SetVoltageCallbackPeriod"))
	p := packet.NewSimpleHeaderPayload(uid, fid, true, pe)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	svcp.SetSubscription(sub)
	svcp.SetResult(&device.EmptyResult{})
	svcp.SetHandler(handler)
	return svcp
}

// SetVoltageCallbackPeriodFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetVoltageCallbackPeriodFuture(brick *bricker.Bricker, connectorname string, uid uint32, pe *device.Period) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetVoltageCallbackPeriod("setvoltagecallbackperiodfuture"+device.GenId(), uid, pe,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetVoltageCallbackPeriod creates a subsctiber to get the callback period value.
func GetVoltageCallbackPeriod(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := function_get_voltage_callback_period
	gvcp := device.New(device.FallbackId(id, "GetVoltageCallbackPeriod"))
	p := packet.NewSimpleHeaderOnly(uid, fid, true)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	gvcp.SetSubscription(sub)
	gvcp.SetResult(&device.Period{})
	gvcp.SetHandler(handler)
	return gvcp
}

// GetVoltageCallbackPeriodFuture is a future pattern version for a synchronized call of the subsctiber.
// If an error occur, the result is nil.
func GetVoltageCallbackPeriodFuture(brick *bricker.Bricker, connectorname string, uid uint32) *device.Period {
	future := make(chan *device.Period)
	defer close(future)
	sub := GetVoltageCallbackPeriod("getvoltagecallbackperiodfuture"+device.GenId(), uid,
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

// VoltagePeriod creates a subscriber for the periodical voltage callback.
// Is only triggered if the voltage changed, since last triggering.
func VoltagePeriod(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := callback_voltage
	vp := device.New(device.FallbackId(id, "VoltagePeriod"))
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, nil, true)
	vp.SetSubscription(sub)
	vp.SetResult(&Voltage{})
	vp.SetHandler(handler)
	return vp
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
