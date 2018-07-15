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

// {"token_type":10,"token":"{\"endpoint\":\"https://android.googleapis.com/gcm/send/fFBQfzHLq0I:APA91bF28ucFXm7ZF2T7sf87dKAARGXHzdK3HbK0rrhJMmPZr42amq6B-QASi-mLzOzZ5qsynyvtEOSNNYbvadNKI5LCxmYMhQXkhoh_fpTB0GsYLBjwpElaV68OmTUzN-AFDgWuqMIpQH5XYDZoYQopg-yHHdsxcQ\",\"expirationTime\":null,\"keys\":{\"p256dh\":\"BJLqPVxd2KNAmW_izYz4ha5hN4ZEzXnNbk4__FC-xhmaa2vZD3RRtvgPNphH8ZSM9wF4_vSTJZLzQ5Iv0byZxrY\",\"auth\":\"nBC8C_1cvhSTlEEelbk9kw\"}}","app_sandbox":{"constructor":-1132882121,"data2":{}}}
// account.registerDevice#637ea878 token_type:int token:string = Bool;
func (s *AccountServiceImpl) AccountRegisterDevice(ctx context.Context, request *mtproto.TLAccountRegisterDevice) (*mtproto.Bool, error) {
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
		// glog.Error("request.TokenType: ", request.TokenType)
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error(err)
		return nil, err
	}

	registered := account.RegisterDevice(md.AuthId, md.UserId, int8(request.TokenType), request.Token)

	glog.Infof("account.registerDevice#637ea878 - reply: {true}")
	return mtproto.ToBool(registered), nil
}
