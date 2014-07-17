// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Implementation of the networking errors.
*/
package errors

/*
Error Codes
*/
const (
	ErrorOK = iota
	ErrorINVALIDPARAMETER
	ErrorFUNCTIONNOTSUPPORTED
	ErrorUNKNOWN
	ErrorHeaderMissing
)

/*
IPConnError is a type for a error from brickd.
*/
type Error struct {
	Type uint8
}

func New(t uint8) *Error {
	return &Error{Type: t}
}

/*
Error gives a string representation as result.
*/
func (e *Error) Error() string {
	switch e.Type {
	case ErrorOK:
		return "Ok"
	case ErrorINVALIDPARAMETER:
		return "Invalid parameter"
	case ErrorFUNCTIONNOTSUPPORTED:
		return "Function not supported"
	case ErrorHeaderMissing:
		return "No header for packet, header needed."
	case ErrorUNKNOWN:
		fallthrough
	default:
		return "Unknown error"
	}
}

/*
String gives a string representation of the error.
*/
func (e *Error) String() string {
	return e.Error()
}
