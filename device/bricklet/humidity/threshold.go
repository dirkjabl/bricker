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

// SetHumidityCallbackThreshold creates the subscriber to set the callback thresold.
// Default value is ('x', 0, 0).
func SetHumidityCallbackThreshold(id string, uid uint32, t *device.Threshold, handler func(device.Resulter, error)) *device.Device {
	fid := function_set_humidity_callback_threshold
	shct := device.New(device.FallbackId(id, "SetHumidityCallbackThreshold"))
	p := packet.NewSimpleHeaderPayload(uid, fid, true, t)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	shct.SetSubscription(sub)
	shct.SetResult(&device.EmptyResult{})
	shct.SetHandler(handler)
	return shct
}

// SetHumidityCallbackThresholdFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetHumidityCallbackThresholdFuture(brick *bricker.Bricker, connectorname string, uid uint32, t *device.Threshold) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetHumidityCallbackThreshold("sethumiditycallbackthresholdfuture"+device.GenId(), uid, t,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetHumidityCallbackThreshold creates the subscriber to get the callback thresold.
func GetHumidityCallbackThreshold(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := function_get_humidity_callback_threshold
	ghct := device.New(device.FallbackId(id, "GetHumidityCallbackThreshold"))
	p := packet.NewSimpleHeaderOnly(uid, fid, true)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	ghct.SetSubscription(sub)
	ghct.SetResult(&device.Threshold{})
	ghct.SetHandler(handler)
	return ghct
}

// GetHumidityCallbackThresholdFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func GetHumidityCallbackThresholdFuture(brick *bricker.Bricker, connectorname string, uid uint32) *device.Threshold {
	future := make(chan *device.Threshold)
	defer close(future)
	sub := GetHumidityCallbackThreshold("gethumiditycallbackthresholdfuture"+device.GenId(), uid,
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

// HumidityReached creates a subscriber for the theshold triggered voltage callback.
func HumidityReached(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := callback_humidity_reached
	vp := device.New(device.FallbackId(id, "HumidityReached"))
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, nil, false)
	vp.SetSubscription(sub)
	vp.SetResult(&Humidity{})
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
