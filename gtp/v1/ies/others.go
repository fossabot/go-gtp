// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewCause creates a new Cause IE.
func NewCause(cause uint8) *IE {
	return newUint8ValIE(Cause, cause)
}

// Cause returns the Cause value if type matches.
func (i *IE) Cause() uint8 {
	if i.Type != Cause {
		return 0
	}
	return i.Payload[0]
}

// NewReorderingRequired creates a new ReorderingRequired IE.
func NewReorderingRequired(required bool) *IE {
	if required {
		return newUint8ValIE(ReorderingRequired, 0xff)
	}
	return newUint8ValIE(ReorderingRequired, 0xfe)
}

// ReorderingRequired returns ReorderingRequired or not if type matches.
func (i *IE) ReorderingRequired() bool {
	if i.Type != ReorderingRequired {
		return false
	}
	return i.Payload[0]&0x01 == 1
}