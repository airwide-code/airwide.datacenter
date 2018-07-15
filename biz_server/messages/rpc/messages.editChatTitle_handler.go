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
	"github.com/airwide-code/airwide.datacenter/biz/base"
	"github.com/airwide-code/airwide.datacenter/biz_server/sync_client"
	"github.com/airwide-code/airwide.datacenter/biz/core/message"
	"github.com/airwide-code/airwide.datacenter/biz/core/user"
	update2 "github.com/airwide-code/airwide.datacenter/biz/core/update"
)

// messages.editChatTitle#dc452855 chat_id:int title:string = Updates;
func (s *MessagesServiceImpl) MessagesEditChatTitle(ctx context.Context, request *mtproto.TLMessagesEditChatTitle) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.editChatTitle#dc452855 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	chatLogic, err := chat.NewChatLogicById(request.ChatId)
	if err != nil {
		glog.Error("messages.editChatTitle#dc452855 - error: ", err)
		return nil, err
	}

	peer := &base.PeerUtil{
		PeerType: base.PEER_CHAT,
		PeerId: chatLogic.GetChatId(),
	}

	err = chatLogic.EditChatTitle(md.UserId, request.Title)
	if err != nil {
		glog.Error("messages.editChatTitle#dc452855 - error: ", err)
		return nil, err
	}

	chatEditMessage := chatLogic.MakeChatEditTitleMessage(md.UserId, request.Title)

	randomId := base.NextSnowflakeId()
	outbox := message.CreateMessageOutboxByNew(md.UserId, peer, randomId, chatEditMessage, func(messageId int32) {
		user.CreateOrUpdateByOutbox(md.UserId, peer.PeerType, peer.PeerId, messageId, false, false)
	})

	syncUpdates := update2.NewUpdatesLogic(md.UserId)
	updateChatParticipants := &mtproto.TLUpdateChatParticipants{Data2: &mtproto.Update_Data{
		Participants: chatLogic.GetChatParticipants().To_ChatParticipants(),
	}}
	syncUpdates.AddUpdate(updateChatParticipants.To_Update())
	syncUpdates.AddUpdateNewMessage(chatEditMessage)
	syncUpdates.AddUsers(user.GetUsersBySelfAndIDList(md.UserId, []int32{md.UserId}))
	syncUpdates.AddChat(chatLogic.ToChat(md.UserId))

	state, _ := sync_client.GetSyncClient().SyncUpdatesData(md.AuthId, md.SessionId, md.UserId, syncUpdates.ToUpdates())
	syncUpdates.PushTopUpdateMessageId(outbox.MessageId, outbox.RandomId)
	syncUpdates.SetupState(state)

	replyUpdates := syncUpdates.ToUpdates()

	inboxList, _ := outbox.InsertMessageToInbox(md.UserId, peer, func(inBoxUserId, messageId int32) {
		user.CreateOrUpdateByInbox(inBoxUserId, base.PEER_CHAT, peer.PeerId, messageId, false)
	})

	for _, inbox := range inboxList {
		updates := update2.NewUpdatesLogic(md.UserId)
		updates.AddUpdate(updateChatParticipants.To_Update())
		updates.AddUpdateNewMessage(inbox.Message)
		updates.AddUsers(user.GetUsersBySelfAndIDList(md.UserId, []int32{md.UserId}))
		updates.AddChat(chatLogic.ToChat(inbox.UserId))
		sync_client.GetSyncClient().PushToUserUpdatesData(inbox.UserId, updates.ToUpdates())
	}

	glog.Infof("messages.editChatTitle#dc452855 - reply: {%v}", replyUpdates)
	return replyUpdates, nil
}
