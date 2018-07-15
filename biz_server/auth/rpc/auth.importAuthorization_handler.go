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

// auth.importAuthorization#e3ef9613 id:int bytes:bytes = auth.Authorization;
func (s *AuthServiceImpl) AuthImportAuthorization(ctx context.Context, request *mtproto.TLAuthImportAuthorization) (*mtproto.Auth_Authorization, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AuthImportAuthorization - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl AuthImportAuthorization logic

	return nil, fmt.Errorf("Not impl AuthImportAuthorization")
}
