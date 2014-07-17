// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lcd20x4

import (
	"bricker"
	"bricker/device"
	"bricker/net/packet"
	"bricker/subscription"
	"bricker/util/hash"
	"fmt"
)

// WriteLine creates a new subscriber to write a text to the LCD (one line).
func WriteLine(id string, uid uint32, ltl *LcdTextLine, handler func(r device.Resulter, e error)) *device.Device {
	fid := function_write_line
	wl := device.New(device.FallbackId(id, "WriteLine"))
	p := packet.NewSimpleHeaderPayload(uid, fid, true, ltl)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	wl.SetSubscription(sub)
	wl.SetResult(&device.EmptyResult{})
	wl.SetHandler(handler)
	return wl
}

// WriteLineFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func WriteLineFuture(brick *bricker.Bricker, connectorname string, uid uint32, ltl *LcdTextLine) bool {
	future := make(chan bool)
	defer close(future)
	sub := WriteLine("writelinefuture"+device.GenId(), uid, ltl, func(r device.Resulter, err error) {
		future <- device.IsEmptyResultOk(r, err)
	})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	v := <-future
	return v
}

// LcdTextLine is the type for a text line to display.
// The text should not be longer than 20 bytes.
type LcdTextLine struct {
	Line uint8
	Pos  uint8
	Text [20]byte
}

// FromPacket creates from a packet a LcdTextLine.
func (ltl *LcdTextLine) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(ltl, p); err != nil {
		return err
	}
	return p.Payload.Decode(ltl)
}

// String fullfill the stringer interface.
func (ltl *LcdTextLine) String() string {
	return fmt.Sprintf("LCD 20x4 Text Line [Line: %d Position: %d Text: %s]", ltl.Line, ltl.Pos, ltl.Text)
}
