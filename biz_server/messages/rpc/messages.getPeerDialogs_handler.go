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
	"github.com/airwide-code/airwide.datacenter/biz/base"
	"github.com/airwide-code/airwide.datacenter/biz/core/user"
	"github.com/airwide-code/airwide.datacenter/biz/core/message"
	"github.com/airwide-code/airwide.datacenter/biz/core/chat"
	update2 "github.com/airwide-code/airwide.datacenter/biz/core/update"
)

// messages.getPeerDialogs#2d9776b9 peers:Vector<InputPeer> = messages.PeerDialogs;
func (s *MessagesServiceImpl) MessagesGetPeerDialogs(ctx context.Context, request *mtproto.TLMessagesGetPeerDialogs) (*mtproto.Messages_PeerDialogs, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getPeerDialogs#2d9776b9 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	peerDialogs := mtproto.NewTLMessagesPeerDialogs()

	messageIdList := []int32{}
	userIdList := []int32{md.UserId}
	chatIdList := []int32{}

	dialogs := user.GetPeersDialogs(md.UserId, request.GetPeers())

	for _, dialog2 := range dialogs {
		// dialog.Peer
		dialog := dialog2.To_Dialog()
		messageIdList = append(messageIdList, dialog.GetTopMessage())
		peer := base.FromPeer(dialog.GetPeer())
		// TODO(@benqi): 先假设只有PEER_USER
		if peer.PeerType == base.PEER_USER {
			userIdList = append(userIdList, peer.PeerId)
		} else if peer.PeerType == base.PEER_SELF {
			userIdList = append(userIdList, md.UserId)
		} else if peer.PeerType == base.PEER_CHAT {
			chatIdList = append(chatIdList, peer.PeerId)
		}
		peerDialogs.Data2.Dialogs = append(peerDialogs.Data2.Dialogs, dialog.To_Dialog())
	}

	glog.Infof("messageIdList - %v", messageIdList)
	if len(messageIdList) > 0 {
		peerDialogs.SetMessages(message.GetMessagesByPeerAndMessageIdList2(md.UserId, messageIdList))
	}

	users := user.GetUsersBySelfAndIDList(md.UserId, userIdList)
	peerDialogs.SetUsers(users)
	//for _, user := range users {
	//	if user.GetId() == md.UserId {
	//		user.SetSelf(true)
	//	} else {
	//		user.SetSelf(false)
	//	}
	//	user.SetContact(true)
	//	user.SetMutualContact(true)
	//	peerDialogs.Data2.Users = append(peerDialogs.Data2.Users, user.To_User())
	//}

	if len(chatIdList) > 0 {
		peerDialogs.Data2.Chats = chat.GetChatListBySelfAndIDList(md.UserId, chatIdList)
	}

	state := update2.GetServerUpdatesState(md.AuthId, md.UserId)
	update2.UpdateAuthStateSeq(md.AuthId, state.GetPts(), 0)

	// state := update2.GetUpdatesState2(md.AuthId, md.UserId)
	peerDialogs.SetState(state.To_Updates_State())

	glog.Infof("messages.getPeerDialogs#2d9776b9 - reply: %s", logger.JsonDebugData(peerDialogs))
	return peerDialogs.To_Messages_PeerDialogs(), nil
}
