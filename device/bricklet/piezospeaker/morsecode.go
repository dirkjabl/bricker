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

const (
	MorsePause = byte(' ') // Pause
	MorseLong  = byte('-') // Long
	MorseShort = byte('.') // Short
)

/*
MorseCode creates a subscriber for output an morse code signal.
The maximum string size is 60.
*/
func MorseCode(id string, uid uint32, m *Morse, handler func(device.Resulter, error)) *device.Device {
	fid := function_morse_code
	mc := device.New(device.FallbackId(id, "MorseCode"))
	p := packet.NewSimpleHeaderPayload(uid, fid, true, m)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	mc.SetSubscription(sub)
	mc.SetResult(&device.EmptyResult{})
	mc.SetHandler(handler)
	return mc
}

// MorseCodeFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func MorseCodeFuture(brick bricker.Bricker, connectorname string, uid uint32, m *Morse) bool {
	future := make(chan bool)
	defer close(future)
	sub := MorseCode("morsecodefuture"+device.GenId(), uid, m,
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
func MorseCodeFinished(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := callback_morse_code_finished
	bf := device.New(device.FallbackId(id, "MorseCodeFinished"))
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, nil, true)
	bf.SetSubscription(sub)
	bf.SetResult(&device.EmptyResult{})
	bf.SetHandler(handler)
	return bf
}

/*
Morse is the representation of the morse code for the MorseCode subscriber.

As characters are only " "(space/pause), "."(dot/short) and "-"(minus/long) allowed.
All other characters are ignored.
It could handle up to 60 character.
*/
type Morse struct {
	Code      [60]byte
	Frequency uint16
}
