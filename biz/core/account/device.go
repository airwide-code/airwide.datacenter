/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package account

import (
	"github.com/airwide-code/airwide.datacenter/biz/dal/dao"
	"github.com/airwide-code/airwide.datacenter/biz/dal/dataobject"
)

//const (
//	TOKEN_TYPE_APNS = 1
//	TOKEN_TYPE_GCM = 2
//	TOKEN_TYPE_MPNS = 3
//	TOKEN_TYPE_SIMPLE_PUSH = 4
//	TOKEN_TYPE_UBUNTU_PHONE = 5
//	TOKEN_TYPE_BLACKBERRY = 6
//	// Android里使用
//	TOKEN_TYPE_INTERNAL_PUSH = 7
//)

func RegisterDevice(authKeyId int64, userId int32, tokenType int8, token string) bool {
	subordinate := dao.GetDevicesDAO(dao.DB_SLAVE)
	main := dao.GetDevicesDAO(dao.DB_MASTER)
	do := subordinate.SelectByToken(tokenType, token)
	if do == nil {
		do = &dataobject.DevicesDO{
			AuthKeyId: authKeyId,
			UserId: userId,
			TokenType: tokenType,
			Token: token,
		}
		do.Id = main.Insert(do)
	} else {
		main.UpdateStateById(0, do.Id)
	}

	return true
}

func UnRegisterDevice(tokenType int8, token string) bool {
	main := dao.GetDevicesDAO(dao.DB_MASTER)
	main.UpdateStateByToken(int8(1), tokenType, token)
	return true
}
