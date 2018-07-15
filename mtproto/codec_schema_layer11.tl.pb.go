/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'codegen_encode_decode.py'
 *
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

// ConstructorList
// RequestList

package mtproto

import (
	// "encoding/binary"
	// "fmt"
	// "github.com/golang/protobuf/proto"
)

///////////////////////////////////////////////////////////////////////////////
// InputFileLocation <--
//  + TL_InputDocumentFileLocation
//

// inputDocumentFileLocation#4e45abe9 id:long access_hash:long = InputFileLocation;
func (m *TLInputDocumentFileLocationLayer11) To_InputFileLocation() *InputFileLocation {
	return &InputFileLocation{
		Constructor: TLConstructor_CRC32_inputDocumentFileLocationLayer11,
		Data2: m.Data2,
	}
}


func (m *TLInputDocumentFileLocationLayer11) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputDocumentFileLocationLayer11) GetId() int64 { return m.Data2.Id }

func (m *TLInputDocumentFileLocationLayer11) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputDocumentFileLocationLayer11) GetAccessHash() int64 { return m.Data2.AccessHash }


func NewTLInputDocumentFileLocationLayer11() * TLInputDocumentFileLocationLayer11 {
	return &TLInputDocumentFileLocationLayer11{ Data2: &InputFileLocation_Data{} }
}

func (m* TLInputDocumentFileLocationLayer11) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputDocumentFileLocationLayer11))

	x.Long(m.GetId())
	x.Long(m.GetAccessHash())

	return x.buf
}

func (m* TLInputDocumentFileLocationLayer11) Decode(dbuf *DecodeBuf) error {
	m.SetId(dbuf.Long())
	m.SetAccessHash(dbuf.Long())

	return dbuf.err
}
