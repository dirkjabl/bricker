// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
bricker is a API for the Tinkerforge Hardware based on the brick daemon (brickd).

A bricker is a manager.
It uses one or more connectors to send and receive packets from brick daemons (real hardware).

The connectors could send and receive events to a specific address.
The packets will encapsulate in events.
The bricker is also a producer of events.
To use this events their will be need consumer (subscriber).
There is a fallback mechanism for events without a consumer (default fallback subscriber).

For using this API you need a running brick daemon (brickd) or some hardware with a brick daemon,
please use an actual version of the daemon.
You get the daemon from http://www.tinkerforge.com/en/doc/Software/Brickd.html#brickd as
source or binary package.

This API based on the documentation of the TCP/IP API from http://www.tinkerforge.com/en/doc/index.html#/software-tcpip-open and so this documentation is also useful.
*/
package bricker

import (
	"github.com/dirkjabl/bricker/connector"
	"github.com/dirkjabl/bricker/event"
	"github.com/dirkjabl/bricker/util/hash"
)

// The bricker type.
// A bricker managed connectors and subscriber.
type Bricker struct {
	connection        map[string]connector.Connector
	first             string
	uids              map[uint32]string
	subscriber        map[hash.Hash]map[string]Subscriber
	choosers          []uint8
	defaultsubscriber Subscriber
}

// New create the bricker.
// The new bricker start direct the service.
// After start, the bricker has no connection and no subscriber.
func New() *Bricker {
	return &Bricker{
		connection: make(map[string]connector.Connector),
		first:      "",
		uids:       make(map[uint32]string),
		subscriber: make(map[hash.Hash]map[string]Subscriber),
		choosers:   make([]uint8, 0)}
}

// Done release all connections and subscriber and release all resources.
func (b *Bricker) Done() {
	// release all subscriber
	for _, subs := range b.subscriber {
		for _, s := range subs {
			b.Unsubscribe(s)
		}
	}
	// Stop all bricker
	for name, _ := range b.connection {
		b.Release(name)
	}
}

// Internal method: read wait for a new event and forward it to the dispatcher.
func (b *Bricker) read(c connector.Connector, n string) {
	var ev *event.Event
	for {
		ev = c.Receive()
		if ev == nil {
			return // done, no more packets
		}
		ev.ConnectorName = n
		go b.dispatch(ev)
	}
}

// Internal method: write takes a event and send it to the right bricker (dispatch).
func (b *Bricker) write(e *event.Event) {
	if e != nil {
		if conn, ok := b.connection[e.ConnectorName]; ok {
			conn.Send(e)
		} else {
			e.Err = NewError(ErrorConnectorNameNotExists)
			go b.dispatch(e)
		}
	}
}

// Internal method: process dispatch the event to the right subscriber.
func (b *Bricker) dispatch(e *event.Event) {
	var h hash.Hash
	if e.Packet == nil { // without a packet, no subscriber could be determined
		go b.process(e, b.defaultsubscriber)
	} else {
		match := false
		for _, chooser := range b.choosers {
			h = hash.New(chooser, e.Packet.Head.Uid, e.Packet.Head.FunctionID)
			if s, ok := b.subscriber[h]; ok {
				match = true && (len(s) > 0)
				go func(ev *event.Event, subs map[string]Subscriber) {
					for _, s := range subs {
						go b.process(e, s)
					}
				}(e, s)
			}
		}
		if !match { // no subscriber hash matched against packet hash
			go b.process(e, b.defaultsubscriber)
		}
	}
}

// Internal method: process notify given subscriber.
func (b *Bricker) process(e *event.Event, sub Subscriber) {
	if sub == nil {
		return // no subscriber, no notify
	} else {
		sub.Notify(e)
		if !sub.Subscription().Callback { // not a callback, call only once
			b.Unsubscribe(sub)
		}
	}
}
