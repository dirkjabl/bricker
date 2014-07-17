// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Payload en-/decoding implentation.
*/
package payload

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

// The payload of a ip packet, maximal 64 bytes.
type Payload []byte

// NewPayload helps to create the new payload object.
func New(d []byte) *Payload {
	var p Payload
	if d != nil {
		p = make(Payload, len(d))
		for i, v := range d {
			p[i] = v
		}
	} else {
		p = Payload{}
	}
	return &p
}

// NewPayloadEncode creates a new payload out of the given data.
func NewPayloadEncode(r interface{}) *Payload {
	p := New(nil)
	err := p.Encode(r)
	if err != nil {
		panic(err)
	}
	return p
}

// Copy makes a real copy of the optional datas.
func (p *Payload) Copy() *Payload {
	if p == nil { // no payload, no copy
		return nil
	}
	n := &Payload{}
	copy(*p, *n)
	return n
}

// Write writes the payload into a given writer.
func (p *Payload) Write(w io.Writer) error {
	if p == nil || len(*p) == 0 { // empty no write
		return nil
	}
	return binary.Write(w, binary.LittleEndian, *p)
}

// Read reads the payload out of a given reader.
func (p *Payload) Read(r io.Reader, l uint8) error {
	if p == nil {
		return errors.New("Error: Payload could not be nil.")
	}
	var err error
	var i int
	if l > 64 {
		return errors.New("Warning: Length of bytes for reading of the payload are to long.")
	}
	if l < 1 { // nothing to read
		return nil
	}
	*p = make(Payload, l)
	i = int(l)
	err = nil
	for i > 0 && err == nil {
		i, err = r.Read(*p)
		i = int(l) - i
	}
	return err
}

// Converts the payload to a byte slice.
func (p *Payload) Bytes() []byte {
	if p == nil {
		return nil
	}
	buf := make([]byte, len(*p))
	for i, v := range *p {
		buf[i] = v
	}
	return buf
}

// Decode converts the bytes of the payload to the given result (structure).
func (p *Payload) Decode(r interface{}) error {
	buf := bytes.NewBuffer(p.Bytes())
	return binary.Read(buf, binary.LittleEndian, r)
}

// Encode converts the given parameter to the payload.
func (p *Payload) Encode(r interface{}) error {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, r)
	if err != nil {
		return err
	}
	err = p.Read(buf, uint8(buf.Len()))
	return err
}

// String fullfill the stringer interface for the payload of a packet.
func (p *Payload) String() string {
	txt := ""
	if p == nil {
		return "<nil>"
	}
	for _, v := range *p {
		txt += fmt.Sprintf("0x%02x ", v)
	}
	return "Payload: [ " + txt + "]"
}
