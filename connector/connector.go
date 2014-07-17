// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Connector is a interface for all real and virtual connection.

This interface is the base connection to the brick daemon (brickd).

The Connector sends or receive packets from the brick daemon. The packets will encapsulate in events.
The Connector is a producer of events. To use this events their will be need consumer (subscriber).
*/
package connector

import (
	"bricker/event"
)

// Interface to the connector. It should send and receive events to or from the hardware.
// The connector works with the packets inside. The connector must be thread safe implementated.
// If no more packets can get, the receive method has to result a nil event.
// By sending a event, the connector has to fix the packet header length.
type Connector interface {
	Send(*event.Event)
	Receive() *event.Event
	Done()
}
