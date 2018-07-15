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
	user2 "github.com/airwide-code/airwide.datacenter/biz/core/user"
	"github.com/airwide-code/airwide.datacenter/biz_server/sync_client"
	updates2 "github.com/airwide-code/airwide.datacenter/biz/core/update"
)

// contacts.unblock#e54100bd id:InputUser = Bool;
func (s *ContactsServiceImpl) ContactsUnblock(ctx context.Context, request *mtproto.TLContactsUnblock) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("contacts.unblock#e54100bd - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		blockedId int32
		id = request.Id
	)

	switch id.GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputUserSelf:
		blockedId = md.UserId
	case mtproto.TLConstructor_CRC32_inputUser:
		// Check access hash
		if ok := user2.CheckAccessHashByUserId(id.GetData2().GetUserId(), id.GetData2().GetAccessHash()); !ok {
			// TODO(@benqi): Add ACCESS_HASH_INVALID codes
			err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
			glog.Error(err, ": is access_hash error")
			return nil, err
		}

		blockedId = id.GetData2().GetUserId()
		// TODO(@benqi): contact exist
	default:
		// mtproto.TLConstructor_CRC32_inputUserEmpty:
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error(err, ": is inputUserEmpty")
		return nil, err
	}

	contactLogic :=contact.MakeContactLogic(md.UserId)
	unBlocked := contactLogic.UnBlockUser(blockedId)

	if unBlocked {
		// Sync unblocked: updateUserBlocked
		updateUserUnBlocked := &mtproto.TLUpdateUserBlocked{Data2: &mtproto.Update_Data{
			UserId: blockedId,
			Blocked: mtproto.ToBool(false),
		}}

		unBlockedUpdates := updates2.NewUpdatesLogic(md.UserId)
		unBlockedUpdates.AddUpdate(updateUserUnBlocked.To_Update())
		unBlockedUpdates.AddUser(user2.GetUserById(md.UserId, blockedId).To_User())

		// TODO(@benqi): handle seq
		sync_client.GetSyncClient().SyncUpdatesData(md.AuthId, md.SessionId, blockedId, unBlockedUpdates.ToUpdates())
	}

	glog.Infof("contacts.unblock#e54100bd - reply: {%v}", unBlocked)
	return mtproto.ToBool(unBlocked), nil
}
