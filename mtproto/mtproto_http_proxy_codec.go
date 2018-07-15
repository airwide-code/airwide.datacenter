/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package mtproto

import (
	"net"
	"fmt"
	"net/http"
	"github.com/golang/glog"
	"encoding/binary"
	"io/ioutil"
	// "time"
	"github.com/airwide-code/airwide.datacenter/baselib/net2"
	// "strings"
	"bytes"
)

type MTProtoHttpProxyCodec struct {
	// conn  *MTProtoHttpProxyConn
	conn net.Conn
}

func NewMTProtoHttpProxyCodec(conn net.Conn) *MTProtoHttpProxyCodec {
	return &MTProtoHttpProxyCodec{
		conn:  conn,
	}
}

func (c* MTProtoHttpProxyCodec) Receive() (interface{}, error) {
	req, err := http.ReadRequest(c.conn.(*net2.BufferedConn).BufioReader())
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	body, _ := ioutil.ReadAll(req.Body) //把  body 内容读入字符串 s

	authKeyId := int64(binary.LittleEndian.Uint64(body))
	msg := NewMTPRawMessage(authKeyId, 0)
	err = msg.Decode(body)
	if err != nil {
		glog.Error(err)
		// conn.Close()
		return nil, err
	}

	return msg, nil
}

func (c* MTProtoHttpProxyCodec) Send(msg interface{}) error {
	// SendToHttpReply(msg, w)
	message, ok := msg.(*MTPRawMessage)
	if !ok {
		err := fmt.Errorf("msg type error, only MTPRawMessage, msg: {%v}", msg)
		glog.Error(err)
		// conn.Close()
		return err
	}

	b := message.Encode()

	rsp := http.Response{
		StatusCode:    200,
		ProtoMajor:    1,
		ProtoMinor:    1,
		Request:       &http.Request{Method: "POST"},
		Header:        http.Header{
			"Access-Control-Allow-Headers": {"origin, content-type"},
			"Access-Control-Allow-Methods": {"POST, OPTIONS"},
			"Access-Control-Allow-Origin":  {"*"},
			"Access-Control-Max-Age":       {"1728000"},
			"Cache-control":                {"no-store"},
			"Connection":                   {"keep-alive"},
			"Content-type":                 {"application/octet-stream"},
			"Pragma":                       {"no-cache"},
			"Strict-Transport-Security":    {"max-age=15768000"},
		},
		ContentLength: int64(len(b)),
		Body:          ioutil.NopCloser(bytes.NewReader(b)),
		Close:         false,
	}

	err := rsp.Write(c.conn)
	if err != nil {
		glog.Error(err)
	}

	return err
}

func (c* MTProtoHttpProxyCodec) Close() error {
	return c.conn.Close()
}

