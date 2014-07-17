// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package piezospeaker

import (
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
	"github.com/dirkjabl/bricker/subscription"
	"github.com/dirkjabl/bricker/util/hash"
)

/*
Beep creates a subscriber for an output of a beep.

The speaker can only approximate the frequency,
it will play the best possible match by applying the calibration.
*/
func Beep(id string, uid uint32, b *Beeps, handler func(device.Resulter, error)) *device.Device {
	fid := function_beep
	bp := device.New(device.FallbackId(id, "Beep"))
	p := packet.NewSimpleHeaderPayload(uid, fid, true, b)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	bp.SetSubscription(sub)
	bp.SetResult(&device.EmptyResult{})
	bp.SetHandler(handler)
	return bp
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
	fid := callback_beep_finished
	bf := device.New(device.FallbackId(id, "BeepFinished"))
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, nil, true)
	bf.SetSubscription(sub)
	bf.SetResult(&device.EmptyResult{})
	bf.SetHandler(handler)
	return bf
}

/*
Beeps is the type for the Beep subscriber.
The duration is set in ms.
Frequency can be set between 585 and 7100 (~ Hz).
*/
type Beeps struct {
	Duration  uint32
	Frequency uint16
}
