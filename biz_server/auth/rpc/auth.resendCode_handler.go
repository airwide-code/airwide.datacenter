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

/*
 Android client auth.resendCode#3ef1a9bf, handler error
	if (error.text != null) {
		if (error.text.contains("PHONE_NUMBER_INVALID")) {
			needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("InvalidPhoneNumber", R.string.InvalidPhoneNumber));
		} else if (error.text.contains("PHONE_CODE_EMPTY") || error.text.contains("PHONE_CODE_INVALID")) {
			needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("InvalidCode", R.string.InvalidCode));
		} else if (error.text.contains("PHONE_CODE_EXPIRED")) {
			onBackPressed();
			setPage(0, true, null, true);
			needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("CodeExpired", R.string.CodeExpired));
		} else if (error.text.startsWith("FLOOD_WAIT")) {
			needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("FloodWait", R.string.FloodWait));
		} else if (error.code != -1000) {
			needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("ErrorOccurred", R.string.ErrorOccurred) + "\n" + error.text);
		}
	}
 */

// auth.resendCode#3ef1a9bf phone_number:string phone_code_hash:string = auth.SentCode;
func (s *AuthServiceImpl) AuthResendCode(ctx context.Context, request *mtproto.TLAuthResendCode) (*mtproto.Auth_SentCode, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AuthResendCode - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// 1. check phone code
	if request.PhoneCodeHash == "" {
		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_BAD_REQUEST), "auth.resendCode#3ef1a9bf: phone code hash empty.")
		glog.Error(err)
		return nil, err
	}

	// 2. check number
	// 客户端发送的手机号格式为: "+86 111 1111 1111"，归一化
	phoneNumber, err := base.CheckAndGetPhoneNumber(request.GetPhoneNumber())
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	// TODO(@benqi): PHONE_NUMBER_FLOOD
	// <string name="PhoneNumberFlood">Sorry, you have deleted and re-created your account too many times recently.
	//    Please wait for a few days before signing up again.</string>
	//

	// PHONE_NUMBER_BANNED: Banned phone number
	banned := auth.CheckBannedByPhoneNumber(phoneNumber)
	if banned {
		err = mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_NUMBER_BANNED), "auth.sendCode#86aef0ec: phone number banned.")
		glog.Error(err)
		return nil, err
	}

	code := auth.MakeCodeDataByHash(md.AuthId, phoneNumber, request.GetPhoneCodeHash())
	err = code.DoReSendCode()
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	// 使用app类型，code统一为123456
	phoneRegistered := auth.CheckPhoneNumberExist(phoneNumber)
	authSentCode := code.ToAuthSentCode(phoneRegistered)

	glog.Infof("auth.resendCode#3ef1a9bf - reply: %s", logger.JsonDebugData(authSentCode))
	return authSentCode.To_Auth_SentCode(), nil
}
