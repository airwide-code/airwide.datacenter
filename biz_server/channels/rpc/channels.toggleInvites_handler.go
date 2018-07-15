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

// channels.toggleInvites#49609307 channel:InputChannel enabled:Bool = Updates;
func (s *ChannelsServiceImpl) ChannelsToggleInvites(ctx context.Context, request *mtproto.TLChannelsToggleInvites) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ChannelsToggleInvites - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ChannelsToggleInvites logic

	return nil, fmt.Errorf("Not impl ChannelsToggleInvites")
}
