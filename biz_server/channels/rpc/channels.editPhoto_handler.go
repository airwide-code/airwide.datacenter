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

// channels.editPhoto#f12e57c9 channel:InputChannel photo:InputChatPhoto = Updates;
func (s *ChannelsServiceImpl) ChannelsEditPhoto(ctx context.Context, request *mtproto.TLChannelsEditPhoto) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ChannelsEditPhoto - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ChannelsEditPhoto logic

	return nil, fmt.Errorf("Not impl ChannelsEditPhoto")
}
