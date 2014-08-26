// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ambientlight

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
)

func GetIlluminance(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetIlluminance"),
		Fid:        function_get_illuminance,
		Uid:        uid,
		Result:     &Illuminance{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetIlluminanceFuture is a future pattern version for a synchronized calll of the subscriber.
// If an error occur, the result is nil.
func GetIlluminanceFuture(brick *bricker.Bricker, connectorname string, uid uint32) *Illuminance {
	future := make(chan *Illuminance)
	defer close(future)
	sub := GetIlluminance("getilluminancefuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *Illuminance = nil
			if err == nil {
				if value, ok := r.(*Illuminance); ok {
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

// Illuminance is a type for the illuminance value.
// The value has a range of 0 to 9000 and is given in Lux/10.
type Illuminance struct {
	Value uint16
}

// FromPacket creates from a packet a illuminance value.
func (i *Illuminance) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(i, p); err != nil {
		return err
	}
	return p.Payload.Decode(i)
}

// String fullfill the stringer interface.
func (i *Illuminance) String() string {
	txt := "Illuminance "
	if i == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("[Value: %d, Illuminance: %03.02 Lux]",
			i.Value, i.Float64())
	}
	return txt
}

// Copy creates a copy of the content.
func (i *Illuminance) Copy() device.Resulter {
	if i == nil {
		return nil
	}
	return &Illuminance{Value: i.Value}
}

// Float64 converts the illuminance value from int16 to float64.
func (i *Illuminance) Float64() float64 {
	f := float64(i.Value) / 10.00
	return f
}

// Float32 converts the illuminance value from int16 to float32.
func (i *Illuminance) Float32() float32 {
	f := float32(i.Value) / 10.00
	return f
}
