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

// channels.readMessageContents#eab5dc38 channel:InputChannel id:Vector<int> = Bool;
func (s *ChannelsServiceImpl) ChannelsReadMessageContents(ctx context.Context, request *mtproto.TLChannelsReadMessageContents) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ChannelsReadMessageContents - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ChannelsReadMessageContents logic

	return nil, fmt.Errorf("Not impl ChannelsReadMessageContents")
}
