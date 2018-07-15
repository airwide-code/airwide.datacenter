/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package mtproto

const (
	MTPROTO_VERSION = 2
)

type TLObject interface {
	Encode() []byte
	Decode(dbuf *DecodeBuf) error
	String() string
}
