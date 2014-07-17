// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Enumerate calls all devices controlled over one brick daemon.
package enumerate

import (
	"bricker/device"
	"bricker/device/identity"
	"bricker/device/name"
	"bricker/net/base58"
	"bricker/net/packet"
	"bricker/subscription"
	"fmt"
)

// Constants for the type of connection of a specific device.
const (
	function_enumerate       = uint8(254)
	callback_enumerate       = uint8(253)
	EnumerationTypeAvailable = iota
	EnumerationTypeNewlyConnected
	EnumerationTypeDisconneted
)

// Enumerate creates the subscriber for the enumeration callback.
// This subscriber collect all information of the connected devices.
// If onlySub is set, no calling packet is generated, only the handler is subscribed.
func Enumerate(id string, onlySub bool, handler func(device.Resulter, error)) *device.Device {
	var p *packet.Packet = nil
	if !onlySub {
		p = packet.NewSimpleHeaderOnly(0, function_enumerate, true)
	}
	e := device.New(device.FallbackId(id, "Enumerate"))
	e.SetSubscription(subscription.NewFid(callback_enumerate, p, true))
	e.SetResult(&Enumeration{})
	e.SetHandler(handler)
	return e
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
	uid := base58.Convert32(base58.Decode(e.Uid))
	cuid := base58.Convert32(base58.Decode(e.ConnectedUid))
	txt := "Enumeration ["
	txt += fmt.Sprintf("UID: %s (%d), ", e.Uid, uid)
	txt += fmt.Sprintf("Connected UID: %s (%d), ", e.ConnectedUid, cuid)
	txt += fmt.Sprintf("Position: %c, ", e.Position)
	txt += fmt.Sprintf("Hardware Version: %d.%d.%d, ", e.HardwareVersion[0],
		e.HardwareVersion[1], e.HardwareVersion[2])
	txt += fmt.Sprintf("Firmware Version: %d.%d.%d, ", e.FirmwareVersion[0],
		e.FirmwareVersion[1], e.FirmwareVersion[2])
	txt += "Name: " + name.Name(e.DeviceIdentifer) + ", "
	txt += e.EnumerationTypeString() + "]"
	return txt
}
