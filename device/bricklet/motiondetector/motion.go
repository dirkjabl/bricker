// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package motiondetector

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
)

/*
GetMotionDetected creates the subscriber for getting a detected motion.

More information about detection an time outs:
http://www.tinkerforge.com/en/doc/Hardware/Bricklets/Motion_Detector.html#motion-detector-bricklet-sensitivity-delay-block-time

There is a blue LED on the bricklet, this LED is on, if the bricklet is in the "motion detected" state.
*/
func GetMotionDetected(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetMotionDetected"),
		Fid:        function_get_motion_detected,
		Uid:        uid,
		Result:     &Motion{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetMotionDetectedFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func GetMotionDetectedFuture(brick bricker.Bricker, connectorname string, uid uint32) *Motion {
	future := make(chan *Motion)
	defer close(future)
	sub := GetMotionDetected("getmotiondetectedfuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *Motion = nil
			if err == nil {
				if value, ok := r.(*Motion); ok {
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

// GetMotionDetectedFutureSimple is a easy to use verion of GetMotionDetectedFuture.
// It returns only true if a motion is detected.
func GetMotionDetectedFutureSimple(brick bricker.Bricker, connectorname string, uid uint32) bool {
	m := GetMotionDetectedFuture(brick, connectorname, uid)
	result := false
	if m != nil {
		result = (m.Value == 1)
	}
	return result
}

// MotionDetected create the subscriber for the motion detected callback.
// There are no result inside, the handler is called, is a motion is detected.
func MotionDetected(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "MotionDetected"),
		Fid:        callback_motion_detected,
		Uid:        uid,
		Handler:    handler,
		IsCallback: true,
		WithPacket: false}.CreateDevice()
}

// Motion type. 1 - for a detection.
type Motion struct {
	Value uint8
}

// FromPacket creates from a packet a Motion value.
func (m *Motion) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(m, p); err != nil {
		return err
	}
	return p.Payload.Decode(m)
}

// String fullfill the stringer interface.
func (m *Motion) String() string {
	txt := "Motion "
	if m == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("[Value: %d Motion Detected: %t]", m.Value, (m.Value == 1))
	}
	return txt
}
