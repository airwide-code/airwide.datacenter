/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package auth

import (
	"github.com/airwide-code/airwide.datacenter/biz/dal/dao"
	"github.com/airwide-code/airwide.datacenter/biz/dal/dataobject"
)

func CheckBannedByPhoneNumber(phoneNumber string) bool {
	params := map[string]interface{}{
		"phone": phoneNumber,
	}
	return dao.GetCommonDAO(dao.DB_SLAVE).CheckExists("banned", params)
}

func CheckPhoneNumberExist(phoneNumber string) bool {
	params := map[string]interface{}{
		"phone": phoneNumber,
	}
	return dao.GetCommonDAO(dao.DB_SLAVE).CheckExists("users", params)
}

func BindAuthKeyAndUser(authKeyId int64, userId int32) {
	do3 := dao.GetAuthUsersDAO(dao.DB_MASTER).SelectByAuthId(authKeyId)
	if do3 == nil {
	    do3 := &dataobject.AuthUsersDO{
			AuthId: authKeyId,
			UserId: userId,
		}
		dao.GetAuthUsersDAO(dao.DB_MASTER).Insert(do3)
	}
}

/*
  auth.checkedPhone#811ea28e phone_registered:Bool = auth.CheckedPhone;
  auth.sentCode#5e002502 flags:# phone_registered:flags.0?true type:auth.SentCodeType phone_code_hash:string next_type:flags.1?auth.CodeType timeout:flags.2?int = auth.SentCode;
  auth.authorization#cd050916 flags:# tmp_sessions:flags.0?int user:User = auth.Authorization;
  auth.exportedAuthorization#df969c2d id:int bytes:bytes = auth.ExportedAuthorization;
*/
