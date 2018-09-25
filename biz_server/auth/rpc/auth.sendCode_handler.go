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
 Android client auth.sendCode#86aef0ec, handler error
 1.
	if (error->error_code == 303) {
		uint32_t migrateToDatacenterId = DEFAULT_DATACENTER_ID;

		static std::vector<std::string> migrateErrors;
		if (migrateErrors.empty()) {
			migrateErrors.push_back("NETWORK_MIGRATE_");
			migrateErrors.push_back("PHONE_MIGRATE_");
			migrateErrors.push_back("USER_MIGRATE_");
		}

		size_t count = migrateErrors.size();
		for (uint32_t a = 0; a < count; a++) {
			std::string &possibleError = migrateErrors[a];
			if (error->error_message.find(possibleError) != std::string::npos) {
				std::string num = error->error_message.substr(possibleError.size(), error->error_message.size() - possibleError.size());
				uint32_t val = (uint32_t) atoi(num.c_str());
				migrateToDatacenterId = val;
			}
		}

		if (migrateToDatacenterId != DEFAULT_DATACENTER_ID) {
			ignoreResult = true;
			moveToDatacenter(migrateToDatacenterId);
		}
	}

 2.
	if (error.text != null) {
		if (error.text.contains("PHONE_NUMBER_INVALID")) {
			needShowInvalidAlert(req.phone_number, false);
		} else if (error.text.contains("PHONE_NUMBER_FLOOD")) {
			needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("PhoneNumberFlood", R.string.PhoneNumberFlood));
		} else if (error.text.contains("PHONE_NUMBER_BANNED")) {
			needShowInvalidAlert(req.phone_number, true);
		} else if (error.text.contains("PHONE_CODE_EMPTY") || error.text.contains("PHONE_CODE_INVALID")) {
			needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("InvalidCode", R.string.InvalidCode));
		} else if (error.text.contains("PHONE_CODE_EXPIRED")) {
			needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("CodeExpired", R.string.CodeExpired));
		} else if (error.text.startsWith("FLOOD_WAIT")) {
			needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("FloodWait", R.string.FloodWait));
		} else if (error.code != -1000) {
			needShowAlert(LocaleController.getString("AppName", R.string.AppName), error.text);
		}
	}
 */

// auth.sendCode#86aef0ec flags:# allow_flashcall:flags.0?true phone_number:string current_number:flags.0?Bool api_id:int api_hash:string = auth.SentCode;
func (s *AuthServiceImpl) AuthSendCode(ctx context.Context, request *mtproto.TLAuthSendCode) (*mtproto.Auth_SentCode, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AuthSendCode - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): To access the telegram network, you must first apply for api_id and api_hash to verify that api_id and api_hash are legal. 
	// 1. check api_id and api_hash

	//// 3. check number
	////    The format of the mobile phone number sent by the client is: "+86 111 1111 1111", normalized 
	phoneNumber, err :=  base.CheckAndGetPhoneNumber(request.GetPhoneNumber())
	if err != nil {
		// PHONE_NUMBER_INVALID
		glog.Error(err)
		return nil, err
	}

	// 2. check allow_flashcall and current_number
	// CurrentNumber: Is it a local phone number? 

	// if allow_flashcall is true then current_number is true
	var currentNumber bool
	if request.GetCurrentNumber() == nil {
		currentNumber = false
	} else {
		currentNumber = mtproto.FromBool(request.GetCurrentNumber())
	}
	//if !currentNumber && request.GetAllowFlashcall() {
	//	err = mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_BAD_REQUEST), "auth.sendCode#86aef0ec: current_number is true but allow_flashcall is false.")
	//	glog.Error(err)
	//	return nil, err
	//}

	// TODO(@benqi): PHONE_NUMBER_FLOOD
	// <string name="PhoneNumberFlood">Sorry, you have deleted and re-created your account too many times recently.
	//    Please wait for a few days before signing up again.</string>
	//

	// glog.Info("phoneNumber: ", phoneNumber)
	// PHONE_NUMBER_BANNED: Banned phone number
	banned := auth.CheckBannedByPhoneNumber(phoneNumber)
	if banned {
		err = mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_NUMBER_BANNED), "auth.sendCode#86aef0ec: phone number banned.")
		glog.Error(err)
		return nil, err
	}

	// TODO(@benqi): MIGRATE datacenter
	// android client:
	//  migrateErrors.push_back("NETWORK_MIGRATE_");
	//  migrateErrors.push_back("PHONE_MIGRATE_");
	//  migrateErrors.push_back("USER_MIGRATE_");
	//
	// https://core.telegram.org/api/datacenter
	// The auth.sendCode method is the basic entry point when registering a new user or authorizing an existing user.
	//   95% of all redirection cases to a different DC will occure when invoking this method.
	//
	// The client does not yet know which DC it will be associated with; therefore,
	//   it establishes an encrypted connection to a random address and sends its query to that address.
	// Having received a phone_number from a client,
	// 	 we can find out whether or not it is registered in the system.
	//   If it is, then, if necessary, instead of sending a text message,
	//   we request that it establish a connection with a different DC first (PHONE_MIGRATE_X error).
	// If we do not yet have a user with this number, we examine its IP-address.
	//   We can use it to identify the closest DC.
	//   Again, if necessary, we redirect the user to a different DC (NETWORK_MIGRATE_X error).
	//
	//if userDO == nil {
	//	// phone registered
	//	// TODO(@benqi): Optimized by phoneNumber and ip 
	//} else {
	//	// TODO(@benqi): Preferred by userId 
	//}

	code := auth.MakeCodeData(md.AuthId, phoneNumber)

	// Check if phoneNumber is abnormal 
	// TODO(@benqi): Check if phoneNumber is abnormal 
	// PhoneNumberFlood
	// FLOOD_WAIT
	phoneRegistered := auth.CheckPhoneNumberExist(phoneNumber)
	err = code.DoSendCode(phoneRegistered, request.AllowFlashcall, currentNumber, request.ApiId, request.ApiHash)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	authSentCode := code.ToAuthSentCode(phoneRegistered)
	glog.Infof("AuthSendCode - reply: %s", logger.JsonDebugData(authSentCode))
	return authSentCode.To_Auth_SentCode(), nil
}
