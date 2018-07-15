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
    "github.com/airwide-code/airwide.datacenter/baselib/grpc_util"
    "github.com/airwide-code/airwide.datacenter/baselib/logger"
)

// help.getScheme#dbb69a9e version:int = Scheme;
func (s *HelpServiceImpl) HelpGetScheme(ctx context.Context, request *mtproto.TLHelpGetScheme) (*mtproto.Scheme, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("help.getScheme#dbb69a9e - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	scheme := mtproto.NewTLScheme()
	scheme.SetSchemeRaw("")
	scheme.SetVersion(1)

    glog.Infof("help.getScheme#dbb69a9e - reply: %s", logger.JsonDebugData(scheme))
    return scheme.To_Scheme(), nil
}
