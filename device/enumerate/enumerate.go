// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Enumerate calls all devices controlled over one brick daemon.
package enumerate

import (
	"fmt"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/device/identity"
	"github.com/dirkjabl/bricker/device/name"
	"github.com/dirkjabl/bricker/net/base58"
	"github.com/dirkjabl/bricker/net/packet"
	"github.com/dirkjabl/bricker/subscription"
)

// Constants for the type of connection of a specific device.
const (
	function_enumerate            = uint8(254)
	callback_enumerate            = uint8(253)
	EnumerationTypeAvailable      = uint8(0)
	EnumerationTypeNewlyConnected = uint8(1)
	EnumerationTypeDisconneted    = uint8(2)
)

// Enumerate creates the subscriber for the enumeration callback.
// This subscriber collect all information of the connected devices.
// If withoutPacket is set, no requesting packet is generated, only the handler is subscribed.
func Enumerate(id string, withoutPacket bool, handler func(device.Resulter, error)) *device.Device {
	var p *packet.Packet = nil
	if !withoutPacket {
		p = packet.NewSimpleHeaderOnly(0, function_enumerate, true)
	}
	return device.NewSubscriptionResulterHandler(device.FallbackId(id, "Enumerate"),
		subscription.NewFid(callback_enumerate, p, true),
		&Enumeration{}, handler)
}

// Enumeration result structure.
type Enumeration struct {
	identity.Identity
	EnumerationType uint8
}

// ResultFromPacket creates a new Enumeration based on a packet.
func (e *Enumeration) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(e, p); err != nil {
		return err
	}
	return p.Payload.Decode(e)
}

// EnumerationTypeString returns the string representation of the actual enumeration type.
func (e *Enumeration) EnumerationTypeString() string {
	switch e.EnumerationType {
	case EnumerationTypeAvailable:
		return "Device is available."
	case EnumerationTypeNewlyConnected:
		return "Device is newly connected."
	case EnumerationTypeDisconneted:
		return "Device is disconnected."
	default:
		return "Unknown."
	}
}

// Stringer interface fulfill.
func (e *Enumeration) String() string {
	txt := "Enumeration "
	if e == nil {
		txt += "[nil]"
	} else {
		uid := base58.Convert32(base58.Decode(e.Uid))
		cuid := base58.Convert32(base58.Decode(e.ConnectedUid))
		txt += fmt.Sprintf("[UID: %s (%d), ", e.Uid, uid)
		txt += fmt.Sprintf("Connected UID: %s (%d), ", e.ConnectedUid, cuid)
		txt += fmt.Sprintf("Position: %c, ", e.Position)
		txt += fmt.Sprintf("Hardware Version: %d.%d.%d, ", e.HardwareVersion[0],
			e.HardwareVersion[1], e.HardwareVersion[2])
		txt += fmt.Sprintf("Firmware Version: %d.%d.%d, ", e.FirmwareVersion[0],
			e.FirmwareVersion[1], e.FirmwareVersion[2])
		txt += "Name: " + name.Name(e.DeviceIdentifer) + ", "
		txt += fmt.Sprintf("State: %s (%d)]", e.EnumerationTypeString(), e.EnumerationType)
	}
	return txt
}
