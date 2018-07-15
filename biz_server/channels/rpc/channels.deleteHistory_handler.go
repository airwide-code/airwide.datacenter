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
    "github.com/airwide-code/airwide.datacenter/mtproto"
    "golang.org/x/net/context"
    "fmt"
    "github.com/airwide-code/airwide.datacenter/baselib/grpc_util"
    "github.com/airwide-code/airwide.datacenter/baselib/logger"
)

// channels.deleteHistory#af369d42 channel:InputChannel max_id:int = Bool;
func (s *ChannelsServiceImpl) ChannelsDeleteHistory(ctx context.Context, request *mtproto.TLChannelsDeleteHistory) (*mtproto.Bool, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("ChannelsDeleteHistory - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    // TODO(@benqi): Impl ChannelsDeleteHistory logic

    return nil, fmt.Errorf("Not impl ChannelsDeleteHistory")
}
