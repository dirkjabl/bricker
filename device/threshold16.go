// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package device

import (
	"fmt"
	"github.com/dirkjabl/bricker/net/packet"
)

// Threshold type for 16bit values.
type Threshold16 struct {
	Option byte
	Min    int16
	Max    int16
}

// FromPacket convert the packet payload to the theshold type.
func (t *Threshold16) FromPacket(p *packet.Packet) error {
	if err := CheckForFromPacket(t, p); err != nil {
		return err
	}
	return p.Payload.Decode(t)
}

// Name convert the threshold option to a readable string.
func (t *Threshold16) Name() string {
	if t == nil { // no object, no option, no option name
		return ""
	}
	return ThresholdName(t.Option)
}

// String fullfill the stringer interface.
func (t *Threshold16) String() string {
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
