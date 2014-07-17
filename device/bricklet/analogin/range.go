// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Collection of subscriber for the Analog In Bricklet.
package analogin

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
	"github.com/dirkjabl/bricker/subscription"
	"github.com/dirkjabl/bricker/util/hash"
)

/*
SetRange creates a subscriber to set the measurement range.
The default value is 0.
  0: Automatically switched
  1: 0V - 6.05V,  1.48mV resolution
  2: 0V - 10.32V,  2.52mV resolution
  3: 0V - 36.30V,  8.86mV resolution
  4: 0V - 45.00V,  11.25mV resolution
  5: 0V - 3.3V,  0.81mV resolution,
*/
func SetRange(id string, uid uint32, r *Range, handler func(device.Resulter, error)) *device.Device {
	fid := function_set_range
	sr := device.New(device.FallbackId(id, "SetRange"))
	p := packet.NewSimpleHeaderPayload(uid, fid, true, r)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	sr.SetSubscription(sub)
	sr.SetResult(&device.EmptyResult{})
	sr.SetHandler(handler)
	return sr
}

// SetRangeFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetRangeFuture(brick bricker.Bricker, connectorname string, uid uint32, r *Range) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetRange("setrangefuture"+device.GenId(), uid, r,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetRange creates a subscriber to get the measurement range value.
func GetRange(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := function_get_range
	gr := device.New(device.FallbackId(id, "GetRange"))
	p := packet.NewSimpleHeaderOnly(uid, fid, true)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	gr.SetSubscription(sub)
	gr.SetResult(&Range{})
	gr.SetHandler(handler)
	return gr
}

// GetRangeFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func GetRangeFuture(brick bricker.Bricker, connectorname string, uid uint32) *Range {
	future := make(chan *Range)
	defer close(future)
	sub := GetRange("getrangefuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *Range = nil
			if err == nil {
				if value, ok := r.(*Range); ok {
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

// Constants for the range.
const (
	RangeAutomaticallySwitched = 0
	Range0V6_05V1_48mV         = 1
	Range0V10_32V2_52mV        = 2
	Range0V36_30V8_86mv        = 3
	Range0V45V11_25mv          = 4
	Range0V3_3V0_81mV          = 5
)

// Range result type
type Range struct {
	Value uint8 // range identifer
}

// FromPacket creates a Range from a packet.
func (r *Range) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(r, p); err != nil {
		return err
	}
	return p.Payload.Decode(r)
}

// Name converts the range identifer value to a readable string.
func (r *Range) Name() string {
	switch r.Value {
	case RangeAutomaticallySwitched:
		return "Automatically switched"
	case Range0V6_05V1_48mV:
		return "0V - 6.05V,  1.48mV resolution"
	case Range0V10_32V2_52mV: // String fullfill the stringer interface.
		return "0V - 10.32V,  2.52mV resolution"
	case Range0V36_30V8_86mv:
		return "0V - 36.30V,  8.86mV resolution"
	case Range0V45V11_25mv:
		return "0V - 45.00V,  11.25mV resolution"
	case Range0V3_3V0_81mV:
		return "0V - 3.3V,  0.81mV resolution"
	default:
		return "Unknown"
	}
}

// String fullfill the stringer interface.
func (r *Range) String() string {
	txt := "Range "
	if r == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("[Value: %s (%d)]", r.Name(), r.Value)
	}
	return txt
}
