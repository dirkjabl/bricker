// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Collection of subscriber for the Dual Relay Bricklet.
package dualrelay

const (
	function_set_state          = uint8(1)
	function_get_state          = uint8(2)
	function_set_monoflop       = uint8(3)
	function_get_monoflop       = uint8(4)
	function_set_selected_state = uint8(6)
	callback_monoflop_done      = uint8(5)
)
