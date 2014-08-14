// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package device

// Threshold states.
const (
	ThresholdTurnedOff  = 'x'
	ThresholdOutside    = 'o'
	ThresholdInside     = 'i'
	ThresholdSmallerMin = '<'
	ThresholdBiggerMin  = '>'
)

// ThresholdName give back a string representation of the threshold option.
func ThresholdName(option byte) string {
	switch option {
	case ThresholdTurnedOff:
		return "Threshold Callback is turned off"
	case ThresholdOutside:
		return "Threshold Callback is triggered when the temperature is outside the min and max values"
	case ThresholdInside:
		return "Callback is triggered when the temperature is inside the min and max values"
	case ThresholdSmallerMin:
		return "Callback is triggered when the temperature is smaller than the min value (max is ignored)"
	case ThresholdBiggerMin:
		return "Callback is triggered when the temperature is greater than the min value (max is ignored)"
	default:
		return "Unknown"
	}
}
