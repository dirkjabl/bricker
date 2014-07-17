// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Collection of subscriber for the LCD 20x4 Bricklet.
package lcd20x4

// Function and callback identifer
const (
	function_write_line               = uint8(1)
	function_clear_display            = uint8(2)
	function_backlight_on             = uint8(3)
	function_backlight_off            = uint8(4)
	function_is_backlight_on          = uint8(5)
	function_set_config               = uint8(6)
	function_get_config               = uint8(7)
	function_is_button_pressed        = uint8(8)
	function_set_custom_character     = uint8(11)
	function_get_custom_character     = uint8(12)
	function_set_default_text         = uint8(13)
	function_get_default_text         = uint8(14)
	function_set_default_text_counter = uint8(15)
	function_get_default_text_counter = uint8(16)
	callback_button_pressed           = uint8(9)
	callback_button_released          = uint8(10)
)
