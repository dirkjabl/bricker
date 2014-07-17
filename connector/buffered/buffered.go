// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Implementation of a connector interface type with buffered channels.
// A buffered working connector. The buffer size for incoming and outgoing events could given as parameter.
package buffered

import (
	"github.com/dirkjabl/bricker/event"
	"github.com/dirkjabl/bricker/net"
	"github.com/dirkjabl/bricker/net/packet"
)

// ConnectorBuffered is the connector with to bufferd channels,
// one for reading (In) and one for Writing(Out).
// Every channel will controled by a go routine.
// The connector puts all of his readed packets into events in the In channel.
// The connector waits for packets to write out to the hardware on the Out channel.
// A close on the Quit channel let the bricker stops all go routines and disconnect to the hardware.
type ConnectorBuffered struct {
	conn *net.Net          // internal, the real connection
	seq  uint8             // internal, actual sequence number
	In   chan *event.Event // input channel, here the bricker put in the readed packets as events
	Out  chan *event.Event // output channel, here the bricker read out the events, which should be send
	Quit chan struct{}     // quit channel, if closed, the bricker stop working and release resources
}

// New creates the connector object with a connection to the given address (addr).
// The function takes to integers for the size of the input and output buffer (channels).
func New(addr string, inbuf, outbuf int) (*ConnectorBuffered, error) {
	conn, err := net.Dial(addr)
	if err != nil {
		return nil, err
	}
	cb := &ConnectorBuffered{conn: conn,
		seq:  0,
		In:   make(chan *event.Event, inbuf),
		Out:  make(chan *event.Event, outbuf),
		Quit: make(chan struct{})}

	go func() { cb.read() }()
	go func() { cb.write() }()

	return cb, nil
}

// NewBrickerUnbuffered creates a connector without bufferd channels.
// It is a buffered bricker with zero buffers.
func NewUnbuffered(addr string) (*ConnectorBuffered, error) {
	return New(addr, 0, 0)
}

// Send puts the given event into the channel for writing the packets to the hardware.
func (cb *ConnectorBuffered) Send(ev *event.Event) {
	cb.Out <- ev
}

// Receive reads a event out of the channel for the readed packets form the hardware.
// If Receive returns a nil event, no more events will follow. The channel is closed.
func (cb *ConnectorBuffered) Receive() *event.Event {
	e, ok := <-cb.In
	if !ok {
		e = nil // done
	}
	return e
}

// Done stops the bricker and release all connections
func (cb *ConnectorBuffered) Done() {
	close(cb.Quit)
	cb.conn.Close()
}

// read is a internal method. Method reads from the hardware connection and put the packet into the event.
func (cb *ConnectorBuffered) read() {
	var err error
	var pck *packet.Packet
	done := false
	defer close(cb.In)
	for {
		if done {
			return
		}
		pck, err = cb.conn.ReadPacket()
		go func(e error, p *packet.Packet) {
			if !done {
				ev := event.NewSimple(e, p)
				select {
				case cb.In <- ev:
				case <-cb.Quit:
					done = true
				}
			}
		}(err, pck)
	}
}

// write is a internal method. Method writes packets to the hardware connection.
func (cb *ConnectorBuffered) write() {
	defer close(cb.Out)
	var ev *event.Event
	for {
		select {
		case ev = <-cb.Out:
			if ev != nil && ev.Err == nil && ev.Packet != nil {
				cb.seq++
				ev.Packet.Head.SetSequence(cb.seq)
				ev.Packet.Head.Length = ev.Packet.ComputeLength()
				cb.conn.WritePacket(ev.Packet)
			}
		case <-cb.Quit:
			return
		}
	}
}
