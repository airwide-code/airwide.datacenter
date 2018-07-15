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
	"github.com/airwide-code/airwide.datacenter/biz/core/contact"
	"github.com/airwide-code/airwide.datacenter/biz/core/user"
)

// contacts.getStatuses#c4a353ee = Vector<ContactStatus>;
func (s *ContactsServiceImpl) ContactsGetStatuses(ctx context.Context, request *mtproto.TLContactsGetStatuses) (*mtproto.Vector_ContactStatus, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("contacts.getStatuses#c4a353ee - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	contactLogic := contact.MakeContactLogic(md.UserId)
	cList := contactLogic.GetContactList()

	statusList := &mtproto.Vector_ContactStatus{
		Datas: make([]*mtproto.ContactStatus, 0, len(cList)),
	}

	for _, c := range cList {
		contactStatus := &mtproto.TLContactStatus{Data2: &mtproto.ContactStatus_Data{
			UserId: c.ContactUserId,
			Status: user.GetUserStatus(c.ContactUserId),
		}}
		statusList.Datas = append(statusList.Datas, contactStatus.To_ContactStatus())
	}

	glog.Infof("contacts.getStatuses#c4a353ee - reply: ", logger.JsonDebugData(statusList))
	return statusList, nil
}
