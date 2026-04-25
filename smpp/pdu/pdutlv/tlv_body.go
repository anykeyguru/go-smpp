// Copyright 2015 go-smpp authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package pdutlv

import (
	"io"
)

// Body is an interface for manipulating binary PDU Tag-Length-Value field data.
type Body interface {
	Len() int
	Raw() interface{}
	String() string
	Bytes() []byte
	SerializeTo(w io.Writer) error
}

// NewTLV returns a Body that wraps the given tag and raw value.
func NewTLV(tag Tag, value []byte) Body {
	return &Field{Tag: tag, Data: value}
}
