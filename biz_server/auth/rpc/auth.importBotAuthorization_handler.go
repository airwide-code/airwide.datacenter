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

// auth.importBotAuthorization#67a3ff2c flags:int api_id:int api_hash:string bot_auth_token:string = auth.Authorization;
func (s *AuthServiceImpl) AuthImportBotAuthorization(ctx context.Context, request *mtproto.TLAuthImportBotAuthorization) (*mtproto.Auth_Authorization, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AuthImportBotAuthorization - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl AuthImportBotAuthorization logic

	return nil, fmt.Errorf("Not impl AuthImportBotAuthorization")
}
