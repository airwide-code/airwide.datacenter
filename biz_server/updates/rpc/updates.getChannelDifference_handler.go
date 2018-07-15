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
	update2 "github.com/airwide-code/airwide.datacenter/biz/core/update"
	"github.com/airwide-code/airwide.datacenter/baselib/base"
)

// updates.getChannelDifference#3173d78 flags:# force:flags.0?true channel:InputChannel filter:ChannelMessagesFilter pts:int limit:int = updates.ChannelDifference;
func (s *UpdatesServiceImpl) UpdatesGetChannelDifference(ctx context.Context, request *mtproto.TLUpdatesGetChannelDifference) (*mtproto.Updates_ChannelDifference, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("updates.getChannelDifference#3173d78 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		lastPts = request.GetPts()
		//otherUpdates []*mtproto.Update
		messages []*mtproto.Message
		//userList []*mtproto.User
		//chatList []*mtproto.Chat
	)

	updateList := update2.GetChannelUpdateListByGtPts(request.GetChannel().GetData2().GetChannelId(), lastPts)

	for _, update := range updateList {
		switch update.GetConstructor() {
		case mtproto.TLConstructor_CRC32_updateNewChannelMessage:
			newMessage := update.To_UpdateNewChannelMessage()
			messages = append(messages, newMessage.GetMessage())
			// otherUpdates = append(otherUpdates, update)

		case mtproto.TLConstructor_CRC32_updateDeleteChannelMessages:
			// readHistoryOutbox := update.To_UpdateReadHistoryOutbox()
			// readHistoryOutbox.SetPtsCount(0)
			// otherUpdates = append(otherUpdates, readHistoryOutbox.To_Update())
		case mtproto.TLConstructor_CRC32_updateEditChannelMessage:
			// readHistoryInbox := update.To_UpdateReadHistoryInbox()
			// readHistoryInbox.SetPtsCount(0)
			// otherUpdates = append(otherUpdates, readHistoryInbox.To_Update())
		case mtproto.TLConstructor_CRC32_updateChannelWebPage:
		default:
			continue
		}
		if update.Data2.GetPts() > lastPts {
			lastPts = update.Data2.GetPts()
		}
	}

	//otherUpdates, boxIDList, lastPts := model.GetUpdatesModel().GetUpdatesByGtPts(md.UserId, request.GetPts())
	//messages := model.GetMessageModel().GetMessagesByPeerAndMessageIdList2(md.UserId, boxIDList)
	// userIdList, chatIdList, _ := message.PickAllIDListByMessages(messages)
	// userList = user.GetUsersBySelfAndIDList(md.UserId, userIdList)
	// chatList = chat.GetChatListBySelfAndIDList(md.UserId, chatIdList)
	//
	//state := &mtproto.TLUpdatesState{Data2: &mtproto.Updates_State_Data{
	//	Pts:         lastPts,
	//	Date:        int32(time.Now().Unix()),
	//	UnreadCount: 0,
	//	// Seq:         int32(model.GetSequenceModel().CurrentSeqId(base2.Int32ToString(md.UserId))),
	//	Seq:         0,
	//}}

	// updates.channelDifference#2064674e flags:# final:flags.0?true pts:int timeout:flags.1?int new_messages:Vector<Message> other_updates:Vector<Update> chats:Vector<Chat> users:Vector<User> = updates.ChannelDifference;
	var difference *mtproto.Updates_ChannelDifference

	//if len(updateList) == 0 {
		difference = &mtproto.Updates_ChannelDifference{
			Constructor: mtproto.TLConstructor_CRC32_updates_channelDifferenceEmpty,
			Data2:  &mtproto.Updates_ChannelDifference_Data{
				Final:   true,
				Pts:     int32(update2.CurrentChannelPtsId(base.Int32ToString(request.GetChannel().GetData2().GetChannelId()))),
				Timeout: 30,
			},
		}
	//} else {
	//	difference = &mtproto.Updates_ChannelDifference{
	//		Constructor: mtproto.TLConstructor_CRC32_updates_channelDifferenceEmpty,
	//		Data2:  &mtproto.Updates_ChannelDifference_Data{
	//			Final:   true,
	//			Pts:     2,
	//			Timeout: 3,
	//		},
	//	}
	//	//difference := &mtproto.TLUpdatesChannelDifference{Data2: &mtproto.Updates_ChannelDifference_Data{
	//	//	Pts: lastPts,
	//	//	Timeout: 3,
	//	//	NewMessages:  messages,
	//	//	OtherUpdates: otherUpdates,
	//	//	Users:        userList,
	//	//	Chats:        chatList,
	//	//	// State:        state.To_Updates_State(),
	//	//}}
	//	//
	//	// TODO(@benqi): remove to received ack handler.
	//	// update2.UpdateAuthStateSeq(md.AuthId, lastPts, 0)
	//}

	glog.Infof("updates.getChannelDifference#3173d78 - reply: %s", logger.JsonDebugData(difference))
	return difference, nil

}
