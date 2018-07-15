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
	slave := dao.GetDevicesDAO(dao.DB_SLAVE)
	master := dao.GetDevicesDAO(dao.DB_MASTER)
	do := slave.SelectByToken(tokenType, token)
	if do == nil {
		do = &dataobject.DevicesDO{
			AuthKeyId: authKeyId,
			UserId: userId,
			TokenType: tokenType,
			Token: token,
		}
		do.Id = master.Insert(do)
	} else {
		master.UpdateStateById(0, do.Id)
	}

	return true
}

func UnRegisterDevice(tokenType int8, token string) bool {
	master := dao.GetDevicesDAO(dao.DB_MASTER)
	master.UpdateStateByToken(int8(1), tokenType, token)
	return true
}
