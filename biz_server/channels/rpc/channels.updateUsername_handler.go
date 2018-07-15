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

// channels.updateUsername#3514b3de channel:InputChannel username:string = Bool;
func (s *ChannelsServiceImpl) ChannelsUpdateUsername(ctx context.Context, request *mtproto.TLChannelsUpdateUsername) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ChannelsUpdateUsername - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ChannelsUpdateUsername logic

	return nil, fmt.Errorf("Not impl ChannelsUpdateUsername")
}
