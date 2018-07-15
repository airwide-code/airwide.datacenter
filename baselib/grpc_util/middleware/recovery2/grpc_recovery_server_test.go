/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package grpc_recovery2

import (
	"testing"
	"google.golang.org/grpc"
	"context"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util/middleware/examples/zproto"
	"net"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"fmt"
)

type ChatTestServiceImpl struct {
}

func (s *ChatTestServiceImpl) Connect(request *zproto.ChatSession, stream zproto.ChatTest_ConnectServer) (err error) {
	return
}

func (s *ChatTestServiceImpl) SendChat(ctx context.Context, request *zproto.ChatMessage) (reply *zproto.VoidRsp2, err error) {
	fmt.Printf("%v.SendChat(_) = _, %v\n", ctx, request)

	switch request.MessageData {
	case "panic":
		panic("very bad thing happened")
	case "nil":
		panic("nil thing happened")
	}
	return &zproto.VoidRsp2{}, nil
}

func unaryRecoveryHandler(ctx context.Context, p interface{}) (err error) {
	return status.Errorf(codes.Unknown, "panic triggered: %v", p)
}

func unaryRecoveryHandler2(ctx context.Context, p interface{}) (err error) {
	return status.Errorf(codes.Unknown, "panic triggered: %v", p)
}

func streamRecoveryHandler(stream grpc.ServerStream, p interface{}) (err error) {
	return
}

func TestRecoveryServer(t *testing.T) {
	lis, err := net.Listen("tcp", "0.0.0.0:22345")
	if err != nil {
		panic(err)
		// glog.Fatalf("failed to listen: %v", err)
	}

	server := NewRecoveryServer2(unaryRecoveryHandler, unaryRecoveryHandler2, streamRecoveryHandler)
	zproto.RegisterChatTestServer(server, &ChatTestServiceImpl{})
	server.Serve(lis)
}
