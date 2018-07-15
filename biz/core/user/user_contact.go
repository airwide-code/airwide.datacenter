/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package user

import (
	"github.com/airwide-code/airwide.datacenter/biz/dal/dao"
	"github.com/airwide-code/airwide.datacenter/mtproto"
)

func GetContactUserIDList(userId int32) []int32 {
	contactsDOList := dao.GetUserContactsDAO(dao.DB_SLAVE).SelectUserContacts(userId)
	idList := make([]int32, 0, len(contactsDOList))

	for _, do := range contactsDOList {
		idList = append(idList, do.ContactUserId)
	}
	return idList
}

func GetStatuseList(selfId int32) []*mtproto.ContactStatus {
	//doList := dao.GetUserContactsDAO(dao.DB_SLAVE).SelectUserContacts(selfId)
	//
	//contactIdList := make([]int32, 0, len(doList))
	//for _, do := range doList {
	//	contactIdList = append(contactIdList, do.ContactUserId)
	//}
	//return nil

	// TODO(@benqi): impl
	return []*mtproto.ContactStatus{}
}
