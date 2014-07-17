// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package device

import (
	"fmt"
	"github.com/dirkjabl/bricker/net/packet"
)

// Type for the debounce period (ms) with which the threshold callback is triggered,
// if the thresold keeps being reached.
type Debounce struct {
	Value uint32
}

// ThresholdFromPacket convert the packet payload to the mode.
func (d *Debounce) FromPacket(p *packet.Packet) error {
	if err := CheckForFromPacket(d, p); err != nil {
		return err
	}
	return p.Payload.Decode(d)
}

// String fullfill the stringer interface.
func (d *Debounce) String() string {
	txt := "Debounce "
	if d == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("[Value: %d]", d.Value)
	}
	return txt
}
