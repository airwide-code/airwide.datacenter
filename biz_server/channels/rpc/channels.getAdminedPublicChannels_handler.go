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

// channels.getAdminedPublicChannels#8d8d82d7 = messages.Chats;
func (s *ChannelsServiceImpl) ChannelsGetAdminedPublicChannels(ctx context.Context, request *mtproto.TLChannelsGetAdminedPublicChannels) (*mtproto.Messages_Chats, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ChannelsGetAdminedPublicChannels - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ChannelsGetAdminedPublicChannels logic

	return nil, fmt.Errorf("Not impl ChannelsGetAdminedPublicChannels")
}
