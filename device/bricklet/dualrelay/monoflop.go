// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dualrelay

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
	"github.com/dirkjabl/bricker/subscription"
	"github.com/dirkjabl/bricker/util/hash"
	misc "github.com/dirkjabl/bricker/util/miscellaneous"
)

// SetMonoflop creates the subscriber to set the monoflop timer value for specifed output relay.
func SetMonoflop(id string, uid uint32, m *Monoflops, handler func(device.Resulter, error)) *device.Device {
	fid := function_set_monoflop
	sm := device.New(device.FallbackId(id, "SetMonoflop"))
	p := packet.NewSimpleHeaderPayload(uid, fid, true, NewMonoflopsRaw(m))
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	sm.SetSubscription(sub)
	sm.SetResult(&device.EmptyResult{})
	sm.SetHandler(handler)
	return sm
}

// SetMonoflopFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetMonoflopFuture(brick *bricker.Bricker, connectorname string, uid uint32, m *Monoflops) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetMonoflop("setmonoflopfuture"+device.GenId(), uid, m,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetMonoflop creates a subscriber for getting the actual monoflop value.
func GetMonoflop(id string, uid uint32, r *Relay, handler func(device.Resulter, error)) *device.Device {
	fid := function_get_monoflop
	gm := device.New(device.FallbackId(id, "GetMonoflop"))
	p := packet.NewSimpleHeaderPayload(uid, fid, true, r)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	gm.SetSubscription(sub)
	gm.SetResult(&Monoflop{})
	gm.SetHandler(handler)
	return gm
}

// GetMonoflopFuture is a future pattern version for a synchronized all of the subscriber.
// If an error occur, the result is nil.
func GetMonoflopFuture(brick *bricker.Bricker, connectorname string, uid uint32, r *Relay) *Monoflop {
	future := make(chan *Monoflop)
	defer close(future)
	sub := GetMonoflop("getmonoflopfuture"+device.GenId(), uid, r,
		func(r device.Resulter, err error) {
			var v *Monoflop = nil
			if err == nil {
				if value, ok := r.(*Monoflop); ok {
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
MonoflopDone creates a subscriber for the monoflop done callback.
This callback is triggered whenever a monoflop timer reaches 0.
*/
func MonoflopDone(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := callback_monoflop_done
	md := device.New(device.FallbackId(id, "MonoflopDone"))
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, nil, true)
	md.SetSubscription(sub)
	md.SetResult(&Value{})
	md.SetHandler(handler)
	return md
}

// Monoflops is a type to set bitmask(4bit) based the time to hold the value.
// The monoflop mechanismus works only with output pins.
// Non output pins will be ignored.
// The time is given in ms.
type Monoflops struct {
	Relay uint8  // 1 or 2
	State bool   // true - on, false - off
	Time  uint32 // ms
}

// MonoflopsRaw is a en/decoding for Monoflops
type MonoflopsRaw struct {
	Relay uint8
	State uint8
	Time  uint32
}

// NewMonoflopsRaw creates a new MonoflopsRaw object from a Monoflops object.
func NewMonoflopsRaw(m *Monoflops) *MonoflopsRaw {
	if m == nil {
		return nil
	}
	mr := new(MonoflopsRaw)
	mr.Relay = m.Relay
	mr.Time = m.Time
	mr.State = misc.BoolToUint8(m.State)
	return mr
}

// Monoflop is the monflop timer value of a specified relay.
type Monoflop struct {
	State         bool   // true - on, false - off
	Time          uint32 // in ms
	TimeRemaining uint32 // in ms
}

// FromPacket creates a Monoflop from a packet.
func (m *Monoflop) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(m, p); err != nil {
		return err
	}
	return p.Payload.Decode(m)
}

// String fullfill the stringer interface.
func (m *Monoflop) String() string {
	txt := "Monoflop "
	if m == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("[State: %t, Time: %d ms, Time Remaining: %d ms]",
			m.State, m.Time, m.TimeRemaining)
	}
	return txt
}

// FromMonoflopRaw converts a en/decoding type MonoflopRaw into Monoflop.
func (m *Monoflop) FromMonoflopRaw(mr *MonoflopRaw) {
	if m == nil || mr == nil {
		return
	}
	m.Time = mr.Time
	m.TimeRemaining = mr.TimeRemaining
	m.State = misc.Uint8ToBool(mr.State)
}

// Relay type to define a specific relay.
type Relay struct {
	Value uint8 // 1 or 2
}

// MonoflopRaw is a de/encoding type for Monoflop.
type MonoflopRaw struct {
	State         uint8
	Time          uint32
	TimeRemaining uint32
}

// Value is a type for the MonoflopDone callback.
// Inside the struct is the relay and the state.
type Value struct {
	Relay uint8 // 1 or 2
	State bool  // true - on, false - false
}

// FromPacket converts the packet payoad to the Value type.
func (v *Value) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(v, p); err != nil {
	}
	vr := new(ValueRaw)
	err := p.Payload.Decode(vr)
	if err == nil && vr != nil {
		v.FromValueRaw(vr)
	}
	return err
}

// FromValueRaw converts a ValueRaw into Value.
func (v *Value) FromValueRaw(vr *ValueRaw) {
	if v == nil || vr == nil {
		return
	}
	v.Relay = vr.Relay
	v.State = misc.Uint8ToBool(vr.State)
}

// String fullfill the stringer interface.
func (v *Value) String() string {
	txt := "Value "
	if v == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("[Relay: %d, State: %t]", v.Relay, v.State)
	}
	return txt
}

// ValueRaw is a encoding/decoding type for Value
type ValueRaw struct {
	Relay uint8
	State uint8
}
