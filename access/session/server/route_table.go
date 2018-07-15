/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package server

import (
	"github.com/airwide-code/airwide.datacenter/mtproto"
)

//const (
//	GENERIC = 0
//	DOWNLOAD = 1
//	UPLOAD = 3
//
//	// Android
//	PUSH = 7
//
//	// 暂时不考虑
//	TEMP = 8
//
//	INVALID_TYPE = -1 // math.MaxInt32
//)
//
// TLObjectType ==> (service, auth, proto)
//
// 路由表
//type RouteTable struct {
//}

const (
	TOKEN_TYPE_APNS         = 1
	TOKEN_TYPE_GCM          = 2
	TOKEN_TYPE_MPNS         = 3
	TOKEN_TYPE_SIMPLE_PUSH  = 4
	TOKEN_TYPE_UBUNTU_PHONE = 5
	TOKEN_TYPE_BLACKBERRY   = 6
	// Android里使用
	TOKEN_TYPE_INTERNAL_PUSH = 7
)

func getConnectionType(tl mtproto.TLObject) int {
	switch tl.(type) {
	case *mtproto.TLAccountRegisterDevice,
		*mtproto.TLAccountUnregisterDevice,
		*mtproto.TLAccountRegisterDeviceLayer74,
		*mtproto.TLAccountUnregisterDeviceLayer74:
		return PUSH
		// check android internal push connection
		//reg, _ := tl.(*mtproto.TLAccountRegisterDevice)
		//if reg.GetTokenType() == TOKEN_TYPE_INTERNAL_PUSH {
		//	// android
		//	return PUSH
		//} else {
		//	return GENERIC
		//}
	case *mtproto.TLUploadSaveFilePart,
		*mtproto.TLUploadSaveBigFilePart:
		// upload connection
		return UPLOAD
	case *mtproto.TLUploadGetFile,
		*mtproto.TLUploadGetWebFile,
		*mtproto.TLUploadGetCdnFile,
		*mtproto.TLUploadReuploadCdnFile,
		*mtproto.TLUploadGetCdnFileHashes:
		// download connection
		return DOWNLOAD
	case *mtproto.TLHelpGetConfig:
		// TODO(@benqi): 可能为TEMP，判断TEMP的规则：
		//  从android client代码看，此连接仅收到help.getConfig消息
	}

	return GENERIC
}

// TL_get_future_salts

// TL_rpc_drop_answer

// TL_auth_exportedAuthorization
// TL_auth_exportAuthorization
// TL_auth_importAuthorization
// TL_auth_sendCode
// TL_auth_cancelCode
// TL_auth_resendCode
// TL_auth_resendCode
// TL_auth_signIn
// TL_auth_requestPasswordRecovery
// TL_auth_checkPassword
// TL_auth_recoverPassword
// TL_auth_signUp
// TL_auth_requestPasswordRecovery
// TL_auth_recoverPassword

// TL_help_getCdnConfig
// TL_help_getConfig

// TL_langpack_getLanguages
// TL_langpack_getDifference
// TL_langpack_getLangPack
// TL_langpack_getStrings
// TL_langpack_getStrings
// TL_langpack_getStrings

// TL_account_getPassword
// TL_account_deleteAccount
// TL_account_deleteAccount
// TL_account_getPassword
// TL_account_updatePasswordSettings
// TL_account_getPasswordSettings
func checkWithoutLogin(tl mtproto.TLObject) bool {
	switch tl.(type) {
	case *mtproto.TLMsgsAck:
		return true
	// TL_get_future_salts
	case *mtproto.TLGetFutureSalts:
		return true

	case *mtproto.TLRpcDropAnswer:
		return true

	case *mtproto.TLAuthCheckPhone,
		*mtproto.TLAuthSendCodeLayer51,
		*mtproto.TLAuthSendCode,
		*mtproto.TLAuthSignIn,
		*mtproto.TLAuthSignUp,
		*mtproto.TLAuthExportedAuthorization,
		*mtproto.TLAuthExportAuthorization,
		*mtproto.TLAuthImportAuthorization,
		*mtproto.TLAuthCancelCode,
		*mtproto.TLAuthResendCode,
		*mtproto.TLAuthRequestPasswordRecovery,
		*mtproto.TLAuthCheckPassword,
		*mtproto.TLAuthRecoverPassword:
		return true

	case *mtproto.TLHelpGetConfig,
		*mtproto.TLHelpGetCdnConfig:
		return true

	case *mtproto.TLLangpackGetLanguages,
		*mtproto.TLLangpackGetDifference,
		*mtproto.TLLangpackGetLangPack,
		*mtproto.TLLangpackGetStrings:
		return true

	case *mtproto.TLAccountGetPassword,
		*mtproto.TLAccountDeleteAccount,
		*mtproto.TLAccountUpdatePasswordSettings,
		*mtproto.TLAccountGetPasswordSettings:
		return true

	case *mtproto.TLInitConnection:
		return true
	}

	// glog.Warning("")
	return false
}

func checkNbfsRpcRequest(tl mtproto.TLObject) bool {
	switch tl.(type) {
	case *mtproto.TLUploadSaveFilePart,
		*mtproto.TLUploadGetFile,
		*mtproto.TLUploadSaveBigFilePart,
		*mtproto.TLUploadGetWebFile,
		*mtproto.TLUploadGetCdnFile,
		*mtproto.TLUploadReuploadCdnFile,
		*mtproto.TLUploadGetCdnFileHashes:
		return true
	}
	return false
}
