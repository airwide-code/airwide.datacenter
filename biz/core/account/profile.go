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

// not found, return 0
func GetUserIdByUserName(name string) int32 {
	do := dao.GetUsersDAO(dao.DB_SLAVE).SelectByUsername(name)
	if do == nil {
		return 0
	}
	return do.Id
}

func ChangeUserNameByUserId(id int32, name string) int64 {
	return dao.GetUsersDAO(dao.DB_MASTER).UpdateUsername(name, id)
}

func UpdateFirstAndLastName(id int32, firstName, lastName string) int64 {
	return dao.GetUsersDAO(dao.DB_MASTER).UpdateFirstAndLastName(firstName, lastName, id)
}

func UpdateAbout(id int32, about string) int64 {
	return dao.GetUsersDAO(dao.DB_MASTER).UpdateAbout(about, id)
}
