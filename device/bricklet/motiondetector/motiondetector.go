// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Collection of subscriber for the Motion Detector Bricklet.
package motiondetector

import (
	"github.com/dirkjabl/bricker/device"
)

const (
	function_get_motion_detected   = uint8(1)
	callback_motion_detected       = uint8(2)
	callback_detection_cycle_ended = uint8(3)
)

/*
DetectionCycleEnded creates a subscriber for the detection cyle ended callback.

The handler is called when after a detection motion the detection cycle ended.
After the detection cycle a new motion can be detected after approximately 2 seconds.
*/
func DetectionCycleEnded(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "DetectionCycleEnded"),
		Fid:        callback_detection_cycle_ended,
		Uid:        uid,
		Handler:    handler,
		IsCallback: true,
		WithPacket: false}.CreateDevice()
}
