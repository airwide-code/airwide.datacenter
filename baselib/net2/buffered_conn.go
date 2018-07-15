/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package net2

import (
	"bufio"
	"net"
)

type BufferedConn struct {
	r        *bufio.Reader
	net.Conn // So that most methods are embedded
}

func NewBufferedConn(c net.Conn) *BufferedConn {
	return &BufferedConn{bufio.NewReader(c), c}
}

func NewBufferedConnSize(c net.Conn, n int) *BufferedConn {
	return &BufferedConn{bufio.NewReaderSize(c, n), c}
}

func (b *BufferedConn) Peek(n int) ([]byte, error) {
	return b.r.Peek(n)
}

func (b *BufferedConn) Read(p []byte) (int, error) {
	return b.r.Read(p)
}

func (b *BufferedConn) Discard(n int) (int, error) {
	return b.r.Discard(n)
}

func (b *BufferedConn) BufioReader() *bufio.Reader {
	return b.r
}
