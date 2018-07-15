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

// channels.togglePreHistoryHidden#eabbb94c channel:InputChannel enabled:Bool = Updates;
func (s *ChannelsServiceImpl) ChannelsTogglePreHistoryHidden(ctx context.Context, request *mtproto.TLChannelsTogglePreHistoryHidden) (*mtproto.Updates, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("ChannelsTogglePreHistoryHidden - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    // TODO(@benqi): Impl ChannelsTogglePreHistoryHidden logic

    return nil, fmt.Errorf("Not impl ChannelsTogglePreHistoryHidden")
}
