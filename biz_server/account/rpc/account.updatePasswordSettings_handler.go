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

// account.updatePasswordSettings#fa7c4b86 current_password_hash:bytes new_settings:account.PasswordInputSettings = Bool;
func (s *AccountServiceImpl) AccountUpdatePasswordSettings(ctx context.Context, request *mtproto.TLAccountUpdatePasswordSettings) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.updatePasswordSettings#fa7c4b86 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	passwordInputSetting := request.NewSettings.To_AccountPasswordInputSettings()

	// TODO(@benqi): check request invalid

	passwordLogic, err := account.MakePasswordData(md.UserId)
	if err == nil {
		err = passwordLogic.UpdatePasswordSetting(request.CurrentPasswordHash,
			passwordInputSetting.GetNewSalt(),
			passwordInputSetting.GetNewPasswordHash(),
			passwordInputSetting.GetHint(),
			passwordInputSetting.GetEmail())

		// 未注册：error_message: "EMAIL_UNCONFIRMED" [STRING],
	}

	if err != nil {
		glog.Error("account.updatePasswordSettings#fa7c4b86 - error: ", err)
		return nil, err
	}

	reply := mtproto.ToBool(true)
	glog.Infof("account.getPassword#548a30f5 - reply: {}", logger.JsonDebugData(reply))
	return reply, nil
}
