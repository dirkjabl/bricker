// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This ist a base virtual connector, this means,
// that this connector do not(!) connect to a real hardware.
package virtual

import (
	"github.com/dirkjabl/bricker/event"
	"github.com/dirkjabl/bricker/util/generator"
	"github.com/dirkjabl/bricker/util/hash"
)

// This is the generator type and it is a function.
// Every generator is a function, which takes an event and results an event.
// If the result is nil, it will be used as "no result".
type GeneratorFunc func(*event.Event) *event.Event

// This is the virtual connector. To use, create one and attach generators.
// The virtual connector takes any event and handle it.
// For handling you need generators. The generator are called if a matching
// packet comes in (per Send()).
// The matching works with hashes (hash.Hash).
// A fallback generator exists and could overwritten.
type Virtual struct {
	receive   chan *event.Event
	generator map[hash.Hash]GeneratorFunc
	fallback  GeneratorFunc
	serial    *generator.Generator
}

// New creates a new virtual connector.
func New() *Virtual {
	v := &Virtual{
		receive:   make(chan *event.Event, 20),
		serial:    generator.New(),
		generator: make(map[hash.Hash]GeneratorFunc)}
	v.DetachFallbackGenerator()
	_ = v.serial.Get() // pull zero
	return v
}

// AttachGenerator add a new generator to the connector.
// If a generator exists with the same hash, it will be overwritten.
func (v *Virtual) AttachGenerator(h hash.Hash, f GeneratorFunc) {
	v.generator[h] = f
}

// AttachFallbackGenerator change the existing fallback generator to
// the new given.
func (v *Virtual) AttachFallbackGenerator(f GeneratorFunc) {
	v.fallback = f
}

// DetachGenerator removes a generator.
func (v *Virtual) DetachGenerator(h hash.Hash) {
	delete(v.generator, h)
}

// DetachFallbackGenerator changes the fallback generator to
// the base fallbck generator without functionality.
func (v *Virtual) DetachFallbackGenerator() {
	v.AttachFallbackGenerator(Fallback)
}

// Send takes the given event and looks for a generator for it.
// If the generator returns a event, it will be put in the receive channel.
func (v *Virtual) Send(e *event.Event) {
	if e == nil { // no event, no processing
		return
	}
	if e.Packet != nil {
		e.Packet.Head.SetSequence(v.serial.Get8())
		e.Packet.Head.Length = e.Packet.ComputeLength()
	}
	f := v.getGen(e)
	r := f(e)
	if r != nil {
		v.receive <- r
	}
}

// Receive reads a event from the virtual connector (synchron).
func (v *Virtual) Receive() *event.Event {
	e, ok := <-v.receive
	if !ok {
		e = nil // done
	}
	return e

}

// Done detach all and reset the fallback generator.
// Close the receive channel.
// The virtual connector should not longer used.
func (v *Virtual) Done() {
	for h, _ := range v.generator {
		v.DetachGenerator(h)
	}
	v.DetachFallbackGenerator()
	close(v.receive)
}

// Fallback is the basic fallback generator.
// It reads a event and do nothing, the result is nil.
func Fallback(e *event.Event) *event.Event {
	return nil
}

// Internal method: getGen find a generator to run with this event.
func (v *Virtual) getGen(e *event.Event) GeneratorFunc {
	var h hash.Hash
	for _, c := range hash.All() {
		h = hash.New(c, e.Packet.Head.Uid, e.Packet.Head.FunctionID)
		if f, ok := v.generator[h]; ok {
			return f
		}
	}
	return Fallback
}
