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
	"time"
	"github.com/airwide-code/airwide.datacenter/biz_server/sync_client"
	message2 "github.com/airwide-code/airwide.datacenter/biz/core/message"
	"github.com/airwide-code/airwide.datacenter/biz/core/user"
)

func makeUpdateEditMessageUpdates(selfUserId int32, message *mtproto.Message) *mtproto.TLUpdates {
	userIdList, _, _ := message2.PickAllIDListByMessages([]*mtproto.Message{message})
	userList := user.GetUsersBySelfAndIDList(selfUserId, userIdList)

	updateNew := &mtproto.TLUpdateEditMessage{Data2: &mtproto.Update_Data{
		Message_1: message,
	}}
	return &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
		Updates: []*mtproto.Update{updateNew.To_Update()},
		Users:   userList,
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}}
}

func setEditMessageData(request *mtproto.TLMessagesEditMessage, message *mtproto.Message) {
	// edit message data
	data2 := message.GetData2()
	if request.GetMessage() != "" {
		data2.Message = request.GetMessage()
	}
	if request.GetReplyMarkup() != nil {
		data2.ReplyMarkup = request.GetReplyMarkup()
	}
	if request.GetEntities() != nil {
		data2.Entities = request.GetEntities()
	}
	data2.EditDate = int32(time.Now().Unix())
}

// messages.editMessage#5d1b8dd flags:# no_webpage:flags.1?true stop_geo_live:flags.12?true peer:InputPeer id:int message:flags.11?string reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> geo_point:flags.13?InputGeoPoint = Updates;
func (s *MessagesServiceImpl) MessagesEditMessage(ctx context.Context, request *mtproto.TLMessagesEditMessage) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.editMessage#5d1b8dd - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// SelectDialogMessageListByMessageId
	editOutbox := message2.GetMessageByPeerAndMessageId(md.UserId, request.GetId())
	// TODO(@benqi): check invalid

	setEditMessageData(request, editOutbox)

	syncUpdates := makeUpdateEditMessageUpdates(md.UserId, editOutbox)
	state, err := sync_client.GetSyncClient().SyncUpdatesData(md.AuthId, md.SessionId, md.UserId, syncUpdates.To_Updates())
	if err != nil {
		return nil, err
	}
	SetupUpdatesState(state, syncUpdates)
	message2.SaveMessage(editOutbox, md.UserId, request.GetId())

	// push edit peer message
	peerEditMessages := message2.GetPeerDialogMessageListByMessageId(md.UserId, request.GetId())
	for i := 0; i < len(peerEditMessages.UserIds); i++ {
		editMessage := peerEditMessages.Messages[i]
		editUserId := peerEditMessages.UserIds[i]

		setEditMessageData(request, editMessage)
		editUpdates := makeUpdateEditMessageUpdates(editUserId, editMessage)
		sync_client.GetSyncClient().PushToUserUpdatesData(editUserId, editUpdates.To_Updates())
		message2.SaveMessage(editMessage, editUserId, editMessage.GetData2().GetId())
	}

	glog.Infof("messages.editMessage#5d1b8dd - reply: %s", logger.JsonDebugData(syncUpdates))
	return syncUpdates.To_Updates(), nil
}
