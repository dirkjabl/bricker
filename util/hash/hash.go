// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Hashs for identify the packets, for which subscriber they are.
package hash

import (
	"crypto/md5"
	"fmt"
)

const (
	ChoosenNothing       = uint8(0)                              // nothing choosen
	ChoosenFunctionID    = uint8(1)                              // FunctionID choosen
	ChoosenUid           = uint8(2)                              // Uid choosen
	ChoosenFunctionIDUid = ChoosenFunctionID | ChoosenUid // FunctionID and Uid choosen
)

// A hash type for subscriptions
type Hash [md5.Size]byte

// New creates a hash with given values based on the choosen ones.
func New(choosen uint8, uid uint32, functionID uint8) Hash {
	t := ""
	if (choosen & ChoosenFunctionID) == ChoosenFunctionID {
		t += fmt.Sprintf("|Function-ID=%d", functionID)
	}
	if (choosen & ChoosenUid) == ChoosenUid {
		t += fmt.Sprintf("|Uid=%d", uid)
	}
	t += "|"
	return md5.Sum([]byte(t))
}

// Equal compares to hashes, if they are equal.
func (a Hash) Equal(b Hash) bool {
	for i, v := range b {
		if a[i] != v {
			return false
		}
	}
	return true
}

// String fullfill the stringer interface.
func (h Hash) String() string {
	t := "Hash ["
	for _, v := range h {
		t += fmt.Sprintf(" %d", v)
	}
	t += " ]"
	return t
}

// All returns a slice with all choosers.
func All() []uint8 {
	return []uint8{ChoosenNothing, ChoosenUid, ChoosenFunctionID, ChoosenFunctionIDUid}
}
