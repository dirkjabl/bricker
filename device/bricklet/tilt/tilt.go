// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Collection of subscriber for the Tilt Bricklet.
package tilt

const (
	function_get_tilt_state                 = uint8(1)
	function_enable_tilt_state_callback     = uint8(2)
	function_disable_tilt_state_callback    = uint8(3)
	function_is_tilt_state_callback_enabled = uint8(4)
	callback_tilt_state                     = uint8(5)
)
