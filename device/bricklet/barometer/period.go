// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package barometer

import (
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
)

// SetAirPressureCallbackPeriod creates the subscriber to set the callback period.
// Default value is 0. A value of 0 deactivates the periodical callbacks.
// AirPressurePeriod is only triggered if the air pressure has changed since the last triggering.
func SetAirPressureCallbackPeriod(id string, uid uint32, pe *device.Period, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SetAirPressureCallbackPeriod"),
		Fid:        function_set_air_pressure_callback_period,
		Uid:        uid,
		Data:       pe,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SetAirPressureCallbackPeriodFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetAirPressureCallbackPeriodFuture(brick *bricker.Bricker, connectorname string, uid uint32, pe *device.Period) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetAirPressureCallbackPeriod("setairpressurecallbackperiodfuture"+device.GenId(), uid, pe,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetAirPressureCallbackPeriod creates a subsctiber to get the callback period value.
func GetAirPressureCallbackPeriod(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetAirPressureCallbackPeriod"),
		Fid:        function_get_air_pressure_callback_period,
		Uid:        uid,
		Result:     &device.Period{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetAirPressureCallbackPeriodFuture is a future pattern version for a synchronized call of the subsctiber.
// If an error occur, the result is nil.
func GetAirPressureCallbackPeriodFuture(brick *bricker.Bricker, connectorname string, uid uint32) *device.Period {
	future := make(chan *device.Period)
	defer close(future)
	sub := GetAirPressureCallbackPeriod("getairpressurecallbackperiod"+device.GenId(), uid,
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

// SetAltitudeCallbackPeriod creates the subscriber to set the callback period.
// Default value is 0. A value of 0 deactivates the periodical callbacks.
// AltitudePeriod is only triggered if the illuminance has changed since the last triggering.
func SetAltitudeCallbackPeriod(id string, uid uint32, pe *device.Period, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SetAltitudeCallbackPeriod"),
		Fid:        function_set_altitude_callback_period,
		Uid:        uid,
		Data:       pe,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SetAltitudeCallbackPeriodFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetAltitudeCallbackPeriodFuture(brick *bricker.Bricker, connectorname string, uid uint32, pe *device.Period) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetAltitudeCallbackPeriod("setaltitudecallbackperiodfuture"+device.GenId(), uid, pe,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetAltitudeCallbackPeriod creates a subscriber to get the callback period value.
func GetAltitudeCallbackPeriod(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetAltitudeCallbackPeriod"),
		Fid:        function_get_altitude_callback_period,
		Uid:        uid,
		Result:     &device.Period{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetAltitudeCallbackPeriodFuture is a future pattern version for a synchronized call of the subsctiber.
// If an error occur, the result is nil.
func GetAltitudeCallbackPeriodFuture(brick *bricker.Bricker, connectorname string, uid uint32) *device.Period {
	future := make(chan *device.Period)
	defer close(future)
	sub := GetAltitudeCallbackPeriod("getaltitudecallbackperiodfuture"+device.GenId(), uid,
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

// AirPressurePeriod creates a subscriber for the periodical air pressure callback.
// Is only triggered if the voltage changed, since last triggering.
func AirPressurePeriod(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "AirPressurePeriod"),
		Fid:        callback_air_pressure,
		Uid:        uid,
		Result:     &AirPressure{},
		Handler:    handler,
		IsCallback: true,
		WithPacket: false}.CreateDevice()
}

// AltitudePeriod creates a subscriber for the periodical altitude callback.
// Is only triggered if the value changed, since last triggering.
func AltitudePeriod(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "AltitudePeriod"),
		Fid:        callback_altitude,
		Uid:        uid,
		Result:     &Altitude{},
		Handler:    handler,
		IsCallback: true,
		WithPacket: false}.CreateDevice()
}
