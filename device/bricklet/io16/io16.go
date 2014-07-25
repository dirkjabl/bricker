// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Collection of subscriber for the IO-16 Bricklet.
package io16

const (
	function_set_port               = uint8(1)
	function_get_port               = uint8(2)
	function_set_port_configuration = uint8(3)
	function_get_port_configuration = uint8(4)
	function_get_edge_count         = uint8(14)
	function_set_port_monoflop      = uint8(10)
	function_get_port_monoflop      = uint8(11)
	function_set_selected_values    = uint8(13)
	function_set_edge_count_config  = uint8(15)
	function_get_edge_count_config  = uint8(16)
	function_set_debounce_period    = uint8(5)
	function_get_debounce_period    = uint8(6)
	function_set_port_interrupt     = uint8(7)
	function_get_port_interrupt     = uint8(8)
	callback_interrupt              = uint8(9)
	callback_monoflop_done          = uint8(12)
)
