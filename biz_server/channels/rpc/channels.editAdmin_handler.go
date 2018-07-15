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

// channels.editAdmin#20b88214 channel:InputChannel user_id:InputUser admin_rights:ChannelAdminRights = Updates;
func (s *ChannelsServiceImpl) ChannelsEditAdmin(ctx context.Context, request *mtproto.TLChannelsEditAdmin) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ChannelsEditAdmin - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ChannelsEditAdmin logic

	return nil, fmt.Errorf("Not impl ChannelsEditAdmin")
}
