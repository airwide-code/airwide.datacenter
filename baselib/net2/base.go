/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package net2

import (
	"fmt"
	"io"
)

type Protocol interface {
	NewCodec(rw io.ReadWriter) (Codec, error)
}

type ProtocolFunc func(rw io.ReadWriter) (Codec, error)

func (pf ProtocolFunc) NewCodec(rw io.ReadWriter) (Codec, error) {
	return pf(rw)
}

type Codec interface {
	Receive() (interface{}, error)
	Send(interface{}) error
	Close() error
}

type MessageBase interface {
	Encode() []byte
	Decode(b []byte) error
}

type ConnectionFactory interface {
	NewConnection(serverName string) TcpConnection
}

type ClearSendChan interface {
	ClearSendChan(<-chan interface{})
}

var (
	protocolRegisters = make(map[string]Protocol)
)

func RegisterProtocol(name string, protocol Protocol) {
	// glog.Info("RegisterProtocol: ", name)
	protocolRegisters[name] = protocol
}

func NewCodecByName(name string, rw io.ReadWriter) (Codec, error) {
	protocol, ok := protocolRegisters[name]
	if !ok {
		return nil, fmt.Errorf("not found protocol name: %s", name)
	}
	return protocol.NewCodec(rw)
}
