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
	user2 "github.com/airwide-code/airwide.datacenter/biz/core/user"
	"github.com/airwide-code/airwide.datacenter/biz_server/sync_client"
	"github.com/airwide-code/airwide.datacenter/biz/core/contact"
	updates2 "github.com/airwide-code/airwide.datacenter/biz/core/update"
)

// contacts.deleteContacts#59ab389e id:Vector<InputUser> = Bool;
func (s *ContactsServiceImpl) ContactsDeleteContacts(ctx context.Context, request *mtproto.TLContactsDeleteContacts) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("contacts.deleteContacts#59ab389e - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// 注意: 目前只支持导入1个并且已经注册的联系人!!!!
	if len(request.Id) == 0 || len(request.Id) > 1 {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error(err, ": query or limit invalid")
		return nil, err
	}

	// TODO(@benqi): Copy from contacts.deleteContact, wrap func!!!
	var (
		deleteId int32
		id = request.Id[0]
	)

	switch id.GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputUserSelf:
		deleteId = md.UserId
	case mtproto.TLConstructor_CRC32_inputUser:
		// Check access hash
		if ok := user2.CheckAccessHashByUserId(id.GetData2().GetUserId(), id.GetData2().GetAccessHash()); !ok {
			// TODO(@benqi): Add ACCESS_HASH_INVALID codes
			err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
			glog.Error(err, ": is access_hash error")
			return nil, err
		}

		deleteId = id.GetData2().GetUserId()
		// TODO(@benqi): contact exist
	default:
		// mtproto.TLConstructor_CRC32_inputUserEmpty:
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error(err, ": is inputUserEmpty")
		return nil, err
	}

	// selfUser := user2.GetUserById(md.UserId, md.UserId)
	deleteUser := user2.GetUserById(md.UserId, deleteId)

	contactLogic := contact.MakeContactLogic(md.UserId)
	needUpdate := contactLogic.DeleteContact(deleteId, deleteUser.GetMutualContact())

	selfUpdates := updates2.NewUpdatesLogic(md.UserId)
	contactLink := &mtproto.TLUpdateContactLink{Data2: &mtproto.Update_Data{
		UserId:      deleteId,
		MyLink:      mtproto.NewTLContactLinkHasPhone().To_ContactLink(),
		ForeignLink: mtproto.NewTLContactLinkHasPhone().To_ContactLink(),
	}}
	selfUpdates.AddUpdate(contactLink.To_Update())
	selfUpdates.AddUser(deleteUser.To_User())
	// TODO(@benqi): handle seq
	sync_client.GetSyncClient().PushToUserUpdatesData(md.UserId, selfUpdates.ToUpdates())

	// TODO(@benqi): 推给联系人逻辑需要再考虑考虑
	if needUpdate {
		// TODO(@benqi): push to contact user update contact link
		contactUpdates := updates2.NewUpdatesLogic(deleteUser.GetId())
		contactLink2 := &mtproto.TLUpdateContactLink{Data2: &mtproto.Update_Data{
			UserId:      md.UserId,
			MyLink:      mtproto.NewTLContactLinkContact().To_ContactLink(),
			ForeignLink: mtproto.NewTLContactLinkContact().To_ContactLink(),
		}}
		contactUpdates.AddUpdate(contactLink2.To_Update())

		selfUser := user2.GetUserById(md.UserId, md.UserId)
		contactUpdates.AddUser(selfUser.To_User())
		// TODO(@benqi): handle seq
		sync_client.GetSyncClient().PushToUserUpdatesData(deleteId, contactUpdates.ToUpdates())
	}

	glog.Infof("contacts.deleteContacts#59ab389e - reply: {true}")
	return mtproto.ToBool(true), nil
}
