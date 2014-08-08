// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package analogin

import (
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
)

// SetVoltageCallbackPeriod creates the subscriber to set the callback period.
// Default value is 0. A value of 0 deactivates the periodical callbacks.
// VoltagePeriod is only triggered if the voltage has changed since the last triggering.
func SetVoltageCallbackPeriod(id string, uid uint32, pe *device.Period, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SetVoltageCallbackPeriod"),
		Fid:        function_set_voltage_callback_period,
		Uid:        uid,
		Data:       pe,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
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
	return device.Generator{
		Id:         device.FallbackId(id, "GetVoltageCallbackPeriod"),
		Fid:        function_get_voltage_callback_period,
		Uid:        uid,
		Result:     &device.Period{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
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

// VoltagePeriod creates a subscriber for the periodical voltage callback.
// Is only triggered if the voltage changed, since last triggering.
func VoltagePeriod(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "VoltagePeriod"),
		Fid:        callback_voltage,
		Uid:        uid,
		Result:     &Voltage{},
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
