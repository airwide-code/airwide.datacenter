/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package mtproto

import (
	"fmt"
	// "encoding/binary"
)

type MessageBase interface {
	Encode() []byte
	Decode(b []byte) error
}

//
//type CodecAble interface {
//	Encode() ([]byte, error)
//	Decode(dbuf *DecodeBuf) error
//}

func NewMTPRawMessage(authKeyId int64, quickAckId int32) *MTPRawMessage {
	return &MTPRawMessage{
		AuthKeyId:  authKeyId,
		QuickAckId: quickAckId,
	}
}

////////////////////////////////////////////////////////////////////////////
// 代理使用
type MTPRawMessage struct {
	AuthKeyId  int64 // 由原始数据解压获得
	QuickAckId int32 // EncryptedMessage，则可能存在

	// 原始数据
	Payload []byte
}

func (m *MTPRawMessage) Encode() []byte {
	return m.Payload
}

func (m *MTPRawMessage) Decode(b []byte) error {
	m.Payload = b
	return nil
}

////////////////////////////////////////////////////////////////////////////
func NewUnencryptedRawMessage() *UnencryptedRawMessage {
	return &UnencryptedRawMessage{
		AuthKeyId: 0,
	}
}

type UnencryptedRawMessage struct {
	// TODO(@benqi): reportAck and quickAck
	// NeedAck bool
	AuthKeyId   int64
	MessageId   int64
	MessageData []byte
}

func (m *UnencryptedRawMessage) Encode() []byte {
	// 一次性分配
	x := NewEncodeBuf(20 + len(m.MessageData))
	x.Long(0)
	m.MessageId = GenerateMessageId()
	x.Long(m.MessageId)
	x.Int(int32(len(m.MessageData)))
	x.Bytes(m.MessageData)
	return x.buf
}

func (m *UnencryptedRawMessage) Decode(b []byte) error {
	dbuf := NewDecodeBuf(b)
	m.MessageId = dbuf.Long()
	messageLen := dbuf.Int()
	if int(messageLen) != dbuf.size-12 {
		return fmt.Errorf("invalid UnencryptedRawMessage len: %d (need %d)", messageLen, dbuf.size-12)
	}
	m.MessageData = dbuf.Bytes(int(messageLen))
	return dbuf.err
}

type EncryptedRawMessage struct {
	// TODO(@benqi): reportAck and quickAck
	// NeedAck bool
	AuthKeyId     int64
	MsgKey        []byte
	EncryptedData []byte
}

func NewEncryptedRawMessage(authKeyId int64) *EncryptedRawMessage {
	return &EncryptedRawMessage{
		AuthKeyId: authKeyId,
	}
}

func (m *EncryptedRawMessage) Encode() []byte {
	// 一次性分配
	x := NewEncodeBuf(24 + len(m.EncryptedData))
	x.Long(m.AuthKeyId)
	x.Bytes(m.MsgKey)
	x.Bytes(m.EncryptedData)
	return x.buf
}

func (m *EncryptedRawMessage) Decode(b []byte) error {
	dbuf := NewDecodeBuf(b)
	m.MsgKey = dbuf.Bytes(16)
	m.EncryptedData = dbuf.Bytes(len(b) - 16)
	return dbuf.err
}
