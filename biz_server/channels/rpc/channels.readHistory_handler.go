/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package rpc

import (
	"github.com/golang/glog"
	"github.com/airwide-code/airwide.datacenter/baselib/logger"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util"
	"github.com/airwide-code/airwide.datacenter/mtproto"
	"golang.org/x/net/context"
)

// channels.readHistory#cc104937 channel:InputChannel max_id:int = Bool;
func (s *ChannelsServiceImpl) ChannelsReadHistory(ctx context.Context, request *mtproto.TLChannelsReadHistory) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("channels.readHistory#cc104937 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ChannelsReadHistory logic

	glog.Infof("channels.readHistory#cc104937 - reply: {true}")
	return mtproto.ToBool(true), nil
}
