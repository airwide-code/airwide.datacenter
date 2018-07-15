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
	"github.com/airwide-code/airwide.datacenter/biz/core/account"
	user2 "github.com/airwide-code/airwide.datacenter/biz/core/user"
)

// auth.recoverPassword#4ea56e92 code:string = auth.Authorization;
func (s *AuthServiceImpl) AuthRecoverPassword(ctx context.Context, request *mtproto.TLAuthRecoverPassword) (*mtproto.Auth_Authorization, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("auth.recoverPassword#4ea56e92 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		err error = nil
	)

	if request.Code == "" {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_CODE_INVALID)
		glog.Error(err)
		return nil, err
	} else {
		err = account.CheckRecoverCode(md.UserId, request.Code)
		if err != nil {
			glog.Error(err)
			return nil, err
		}
	}

	user := user2.GetUserById(md.UserId, md.UserId)
	authAuthorization := &mtproto.TLAuthAuthorization{Data2: &mtproto.Auth_Authorization_Data{
		User: user.To_User(),
	}}

	glog.Infof("auth.recoverPassword#4ea56e92 - reply: %s\n", logger.JsonDebugData(authAuthorization))
	return authAuthorization.To_Auth_Authorization(), nil
}
