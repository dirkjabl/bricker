// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Collection of subscriber for the Barometer Bricklet.
package barometer

const (
	function_get_air_pressure                    = uint8(1)
	function_get_altitude                        = uint8(2)
	function_set_reference_air_pressure          = uint8(13)
	function_get_reference_air_pressure          = uint8(19)
	function_get_chip_temperature                = uint8(14)
	function_set_averaging                       = uint8(20)
	function_get_averaging                       = uint8(21)
	function_set_air_pressure_callback_period    = uint8(3)
	function_get_air_pressure_callback_period    = uint8(4)
	function_set_altitude_callback_period        = uint8(5)
	function_get_altitude_callback_period        = uint8(6)
	function_set_air_pressure_callback_threshold = uint8(7)
	function_get_air_pressure_callback_threshold = uint8(8)
	function_set_altitude_callback_threshold     = uint8(9)
	function_get_altitude_callback_threshold     = uint8(10)
	function_set_debounce_period                 = uint8(11)
	function_get_debounce_period                 = uint8(12)
	callback_air_pressure                        = uint8(15)
	callback_altitude                            = uint8(16)
	callback_air_pressure_reached                = uint8(17)
	callback_altitude_reached                    = uint8(18)
)
