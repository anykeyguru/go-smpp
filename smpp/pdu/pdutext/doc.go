// Copyright 2015 go-smpp authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package pdutext provides text codecs for the short_message PDU field.
//
// See section 2.2.2 of the SMPP 3.4 over GSM/UMTS Implementation Guide
// for background: http://opensmpp.org/specs/smppv34_gsmumts_ig_v10.pdf
//
// The following codecs are provided, each implementing the Codec
// interface and mapping to a value of the data_coding PDU field:
//
//	Raw         (0x00) octets passed through unchanged.
//	GSM7        (0x00) GSM 7-bit default alphabet, unpacked (one septet per octet).
//	GSM7Packed  (0x00) GSM 7-bit default alphabet, packed into octets per 3GPP TS 23.038.
//	Latin1      (0x03) Windows-1252 (CP1252), a superset of ISO-8859-1.
//	ISO88595    (0x06) ISO-8859-5 (Cyrillic).
//	UCS2        (0x08) UTF-16 big-endian.
//
// Latin1 uses Windows-1252 rather than ISO-8859-1; the two differ in the
// 0x80-0x9F range. See:
// http://www.i18nqa.com/debug/table-iso8859-1-vs-windows-1252.html
package pdutext
