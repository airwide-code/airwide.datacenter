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

func NewTLMessagesReadHistoryLayer2() * TLMessagesReadHistoryLayer2 {
	return &TLMessagesReadHistoryLayer2{}
}

func (m* TLMessagesReadHistoryLayer2) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_readHistory))

	x.Bytes(m.Peer.Encode())
	x.Int(m.MaxId)
	x.Int(m.Offset)

	return x.buf
}

func (m* TLMessagesReadHistoryLayer2) Decode(dbuf *DecodeBuf) error {
	m1 := &InputPeer{}
	m1.Decode(dbuf)
	m.Peer = m1
	m.MaxId = dbuf.Int()
	m.Offset = dbuf.Int()

	return dbuf.err
}
