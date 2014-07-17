// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bricker

import (
	"bricker/event"
	"bricker/subscription"
)

// Interface for all Subscriber.
// Every Subscriber needs a subscription, which identifies the events, the subscriber wants to handle.
// A subscriber must be bound to a bricker.
type Subscriber interface {
	Id() string
	Subscription() *subscription.Subscription
	Notify(*event.Event)
}

// Subscriber register a subscriber. Internaly it use the subscription of the subscriber.
func (b *Bricker) Subscribe(s Subscriber, dest interface{}) error {
	hash := s.Subscription().Hash()
	if v, ok := b.subscriber[hash]; ok {
		if _, ok := v[s.Id()]; ok {
			return NewError(ErrorSubscriberExists)
		}
		v[s.Id()] = s
	} else {
		b.subscriber[hash] = map[string]Subscriber{s.Id(): s}
	}
	b.insertChooser(s.Subscription().Choosen)
	if s.Subscription().Request != nil { // only send a event, if a packet is given
		ev := event.NewPacket(s.Subscription().Request)
		ev.ConnectorName = b.computeConnectorsName(dest)
		go b.write(ev)
	}
	return nil
}

// Unsubscribe release a registered subscriber identified with the subscription.
func (b *Bricker) Unsubscribe(s Subscriber) error {
	hash := s.Subscription().Hash()
	subs, ok := b.subscriber[hash]
	if !ok {
		return NewError(ErrorNoSubscriberToRelease)
	}
	_, ok = subs[s.Id()]
	if ok {
		delete(subs, s.Id())
	} else {
		return NewError(ErrorNoSubscriberToRelease)
	}
	if len(subs) == 0 { // delete empty map
		delete(b.subscriber, hash)
	}
	return nil
}

// Internal method: insertChooser add a new chooser to the slice of chooser
func (b *Bricker) insertChooser(n uint8) {
	for _, v := range b.choosers {
		if v == n {
			return // already inserted
		}
	}
	b.choosers = append(b.choosers, n)
}
