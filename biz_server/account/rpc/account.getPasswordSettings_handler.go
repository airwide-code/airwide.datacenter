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

// account.getPasswordSettings#bc8d11bb current_password_hash:bytes = account.PasswordSettings;
func (s *AccountServiceImpl) AccountGetPasswordSettings(ctx context.Context, request *mtproto.TLAccountGetPasswordSettings) (*mtproto.Account_PasswordSettings, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.getPasswordSettings#bc8d11bb - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	passwordLogic, err := account.MakePasswordData(md.UserId)
	if err != nil {
		glog.Error("account.getPassword#548a30f5 - error: ", err)
		return nil, err
	}

	settings, err := passwordLogic.GetPasswordSetting(request.GetCurrentPasswordHash())
	if err != nil {
		glog.Error("account.getPassword#548a30f5 - error: ", err)
		return nil, err
	}

	glog.Infof("account.getPasswordSettings#bc8d11bb - reply: %s", logger.JsonDebugData(settings))
	return settings, nil
}
