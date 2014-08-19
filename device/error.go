// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package device

const (
	ErrorUnknownPacket = iota
	ErrorUnknown
	ErrorNotMatchingSubscription
	ErrorNoMemoryForResult
	ErrorNoPacketToConvert
	ErrorNoEvent
)

// Error type for encoding or decoding packets for devices like bricks or bricklets.
type DeviceError struct {
	Code uint8
}

// Create simple the error
func NewDeviceError(code uint8) DeviceError {
	return DeviceError{code}
}

// Error gives a string representation back for the error code
func (e DeviceError) Error() string {
	switch e.Code {
	case ErrorUnknownPacket:
		return "Unknown ipconn packet to encode."
	case ErrorNotMatchingSubscription:
		return "Not the right subscription for matching."
	case ErrorNoMemoryForResult:
		return "No memory for the result values."
	case ErrorNoPacketToConvert:
		return "No packet for converting or notify."
	case ErrorNoEvent:
		return "No event for converting or notify."
	case ErrorUnknown:
		fallthrough
	default:
		return "Unknown error."
	}
}
