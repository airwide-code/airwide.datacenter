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
)

// account.resetAuthorization#df77f3bc hash:long = Bool;
func (s *AccountServiceImpl) AccountResetAuthorization(ctx context.Context, request *mtproto.TLAccountResetAuthorization) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.resetAuthorization#df77f3bc - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	authKeyId := account.GetAuthKeyIdByHash(md.UserId, request.GetHash())
	if authKeyId == 0 {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error("account.resetAuthorization#df77f3bc - not found hash ", err)
		return nil, err
	}

	// TODO(@benqi): found session, kick off.
	account.DeleteAuthorization(authKeyId)

	glog.Infof("account.checkUsername#2714d86c - reply: {true}")
	return mtproto.ToBool(true), nil
}
