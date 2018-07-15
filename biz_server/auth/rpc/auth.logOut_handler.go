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

// auth.logOut#5717da40 = Bool;
func (s *AuthServiceImpl) AuthLogOut(ctx context.Context, request *mtproto.TLAuthLogOut) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("auth.logOut#5717da40 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl AuthLogOut logic

	glog.Info("auth.logOut#5717da40 - reply: {true}")
	return mtproto.ToBool(true), nil
}
