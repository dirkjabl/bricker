// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package device

import (
	"bricker/net/packet"
	"fmt"
)

// Type for callback period.
type Period struct {
	Value uint32
}

// FromPacket converts a packet to a period type.
func (pe *Period) FromPacket(p *packet.Packet) error {
	if err := CheckForFromPacket(pe, p); err != nil {
		return err
	}
	return p.Payload.Decode(pe)
}

// String fullfill the stringer interface.
func (p *Period) String() string {
	return fmt.Sprintf("Period [%d ms]", p.Value)
}
