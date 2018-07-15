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
	"github.com/airwide-code/airwide.datacenter/biz_server/sync_client"
	"github.com/airwide-code/airwide.datacenter/biz/core/chat"
	update2 "github.com/airwide-code/airwide.datacenter/biz/core/update"
)

// messages.toggleChatAdmins#ec8bd9e1 chat_id:int enabled:Bool = Updates;
func (s *MessagesServiceImpl) MessagesToggleChatAdmins(ctx context.Context, request *mtproto.TLMessagesToggleChatAdmins) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.toggleChatAdmins#ec8bd9e1 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	chatLogic, err := chat.NewChatLogicById(request.ChatId)
	if err != nil {
		glog.Error("toggleChatAdmins#ec8bd9e1 - error: ", err)
		return nil, err
	}

	err = chatLogic.ToggleChatAdmins(md.UserId, mtproto.FromBool(request.GetEnabled()))
	if err != nil {
		glog.Error("toggleChatAdmins#ec8bd9e1 - error: ", err)
		return nil, err
	}

	syncUpdates := update2.NewUpdatesLogic(md.UserId)
	//updateChatParticipants := &mtproto.TLUpdateChatParticipants{Data2: &mtproto.Update_Data{
	//	Participants: chatLogic.GetChatParticipants().To_ChatParticipants(),
	//}}
	//syncUpdates.AddUpdate(updateChatParticipants.To_Update())
	syncUpdates.AddChat(chatLogic.ToChat(md.UserId))

	replyUpdates := syncUpdates.ToUpdates()

	updateChatAdmins := &mtproto.TLUpdateChatAdmins{Data2: &mtproto.Update_Data{
		ChatId:  chatLogic.GetChatId(),
		Enabled: request.GetEnabled(),
		Version: chatLogic.GetVersion(),
	}}

	sync_client.GetSyncClient().PushToUserNotMeUpdateShortData(md.AuthId, md.SessionId, md.UserId, updateChatAdmins.To_Update())


	idList := chatLogic.GetChatParticipantIdList()
	for _, id := range idList {
		sync_client.GetSyncClient().PushToUserUpdateShortData(id, updateChatAdmins.To_Update())
	}

	glog.Infof("messages.toggleChatAdmins#ec8bd9e1 - reply: {%v}", replyUpdates)
	return replyUpdates, nil
}
