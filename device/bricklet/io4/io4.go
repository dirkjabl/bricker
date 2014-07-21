// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Collection of subscriber for the IO-4 Bricklet.
package io4

import (
	"fmt"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
)

const (
	function_set_value             = uint8(1)
	function_get_value             = uint8(2)
	function_set_configuration     = uint8(3)
	function_get_configuration     = uint8(4)
	function_get_edge_count        = uint8(14)
	function_set_monoflop          = uint8(10)
	function_get_monoflop          = uint8(11)
	function_set_selected_values   = uint8(13)
	function_set_edge_count_config = uint8(15)
	function_get_edge_count_config = uint8(16)
	function_set_debounce_period   = uint8(5)
	function_get_debounce_period   = uint8(6)
	function_set_interrupt         = uint8(7)
	function_get_interrupt         = uint8(8)
	callback_interrupt             = uint8(9)
	callback_monoflop_done         = uint8(12)
	// Bitmap-Mask
	Bit_0 = uint8(0x1) // Output mask bit 0
	Bit_1 = uint8(0x2) // Output mask bit 1
	Bit_2 = uint8(0x4) // Output mask bit 2
	Bit_3 = uint8(0x8) // Output mask bit 3
	// Direction-Character
	Direction_Output = 'o'
	Direction_Input  = 'i'
	// Edge count types
	EdgeCountType_Rising  = uint8(0) // default
	EdgeCountType_Falling = uint8(1)
	EdgeCountType_Both    = uint8(2)
)

// Pin is a type to select a special Pin.
type Pin struct {
	Value uint8
}

// Values is a type for setting or getting values.
// The bitmasks are 4bit wide.
type Values struct {
	SelectionMask uint8
	ValueMask     uint8
}

// FromPacket creates the values bitmasks from a packet.
func (v *Values) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(v, p); err != nil {
		return err
	}
	return p.Payload.Decode(v)
}

// String fullfill the stringer interface.
func (v *Values) String() string {
	txt := "Values "
	if v == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("[Selection Mask: %d (%s), Value Mask: %d (%s)]",
			v.SelectionMask, MaskToString(v.SelectionMask),
			v.ValueMask, MaskToString(v.ValueMask))
	}
	return txt
}

// MaskToString converts for better printing a mask (uin8 - 4bit) to a string.
func MaskToString(mask uint8) string {
	m := "0000"
	if (mask & Bit_3) == Bit_3 {
		m += "1"
	} else {
		m += "0"
	}
	if (mask & Bit_2) == Bit_2 {
		m += "1"
	} else {
		m += "0"
	}
	if (mask & Bit_1) == Bit_1 {
		m += "1"
	} else {
		m += "0"
	}
	if (mask & Bit_0) == Bit_0 {
		m += "1"
	} else {
		m += "0"
	}
	return m
}
