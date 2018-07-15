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
	"github.com/airwide-code/airwide.datacenter/biz/core/contact"
	"github.com/airwide-code/airwide.datacenter/biz_server/sync_client"
	updates2 "github.com/airwide-code/airwide.datacenter/biz/core/update"
)

// contacts.block#332b49fc id:InputUser = Bool;
func (s *ContactsServiceImpl) ContactsBlock(ctx context.Context, request *mtproto.TLContactsBlock) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("contacts.block#332b49fc - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		blockId int32
		id = request.Id
	)

	switch id.GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputUserSelf:
		blockId = md.UserId
	case mtproto.TLConstructor_CRC32_inputUser:
		// Check access hash
		if ok := user2.CheckAccessHashByUserId(id.GetData2().GetUserId(), id.GetData2().GetAccessHash()); !ok {
			// TODO(@benqi): Add ACCESS_HASH_INVALID codes
			err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
			glog.Error(err, ": is access_hash error")
			return nil, err
		}

		blockId = id.GetData2().GetUserId()
		// TODO(@benqi): contact exist
	default:
		// mtproto.TLConstructor_CRC32_inputUserEmpty:
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error(err, ": is inputUserEmpty")
		return nil, err
	}

	contactLogic :=contact.MakeContactLogic(md.UserId)
	blocked := contactLogic.BlockUser(blockId)

	if blocked {
		// Sync unblocked: updateUserBlocked
		updateUserBlocked := &mtproto.TLUpdateUserBlocked{Data2: &mtproto.Update_Data{
			UserId: blockId,
			Blocked: mtproto.ToBool(true),
		}}

		blockedUpdates := updates2.NewUpdatesLogic(md.UserId)
		blockedUpdates.AddUpdate(updateUserBlocked.To_Update())
		blockedUpdates.AddUser(user2.GetUserById(md.UserId, blockId).To_User())

		// TODO(@benqi): handle seq
		sync_client.GetSyncClient().SyncUpdatesData(md.AuthId, md.SessionId, blockId, blockedUpdates.ToUpdates())
	}

	// Blocked会影响收件箱
	glog.Infof("contacts.block#332b49fc - reply: {%v}", blocked)
	return mtproto.ToBool(blocked), nil
}
