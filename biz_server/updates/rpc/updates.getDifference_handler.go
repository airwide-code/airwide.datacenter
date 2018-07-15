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
	// base2 "github.com/airwide-code/airwide.datacenter/baselib/helper"
	update2 "github.com/airwide-code/airwide.datacenter/biz/core/update"
	"github.com/airwide-code/airwide.datacenter/biz/core/user"
	"github.com/airwide-code/airwide.datacenter/biz/core/message"
	"github.com/airwide-code/airwide.datacenter/biz/core/chat"
)

// updates.getDifference#25939651 flags:# pts:int pts_total_limit:flags.0?int date:int qts:int = updates.Difference;
func (s *UpdatesServiceImpl) UpdatesGetDifference(ctx context.Context, request *mtproto.TLUpdatesGetDifference) (*mtproto.Updates_Difference, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("updates.getDifference#25939651 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		lastPts = request.GetPts()
		otherUpdates []*mtproto.Update
		messages []*mtproto.Message
		userList []*mtproto.User
		chatList []*mtproto.Chat
	)

	updateList := update2.GetUpdateListByGtPts(md.UserId, lastPts)

	for _, update := range updateList {
		switch update.GetConstructor() {
		case mtproto.TLConstructor_CRC32_updateNewMessage:
			newMessage := update.To_UpdateNewMessage()
			messages = append(messages, newMessage.GetMessage())
			// otherUpdates = append(otherUpdates, update)

		case mtproto.TLConstructor_CRC32_updateReadHistoryOutbox:
			readHistoryOutbox := update.To_UpdateReadHistoryOutbox()
			readHistoryOutbox.SetPtsCount(0)
			otherUpdates = append(otherUpdates, readHistoryOutbox.To_Update())
		case mtproto.TLConstructor_CRC32_updateReadHistoryInbox:
			readHistoryInbox := update.To_UpdateReadHistoryInbox()
			readHistoryInbox.SetPtsCount(0)
			otherUpdates = append(otherUpdates, readHistoryInbox.To_Update())
		default:
			continue
		}
		if update.Data2.GetPts() > lastPts {
			lastPts = update.Data2.GetPts()
		}
	}

	//otherUpdates, boxIDList, lastPts := model.GetUpdatesModel().GetUpdatesByGtPts(md.UserId, request.GetPts())
	//messages := model.GetMessageModel().GetMessagesByPeerAndMessageIdList2(md.UserId, boxIDList)
	userIdList, chatIdList, _ := message.PickAllIDListByMessages(messages)
	userList = user.GetUsersBySelfAndIDList(md.UserId, userIdList)
	chatList = chat.GetChatListBySelfAndIDList(md.UserId, chatIdList)

	state := &mtproto.TLUpdatesState{Data2: &mtproto.Updates_State_Data{
		Pts:         lastPts,
		Date:        int32(time.Now().Unix()),
		UnreadCount: 0,
		// Seq:         int32(model.GetSequenceModel().CurrentSeqId(base2.Int32ToString(md.UserId))),
		Seq:         0,
	}}
	difference := &mtproto.TLUpdatesDifference{Data2: &mtproto.Updates_Difference_Data{
		NewMessages:  messages,
		OtherUpdates: otherUpdates,
		Users:        userList,
		Chats:        chatList,
		State:        state.To_Updates_State(),
	}}

	// TODO(@benqi): remove to received ack handler.
	update2.UpdateAuthStateSeq(md.AuthId, lastPts, 0)

	glog.Infof("updates.getDifference#25939651 - reply: %s", logger.JsonDebugData(difference))
	return difference.To_Updates_Difference(), nil
}
