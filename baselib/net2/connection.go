/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package net2

type Connection interface {
	GetConnID() uint64
	IsClosed() bool
	Close() error
	Codec() Codec
	Receive() (interface{}, error)
	Send(msg interface{}) error
}

type closeCallback interface {
	// func(Connection)
	OnConnectionClosed(Connection)
}
