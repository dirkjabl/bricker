// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package analogin

import (
	"bricker"
	"bricker/device"
	"bricker/net/packet"
	"bricker/subscription"
	"bricker/util/hash"
)

// SetVoltageCallbackThreshold creates the subscriber to set the callback thresold.
// Default value is ('x', 0, 0).
func SetVoltageCallbackThreshold(id string, uid uint32, t *device.Threshold, handler func(device.Resulter, error)) *device.Device {
	fid := function_set_voltage_callback_threshold
	svct := device.New(device.FallbackId(id, "SetVoltageCallbackThreshold"))
	p := packet.NewSimpleHeaderPayload(uid, fid, true, t)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	svct.SetSubscription(sub)
	svct.SetResult(&device.EmptyResult{})
	svct.SetHandler(handler)
	return svct
}

// SetVoltageCallbackThresholdFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetVoltageCallbackThresholdFuture(brick *bricker.Bricker, connectorname string, uid uint32, t *device.Threshold) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetVoltageCallbackThreshold("setvoltagecallbackthresholdfuture"+device.GenId(), uid, t,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetVoltageCallbackThreshold creates the subscriber to get the callback thresold.
func GetVoltageCallbackThreshold(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := function_get_voltage_callback_threshold
	gvct := device.New(device.FallbackId(id, "GetVoltageCallbackThreshold"))
	p := packet.NewSimpleHeaderOnly(uid, fid, true)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	gvct.SetSubscription(sub)
	gvct.SetResult(&device.Threshold{})
	gvct.SetHandler(handler)
	return gvct
}

// GetVoltageCallbackThresholdFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func GetVoltageCallbackThresholdFuture(brick *bricker.Bricker, connectorname string, uid uint32) *device.Threshold {
	future := make(chan *device.Threshold)
	defer close(future)
	sub := GetVoltageCallbackThreshold("getvoltagecallbackthresholdfuture"+device.GenId(), uid,
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

// SetAnalogValueCallbackThreshold creates the subscriber to set the callback thresold.
// Default value is ('x', 0, 0).
func SetAnalogValueCallbackThreshold(id string, uid uint32, t *device.Threshold, handler func(device.Resulter, error)) *device.Device {
	fid := function_set_analog_value_callback_threshold
	savct := device.New(device.FallbackId(id, "SetAnalogValueCallbackThreshold"))
	p := packet.NewSimpleHeaderPayload(uid, fid, true, t)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	savct.SetSubscription(sub)
	savct.SetResult(&device.EmptyResult{})
	savct.SetHandler(handler)
	return savct
}

// SetAnalogValueCallbackThresholdFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetAnalogValueCallbackThresholdFuture(brick *bricker.Bricker, connectorname string, uid uint32, t *device.Threshold) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetAnalogValueCallbackThreshold("setanalogvaluecallbackthresholdfuture"+device.GenId(), uid, t,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetAnalogValueCallbackThreshold creates the subscriber to get the callback thresold.
func GetAnalogValueCallbackThreshold(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := function_get_analog_value_callback_threshold
	gavct := device.New(device.FallbackId(id, "GetAnalogValueCallbackThreshold"))
	p := packet.NewSimpleHeaderOnly(uid, fid, true)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	gavct.SetSubscription(sub)
	gavct.SetResult(&device.Threshold{})
	gavct.SetHandler(handler)
	return gavct
}

// GetAnalogValueCallbackThresholdFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func GetAnalogValueCallbackThresholdFuture(brick *bricker.Bricker, connectorname string, uid uint32) *device.Threshold {
	future := make(chan *device.Threshold)
	defer close(future)
	sub := GetAnalogValueCallbackThreshold("getanalogvaluecallbackthresholdfuture"+device.GenId(), uid,
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

// VoltageReached creates a subscriber for the theshold triggered voltage callback.
func VoltageReached(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := callback_voltage_reached
	vp := device.New(device.FallbackId(id, "VoltageReached"))
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, nil, false)
	vp.SetSubscription(sub)
	vp.SetResult(&Voltage{})
	vp.SetHandler(handler)
	return vp
}

// AnalogValueReached creates a subscriber for the theshold triggered voltage callback.
func AnalogValueReached(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := callback_analog_value_reached
	avp := device.New(device.FallbackId(id, "AnalogValueReached"))
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, nil, false)
	avp.SetSubscription(sub)
	avp.SetResult(&AnalogValue{})
	avp.SetHandler(handler)
	return avp
}
