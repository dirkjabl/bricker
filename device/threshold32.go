// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package device

import (
	"fmt"
	"github.com/dirkjabl/bricker/net/packet"
)

// Theshold is a own type definition. Here the values for min and max are 32bit sized.
type Threshold32 struct {
	Option byte
	Min    int32
	Max    int32
}

// FromPacket convert the packet payload to the theshold type.
func (t *Threshold32) FromPacket(p *packet.Packet) error {
	if err := CheckForFromPacket(t, p); err != nil {
		return err
	}
	return p.Payload.Decode(t)
}

// Name converts the threshold option to a readable string.
func (t *Threshold32) Name() string {
	if t == nil { // no object, no option, no option name
		return ""
	}
	return ThresholdName(t.Option)
}

// String fullfill the stringer interface.
func (t *Threshold32) String() string {
	txt := "Threshold "
	if t == nil {
		return txt + "[nil]"
	}
	txt += "[Option: " + t.Name()
	if t.Option == ThresholdOutside || t.Option == ThresholdInside {
		txt += fmt.Sprintf(", Min: %d, Max: %d", t.Min, t.Max)
	} else if t.Option == ThresholdBiggerMin || t.Option == ThresholdSmallerMin {
		txt += fmt.Sprintf(", Min: %d", t.Min)
	}
	return txt + "]"
}

// Copy creates a copy of the content.
func (t *Threshold32) Copy() Resulter {
	if t == nil {
		return nil
	}
	return &Threshold32{
		Option: t.Option,
		Min:    t.Min,
		Max:    t.Max}
}
