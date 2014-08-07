// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package device

import (
	"github.com/dirkjabl/bricker/net/packet"
	"github.com/dirkjabl/bricker/subscription"
	"github.com/dirkjabl/bricker/util/hash"
)

// Creates a device with a given subscription, resulter and handler and an id.
func NewSubscriptionResulterHandler(id string, sub *subscription.Subscription, result Resulter, handler func(Resulter, error)) *Device {
	d := New(FallbackId(id, "Device"))
	d.SetSubscription(sub)
	d.SetResult(result)
	d.SetHandler(handler)
	return d
}

/*
Creates a new device with a subscription, simple way to create a "header only" subscriber.
The subscrition is bind to the Function ID and the UID.
  id - id for the new device.
  uid - UID of the real hardware device.
  fid - Function Id with specifies the function to call.
  iscb - is this a callback or only a once called function.
  result - the result object (which have to fullfill the Resulter interface)
  handler - the callback/event handler function
*/
func NewHeaderOnlyWithResult(id string, uid uint32, fid uint8, iscb bool, result Resulter, handler func(Resulter, error)) *Device {
	p := packet.NewSimpleHeaderOnly(uid, fid, true)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, iscb)
	return NewSubscriptionResulterHandler(id, sub, result, handler)
}

// Creates a new device with a subscription but without a result object (inside the resulting packet payload).
func NewHeaderOnlyEmptyResult(id string, uid uint32, fid uint8, iscb bool, handler func(Resulter, error)) *Device {
	return NewHeaderOnlyWithResult(id, uid, fid, iscb, &EmptyResult{}, handler)
}

/*
Creates a new device with a subscription, simple way to create a "header and payload only" subscriber.
The subscrition is bind to the Function ID and the UID.
  id - id for the new device.
  uid - UID of the real hardware device.
  fid - Function Id with specifies the function to call.
  iscb - is this a callback or only a once called function.
  data - the data to send with the header.
  result - the result object (which have to fullfill the Resulter interface)
  handler - the callback/event handler function
*/
func NewHeaderPayloadWithResult(id string, uid uint32, fid uint8, iscb bool, data interface{}, result Resulter, handler func(Resulter, error)) *Device {
	p := packet.NewSimpleHeaderPayload(uid, fid, true, data)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, iscb)
	return NewSubscriptionResulterHandler(id, sub, result, handler)
}

// Creates a new device with a subscription but without a result object (inside the resulting packet payload).
func NewHeaderPayloadEmptyResult(id string, uid uint32, fid uint8, iscb bool, data interface{}, handler func(Resulter, error)) *Device {
	return NewHeaderPayloadWithResult(id, uid, fid, iscb, data, &EmptyResult{}, handler)
}

/*
Creates a new device with a subscription,
simple way to create a "header only" subscriber without a packet to send.
This is the typical callback subscriber.
The subscrition is bind to the Function ID and the UID.
  id - id for the new device.
  uid - UID of the real hardware device.
  fid - Function Id with specifies the function to call.
  iscb - is this a callback or only a once called function.
  result - the result object (which have to fullfill the Resulter interface)
  handler - the callback/event handler function
*/
func NewHeaderNoPacketWithResult(id string, uid uint32, fid uint8, iscb bool, result Resulter, handler func(Resulter, error)) *Device {
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, nil, iscb)
	return NewSubscriptionResulterHandler(id, sub, result, handler)
}
