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

// channels.getMessages#93d7b347 channel:InputChannel id:Vector<int> = messages.Messages;
func (s *ChannelsServiceImpl) ChannelsGetMessages(ctx context.Context, request *mtproto.TLChannelsGetMessages) (*mtproto.Messages_Messages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("channels.getMessages#93d7b347 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ChannelsGetMessages logic

	return nil, fmt.Errorf("Not impl ChannelsGetMessages")
}
