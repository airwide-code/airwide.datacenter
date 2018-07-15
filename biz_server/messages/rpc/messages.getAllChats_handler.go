/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package rpc

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/airwide-code/airwide.datacenter/baselib/logger"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util"
	"github.com/airwide-code/airwide.datacenter/mtproto"
	"golang.org/x/net/context"
)

// messages.getAllChats#eba80ff0 except_ids:Vector<int> = messages.Chats;
func (s *MessagesServiceImpl) MessagesGetAllChats(ctx context.Context, request *mtproto.TLMessagesGetAllChats) (*mtproto.Messages_Chats, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesGetAllChats - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesGetAllChats logic

	return nil, fmt.Errorf("Not impl MessagesGetAllChats")
}
