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

// channels.reportSpam#fe087810 channel:InputChannel user_id:InputUser id:Vector<int> = Bool;
func (s *ChannelsServiceImpl) ChannelsReportSpam(ctx context.Context, request *mtproto.TLChannelsReportSpam) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ChannelsReportSpam - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ChannelsReportSpam logic

	return nil, fmt.Errorf("Not impl ChannelsReportSpam")
}
