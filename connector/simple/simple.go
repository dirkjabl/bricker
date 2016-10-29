// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Simple implementation of a connector interface type with mutex locks.
package simple

import (
	"github.com/dirkjabl/bricker/connector"
	"github.com/dirkjabl/bricker/event"
	"github.com/dirkjabl/bricker/net"
	"sync"
)

// The simple connector type.
type ConnectorSimple struct {
	conn  *net.Net
	seq   *connector.Sequence
	rlock *sync.Mutex
	wlock *sync.Mutex
}

// New creates a simple connector with read and write locks.
func New(addr string) (*ConnectorSimple, error) {
	conn, err := net.Dial(addr)
	if err != nil {
		return nil, err
	}
	cs := &ConnectorSimple{
		conn:  conn,
		rlock: new(sync.Mutex),
		wlock: new(sync.Mutex),
		seq:   new(connector.Sequence)}
	return cs, nil
}

// Send take the packet out of the event, and write it with a write lock to the hardware connection.
func (cs *ConnectorSimple) Send(ev *event.Event) {
	if ev == nil || ev.Packet == nil { // no packet, no send
		return
	}
	cs.wlock.Lock()
	defer cs.wlock.Unlock()
	ev.Packet.Head.SetSequence(cs.seq.GetSequence())
	ev.Packet.Head.Length = ev.Packet.ComputeLength()
	cs.conn.WritePacket(ev.Packet)
}

// Receive reads a packet from the hardware connection with a read lock, put it in a event and return it.
func (cs *ConnectorSimple) Receive() *event.Event {
	cs.rlock.Lock()
	defer cs.rlock.Unlock()
	pck, err := cs.conn.ReadPacket()
	ev := event.NewSimple(err, pck)
	return ev
}

// Done closes the hardware connection
func (cs *ConnectorSimple) Done() {
	cs.conn.Close()
}
