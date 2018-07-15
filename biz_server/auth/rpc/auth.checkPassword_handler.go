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

/*
	if (error.text.equals("PASSWORD_HASH_INVALID")) {
		onPasscodeError(true);
	} else if (error.text.startsWith("FLOOD_WAIT")) {
		int time = Utilities.parseInt(error.text);
		String timeString;
		if (time < 60) {
			timeString = LocaleController.formatPluralString("Seconds", time);
		} else {
			timeString = LocaleController.formatPluralString("Minutes", time / 60);
		}
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.formatString("FloodWaitTime", R.string.FloodWaitTime, timeString));
	} else {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), error.text);
	}
 */

// 客户端调用auth.signIn时返回SESSION_PASSWORD_NEEDED时会触发

// auth.checkPassword#a63011e password_hash:bytes = auth.Authorization;
func (s *AuthServiceImpl) AuthCheckPassword(ctx context.Context, request *mtproto.TLAuthCheckPassword) (*mtproto.Auth_Authorization, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("auth.checkPassword#a63011e - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		err error
	)

	if len(request.PasswordHash) == 0 {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PASSWORD_HASH_INVALID)
		glog.Error(err)
		return nil, err
	}

	passwordLogic, err := account.MakePasswordData(md.UserId)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	ok := passwordLogic.CheckPassword(request.PasswordHash)
	if !ok {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PASSWORD_HASH_INVALID)
		glog.Error(err)
		return nil, err
	}

	user := user2.GetUserById(md.UserId, md.UserId)
	authAuthorization := &mtproto.TLAuthAuthorization{Data2: &mtproto.Auth_Authorization_Data{
		User: user.To_User(),
	}}

	glog.Infof("auth.checkPassword#a63011e - reply: %s\n", logger.JsonDebugData(authAuthorization))
	return authAuthorization.To_Auth_Authorization(), nil
}
