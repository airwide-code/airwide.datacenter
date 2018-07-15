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
	"github.com/airwide-code/airwide.datacenter/biz/core/message"
	"github.com/airwide-code/airwide.datacenter/biz/core/user"
	"github.com/airwide-code/airwide.datacenter/biz/core/chat"
)

// messages.getMessages#4222fa74 id:Vector<int> = messages.Messages;
func (s *MessagesServiceImpl) MessagesGetMessages(ctx context.Context, request *mtproto.TLMessagesGetMessages) (*mtproto.Messages_Messages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getMessages#4222fa74 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	messages :=message.GetMessagesByPeerAndMessageIdList2(md.UserId, request.Id)
	userIdList, chatIdList, _ := message.PickAllIDListByMessages(messages)
	userList := user.GetUsersBySelfAndIDList(md.UserId, userIdList)
	chatList := chat.GetChatListBySelfAndIDList(md.UserId, chatIdList)

	messagesMessages := &mtproto.TLMessagesMessages{Data2: &mtproto.Messages_Messages_Data{
		Messages: messages,
		Users: userList,
		Chats: chatList,
	}}

	glog.Infof("messages.getMessages#4222fa74 - reply: %s", logger.JsonDebugData(messagesMessages))
	return messagesMessages.To_Messages_Messages(), nil
}
