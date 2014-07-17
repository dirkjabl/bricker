// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bricker

import (
	"bricker/connector"
)

// AttachConnector adds a named connector to the bricker.
// The name must be unique and should not used before.
func (b *Bricker) Attach(c connector.Connector, n string) error {
	if _, ok := b.connection[n]; ok { // name exists, no add
		return NewError(ErrorConnectorNameExists)
	}
	b.connection[n] = c
	if b.first == "" {
		b.first = n
	}
	go b.read(c, n) // start working for incoming events
	return nil
}

// ReleaseConnector take a connector from the bricker.
func (b *Bricker) Release(n string) error {
	if _, ok := b.connection[n]; !ok { // name does not exists
		return NewError(ErrorNoConnectorToRelease)
	}
	if n == b.first {
		b.first = ""
	} // TODO: search for a new connector as first
	delete(b.connection, n)
	return nil
}

// Internal method: computeConnectorsName try to compute the connectors name from the given parameter.
// The parameter could be a string with the name, a uid of a device (uint32) or nil, then the
// first registered connector will be used.
func (b *Bricker) computeConnectorsName(d interface{}) string {
	switch value := d.(type) {
	case string:
		return value
	case uint32:
		if n, ok := b.uids[value]; ok {
			return n
		} // when not ok, than first connector, if possible
	}
	return b.first
}
