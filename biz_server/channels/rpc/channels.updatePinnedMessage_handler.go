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

// channels.updatePinnedMessage#a72ded52 flags:# silent:flags.0?true channel:InputChannel id:int = Updates;
func (s *ChannelsServiceImpl) ChannelsUpdatePinnedMessage(ctx context.Context, request *mtproto.TLChannelsUpdatePinnedMessage) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ChannelsUpdatePinnedMessage - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ChannelsUpdatePinnedMessage logic

	return nil, fmt.Errorf("Not impl ChannelsUpdatePinnedMessage")
}
