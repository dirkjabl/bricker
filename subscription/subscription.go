// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
The subscription describes the intressts of a subscriber.
Every subscriber need this data.
The subscription holds the information, which has to send (as packet) for notifyable results.
Also it has the information, on which events (notifyable events) the subscriber wants to be informed.
*/
package subscription

import (
	"bricker/net/packet"
	"bricker/util/hash"
	"fmt"
)

// Identify the events a subscriber handle.
// The flag Callback is Tinkerforge specific. Is this not a callback, then there is only one result to get.
// The Request holds the packet, which is to send to get events.
type Subscription struct {
	Choosen    uint8          // Choose which values are reconized for matching
	Uid        uint32         // Value uid
	FunctionID uint8          // Value Function-ID
	Request    *packet.Packet // ip packet
	Callback   bool           // Is this subscription a callback (get more as one result) or not (one result)
}

// NewSubscription creates a new subscription with all informations.
func New(c uint8, u uint32, f uint8, p *packet.Packet, cb bool) *Subscription {
	return &Subscription{
		Choosen:    c,
		Uid:        u,
		FunctionID: f,
		Request:    p,
		Callback:   cb}
}

// NewSubscriptionFid creates a new subscription with the choosen values, a function identifyer and a packet.
func NewFid(f uint8, p *packet.Packet, cb bool) *Subscription {
	return New(hash.ChoosenFunctionID, 0, f, p, cb)
}

// Hash creates the actual hash sum of the subscription
func (s *Subscription) Hash() hash.Hash {
	return hash.New(s.Choosen, s.Uid, s.FunctionID)
}

// CompareHash test the given hash with the computed hash of the subscription.
func (s *Subscription) CompareHash(h hash.Hash) bool {
	return h.Equal(s.Hash())
}

// String fullfill the stringer interface.
func (s *Subscription) String() string {
	t := "Subscription [Choosen: "
	if (s.Choosen | hash.ChoosenFunctionID) == 1 {
		t += "FunctionID "
	}
	if (s.Choosen | hash.ChoosenUid) == 1 {
		t += "Uid"
	}
	t += fmt.Sprintf(", Uid: %d, Function-ID: %d", s.Uid, s.FunctionID)
	if s.Request != nil {
		t += fmt.Sprintf(", Request-Packet: %s", s.Request.String())
	} else {
		t += fmt.Sprintf(", Request-Packet: nil")
	}
	t += "]"
	return t
}
