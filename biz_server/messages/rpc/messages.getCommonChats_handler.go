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

// messages.getCommonChats#d0a48c4 user_id:InputUser max_id:int limit:int = messages.Chats;
func (s *MessagesServiceImpl) MessagesGetCommonChats(ctx context.Context, request *mtproto.TLMessagesGetCommonChats) (*mtproto.Messages_Chats, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesGetCommonChats - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesGetCommonChats logic

	return nil, fmt.Errorf("Not impl MessagesGetCommonChats")
}
