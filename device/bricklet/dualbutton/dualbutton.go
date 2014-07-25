// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Collection of subscriber for the Dual Button Bricklet.
package dualbutton

const (
	function_set_led_state          = uint8(1)
	function_get_led_state          = uint8(2)
	function_get_button_state       = uint8(3)
	function_set_selected_led_state = uint8(5)
	callback_state_changed          = uint8(4)
	// Led States
	LedStateAutoToggleOn  = uint8(0) // Auto toggle enabled and LED on.
	LedStateAutoToggleOff = uint8(1) // Auto toggle enablde and LED off.
	LedStateOn            = uint8(2) // LED on (auto toggle is disabled).
	LedStateOff           = uint8(3) // LED off (auto toggle is disabled).
	// Button States
	ButtonStatePressed  = uint8(0) // Button pressed.
	ButtonStateReleased = uint8(1) // Button released.
)

// LedStateName results a string representation of the given led state.
func LedStateName(s uint8) string {
	switch s {
	case LedStateAutoToggleOn:
		return "Auto toggle enabled and LED on"
	case LedStateAutoToggleOff:
		return "Auto toggle enablde and LED off"
	case LedStateOn:
		return "LED on (auto toggle is disabled)"
	case LedStateOff:
		return "LED off (auto toggle is disabled)"
	default:
		return "Unknown"
	}
}

// ButtonStateName give back the corresponding string representation of the given button state.
func ButtonStateName(s uint8) string {
	switch s {
	case ButtonStatePressed:
		return "Button Pressed"
	case ButtonStateReleased:
		return "Button Released"
	default:
		return "Unknown"
	}
}
