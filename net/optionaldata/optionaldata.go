// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Implementation of the en-/decoding of the optional data from a packet.
*/
package optionaldata

import (
	"encoding/binary"
	"fmt"
	"io"
)

// The optinal data of a ip packet.
type OptionalData []byte

// NewPayload helps to create a new payload object.
func New(d []byte) *OptionalData {
	var o OptionalData
	if d != nil {
		o = make(OptionalData, len(d))
		for i, v := range d {
			o[i] = v
		}
	} else {
		o = OptionalData{}
	}
	return &o
}

// Copy makes a real copy of the optional datas.
func (o *OptionalData) Copy() *OptionalData {
	if o == nil { // no optionaldata, no copy
		return nil
	}
	n := &OptionalData{}
	copy(*o, *n)
	return n
}

// Write writes the optinal data into a given writer.
func (o *OptionalData) Write(w io.Writer) error {
	if o == nil || len(*o) == 0 { // empty no write
		return nil
	}
	return binary.Write(w, binary.LittleEndian, *o)
}

// Read reads the optinal data out of a given reader.
func (o *OptionalData) Read(r io.Reader, l uint8) error {
	var err error
	var i int
	if l < 1 { // nothing to read
		return nil
	}
	*o = make(OptionalData, l)
	i = int(l)
	err = nil
	for i > 0 && err == nil {
		i, err = r.Read(*o)
		i = int(l) - i
	}
	return err
}

// Converts the optional data to a byte slice.
func (o *OptionalData) Bytes() []byte {
	if o == nil {
		return nil
	}
	buf := make([]byte, len(*o))
	for i, v := range *o {
		buf[i] = v
	}
	return buf
}

// String fulfill the stringer interface for the payload of a packet.
func (o *OptionalData) String() string {
	txt := ""
	if o == nil {
		return "<nil>"
	}
	for _, v := range *o {
		txt += fmt.Sprintf("0x%02x ", v)
	}
	return "Optional Data: [ " + txt + "]"
}
