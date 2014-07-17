// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Generator type for serial numbers.
package generator

// Generator type
type Generator struct {
	queue chan uint32
}

// New creates a generator for ids (uint32)
func New() *Generator {
	g := new(Generator)
	g.queue = make(chan uint32)
	go func(gen *Generator) {
		var i uint32 = 0
		for {
			gen.queue <- i
			i++
		}
	}(g)
	return g
}

// Get computes a new id.
func (g *Generator) Get() uint32 {
	return <-g.queue
}

// Get8 computes a new id (uint8)
func (g *Generator) Get8() uint8 {
	id := <-g.queue
	if id > 255 {
		id %= 256
	}
	return uint8(id)
}
