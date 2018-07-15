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

// help.setBotUpdatesStatus#ec22cfcd pending_updates_count:int message:string = Bool;
func (s *HelpServiceImpl) HelpSetBotUpdatesStatus(ctx context.Context, request *mtproto.TLHelpSetBotUpdatesStatus) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("HelpSetBotUpdatesStatus - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl HelpSetBotUpdatesStatus logic

	return nil, fmt.Errorf("Not impl HelpSetBotUpdatesStatus")
}
