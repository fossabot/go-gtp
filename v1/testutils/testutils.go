// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

// Package testutils is an internal package to be used for unit tests. Don't use this.
package testutils

import (
	"testing"

	"github.com/pascaldekloe/goe/verify"
	"github.com/wmnsk/go-gtp/v1/messages"
)

// Serializeable is just for testing v2.Messages. Don't use this.
type Serializeable interface {
	Serialize() ([]byte, error)
	Len() int
}

// TestCase is just for testing v2.Messages. Don't use this.
type TestCase struct {
	Description string
	Structured  Serializeable
	Serialized  []byte
}

// DecodeFunc is just for testing v2.Messages. Don't use this.
type DecodeFunc func([]byte) (Serializeable, error)

// TestBearerInfo is just for testing v2.Messages. Don't use this.
var TestBearerInfo = struct {
	TEID uint32
	Seq  uint16
}{0x11223344, 0x00000001}

// Run is just for testing v2.Messages. Don't use this.
func Run(t *testing.T, cases []TestCase, decode DecodeFunc) {
	t.Helper()

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			t.Run("Decode", func(t *testing.T) {
				v, err := decode(c.Serialized)
				if err != nil {
					t.Fatal(err)
				}

				if got, want := v, c.Structured; !verify.Values(t, "", got, want) {
					t.Fail()
				}
			})

			t.Run("Serialize", func(t *testing.T) {
				b, err := c.Structured.Serialize()
				if err != nil {
					t.Fatal(err)
				}

				if got, want := b, c.Serialized; !verify.Values(t, "", got, want) {
					t.Fail()
				}
			})

			t.Run("Len", func(t *testing.T) {
				if got, want := c.Structured.Len(), len(c.Serialized); got != want {
					t.Fatalf("got %v want %v", got, want)
				}
			})

			t.Run("Interface", func(t *testing.T) {
				// Ignore *Header and Generic in this tests.
				if _, ok := c.Structured.(*messages.Header); ok {
					return
				}

				if _, ok := c.Structured.(*messages.Generic); ok {
					return
				}

				decoded, err := messages.Decode(c.Serialized)
				if err != nil {
					t.Fatal(err)
				}

				if got, want := decoded.Version(), c.Structured.(messages.Message).Version(); got != want {
					t.Fatalf("got %v want %v", got, want)
				}
				if got, want := decoded.MessageType(), c.Structured.(messages.Message).MessageType(); got != want {
					t.Fatalf("got %v want %v", got, want)
				}
				if got, want := decoded.MessageTypeName(), c.Structured.(messages.Message).MessageTypeName(); got != want {
					t.Fatalf("got %v want %v", got, want)
				}
				if got, want := decoded.TEID(), c.Structured.(messages.Message).TEID(); got != want {
					t.Fatalf("got %v want %v", got, want)
				}
			})
		})
	}
}
