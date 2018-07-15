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
	"github.com/airwide-code/airwide.datacenter/biz/core"
)

// Layer74
// account.registerDevice#1389cc token_type:int token:string app_sandbox:Bool other_uids:Vector<int> = Bool;
func (s *AccountServiceImpl) AccountRegisterDeviceLayer74(ctx context.Context, request *mtproto.TLAccountRegisterDeviceLayer74) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.registerDevice#637ea878 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// Check token format by token_type
	// TODO(@benqi): check token format by token_type
	if request.Token == "" {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error(err)
		return nil, err
	}

	// TODO(@benqi): check toke_type invalid
	if request.TokenType < core.TOKEN_TYPE_APNS || request.TokenType > core.TOKEN_TYPE_MAXSIZE {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error(err)
		return nil, err
	}

	registered := account.RegisterDevice(md.AuthId, md.UserId, int8(request.TokenType), request.Token)

	glog.Infof("account.registerDevice#637ea878 - reply: {true}")
	return mtproto.ToBool(registered), nil
}
