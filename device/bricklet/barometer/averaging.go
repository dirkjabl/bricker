// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package barometer

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
)

/*
SetAveraging creates a subscriber to set the different averaging parameters.
There is no moving average for the temperature.

The maximum length for the pressure average is 10,
for the temperature average is 255 and for the moving average is 25.
The default values are 10 for the normal averages and 25 for the moving average.

Setting the all three parameters to 0 will turn the averaging completely off.

This brings the data without delay but with much more noise.
*/
func SetAveraging(id string, uid uint32, a *Average, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SetAveraging"),
		Fid:        function_set_averaging,
		Uid:        uid,
		Data:       a,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SetAveragingFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetAveragingFuture(brick bricker.Bricker, connectorname string, uid uint32, a *Average) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetAveraging("setaveragingfuture"+device.GenId(), uid, a,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetAveraging creates a subscriber to get the different averaging values.
func GetAveraging(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetAveraging"),
		Fid:        function_get_averaging,
		Uid:        uid,
		Result:     &Average{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetAveragingFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func GetAveragingFuture(brick bricker.Bricker, connectorname string, uid uint32) *Average {
	future := make(chan *Average)
	defer close(future)
	sub := GetAveraging("getaveragingfuture"+device.GenId(), uid,
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

// Average is the type for the length of a averaging for the voltage value.
type Average struct {
	MovingPressure uint8
	Pressure       uint8
	Temperature    uint8
}

// FromPacket creates from a packet a Average.
func (a *Average) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(a, p); err != nil {
		return err
	}
	return p.Payload.Decode(a)
}

// String fullfill the stringer interface.
func (a *Average) String() string {
	txt := "Average "
	if a == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("[Moving Pressure: %d, Pressure: %d, Temperature: %d]",
			a.MovingPressure, a.Pressure, a.Temperature)
	}
	return txt
}
