// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package device

import (
	"fmt"
	"github.com/dirkjabl/bricker/net/packet"
	"github.com/dirkjabl/bricker/subscription"
	"github.com/dirkjabl/bricker/util/hash"
)

/*
It is a generator type to create a device (subscriber) for a given data constellation.
*/
type Generator struct {
	Id         string                // device Id, if not given, the generator use the word "Device" as base.
	Uid        uint32                // UID of the device.
	Fid        uint8                 // Function Id on the device to call the right function.
	Data       interface{}           // Data for a calling (to send) packet.
	Result     Resulter              // Result object for the result of the callback, it have to fullfill the Resulter interface.
	Handler    func(Resulter, error) // The callback/event handler function to call on an event.
	IsCallback bool                  // This is a callback and comes often, not only once.
	WithPacket bool                  // This subscriber should create a calling (to send) packet.
}

/*
CreateDevice creates a subscriber (device) on base of the generator object data.
This should make the source code of the subscriber creator methods more readable and easy.
The subscrition is bind to the Function ID and the UID.
A generator object should not be changed after creation, so no pointer version exists.

If WithPacket is false, no packet for sending to the device will be created.
When a packet is created, it will only have data inside the payload, if Data is filled.

The Result type of the subscriber will be EmptyResult if no Result is given.
*/
func (g Generator) CreateDevice() *Device {
	id := FallbackId(g.Id, "Device")
	var r Resulter = g.Result
	var p *packet.Packet = nil
	if g.WithPacket {
		if g.Data == nil {
			p = packet.NewSimpleHeaderOnly(g.Uid, g.Fid, true)
		} else {
			p = packet.NewSimpleHeaderPayload(g.Uid, g.Fid, true, g.Data)
		}
	}
	if g.Result == nil {
		r = &EmptyResult{}
	}
	sub := subscription.New(hash.ChoosenFunctionIDUid, g.Uid, g.Fid, p, g.IsCallback)
	return NewSubscriptionResulterHandler(id, sub, r, g.Handler)
}

// String fullfill the stringer interface.
func (g Generator) String() string {
	txt := "Generator ["
	txt += fmt.Sprintf("Id: %s, UID: %d, Function ID: %d, ", g.Id, g.Uid, g.Fid)
	txt += fmt.Sprintf("Has Data: %t, ", (g.Data != nil))
	txt += fmt.Sprintf("Has Resulter: %t, ", (g.Result != nil))
	txt += fmt.Sprintf("Is Callback: %t, With Packet: %t", g.IsCallback, g.WithPacket)
	txt += "]"
	return txt
}
