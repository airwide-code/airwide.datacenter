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
	"github.com/golang/glog"
	"io"
)

type TestCodec struct {
	*bufio.Reader
	io.Writer
	io.Closer
	mt string
}

func (c *TestCodec) Send(msg interface{}) error {
	buf := []byte(msg.(string))
	if _, err := c.Writer.Write(buf); err != nil {
		return err
	}

	return nil
}

func (c *TestCodec) Receive() (interface{}, error) {
	line, err := c.Reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	return line, err
}

func (c *TestCodec) Close() error {
	return c.Closer.Close()
}

func (c *TestCodec) ClearSendChan(ic <-chan interface{}) {
	glog.Info(`TestCodec ClearSendChan, `, ic)
}

//////////////////////////////////////////////////////////////////////////////////////////
type TestProto struct {
}

func (b *TestProto) NewCodec(rw io.ReadWriter) (cc Codec, err error) {
	c := &TestCodec{
		Reader: bufio.NewReader(rw),
		Writer: rw.(io.Writer),
		Closer: rw.(io.Closer),
	}
	return c, nil
}
