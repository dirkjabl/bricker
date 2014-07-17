// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The head object has all methods to work with the header of a packet.
package head

import (
	"bricker/net/errors"
	"encoding/binary"
	"fmt"
	"io"
)

// Header of the IP packet for the connection.
type Head struct {
	Uid                   uint32
	Length                uint8
	FunctionID            uint8
	SequenceAndOptions    uint8
	ErrorCodeAndFutureUse uint8
}

func New(uid uint32, length, functionID, seqAndOpt, errCode uint8) *Head {
	return &Head{Uid: uid, Length: length, FunctionID: functionID, SequenceAndOptions: seqAndOpt, ErrorCodeAndFutureUse: errCode}
}

// Copy makes a real copy of a header.
func (h *Head) Copy() *Head {
	if h == nil { // nothing to copy
		return nil
	}
	return &Head{Uid: h.Uid, Length: h.Length, FunctionID: h.FunctionID,
		SequenceAndOptions: h.SequenceAndOptions, ErrorCodeAndFutureUse: h.ErrorCodeAndFutureUse}
}

// Sequence read out the actual sequence of the header.
func (h *Head) Sequence() uint8 {
	return (h.SequenceAndOptions >> 4) & 0x0f
}

// SetSequence sets a new sequence to the header.
func (h *Head) SetSequence(seq uint8) {
	h.SequenceAndOptions |= (seq << 4) & 0xf0
}

// OptionResponseExpected read out, if this header is configured to expect a response after sending.
func (h *Head) OptionResponseExpected() bool {
	return (((h.SequenceAndOptions >> 3) & 1) == 1)
}

// SetOptionResponseExpected writes into the head, if or if not a response is expected.
func (h *Head) SetOptionResponseExpected(exp bool) {
	var v uint8
	if exp {
		v = 1
	} else {
		v = 0
	}
	h.SequenceAndOptions |= (v << 3) & 8
}

// OptionOther read out the other options from the header.
func (h *Head) OptionOther() uint8 {
	return h.SequenceAndOptions & 192
}

// ErrorCodeNbr read the error code number from the header.
func (h *Head) ErrorCodeNbr() uint8 {
	return (h.ErrorCodeAndFutureUse >> 6) & 3
}

// ErrorCode gets the actual error code number in a go error.
func (h *Head) ErrorCode() *errors.Error {
	return errors.New(h.ErrorCodeNbr())
}

// FutureUse reads the flags which reserved for future use.
func (h *Head) FutureUse() uint8 {
	return h.ErrorCodeAndFutureUse & 31
}

// Write writes the binary representation of the header in a given writer.
func (h *Head) Write(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, h)
}

// Read reads the binary representation of the header in the acutal header from the given reader.
func (h *Head) Read(r io.Reader) error {
	return binary.Read(r, binary.LittleEndian, h)
}

/*
String gives a representation of the single entries.
*/
func (h *Head) String() string {
	txt := ""
	txt += fmt.Sprintf("Uid: %d, ", h.Uid)
	txt += fmt.Sprintf("Length: %d, ", h.Length)
	txt += fmt.Sprintf("Function-ID: %d, ", h.FunctionID)
	txt += fmt.Sprintf("Sequence: %d, ", h.Sequence())
	txt += fmt.Sprintf("Response Expected: %v, ", h.OptionResponseExpected())
	txt += fmt.Sprintf("Other Options: 0x%02x, ", h.OptionOther())
	txt += fmt.Sprintf("Error Code: %v, ", h.ErrorCode())
	txt += fmt.Sprintf("Future Use: 0x%02x", h.FutureUse())
	return fmt.Sprintf("Header: [%s]", txt)
}
