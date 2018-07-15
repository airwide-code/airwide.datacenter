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
	"github.com/airwide-code/airwide.datacenter/biz/core/chat"
	"github.com/airwide-code/airwide.datacenter/biz_server/sync_client"
	"github.com/airwide-code/airwide.datacenter/biz/core/user"
	update2 "github.com/airwide-code/airwide.datacenter/biz/core/update"
)

// messages.editChatAdmin#a9e69f2e chat_id:int user_id:InputUser is_admin:Bool = Bool;
func (s *MessagesServiceImpl) MessagesEditChatAdmin(ctx context.Context, request *mtproto.TLMessagesEditChatAdmin) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.editChatAdmin#a9e69f2e - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		userId int32
		isAdmin = mtproto.FromBool(request.GetIsAdmin())
		err error
	)

	switch request.GetUserId().GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputUser:
		// TODO(@benqi): check user_id's access_hash
		userId = request.GetUserId().GetData2().GetUserId()
	default:
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error("messages.editChatAdmin#a9e69f2e - invalid user_id, err: ", err)
		return nil, err
	}

	chatLogic, err := chat.NewChatLogicById(request.ChatId)
	if err != nil {
		glog.Error("messages.editChatAdmin#a9e69f2e - error: ", err)
		return nil, err
	}

	err = chatLogic.EditChatAdmin(md.UserId, userId, isAdmin)
	if err != nil {
		glog.Error("messages.editChatAdmin#a9e69f2e - error: ", err)
		return nil, err
	}

	updateChatParticipants := &mtproto.TLUpdateChatParticipants{Data2: &mtproto.Update_Data{
		Participants: chatLogic.GetChatParticipants().To_ChatParticipants(),
	}}

	idList := chatLogic.GetChatParticipantIdList()
	for _, id := range idList {
		updates := update2.NewUpdatesLogic(md.UserId)
		updates.AddUpdate(updateChatParticipants.To_Update())
		updates.AddUsers(user.GetUsersBySelfAndIDList(id, idList))
		updates.AddChat(chatLogic.ToChat(md.UserId))
		sync_client.GetSyncClient().PushToUserUpdatesData(id, updates.ToUpdates())
	}

	glog.Infof("messages.editChatAdmin#a9e69f2e - reply: {true}")
	return mtproto.ToBool(true), nil
}
