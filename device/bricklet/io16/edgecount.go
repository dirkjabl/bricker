// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io16

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
	misc "github.com/dirkjabl/bricker/util/miscellaneous"
)

// GetEdgeCount creates a subscriber to get the actual value of the edge counter.
// Supports only edge counts on port a.
// With ResetCounter set to true, the edge counter is reseted directly after call.
func GetEdgeCount(id string, uid uint32, ec *EdgeCount, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetEdgeCount"),
		Fid:        function_get_edge_count,
		Uid:        uid,
		Result:     &EdgeCounts{},
		Data:       NewEdgeCountRaw(ec),
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetEdgeCountFuture is a future pattern version for a synchronized all of the subscriber.
// If an error occur, the result is nil.
func GetEdgeCountFuture(brick *bricker.Bricker, connectorname string, uid uint32, ec *EdgeCount) *EdgeCounts {
	future := make(chan *EdgeCounts)
	defer close(future)
	sub := GetEdgeCount("getedgecountfuture"+device.GenId(), uid, ec,
		func(r device.Resulter, err error) {
			var v *EdgeCounts = nil
			if err == nil {
				if value, ok := r.(*EdgeCounts); ok {
					v = value
				}
			}
			future <- v
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return nil
	}
	return <-future
}

// SetEdgeCountConfig creates the subscriber to configure the edge counter for the selected pins.
// The debounce time is given in ms (default 100ms).
// Configuring an edge counter resets its value to 0..
// Default edge type is 0 (rising).
func SetEdgeCountConfig(id string, uid uint32, e *EdgeCountConfigs, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SetEdgeCountConfig"),
		Fid:        function_set_edge_count_config,
		Uid:        uid,
		Data:       e,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SetEdgeCountConfigFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetEdgeCountConfigFuture(brick *bricker.Bricker, connectorname string, uid uint32, e *EdgeCountConfigs) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetEdgeCountConfig("setedgecountconfigfuture"+device.GenId(), uid, e,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetEdgeCountConfig creates a subscriber for getting the actual edge count configurations.
func GetEdgeCountConfig(id string, uid uint32, pin *Pin, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetEdgeCountConfig"),
		Fid:        function_get_edge_count_config,
		Uid:        uid,
		Result:     &EdgeCountConfig{},
		Data:       pin,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetEdgeCountConfigFuture is a future pattern version for a synchronized all of the subscriber.
// If an error occur, the result is nil.
func GetEdgeCountConfigFuture(brick *bricker.Bricker, connectorname string, uid uint32, pin *Pin) *EdgeCountConfig {
	future := make(chan *EdgeCountConfig)
	defer close(future)
	sub := GetEdgeCountConfig("getedgecountconfigfuture"+device.GenId(), uid, pin,
		func(r device.Resulter, err error) {
			var v *EdgeCountConfig = nil
			if err == nil {
				if value, ok := r.(*EdgeCountConfig); ok {
					v = value
				}
			}
			future <- v
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return nil
	}
	return <-future
}

// EdgeCount is the type for GetEdgeCount.
type EdgeCount struct {
	Pin          uint8 // selected pin
	ResetCounter bool  // reset the counter directly after call
}

// EdgeCountRaw is a de/encoding type for EdgeCount.
type EdgeCountRaw struct {
	Pin          uint8
	ResetCounter uint8
}

// NewEdgeCountRaw creates a EdgeCountRaw from a EdgeCount.
func NewEdgeCountRaw(ec *EdgeCount) *EdgeCountRaw {
	if ec == nil {
		return nil
	}
	ecr := new(EdgeCountRaw)
	ecr.Pin = ec.Pin
	ecr.ResetCounter = misc.BoolToUint8(ec.ResetCounter)
	return ecr
}

// The value of the EdgeCount
type EdgeCounts struct {
	Value uint32
}

// Converts a packet to a EdgeCounts type.
func (e *EdgeCounts) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(e, p); err != nil {
		return err
	}
	return p.Payload.Decode(e)
}

// String fullfill the stringer interface.
func (e *EdgeCounts) String() string {
	txt := "EdgeCounts "
	if e == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("Value: %d", e.Value)
	}
	return txt
}

// EdgeCountConfig type for configurate the edge count.
type EdgeCountConfig struct {
	Type     uint8
	Debounce uint8 // in ms
}

// FromPacket creates a edge count configurations from a packet.
func (ecc *EdgeCountConfig) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(ecc, p); err != nil {
		return err
	}
	return p.Payload.Decode(ecc)
}

// String fullfill the stringer interface.
func (ecc *EdgeCountConfig) String() string {
	txt := "Edge Count Configuration "
	if ecc == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("[Edge Type: %d (%s), Debounce: %d ms]",
			ecc.Type, EdgeTypeName(ecc.Type), ecc.Debounce)
	}
	return txt
}

// EdgeTypeName converts the numeric edge type to a string reprensentation.
func EdgeTypeName(t uint8) string {
	switch t {
	case EdgeCountType_Rising:
		return "Rising"
	case EdgeCountType_Falling:
		return "Falling"
	case EdgeCountType_Both:
		return "Both"
	default:
		return "Unknown"
	}
}

// EdgeCountConfigs type for set a edge count configuration.
type EdgeCountConfigs struct {
	Pin uint8
	EdgeCountConfig
}
