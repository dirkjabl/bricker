// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tilt

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
	"github.com/dirkjabl/bricker/subscription"
	"github.com/dirkjabl/bricker/util/hash"
	misc "github.com/dirkjabl/bricker/util/miscellaneous"
)

// GetTiltState creates a subscriber to get the current tilt state.
func GetTiltState(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := function_get_tilt_state
	gts := device.New(device.FallbackId(id, "GetTiltState"))
	p := packet.NewSimpleHeaderOnly(uid, fid, true)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	gts.SetSubscription(sub)
	gts.SetResult(&TiltState{})
	gts.SetHandler(handler)
	return gts
}

// GetTiltStateFuture is a future pattern version for a synchronized calll of the subscriber.
// If an error occur, the result is nil.
func GetTiltStateFuture(brick *bricker.Bricker, connectorname string, uid uint32) *TiltState {
	future := make(chan *TiltState)
	defer close(future)
	sub := GetTiltState("gettiltstatefuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *TiltState = nil
			if err == nil {
				if value, ok := r.(*TiltState); ok {
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

// EnableTiltStateCallback creates a subscriber to enable the TiltStateChanged callback.
func EnableTiltStateCallback(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := function_enable_tilt_state_callback
	etsc := device.New(device.FallbackId(id, "EnableTiltStateCallback"))
	p := packet.NewSimpleHeaderOnly(uid, fid, true)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	etsc.SetSubscription(sub)
	etsc.SetResult(&device.EmptyResult{})
	etsc.SetHandler(handler)
	return etsc
}

// EnableTiltStateCallbackFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func EnableTiltStateCallbackFuture(brick *bricker.Bricker, connectorname string, uid uint32) bool {
	future := make(chan bool)
	defer close(future)
	sub := EnableTiltStateCallback("enabletiltstatecallbackfuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// DisableTiltStateCallback creates a subscriber to disable the TiltStateChanged callback.
func DisableTiltStateCallback(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := function_disable_tilt_state_callback
	dtsc := device.New(device.FallbackId(id, "EnableTiltStateCallback"))
	p := packet.NewSimpleHeaderOnly(uid, fid, true)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	dtsc.SetSubscription(sub)
	dtsc.SetResult(&device.EmptyResult{})
	dtsc.SetHandler(handler)
	return dtsc
}

// DisableTiltStateCallbackFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func DisableTiltStateCallbackFuture(brick *bricker.Bricker, connectorname string, uid uint32) bool {
	future := make(chan bool)
	defer close(future)
	sub := DisableTiltStateCallback("disabletiltstatecallbackfuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// IsTiltStateCallbackEnabled creates a subscriber for calling, if the TiltStateChanged callback is enabled.
func IsTiltStateCallbackEnabled(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := function_is_tilt_state_callback_enabled
	itsce := device.New(device.FallbackId(id, "IsTiltStateCallbackEnabled"))
	p := packet.NewSimpleHeaderOnly(uid, fid, true)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	itsce.SetSubscription(sub)
	itsce.SetResult(&Enabled{})
	itsce.SetHandler(handler)
	return itsce
}

// IsTiltStateCallbackEnabledFuture is a future pattern version for a synchronized calll of the subscriber.
// If an error occur, the result is nil.
func IsTiltStateCallbackEnabledFuture(brick *bricker.Bricker, connectorname string, uid uint32) *Enabled {
	future := make(chan *Enabled)
	defer close(future)
	sub := IsTiltStateCallbackEnabled("istiltstatecallbackenabledfuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *Enabled = nil
			if err == nil {
				if value, ok := r.(*Enabled); ok {
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

// TiltStateChanged creates a subscriber which is called every time the tilt state changed.
func TiltStateChanged(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	fid := callback_tilt_state
	tsc := device.New(device.FallbackId(id, "TiltStateChanged"))
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, nil, true)
	tsc.SetSubscription(sub)
	tsc.SetResult(&TiltState{})
	tsc.SetHandler(handler)
	return tsc
}

/*
TiltState is a type for the tilt state.

The state can either be

    0 = Closed: The ball in the tilt switch closes the circuit.
    1 = Open: The ball in the tilt switch does not close the circuit.
    2 = Closed Vibrating: The tilt switch is in motion (rapid change between open and close).
*/
type TiltState struct {
	Value uint8
}

// FromPacket create from a packet a TiltState.
func (t *TiltState) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(t, p); err != nil {
		return err
	}
	return p.Payload.Decode(t)
}

// Name gives a readable representation of the tilt state as string.
func (t *TiltState) Name() string {
	if t == nil {
		return ""
	}
	switch t.Value {
	case TiltClosed:
		return "Closed: The ball in the tilt switch closes the circuit."
	case TiltOpen:
		return "Open: The ball in the tilt switch does not close the circuit."
	case TiltClosedVibrating:
		return "Closed Vibrating: The tilt switch is in motion (rapid change between open and close)."
	default:
		return "Unknown"
	}
}

// String fullfill the stringer interface.
func (t *TiltState) String() string {
	txt := "Tilt state "
	if t == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("[Value: %s (%d)]", t.Name(), t.Value)
	}
	return txt
}

// Enabled is a type for showing if the TiltStateChanged callback is enabled or disabled.
type Enabled struct {
	Value bool // true - enabled, false - disabled
}

// FromPacket converts the packet payload to the Enabled type.
func (e *Enabled) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(e, p); err != nil {
		return err
	}
	er := new(EnabledRaw)
	err := p.Payload.Decode(er)
	if err == nil && er != nil {
		e.FromEnabledRaw(er)
	}
	return err
}

// String fullfill the stringer interface.
func (e *Enabled) String() string {
	txt := "Enabled "
	if e == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("[Value: %t]", e.Value)
	}
	return txt
}

// FromEnabledRaw converts the EnabledRaw into a Enabled.
func (e *Enabled) FromEnabledRaw(er *EnabledRaw) {
	if e == nil || er == nil {
		return
	}
	e.Value = misc.Uint8ToBool(er.Value)
}

// EnabledRaw is the real de/encoding type for a Enabled.
type EnabledRaw struct {
	Value uint8
}
