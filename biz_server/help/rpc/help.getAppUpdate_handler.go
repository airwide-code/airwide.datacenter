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

// help.getAppUpdate#ae2de196 = help.AppUpdate;
func (s *HelpServiceImpl) HelpGetAppUpdate(ctx context.Context, request *mtproto.TLHelpGetAppUpdate) (*mtproto.Help_AppUpdate, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("HelpGetAppUpdate - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl HelpGetAppUpdate logic
	reply := &mtproto.TLHelpNoAppUpdate{Data2: &mtproto.Help_AppUpdate_Data{}}

	glog.Infof("HelpGetAppUpdate - reply: %s\n", logger.JsonDebugData(reply))
	return reply.To_Help_AppUpdate(), nil
}
