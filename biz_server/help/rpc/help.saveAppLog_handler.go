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
	"github.com/airwide-code/airwide.datacenter/baselib/logger"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util"
	"github.com/airwide-code/airwide.datacenter/mtproto"
	"golang.org/x/net/context"
)

// help.saveAppLog#6f02f748 events:Vector<InputAppEvent> = Bool;
func (s *HelpServiceImpl) HelpSaveAppLog(ctx context.Context, request *mtproto.TLHelpSaveAppLog) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("HelpSaveAppLog - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl HelpSaveAppLog logic
	glog.Infof("HelpGetConfig - reply: {true}")
	return mtproto.ToBool(true), nil
}
