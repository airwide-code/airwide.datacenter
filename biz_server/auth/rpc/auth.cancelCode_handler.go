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
	"github.com/airwide-code/airwide.datacenter/biz/base"
	"github.com/airwide-code/airwide.datacenter/biz/core/auth"
)

// 客户端不处理错误码

// auth.cancelCode#1f040578 phone_number:string phone_code_hash:string = Bool;
func (s *AuthServiceImpl) AuthCancelCode(ctx context.Context, request *mtproto.TLAuthCancelCode) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("auth.cancelCode#1f040578 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// 1. check phone code
	if request.PhoneCodeHash == "" {
		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_BAD_REQUEST), "auth.resendCode#3ef1a9bf: phone code hash empty.")
		return nil, err
	}

	// 2. check number
	// 客户端发送的手机号格式为: "+86 111 1111 1111"，归一化
	phoneNumber, err := base.CheckAndGetPhoneNumber(request.GetPhoneNumber())
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	code := auth.MakeCancelCodeData(md.AuthId, phoneNumber, request.PhoneCodeHash)
	canceled := mtproto.ToBool(code.DoCancelCode())

	glog.Infof("auth.cancelCode#1f040578 -  - reply: %s", logger.JsonDebugData(canceled))
	return canceled, nil
}
