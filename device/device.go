// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* Every device is a collection of subscriptions and
every of them should fullfill the subscriber interface. */
package device

import (
	"fmt"
	"github.com/dirkjabl/bricker/event"
	"github.com/dirkjabl/bricker/subscription"
	"github.com/dirkjabl/bricker/util/generator"
)

// internal generator for ids
var gen *generator.Generator

// init device
func init() {
	gen = generator.New()
}

// Base type of a device.
type Device struct {
	id           string
	subscription *subscription.Subscription
	result       Resulter
	handler      func(Resulter, error)
}

// New creates a device.
func New(id string) *Device {
	d := &Device{id: id}
	return d
}

// Creates a device with a given subscription, resulter and handler and an id.
func NewSubscriptionResulterHandler(id string, sub *subscription.Subscription, result Resulter, handler func(Resulter, error)) *Device {
	d := New(FallbackId(id, "Device"))
	d.SetSubscription(sub)
	d.SetResult(result)
	d.SetHandler(handler)
	return d
}

// Subscription returns the stored subscription. (Getter)
func (d *Device) Subscription() *subscription.Subscription {
	if d == nil {
		return nil
	}
	return d.subscription
}

// SetSubscription stores a given subscription. (Setter)
func (d *Device) SetSubscription(sub *subscription.Subscription) {
	if d != nil {
		d.subscription = sub
	}
}

// Result returns the stored result. (Getter)
func (d *Device) Result() Resulter {
	if d == nil {
		return nil
	}
	return d.result
}

// SetResult set the given result. (Setter)
func (d *Device) SetResult(r Resulter) {
	if d != nil {
		d.result = r
	}
}

// Handler get the stored handler routine. (Getter)
func (d *Device) Handler() func(Resulter, error) {
	if d == nil {
		return nil
	}
	return d.handler
}

// SetHandler stores a given handler routine. (Setter)
func (d *Device) SetHandler(h func(Resulter, error)) {
	if d != nil {
		d.handler = h
	}
}

// Id returns the subscriber id. (Getter)
func (d *Device) Id() string {
	if d == nil {
		return ""
	}
	return d.id
}

// Generator for the id.
func GenId() string {
	return fmt.Sprintf("%04d", gen.Get())
}

// FallbackId is a fallback test for ids, creates one when needed.
func FallbackId(id, fb string) string {
	if id == "" {
		id = fb + " " + GenId()
	}
	return id
}

// Notify process about the result event.
// The handler routine (callback or event listner) gets only an error value.
// The result of the event is in the resulter value of the device.
func (d *Device) Notify(e *event.Event) {
	var err error = e.Err
	if e.Packet != nil && e.Err == nil && e.Packet.Head.FunctionID == d.Subscription().FunctionID {
		err = d.Result().FromPacket(e.Packet)
		d.Handler()(d.Result(), err)
	} else {
		if err == nil {
			err = NewDeviceError(ErrorNotMatchingSubscription)
		}
		d.Handler()(nil, err)
	}
}

// String fullfill the stringer interface.
func (d *Device) String() string {
	return fmt.Sprintf("Device [Id: %s, Subscription: %v, Result: %v]", d.id, d.subscription, d.result)
}
