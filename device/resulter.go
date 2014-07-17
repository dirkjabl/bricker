// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package device

import (
	"github.com/dirkjabl/bricker/net/errors"
	"github.com/dirkjabl/bricker/net/packet"
)

// Type for a event result.
type Resulter interface {
	FromPacket(*packet.Packet) error
	String() string
}

// CheckForFromPacket tests if a needed parameter is nil and returns an error.
func CheckForFromPacket(r Resulter, p *packet.Packet) error {
	if r == nil {
		return NewDeviceError(ErrorNoMemoryForResult)
	}
	if p == nil {
		return NewDeviceError(ErrorNoPacketToConvert)
	}
	return nil
}

// Emptry result type. Needful for Events/Result, where only the errors are checkable.
type EmptyResult struct{}

// FromPacket converts a packet.
// It checks only the errors.
func (er *EmptyResult) FromPacket(p *packet.Packet) error {
	err := CheckForFromPacket(er, p)
	if err != nil {
		return err
	}
	if p.Head != nil {
		err := p.Head.ErrorCode()
		if err != nil && err.Type != errors.ErrorOK {
			return err
		}
	}
	return nil
}

// String fullfill stringer interface.
func (er *EmptyResult) String() string {
	return "EmptyResult []"
}

// IsEmptyResultOk checks if an error occur by an empty result
func IsEmptyResultOk(r Resulter, e error) bool {
	var v bool = false
	if e == nil {
		if _, ok := r.(*EmptyResult); ok {
			return true
		}
	}
	return v
}
