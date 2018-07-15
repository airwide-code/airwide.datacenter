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

// help.getAppUpdate#c812ac7e device_model:string system_version:string app_version:string lang_code:string = help.AppUpdate;
func (s *HelpServiceImpl) HelpGetAppUpdateLayer46(ctx context.Context, request *mtproto.TLHelpGetAppUpdateLayer46) (*mtproto.Help_AppUpdate, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("help.getAppUpdate#c812ac7e - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl HelpGetAppUpdate logic
	reply := &mtproto.TLHelpNoAppUpdate{Data2: &mtproto.Help_AppUpdate_Data{}}

	glog.Infof("help.getAppUpdate#c812ac7e - reply: %s\n", logger.JsonDebugData(reply))
	return reply.To_Help_AppUpdate(), nil
}
