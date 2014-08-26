// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package piezospeaker

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
	misc "github.com/dirkjabl/bricker/util/miscellaneous"
)

/*
Calibrate creates a subscriber to start a calibration.
This is only needed after reflashing the bricklet plugin.

The calibration plays each tone and measures the exact frequency back.
The result of the calibration is a mapping setting value and frequency.
This mapping is stored in the EEPROM and loaded on startup.

The callback result is true after calibration.
*/
func Calibrate(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "Calibrate"),
		Fid:        function_calibrate,
		Uid:        uid,
		Result:     &Calibration{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// CalibrateFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func CalibrateFuture(brick *bricker.Bricker, connectorname string, uid uint32) *Calibration {
	future := make(chan *Calibration)
	defer close(future)
	sub := Calibrate("calibratefuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *Calibration = nil
			if err == nil {
				if value, ok := r.(*Calibration); ok {
					v = value
				}
			}
			future <- v
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return nil
	}
	return <-future
}

// CalibrateFutureSimple is a simple future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func CalibrateFutureSimple(brick *bricker.Bricker, connectorname string, uid uint32) bool {
	c := CalibrateFuture(brick, connectorname, uid)
	if c == nil {
		return false
	}
	return c.Done
}

// Calibration is the type for the Calibrate result.
type Calibration struct {
	Done bool // is calibration done
}

// FromPacket converts the packet payload to the Cursor type.
func (c *Calibration) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(c, p); err != nil {
		return err
	}
	rc := new(CalibrationRaw)
	err := p.Payload.Decode(rc)
	if err == nil && rc != nil {
		c.FromCalibrationRaw(rc)
	}
	return err
}

// FromCalibrationRaw converts from a CalibrationRaw to a Calibration type.
func (c *Calibration) FromCalibrationRaw(rc *CalibrationRaw) {
	if rc == nil || c == nil {
		return
	}
	c.Done = misc.Uint8ToBool(rc.Done)
}

// String fullfill the stringer interface.
func (c *Calibration) String() string {
	txt := "Calibration "
	if c == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("[Done: %v]", c.Done)
	}
	return txt
}

// Copy creates a copy of the content.
func (c *Calibration) Copy() device.Resulter {
	if c == nil {
		return nil
	}
	return &Calibration{Done: c.Done}
}

// CalibrationRaw is the real de/encoding type for a Calibration type.
type CalibrationRaw struct {
	Done uint8
}
