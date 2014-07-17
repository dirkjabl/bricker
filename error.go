// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bricker

// All known errors for a bricker.
const (
	ErrorUnknown = iota
	ErrorSubscriberExists
	ErrorNoSubscriberToRelease
	ErrorConnectorNameExists
	ErrorConnectorNameNotExists
	ErrorNoConnectorToRelease
)

// Error type for bricker.
type Error struct {
	Code uint8
}

// NewError create the error object.
func NewError(code uint8) Error {
	return Error{code}
}

// Error gives a string representation for the error code.
func (e Error) Error() string {
	switch e.Code {
	case ErrorConnectorNameExists:
		return "Connector with this name exists already."
	case ErrorSubscriberExists:
		return "Subscriber with this subscription exists already."
	case ErrorConnectorNameNotExists:
		return "Connector with this name does not exists."
	case ErrorNoConnectorToRelease:
		return "No connector with this name could be released."
	case ErrorNoSubscriberToRelease:
		return "No subscriber with this subscription could be released."
	case ErrorUnknown:
		fallthrough
	default:
		return "Unknown error."
	}
}
