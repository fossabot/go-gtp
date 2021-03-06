// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package v2

import (
	"net"

	"github.com/wmnsk/go-gtp/v2/ies"
)

// QoSProfile is a QoS-related information that belongs to a Bearer.
type QoSProfile struct {
	PCI, PVI bool
	PL, QCI  uint8
	// Max bit rate for Uplink and Donwlink
	MBRUL, MBRDL uint64
	// Guaranteed bit rate for Uplink and Donwlink
	GBRUL, GBRDL uint64
}

// Bearer is a GTPv2 bearer.
type Bearer struct {
	raddr           net.Addr
	teidIn, teidOut uint32

	EBI               uint8
	SubscriberIP, APN string
	ChargingID        uint32
	*QoSProfile
}

// NewBearer creates a new Bearer.
func NewBearer(ebi uint8, apn string, qos *QoSProfile) *Bearer {
	return &Bearer{
		EBI: ebi, APN: apn, QoSProfile: qos,
	}
}

// Modify is just an alias of (*Conn) ModifyBearer.
func (b *Bearer) Modify(c *Conn, ie ...*ies.IE) error {
	return c.ModifyBearer(b.teidOut, ie...)
}

// RemoteAddress returns the remote address associated with Bearer.
func (b *Bearer) RemoteAddress() net.Addr {
	return b.raddr
}

// SetRemoteAddress sets the remote address associated with Bearer.
func (b *Bearer) SetRemoteAddress(raddr net.Addr) {
	b.raddr = raddr
}

// IncomingTEID returns the incoming TEID associated with Bearer.
func (b *Bearer) IncomingTEID() uint32 {
	return b.teidIn
}

// SetIncomingTEID sets the incoming TEID associated with Bearer.
func (b *Bearer) SetIncomingTEID(teid uint32) {
	b.teidIn = teid
}

// OutgoingTEID returns the outgoing TEID associated with Bearer.
func (b *Bearer) OutgoingTEID() uint32 {
	return b.teidOut
}

// SetOutgoingTEID sets the outgoing TEID associated with Bearer.
func (b *Bearer) SetOutgoingTEID(teid uint32) {
	b.teidOut = teid
}
