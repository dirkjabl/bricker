// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package piezospeaker

import (
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
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
	return device.Generator{
		Id:         device.FallbackId(id, "MorseCode"),
		Fid:        function_morse_code,
		Uid:        uid,
		Data:       m,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
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
	return device.Generator{
		Id:         device.FallbackId(id, "MorseCodeFinished"),
		Fid:        callback_morse_code_finished,
		Uid:        uid,
		Handler:    handler,
		IsCallback: true,
		WithPacket: false}.CreateDevice()
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
