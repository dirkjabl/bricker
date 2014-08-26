// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Type for a device identity result. Every device should have a subscriber getidentity.
package identity

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/device/name"
	"github.com/dirkjabl/bricker/net/base58"
	"github.com/dirkjabl/bricker/net/packet"
)

const (
	function_get_identity = uint8(255)
)

// GetIdentity creates the subscriber to get the identity of a device.
func GetIdentity(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetIdentity"),
		Fid:        function_get_identity,
		Uid:        uid,
		Result:     &Identity{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// Future is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func GetIdentityFuture(brick bricker.Bricker, connectorname string, uid uint32) *Identity {
	future := make(chan *Identity)
	defer close(future)
	sub := GetIdentity("getidentityfuture", uid,
		func(r device.Resulter, err error) {
			var v *Identity = nil
			if err == nil {
				if value, ok := r.(*Identity); ok {
					v = value
				}
			}
			future <- v
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return nil
	}
	v := <-future
	return v
}

// Result type for a getidentity subscriber.
type Identity struct {
	Uid             [8]byte
	ConnectedUid    [8]byte
	Position        byte
	HardwareVersion [3]uint8
	FirmwareVersion [3]uint8
	DeviceIdentifer uint16
}

// FromPacket fill up the values of the identity object from a net packet.
// Fullfill the resulter interface.
func (i *Identity) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(i, p); err != nil {
		return err
	}
	return p.Payload.Decode(i)
}

// Stringer interface fulfill.
func (i *Identity) String() string {
	uid := base58.Convert32(base58.Decode(i.Uid))
	cuid := base58.Convert32(base58.Decode(i.ConnectedUid))
	txt := "Identity ["
	txt += fmt.Sprintf("UID: %s (%d), ", i.Uid, uid)
	txt += fmt.Sprintf("Connected UID: %s (%d), ", i.ConnectedUid, cuid)
	txt += fmt.Sprintf("Position: %c, ", i.Position)
	txt += fmt.Sprintf("Hardware Version: %d.%d.%d, ", i.HardwareVersion[0],
		i.HardwareVersion[1], i.HardwareVersion[2])
	txt += fmt.Sprintf("Firmware Version: %d.%d.%d, ", i.FirmwareVersion[0],
		i.FirmwareVersion[1], i.FirmwareVersion[2])
	txt += "Name: " + name.Name(i.DeviceIdentifer) + "]"
	return txt
}

// Copy creates a copy of the content.
func (i *Identity) Copy() device.Resulter {
	if i == nil {
		return nil
	}
	return &Identity{
		Uid:             i.Uid,
		ConnectedUid:    i.ConnectedUid,
		Position:        i.Position,
		HardwareVersion: i.HardwareVersion,
		FirmwareVersion: i.FirmwareVersion,
		DeviceIdentifer: i.DeviceIdentifer}
}

// IntUid returns the Uid from the identity values as uint32.
// Is the identity not filled, the return will be 0.
func (i *Identity) IntUid() uint32 {
	if i == nil {
		return 0
	}
	return base58.Convert32(base58.Decode(i.Uid))
}

// Is checks if the given device identifer equals the existing identifer.
func (i *Identity) Is(deviceidentifer uint16) bool {
	if i == nil {
		return false
	}
	return (deviceidentifer == i.DeviceIdentifer)
}
