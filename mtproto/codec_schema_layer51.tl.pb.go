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

func NewTLAuthSendCodeLayer51() * TLAuthSendCodeLayer51 {
	return &TLAuthSendCodeLayer51{}
}

func (m* TLAuthSendCodeLayer51) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_sendCodeLayer51))

	// flags
	var flags uint32 = 0
	if m.AllowFlashcall == true { flags |= 1 << 0 }
	if m.CurrentNumber != nil { flags |= 1 << 0 }
	x.UInt(flags)


	x.String(m.PhoneNumber)
	if m.CurrentNumber != nil {
		x.Bytes(m.CurrentNumber.Encode())
	}
	x.Int(m.ApiId)
	x.String(m.ApiHash)
	x.String(m.LangCode)

	return x.buf
}

func (m* TLAuthSendCodeLayer51) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	_ = flags
	if (flags & (1 << 0)) != 0 { m.AllowFlashcall = true }
	m.PhoneNumber = dbuf.String()
	if (flags & (1 << 0)) != 0 {
		m4 := &Bool{}
		m4.Decode(dbuf)
		m.CurrentNumber = m4
	}
	m.ApiId = dbuf.Int()
	m.ApiHash = dbuf.String()
	m.LangCode = dbuf.String()

	return dbuf.err
}
