// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
TCP/IP Networking connection.
*/
package net

import (
	"github.com/dirkjabl/bricker/net/packet"
	"net"
)

// IPConn holds the connection and the address for that connection and has methods to read and write packets.
// No locks or anything to make it thread save. This is the raw structure for communication.
type Net struct {
	Address string
	Conn    *net.TCPConn
}

// Dial is a shortcut to IPConn.Dial.
func Dial(addr string) (*Net, error) {
	conn := new(Net)
	conn.Address = addr
	err := conn.Dial()
	return conn, err
}

// Dial makes a connection to the given IPConn object.
func (c *Net) Dial() error {
	var conn *net.TCPConn

	taddr, err := net.ResolveTCPAddr("tcp", c.Address)
	if err != nil {
		return err
	}
	conn, err = net.DialTCP("tcp", nil, taddr)
	if err != nil {
		return err
	}
	c.Conn = conn
	return nil
}

// WritePacket sends one packet over the network connection to the brickd.
func (c *Net) WritePacket(p *packet.Packet) error {
	return p.Write(c.Conn)
}

// ReadPacket receive one packet from the network connection (brickd).
func (c *Net) ReadPacket() (*packet.Packet, error) {
	p, err := packet.ReadNew(c.Conn)
	return p, err
}

// Close disconnected the connection.
func (c *Net) Close() {
	c.Conn.Close()
}
