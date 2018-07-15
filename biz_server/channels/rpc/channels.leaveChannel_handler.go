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

// channels.leaveChannel#f836aa95 channel:InputChannel = Updates;
func (s *ChannelsServiceImpl) ChannelsLeaveChannel(ctx context.Context, request *mtproto.TLChannelsLeaveChannel) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ChannelsLeaveChannel - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ChannelsLeaveChannel logic

	return nil, fmt.Errorf("Not impl ChannelsLeaveChannel")
}
