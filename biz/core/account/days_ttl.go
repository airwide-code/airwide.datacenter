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
)

func SetAccountDaysTTL(userId int32, ttl int32) {
	dao.GetUsersDAO(dao.DB_MASTER).UpdateAccountDaysTTL(ttl, userId)
}

func GetAccountDaysTTL(userId int32) int32 {
	do := dao.GetUsersDAO(dao.DB_SLAVE).SelectAccountDaysTTL(userId)
	return do.AccountDaysTtl
}
