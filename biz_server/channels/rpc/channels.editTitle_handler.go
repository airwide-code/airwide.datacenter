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

// channels.editTitle#566decd0 channel:InputChannel title:string = Updates;
func (s *ChannelsServiceImpl) ChannelsEditTitle(ctx context.Context, request *mtproto.TLChannelsEditTitle) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ChannelsEditTitle - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ChannelsEditTitle logic

	return nil, fmt.Errorf("Not impl ChannelsEditTitle")
}
