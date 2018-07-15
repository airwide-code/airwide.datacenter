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

// channels.getChannels#a7f6bbb id:Vector<InputChannel> = messages.Chats;
func (s *ChannelsServiceImpl) ChannelsGetChannels(ctx context.Context, request *mtproto.TLChannelsGetChannels) (*mtproto.Messages_Chats, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ChannelsGetChannels - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ChannelsGetChannels logic

	return nil, fmt.Errorf("Not impl ChannelsGetChannels")
}
