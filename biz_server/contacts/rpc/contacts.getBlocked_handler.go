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
	"github.com/airwide-code/airwide.datacenter/biz/core/user"
	"github.com/airwide-code/airwide.datacenter/biz/core/contact"
)

// contacts.blocked#1c138d15 blocked:Vector<ContactBlocked> users:Vector<User> = contacts.Blocked;
// contacts.blockedSlice#900802a1 count:int blocked:Vector<ContactBlocked> users:Vector<User> = contacts.Blocked;
//
// contacts.getBlocked#f57c350f offset:int limit:int = contacts.Blocked;
func (s *ContactsServiceImpl) ContactsGetBlocked(ctx context.Context, request *mtproto.TLContactsGetBlocked) (*mtproto.Contacts_Blocked, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("contacts.getBlocked#f57c350f - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	contactLogic := contact.MakeContactLogic(md.UserId)
	blockedList := contactLogic.GetBlockedList(request.Offset, request.Limit)

	// TODO(@benqi): impl blockedSlice

	contactsBlocked := &mtproto.TLContactsBlocked{Data2: &mtproto.Contacts_Blocked_Data{
		Blocked: blockedList,
	}}
	// .NewTLContactsBlocked()
	if len(blockedList) > 0 {
		blockedIdList := make([]int32, 0, len(blockedList))
		userIdList := make([]int32, 0, len(blockedList))
		for _, c := range blockedList {
			userIdList = append(userIdList, c.GetData2().GetUserId())
		}

		users := user.GetUsersBySelfAndIDList(md.UserId, blockedIdList)
		contactsBlocked.SetUsers(users)
	}

	glog.Infof("contacts.getBlocked#f57c350f - reply: %s\n", logger.JsonDebugData(contactsBlocked))
	return contactsBlocked.To_Contacts_Blocked(), nil
}
