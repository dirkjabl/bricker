// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package temperature

import (
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
	"github.com/dirkjabl/bricker/subscription"
	"github.com/dirkjabl/bricker/util/hash"
)

// SetDebouncePeriod creates the subscriber to get the debounce period.
// The default value is 100.
func SetDebouncePeriod(id string, uid uint32, d *device.Debounce, handler func(device.Resulter, error)) *device.Device {
	fid := function_set_debounce_period
	sdp := device.New(device.FallbackId(id, "SetDebouncePeriod"))
	p := packet.NewSimpleHeaderPayload(uid, fid, true, d)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	sdp.SetSubscription(sub)
	sdp.SetResult(&device.EmptyResult{})
	sdp.SetHandler(handler)
	return sdp
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
	fid := function_get_debounce_period
	gdp := device.New(device.FallbackId(id, "GetDebouncePeriod"))
	p := packet.NewSimpleHeaderOnly(uid, fid, true)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	gdp.SetSubscription(sub)
	gdp.SetResult(&device.Debounce{})
	gdp.SetHandler(handler)
	return gdp
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
