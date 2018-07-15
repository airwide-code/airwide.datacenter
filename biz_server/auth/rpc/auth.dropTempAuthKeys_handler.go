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

// auth.dropTempAuthKeys#8e48a188 except_auth_keys:Vector<long> = Bool;
func (s *AuthServiceImpl) AuthDropTempAuthKeys(ctx context.Context, request *mtproto.TLAuthDropTempAuthKeys) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AuthDropTempAuthKeys - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl AuthDropTempAuthKeys logic

	return nil, fmt.Errorf("Not impl AuthDropTempAuthKeys")
}
