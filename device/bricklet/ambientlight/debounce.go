// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ambientlight

import (
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
)

/*
SetDebouncePeriod creates the subscriber to get the debounce period.
The default value is 100 (ms).
This sets the period in ms in which the threshold callbacks are triggered,
only if the threshold are being reached.
*/
func SetDebouncePeriod(id string, uid uint32, d *device.Debounce, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SetDebouncePeriod"),
		Fid:        function_set_debounce_period,
		Uid:        uid,
		Data:       d,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SetDebouncePeriodFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetDebouncePeriodFuture(brick *bricker.Bricker, connectorname string, uid uint32, d *device.Debounce) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetDebouncePeriod("setdebounceperiodfuture"+device.GenId(), uid, d,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetDebouncePeriod creates the subscriber to set the debounce period.
func GetDebouncePeriod(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetDebouncePeriod"),
		Fid:        function_get_debounce_period,
		Uid:        uid,
		Result:     &device.Debounce{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetDebouncePeriodFuture is a future pattern version for a synchronized all of the subscriber.
// If an error occur, the result is nil.
func GetDebouncePeriodFuture(brick *bricker.Bricker, connectorname string, uid uint32) *device.Debounce {
	future := make(chan *device.Debounce)
	defer close(future)
	sub := GetDebouncePeriod("getdebounceperiodfuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *device.Debounce = nil
			if err == nil {
				if value, ok := r.(*device.Debounce); ok {
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
