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

// SetCustomCharacter creates a subsriber to set a custom character.
func SetCustomCharacter(id string, uid uint32, c *CustomCharacter, handler func(device.Resulter, error)) *device.Device {
	fid := function_set_custom_character
	scc := device.New(device.FallbackId(id, "SetCustomCharacter"))
	p := packet.NewSimpleHeaderPayload(uid, fid, true, c)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	scc.SetSubscription(sub)
	scc.SetResult(&device.EmptyResult{})
	scc.SetHandler(handler)
	return scc
}

// SetCustomCharacterFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetCustomCharacterFuture(brick *bricker.Bricker, connectorname string, uid uint32, c *CustomCharacter) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetCustomCharacter("setcustomcharacterfuture"+device.GenId(), uid, c,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	b := <-future
	return b
}

// GetCustomCharacter creates a subscriber to get a stored custom character at the given index.
func GetCustomCharacter(id string, uid uint32, index uint8, handler func(device.Resulter, error)) *device.Device {
	fid := function_get_custom_character
	gcc := device.New(device.FallbackId(id, "GetCustomCharacter"))
	p := packet.NewSimpleHeaderPayload(uid, fid, true, index)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	gcc.SetSubscription(sub)
	gcc.SetResult(&Character{})
	gcc.SetHandler(handler)
	return gcc
}

// CustomCharacter is the type for a custom character.
// There could store up to 8 custom character.
//
// The characters can later be written with WriteLine
// by using the characters with the byte representation 8 ("x08") to 15 ("x0F").
//
// Custom characters are stored by the LCD in RAM, so they have to be set after each startup.
type CustomCharacter struct {
	Index uint8
	Char  Character
}

func (c *CustomCharacter) String() string {
	txt := "LCD 20x4 Custom Character "
	if c != nil {
		txt += "[" + c.Char.String() + "]"
	} else {
		txt += "[nil]"
	}
	return txt
}

// Character stores the pixel data for a single custom character.
// It is a array with 8 lines of 5 pixel (5x8).
// Element 1 (index 0) is the first line, element 8 (index 7) the final line.
type Character [8]uint8

// FromPacket converts from a packet payload to type Character.
func (c *Character) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(c, p); err != nil {
		return err
	}
	return p.Payload.Decode(c)
}

// String fullfill the stringer interface.
func (c *Character) String() string {
	txt := "LCD 20x4 Character "
	if c != nil {
		txt += "["
		for i, v := range c {
			txt += fmt.Sprintf(" %d:0x%x", i, v)
		}
		txt += "]"
	} else {
		txt += "[nil]"
	}
	return txt
}
