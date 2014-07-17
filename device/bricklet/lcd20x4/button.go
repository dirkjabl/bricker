// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lcd20x4

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
	"github.com/dirkjabl/bricker/subscription"
	"github.com/dirkjabl/bricker/util/hash"
)

// IsButtonPressed creates a subscriber to get the information, if a specific button is pressed.
func IsButtonPressed(id string, uid uint32, button *Button, handler func(device.Resulter, error)) *device.Device {
	fid := function_is_button_pressed
	ibp := device.New(device.FallbackId(id, "IsButtonPressed"))
	p := packet.NewSimpleHeaderPayload(uid, fid, true, button)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	ibp.SetSubscription(sub)
	ibp.SetResult(&Pressed{})
	ibp.SetHandler(handler)
	return ibp
}

// IsButtonPressedFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func IsButtonPressedFuture(brick *bricker.Bricker, connectorname string, uid uint32, button *Button) *Pressed {
	future := make(chan *Pressed)
	defer close(future)
	sub := IsButtonPressed("isbuttonpressedfuture"+device.GenId(), uid, button,
		func(r device.Resulter, err error) {
			var v *Pressed = nil
			if err == nil {
				if value, ok := r.(*Pressed); ok {
					v = value
				}
			}
			future <- v
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return nil
	}
	v := <-future
	return v
}

// IsButtonPressedFutureSimple calls the IsButtonPressedFuture method with a simple boolean result.
// If it fails, the result is false.
func IsButtonPressedFutureSimple(brick *bricker.Bricker, connectorname string, uid uint32, button *Button) bool {
	p := IsButtonPressedFuture(brick, connectorname, uid, button)
	if p == nil {
		return false
	}
	return p.IsPressed
}

// ButtonPressed creates a subscriber for the button pressed callback.
func ButtonPressed(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := callback_button_pressed
	bp := device.New(device.FallbackId(id, "ButtonPressed"))
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, nil, true)
	bp.SetSubscription(sub)
	bp.SetResult(&Button{})
	bp.SetHandler(handler)
	return bp
}

// ButtonReleased creates a subscriber for the button release callback.
func ButtonReleased(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := callback_button_released
	br := device.New(device.FallbackId(id, "ButtonReleased"))
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, nil, true)
	br.SetSubscription(sub)
	br.SetResult(&Button{})
	br.SetHandler(handler)
	return br
}

// Button type.
// For calling, if the button is pressed, or the button callbacks.
type Button struct {
	Number uint8
}

// FromPacket converts the packet payload to the Button type.
func (b *Button) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(b, p); err != nil {
		return err
	}
	return p.Payload.Decode(b)
}

// String fullfill the stringer interface.
func (b *Button) String() string {
	txt := "Button "
	if b != nil {
		txt += fmt.Sprintf("[Number: %d]", b.Number)
	} else {
		txt += "[nil]"
	}
	return txt
}

// Pressed is a type for the return of the IsButtonPressed subscriber.
type Pressed struct {
	IsPressed bool // is the button pressed
}

// FromPacket converts the packet payload to the Pressed type.
func (pr *Pressed) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(pr, p); err != nil {
		return err
	}
	prr := NewPressedRawFromPressed(pr)
	err := p.Payload.Decode(prr)
	if err == nil {
		pr.FromPressedRaw(prr)
	}
	return err
}

// String fullfill the stringer interface.
func (pr *Pressed) String() string {
	txt := "Pressed ["
	if pr != nil {
		txt += fmt.Sprintf("IsPressed: %t", pr.IsPressed)
	} else {
		txt += "nil"
	}
	return txt + "]"
}

// FromPressedRaw converts a PressedRaw type to a Pressed type.
func (pr *Pressed) FromPressedRaw(prr *PressedRaw) {
	if pr == nil || prr == nil {
		return
	}
	pr.IsPressed = (prr.IsPressed & 0x01) == 0x01
}

// PressedRaw is a type for raw coding of the pressed state.
type PressedRaw struct {
	IsPressed uint8
}

// NewPressedRawFromPressed is a simple constructor for a PressedRaw from a Pressed type.
func NewPressedRawFromPressed(pr *Pressed) *PressedRaw {
	prr := new(PressedRaw)
	prr.FromPressed(pr)
	return prr
}

// FromPressed simple converter from Pressed type to PressedRaw.
func (prr *PressedRaw) FromPressed(pr *Pressed) {
	if pr == nil {
		return
	}
	if pr.IsPressed {
		prr.IsPressed = 0x01
	} else {
		prr.IsPressed = 0x00
	}
}
