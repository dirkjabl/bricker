// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lcd20x4

import (
	"bricker"
	"bricker/device"
	"bricker/net/packet"
	"bricker/subscription"
	"bricker/util/hash"
	"fmt"
)

// SetDefaultText creates a new subscriber to set the default text line (for one line 0-3).
// Here the line starts on 0 up to 19 characters (20 bytes).
// The default text will be showed, if the default text counter timed out (look at SetDefaultTextCounter).
func SetDefaultText(id string, uid uint32, dtl *DefaultTextLine, handler func(r device.Resulter, e error)) *device.Device {
	fid := function_set_default_text
	sdt := device.New(device.FallbackId(id, "SetDefaultText"))
	p := packet.NewSimpleHeaderPayload(uid, fid, true, dtl)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	sdt.SetSubscription(sub)
	sdt.SetResult(&device.EmptyResult{})
	sdt.SetHandler(handler)
	return sdt
}

// SetDefaultTextFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetDefaultTextFuture(brick *bricker.Bricker, connectorname string, uid uint32, dtl *DefaultTextLine) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetDefaultText("setdefaulttextfuture"+device.GenId(), uid, dtl,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	v := <-future
	return v
}

// GetDefaultText creates a new subscriber to get the default text on the given line.
func GetDefaultText(id string, uid uint32, l *Line, handler func(r device.Resulter, e error)) *device.Device {
	fid := function_get_default_text
	gdt := device.New(device.FallbackId(id, "GetDefaultText"))
	p := packet.NewSimpleHeaderPayload(uid, fid, true, l)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	gdt.SetSubscription(sub)
	gdt.SetResult(&Text{})
	gdt.SetHandler(handler)
	return gdt
}

// GetDefaultTextFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func GetDefaultTextFuture(brick *bricker.Bricker, connectorname string, uid uint32, l *Line) *Text {
	future := make(chan *Text)
	defer close(future)
	sub := GetDefaultText("getdefaulttextfuture"+device.GenId(), uid, l,
		func(r device.Resulter, err error) {
			var v *Text = nil
			if err != nil {
				if value, ok := r.(*Text); ok {
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

// SetDefaultTextCounter creates a subscribte to set the default text output counter (timeout).
// The counter is in mili seconds (ms). If the counter larger than 0 it decrement every ms.
// If the counter timed out (reaches 0) than the default text will be shown.
// A negative value for the counter stops the timer. Default value is -1.
func SetDefaultTextCounter(id string, uid uint32, c *Counter, handler func(r device.Resulter, e error)) *device.Device {
	fid := function_set_default_text_counter
	sdt := device.New(device.FallbackId(id, "SetDefaultTextCounter"))
	p := packet.NewSimpleHeaderPayload(uid, fid, true, c)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	sdt.SetSubscription(sub)
	sdt.SetResult(&device.EmptyResult{})
	sdt.SetHandler(handler)
	return sdt
}

// SetDefaultTextCounterFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetDefaultTextCounterFuture(brick *bricker.Bricker, connectorname string, uid uint32, c *Counter) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetDefaultTextCounter("setdefaulttextcounterfuture"+device.GenId(), uid, c,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	v := <-future
	return v
}

// GetDefaultTextCounter creates a subscriber to get the value from the counter.
// Default value is -1.
func GetDefaultTextCounter(id string, uid uint32, handler func(r device.Resulter, e error)) *device.Device {
	fid := function_get_default_text
	gdtc := device.New(device.FallbackId(id, "GetDefaultTextCounter"))
	p := packet.NewSimpleHeaderOnly(uid, fid, true)
	sub := subscription.New(hash.ChoosenFunctionIDUid, uid, fid, p, false)
	gdtc.SetSubscription(sub)
	gdtc.SetResult(&Counter{})
	gdtc.SetHandler(handler)
	return gdtc
}

// GetDefaultTextCounterFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func GetDefaultTextCounterFuture(brick *bricker.Bricker, connectorname string, uid uint32) *Counter {
	future := make(chan *Counter)
	defer close(future)
	sub := GetDefaultTextCounter("getdefaulttextcounterfuture"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v *Counter = nil
			if err != nil {
				if value, ok := r.(*Counter); ok {
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

// DefaultTextLine is the type for a full text line to display.
// The text should not be longer than 20 bytes.
type DefaultTextLine struct {
	Line uint8
	Text [20]byte
}

// FromPacket creates from a packet a DefaultTextLine.
func (dtl *DefaultTextLine) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(dtl, p); err != nil {
		return err
	}
	return p.Payload.Decode(dtl)
}

// String fullfill the stringer interface.
func (dtl *DefaultTextLine) String() string {
	txt := "LCD 20x4 Default Text Line "
	if dtl != nil {
		txt += fmt.Sprintf("[Line: %d Text: %s]", dtl.Line, dtl.Text)
	} else {
		txt += "[nil]"
	}
	return txt
}

// Line type for the line number of the default text.
type Line struct {
	Number uint8
}

// FromPacket creates from a packet a Line type.
func (l *Line) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(l, p); err != nil {
		return err
	}
	return p.Payload.Decode(l)
}

// String fullfill the stringer interface.
func (l *Line) String() string {
	txt := "LCD 20x4 Line "
	if l != nil {
		txt += fmt.Sprintf("[Number: %d]", l.Number)
	} else {
		txt += "[nil]"
	}
	return txt
}

// Text is the type for a text line (the characters).
// The text should not be longer than 20 bytes.
type Text struct {
	Text [20]byte
}

// FromPacket creates from a packet a Text type.
func (t *Text) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(t, p); err != nil {
		return err
	}
	return p.Payload.Decode(t)
}

// String fullfill the stringer interface.
func (t *Text) String() string {
	txt := "LCD 20x4 Text "
	if t != nil {
		txt += fmt.Sprintf("[Text: %s]", t.Text)
	} else {
		txt += "[nil]"
	}
	return txt
}

// Counter is the value type for the default text time out timer.
type Counter struct {
	Value int32 // every not negative value starts the timer
}

// FromPacket creates from a packet a Text type.
func (c *Counter) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(c, p); err != nil {
		return err
	}
	return p.Payload.Decode(c)
}

// String fullfill the stringer interface.
func (c *Counter) String() string {
	txt := "LCD 20x4 Counter "
	if c != nil {
		txt += fmt.Sprintf("[Value: %d]", c.Value)
	} else {
		txt += "[nil]"
	}
	return txt
}
