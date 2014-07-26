// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dualbutton

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
	"github.com/dirkjabl/bricker/subscription"
	"github.com/dirkjabl/bricker/util/hash"
)

// GetButtonState creates a subscriber to get the button states.
func GetButtonState(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := function_get_button_state
	gbs := device.New(device.FallbackId(id, "GetButtonState"))
	p := packet.NewSimpleHeaderOnly(uid, fid, true)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	gbs.SetSubscription(sub)
	gbs.SetResult(&ButtonState{})
	gbs.SetHandler(handler)
	return gbs
}

// GetButtonStateFuture is a future pattern version for a synchronized all of the subscriber.
// If an error occur, the result is nil.
func GetButtonStateFuture(brick *bricker.Bricker, connectorname string, uid uint32) *ButtonState {
	future := make(chan *ButtonState)
	defer close(future)
	sub := GetButtonState("getbuttonstatefuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *ButtonState = nil
			if err == nil {
				if value, ok := r.(*ButtonState); ok {
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

/*
ButtonState is the type for the state of the buttons.

    0 - button pressed
    1 - button released
*/
type ButtonState struct {
	ButtonLeft  uint8 // button left
	ButtonRight uint8 // button right
}

// FromPacket creates a ButtonState from a packet.
func (bs *ButtonState) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(bs, p); err != nil {
		return err
	}
	return p.Payload.Decode(bs)
}

// String fullfill the stringer interface.
func (bs *ButtonState) String() string {
	txt := "Button state "
	if bs == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("[Button left: %s (%d), Button right: %s (%d)]",
			ButtonStateName(bs.ButtonLeft), bs.ButtonLeft,
			ButtonStateName(bs.ButtonRight), bs.ButtonRight)
	}
	return txt
}
