// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package piezobuzzer

import (
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
)

/*
Beep creates a subscriber for an output of a beep.

The speaker can only approximate the frequency,
it will play the best possible match by applying the calibration.
*/
func Beep(id string, uid uint32, b *Beeps, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "Beep"),
		Fid:        function_beep,
		Uid:        uid,
		Data:       b,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// BeepFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func BeepFuture(brick bricker.Bricker, connectorname string, uid uint32, b *Beeps) bool {
	future := make(chan bool)
	defer close(future)
	sub := Beep("beepfuture"+device.GenId(), uid, b,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// BeepFinished creates a subscriber which is triggered if a Beep subscriber is finished.
// No data are submitted.
func BeepFinished(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "BeepFinished"),
		Fid:        callback_beep_finished,
		Uid:        uid,
		Handler:    handler,
		IsCallback: true,
		WithPacket: false}.CreateDevice()
}

/*
Beeps is the type for the Beep subscriber.
The duration is set in ms.
*/
type Beeps struct {
	Duration uint32
}
