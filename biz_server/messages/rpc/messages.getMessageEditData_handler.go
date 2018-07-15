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

// messages.getMessageEditData#fda68d36 peer:InputPeer id:int = messages.MessageEditData;
func (s *MessagesServiceImpl) MessagesGetMessageEditData(ctx context.Context, request *mtproto.TLMessagesGetMessageEditData) (*mtproto.Messages_MessageEditData, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesGetMessageEditData - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesGetMessageEditData logic

	return nil, fmt.Errorf("Not impl MessagesGetMessageEditData")
}
