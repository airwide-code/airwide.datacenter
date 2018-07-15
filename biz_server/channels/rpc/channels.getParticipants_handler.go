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

// channels.getParticipants#123e05e9 channel:InputChannel filter:ChannelParticipantsFilter offset:int limit:int hash:int = channels.ChannelParticipants;
func (s *ChannelsServiceImpl) ChannelsGetParticipants(ctx context.Context, request *mtproto.TLChannelsGetParticipants) (*mtproto.Channels_ChannelParticipants, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("channels.getParticipants#123e05e9 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ChannelsGetParticipants logic

	return nil, fmt.Errorf("Not impl ChannelsGetParticipants")
}
