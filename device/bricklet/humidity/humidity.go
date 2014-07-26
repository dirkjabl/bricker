// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Collection of subscriber for the Humidity Bricklet.
package humidity

const (
	function_get_humidity                        = uint8(1)
	function_get_analog_value                    = uint8(2)
	function_set_humidity_callback_period        = uint8(3)
	function_get_humidity_callback_period        = uint8(4)
	function_set_analog_value_callback_period    = uint8(5)
	function_get_analog_value_callback_period    = uint8(6)
	function_set_humidity_callback_threshold     = uint8(7)
	function_get_humidity_callback_threshold     = uint8(8)
	function_set_analog_value_callback_threshold = uint8(9)
	function_get_analog_value_callback_threshold = uint8(10)
	function_set_debounce_period                 = uint8(11)
	function_get_debounce_period                 = uint8(12)
	callback_humidity                            = uint8(13)
	callback_analog_value                        = uint8(14)
	callback_humidity_reached                    = uint8(15)
	callback_analog_value_reached                = uint8(16)
)
