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
	chat2 "github.com/airwide-code/airwide.datacenter/biz/core/chat"
	"github.com/airwide-code/airwide.datacenter/biz/core/user"
)

// messages.getFullChat#3b831c66 chat_id:int = messages.ChatFull;
func (s *MessagesServiceImpl) MessagesGetFullChat(ctx context.Context, request *mtproto.TLMessagesGetFullChat) (*mtproto.Messages_ChatFull, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getFullChat#3b831c66 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): chat_id is channel

	chatLogic, err := chat2.NewChatLogicById(request.GetChatId())
	if err != nil {
		glog.Error("messages.getFullChat#3b831c66 - error: ", err)
		return nil, err
	}

	idList := chatLogic.GetChatParticipantIdList()
	messagesChatFull := &mtproto.TLMessagesChatFull{Data2: &mtproto.Messages_ChatFull_Data{
		FullChat: 	chat2.GetChatFullBySelfId(md.UserId, chatLogic).To_ChatFull(),
		Chats:      []*mtproto.Chat{chatLogic.ToChat(md.UserId)},
		Users: user.GetUsersBySelfAndIDList(md.UserId, idList),
	}}

	glog.Infof("messages.getFullChat#3b831c66 - reply: %s", logger.JsonDebugData(messagesChatFull))
	return messagesChatFull.To_Messages_ChatFull(), nil
}
