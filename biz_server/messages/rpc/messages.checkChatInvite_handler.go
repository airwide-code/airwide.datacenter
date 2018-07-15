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

// messages.checkChatInvite#3eadb1bb hash:string = ChatInvite;
func (s *MessagesServiceImpl) MessagesCheckChatInvite(ctx context.Context, request *mtproto.TLMessagesCheckChatInvite) (*mtproto.ChatInvite, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesCheckChatInvite - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesCheckChatInvite logic

	return nil, fmt.Errorf("Not impl MessagesCheckChatInvite")
}
