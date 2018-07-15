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

// channels.toggleSignatures#1f69b606 channel:InputChannel enabled:Bool = Updates;
func (s *ChannelsServiceImpl) ChannelsToggleSignatures(ctx context.Context, request *mtproto.TLChannelsToggleSignatures) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ChannelsToggleSignatures - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ChannelsToggleSignatures logic

	return nil, fmt.Errorf("Not impl ChannelsToggleSignatures")
}
