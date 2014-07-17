// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Collection of subscriber for the Piezo Speaker Bricklet.
package piezospeaker

const (
	function_beep                = uint8(1)
	function_morse_code          = uint8(2)
	function_calibrate           = uint8(3)
	callback_beep_finished       = uint8(4)
	callback_morse_code_finished = uint8(5)
)
