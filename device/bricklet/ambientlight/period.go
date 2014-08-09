// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ambientlight

import (
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
)

// SetIlluminanceCallbackPeriod creates the subscriber to set the callback period.
// Default value is 0. A value of 0 deactivates the periodical callbacks.
// IlluminancePeriod is only triggered if the illuminance has changed since the last triggering.
func SetIlluminanceCallbackPeriod(id string, uid uint32, pe *device.Period, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SetIlluminanceCallbackPeriod"),
		Fid:        function_set_illuminance_callback_period,
		Uid:        uid,
		Data:       pe,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SetIlluminanceCallbackPeriodFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetIlluminanceCallbackPeriodFuture(brick *bricker.Bricker, connectorname string, uid uint32, pe *device.Period) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetIlluminanceCallbackPeriod("setilluminancecallbackperiodfuture"+device.GenId(), uid, pe,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetIlluminanceCallbackPeriod creates a subsctiber to get the callback period value.
func GetIlluminanceCallbackPeriod(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetIlluminanceCallbackPeriod"),
		Fid:        function_get_illuminance_callback_period,
		Uid:        uid,
		Result:     &device.Period{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetIlluminanceCallbackPeriodFuture is a future pattern version for a synchronized call of the subsctiber.
// If an error occur, the result is nil.
func GetIlluminanceCallbackPeriodFuture(brick *bricker.Bricker, connectorname string, uid uint32) *device.Period {
	future := make(chan *device.Period)
	defer close(future)
	sub := GetIlluminanceCallbackPeriod("getilluminancecallbackperiodfuture"+device.GenId(), uid,
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
	return device.Generator{
		Id:         device.FallbackId(id, "SetAnalogValueCallbackPeriod"),
		Fid:        function_set_analog_value_callback_period,
		Uid:        uid,
		Data:       pe,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
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
	return device.Generator{
		Id:         device.FallbackId(id, "GetAnalogValueCallbackPeriod"),
		Fid:        function_get_analog_value_callback_period,
		Uid:        uid,
		Result:     &device.Period{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
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

// IlluminancePeriod creates a subscriber for the periodical illuminance callback.
// Is only triggered if the voltage changed, since last triggering.
func IlluminancePeriod(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "IlluminancePeriod"),
		Fid:        callback_illuminance,
		Uid:        uid,
		Result:     &Illuminance{},
		Handler:    handler,
		IsCallback: true,
		WithPacket: false}.CreateDevice()
}

// AnalogValuePeriod creates a subscriber for the periodical analog value callback.
// Is only triggered if the value changed, since last triggering.
func AnalogValuePeriod(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "AnalogValuePeriod"),
		Fid:        callback_analog_value,
		Uid:        uid,
		Result:     &AnalogValue{},
		Handler:    handler,
		IsCallback: true,
		WithPacket: false}.CreateDevice()
}
