// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package moisture

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
)

/*
SetMovingAverage creates a subscriber to set a length of the moving average.
The default value is 100.
*/
func SetMovingAverage(id string, uid uint32, a *Average, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SetMovingAverage"),
		Fid:        function_set_moving_average,
		Uid:        uid,
		Data:       a,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SetMovingAverageFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetMovingAverageFuture(brick *bricker.Bricker, connectorname string, uid uint32, a *Average) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetMovingAverage("setmovingaveragefuture"+device.GenId(), uid, a,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetMovingAverage creates a subscriber to get the length of the moving average.
func GetMovingAverage(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetMovingAverage"),
		Fid:        function_get_moving_average,
		Uid:        uid,
		Result:     &Average{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetMovingAverageFuture is a future pattern version for a synchronized calll of the subscriber.
// If an error occur, the result is nil.
func GetMovingAverageFuture(brick *bricker.Bricker, connectorname string, uid uint32) *Average {
	future := make(chan *Average)
	defer close(future)
	sub := GetMovingAverage("getmovingaveragefuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *Average = nil
			if err == nil {
				if value, ok := r.(*Average); ok {
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

/*
Average is the type for the averaging length.
A value 0 is turn the averaging completely off.
With less averaging the value of the moisture value contains more noise.
The range for the averaging is 0-100.
More information in the Tinkerforge documentation:
http://www.tinkerforge.com/en/doc/Software/Bricklets/Moisture_Bricklet_TCPIP.html#BrickletMoisture.get_moisture_value
*/
type Average struct {
	Value uint8
}

// FromPacket create from a packet a Average.
func (a *Average) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(a, p); err != nil {
		return err
	}
	return p.Payload.Decode(a)
}

// String fullfill stringer interface.
func (a *Average) String() string {
	txt := "Average "
	if a == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("Value: %d", a.Value)
	}
	return txt
}
