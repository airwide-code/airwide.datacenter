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
)

// messages.getChats#3c6aa187 id:Vector<int> = messages.Chats;
func (s *MessagesServiceImpl) MessagesGetChats(ctx context.Context, request *mtproto.TLMessagesGetChats) (*mtproto.Messages_Chats, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getChats#3c6aa187 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): messages.chatsSlice
	chats := &mtproto.TLMessagesChats{Data2: &mtproto.Messages_Chats_Data{
		Chats: chat.GetChatListBySelfAndIDList(md.UserId, request.GetId()),
	}}

	glog.Infof("messages.getChats#3c6aa187 - reply: %s", chats)
	return chats.To_Messages_Chats(), nil
}
