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

// channels.editBanned#bfd915cd channel:InputChannel user_id:InputUser banned_rights:ChannelBannedRights = Updates;
func (s *ChannelsServiceImpl) ChannelsEditBanned(ctx context.Context, request *mtproto.TLChannelsEditBanned) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ChannelsEditBanned - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ChannelsEditBanned logic

	return nil, fmt.Errorf("Not impl ChannelsEditBanned")
}
