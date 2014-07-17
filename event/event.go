// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
This package implements the event type.
The event holds the packet, error, timestamp and a name of the connector.

An error is only set, if an error occured.
The packet could be a nil pointer.
Also the name of the connector is only set,
if the event is handled or created from a connector.

Before you use anything from the event, check if it exists.
*/
package event

import (
	"bricker/net/packet"
	"fmt"
	"time"
)

// Event for the bricker.
//
// Save the packet, error (if occured) and a timestamp.
type Event struct {
	Err           error
	TimeStamp     time.Time
	Packet        *packet.Packet
	ConnectorName string
}

// NewEvent creates an event with all content.
func New(err error, ts time.Time, p *packet.Packet) *Event {
	return &Event{Err: err, TimeStamp: ts, Packet: p}
}

// NewEventSimple creates an event with only a packet and an error.
// The timestamp will be created.
func NewSimple(err error, p *packet.Packet) *Event {
	return New(err, time.Now(), p)
}

// NewEventPacket creates an event with only a packet.
// The timestamp will be created and the error is set to nil.
func NewPacket(p *packet.Packet) *Event {
	return New(nil, time.Now(), p)
}

// NewError creates an event with only an error.
// The timestamp will be created and the packet is set to nil.
func NewError(e error) *Event {
	return New(e, time.Now(), nil)
}

// Copy makes a real deep copy of the event.
func (e *Event) Copy() *Event {
	if e == nil { // no event, no copy
		return nil
	}
	return &Event{Err: e.Err, TimeStamp: e.TimeStamp, Packet: e.Packet.Copy(), ConnectorName: e.ConnectorName}
}

// String fullfill the stringer interface.
func (e *Event) String() string {
	txt := "Event "
	if e == nil {
		txt += "[nil]"
	} else {
		if e.ConnectorName != "" {
			txt += "ConnectorName: " + e.ConnectorName + ", "
		}
		txt += "TimeStamp:" + e.TimeStamp.Format(time.RFC3339Nano) + ", "
		txt += fmt.Sprintf("%v, ", e.Packet)
		if e.Err != nil {
			txt += "Error: " + e.Err.Error()
		} else {
			txt += "Error: <nil>"
		}
	}
	return txt
}
