// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Packet implementation.

Creates, copy, modify, read, write and print packets for networking.
*/
package packet

import (
	"fmt"
	"github.com/dirkjabl/bricker/net/errors"
	"github.com/dirkjabl/bricker/net/head"
	"github.com/dirkjabl/bricker/net/optionaldata"
	"github.com/dirkjabl/bricker/net/payload"
	"io"
)

// TCP/IP packet type.
type Packet struct {
	Head         *head.Head
	Payload      *payload.Payload
	OptionalData *optionaldata.OptionalData
}

// New create a new Packet. If a header is given it computes the length.
func New(h *head.Head, p *payload.Payload, o *optionaldata.OptionalData) *Packet {
	packet := &Packet{Head: h, Payload: p, OptionalData: o}
	if packet.Head != nil {
		packet.Head.Length = packet.ComputeLength()
	}
	return packet
}

// Copy makes a real deep copy of the packet.
func (p *Packet) Copy() *Packet {
	n := &Packet{}
	n.Head = p.Head.Copy()
	n.Payload = p.Payload.Copy()
	n.OptionalData = p.OptionalData.Copy()
	return n
}

/*
ComputeLength computes the actuel length of the IP packet.
*/
func (p *Packet) ComputeLength() uint8 {
	var l int = 0
	if p.Head != nil {
		l = 8 // length of the header
	}
	if p.Payload != nil {
		l += len(*p.Payload)
	}
	if p.OptionalData != nil {
		l += len(*p.OptionalData)
	}
	return uint8(l)
}

// Write writes the parts of a IP packet.
func (p *Packet) Write(w io.Writer) error {
	if p.Head == nil {
		return errors.New(errors.ErrorHeaderMissing)
	}
	err := p.Head.Write(w)
	if err != nil {
		return err
	}
	if p.Head.Length > 8 && p.Payload != nil {
		err = p.Payload.Write(w)
		if err != nil {
			return err
		}
	}
	if p.Head.Length > 72 && p.OptionalData != nil {
		err = p.OptionalData.Write(w)
	}
	return err
}

// Read reads a ip paket in parts, the existing packet will be overwritten.
func (p *Packet) Read(r io.Reader) error {
	p.Head, p.Payload, p.OptionalData = &head.Head{}, &payload.Payload{}, &optionaldata.OptionalData{}
	err := p.Head.Read(r)
	if err != nil {
		return err
	}
	l := p.Head.Length - 8
	if l > 0 {
		pl := l
		if pl > 64 {
			pl = 64
			l = l - pl
		} else {
			l = 0 // smaller than 64, everything is payload, delete optinal data space
			p.OptionalData = nil
		}
		err = p.Payload.Read(r, pl)
		if err != nil {
			return err
		}
		if l > 0 {
			err = p.OptionalData.Read(r, l)
			if err != nil {
				return err
			}
		}
	} else { // no payload or optional data, free space
		p.Payload = nil
		p.OptionalData = nil
	}
	if p != nil && p.Head != nil && p.Head.ErrorCodeNbr() != errors.ErrorOK {
		return p.Head.ErrorCode()
	}
	return nil
}

// String fulfill the Stringer Interface
func (p *Packet) String() string {
	return fmt.Sprintf("[%v, %v, %v]", p.Head, p.Payload, p.OptionalData)
}

// ReadPacket simplify the read from a io.Reader, it creates the new packet
func ReadNew(r io.Reader) (*Packet, error) {
	p := &Packet{}
	err := p.Read(r)
	if err != nil {
		p = nil // delete packet, if a error occur
	}
	return p, err
}

// NewSimpleHeaderOnly create a packet with only a header without special options (simple).
func NewSimpleHeaderOnly(uid uint32, fid uint8, expect bool) *Packet {
	p := New(head.New(uid, 8, fid, 0, 0), nil, nil)
	p.Head.SetOptionResponseExpected(expect)
	return p
}

// NewSimpleHeaderPayload create a packet with a header and a payload (simple).
func NewSimpleHeaderPayload(uid uint32, fid uint8, expect bool, r interface{}) *Packet {
	p := New(head.New(uid, 8, fid, 0, 0), nil, nil)
	p.Payload = payload.NewPayloadEncode(r)
	p.Head.SetOptionResponseExpected(expect)
	p.Head.Length = p.ComputeLength()
	return p
}
