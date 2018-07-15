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

// auth.exportAuthorization#e5bfffcd dc_id:int = auth.ExportedAuthorization;
func (s *AuthServiceImpl) AuthExportAuthorization(ctx context.Context, request *mtproto.TLAuthExportAuthorization) (*mtproto.Auth_ExportedAuthorization, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AuthExportAuthorization - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl AuthExportAuthorization logic

	return nil, fmt.Errorf("Not impl AuthExportAuthorization")
}
