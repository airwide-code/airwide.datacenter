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

// account.unregisterDevice#65c55b40 token_type:int token:string = Bool;
func (s *AccountServiceImpl) AccountUnregisterDevice(ctx context.Context, request *mtproto.TLAccountUnregisterDevice) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.unregisterDevice#65c55b40 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// Check token invalid
	// TODO(@benqi): check token format by token_type
	if request.Token == "" {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error(err)
		return nil, err
	}

	// Check token format by token_type
	if request.TokenType < core.TOKEN_TYPE_APNS || request.TokenType > core.TOKEN_TYPE_INTERNAL_PUSH {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error(err)
		return nil, err
	}

	unregistered := account.UnRegisterDevice(int8(request.TokenType), request.Token)

	glog.Infof("account.unregisterDevice#65c55b40 - reply: {%v}\n", unregistered)
	return mtproto.ToBool(unregistered), nil
}
