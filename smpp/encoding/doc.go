// Copyright 2015 go-smpp authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package encoding provides a golang.org/x/text/encoding compatible
// implementation of the GSM 7-bit default alphabet, in both unpacked
// and packed forms, along with helpers for validating that a string
// or byte buffer fits the alphabet.
//
// See 3GPP TS 23.038 for the alphabet and packing rules.
package encoding
