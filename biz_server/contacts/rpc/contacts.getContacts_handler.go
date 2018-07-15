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

// contacts.getContacts#c023849f hash:int = contacts.Contacts;
func (s *ContactsServiceImpl) ContactsGetContacts(ctx context.Context, request *mtproto.TLContactsGetContacts) (*mtproto.Contacts_Contacts, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("contacts.getContacts#c023849f - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		contacts *mtproto.Contacts_Contacts
	)
	contactLogic := contact.MakeContactLogic(md.UserId)

	contactList := contactLogic.GetContactList()
	// 避免查询数据库时IN()条件为empty
	if len(contactList) > 0 {
		idList := make([]int32, 0, len(contactList))
		cList := make([]*mtproto.Contact, 0, len(contactList))
		for _, c := range contactList {
			idList = append(idList, c.ContactUserId)
			c2 := &mtproto.Contact{
				Constructor: mtproto.TLConstructor_CRC32_contact,
				Data2: &mtproto.Contact_Data{
					UserId: c.ContactUserId,
					Mutual: mtproto.ToBool(c.Mutual == 1),
				},
			}
			cList = append(cList, c2)
		}

		glog.Infof("contactIdList - {%v}", idList)

		users := user.GetUsersBySelfAndIDList(md.UserId, idList)
		contacts = &mtproto.Contacts_Contacts{
			Constructor: mtproto.TLConstructor_CRC32_contacts_contacts,
			Data2: &mtproto.Contacts_Contacts_Data{
				Contacts:   cList,
				SavedCount: 0,
				Users:      users,
			},
		}
	} else {
		contacts = mtproto.NewTLContactsContacts().To_Contacts_Contacts()
	}

	glog.Infof("contacts.getContacts#c023849f - reply: %s\n", logger.JsonDebugData(contacts))
	return contacts, nil
}
