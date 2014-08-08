// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lcd20x4

import (
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
)

// ClearDisplay is a subscriber to clear the LCD display.
func ClearDisplay(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "ClearDisplay"),
		Fid:        function_clear_display,
		Uid:        uid,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// ClearDisplayFuture is the future version of the ClearDisplay subscriber.
func ClearDisplayFuture(brick *bricker.Bricker, connectorname string, uid uint32) bool {
	future := make(chan bool)
	sub := ClearDisplay("cleardisplayfuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	b := <-future
	close(future)
	return b
}
