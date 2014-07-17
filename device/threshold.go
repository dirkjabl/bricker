// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package device

import (
	"bricker/net/packet"
	"fmt"
)

// Threshold states.
const (
	ThresholdTurnedOff  = 'x'
	ThresholdOutside    = 'o'
	ThresholdInside     = 'i'
	ThresholdSmallerMin = '<'
	ThresholdBiggerMin  = '>'
)

// Threshold type.
type Threshold struct {
	Option byte
	Min    int16
	Max    int16
}

// FromPacket convert the packet payload to the theshold type.
func (t *Threshold) FromPacket(p *packet.Packet) error {
	if err := CheckForFromPacket(t, p); err != nil {
		return err
	}
	return p.Payload.Decode(t)
}

// Name convert the threshold option to a readable string.
func (t *Threshold) Name() string {
	switch t.Option {
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

// String fullfill the stringer interface.
func (t *Threshold) String() string {
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
