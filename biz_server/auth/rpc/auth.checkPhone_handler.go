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
	"github.com/airwide-code/airwide.datacenter/biz/core/user"
	"github.com/airwide-code/airwide.datacenter/biz/base"
)

// tdesktop客户端会调用，android客户端未使用

// auth.checkPhone#6fe51dfb phone_number:string = auth.CheckedPhone;
func (s *AuthServiceImpl) AuthCheckPhone(ctx context.Context, request *mtproto.TLAuthCheckPhone) (*mtproto.Auth_CheckedPhone, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("auth.checkPhone#6fe51dfb - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	phoneNumber, err := base.CheckAndGetPhoneNumber(request.GetPhoneNumber())
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	registered := user.CheckPhoneNumberExist(phoneNumber)
	checkedPhone := mtproto.TLAuthCheckedPhone{Data2: &mtproto.Auth_CheckedPhone_Data{
		PhoneRegistered: mtproto.ToBool(registered),
	}}

	glog.Infof("uth.checkPhone#6fe51dfb - reply: %s\n", checkedPhone)
	return checkedPhone.To_Auth_CheckedPhone(), nil
}
