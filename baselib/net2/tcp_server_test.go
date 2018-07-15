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
	"github.com/golang/glog"
	"net"
)

type TestPingPongServer struct {
	server      *TcpServer
	serverName  string
	workLoadCnt int
}

func NewTestServer(listener net.Listener, serverName, protoName string, chanSize int, maxConn int) *TestPingPongServer {
	s := &TestPingPongServer{}
	s.server = NewTcpServer(
		TcpServerArgs{
			Listener:                listener,
			ServerName:              serverName,
			ProtoName:               protoName,
			SendChanSize:            chanSize,
			ConnectionCallback:      s,
			MaxConcurrentConnection: maxConn,
		})
	s.serverName = serverName
	s.workLoadCnt = 0
	return s
}

func (s *TestPingPongServer) Serve() {
	s.server.Serve()
}

func (s *TestPingPongServer) Stop() {
	s.server.Stop()
}

func (s *TestPingPongServer) isReady() bool {
	return s.server.running
}

func (s *TestPingPongServer) OnNewConnection(conn *TcpConnection) {
	glog.Infof("server OnNewConnection %v", conn.String())
}

func (s *TestPingPongServer) OnConnectionDataArrived(conn *TcpConnection, msg interface{}) (err error) {
	glog.Infof("%s server receive peer(%v) data: %v", s.serverName, conn.RemoteAddr(), msg)
	err = conn.Send(fmt.Sprintf("pong --> %s", msg))
	s.workLoadCnt++
	return err
}

func (s *TestPingPongServer) OnConnectionClosed(conn *TcpConnection) {
	glog.Infof("server OnConnectionClosed - %v", conn.RemoteAddr())
}
