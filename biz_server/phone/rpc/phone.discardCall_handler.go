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
	"github.com/airwide-code/airwide.datacenter/biz/core/user"
	update2 "github.com/airwide-code/airwide.datacenter/biz/core/update"
	"github.com/airwide-code/airwide.datacenter/biz/core/phone_call"
	"github.com/airwide-code/airwide.datacenter/biz_server/sync_client"
	"github.com/airwide-code/airwide.datacenter/biz/base"
	"time"
	message2 "github.com/airwide-code/airwide.datacenter/biz/core/message"
)

// phone.discardCall#78d413a6 peer:InputPhoneCall duration:int reason:PhoneCallDiscardReason connection_id:long = Updates;
func (s *PhoneServiceImpl) PhoneDiscardCall(ctx context.Context, request *mtproto.TLPhoneDiscardCall) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("phone.discardCall#78d413a6 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	//// TODO(@benqi): check peer
	peer := request.GetPeer().To_InputPhoneCall()

	callSession, err := phone_call.MakePhoneCallLogcByLoad(peer.GetId())
	if err != nil {
		glog.Errorf("invalid peer: {%v}, err: %v", peer, err)
		return nil, err
	}

	phoneCallDiscarded := &mtproto.TLPhoneCallDiscarded{Data2: &mtproto.PhoneCall_Data{
		Id: callSession.Id,
		NeedDebug: true,
		Reason: request.GetReason(),
		Duration: request.GetDuration(),
	}}

	// var toId int32
	// = md.UserId
	// if md.UserId == callSession.AdminId {
	//	toId = callSession.ParticipantId
	// } else {
	//	toId = callSession.AdminId
	// }

	// glog.Info("toId: ", toId)

	/////////////////////////////////////////////////////////////////////////////////
	// updatesData := update2.NewUpdatesLogic(md.UserId)
	adminUpdatesData := update2.NewUpdatesLogic(callSession.AdminId)
	participantUpdatesData := update2.NewUpdatesLogic(callSession.ParticipantId)

	// 1. add phoneCallRequested
	updatePhoneCall := &mtproto.TLUpdatePhoneCall{Data2: &mtproto.Update_Data{
		PhoneCall: phoneCallDiscarded.To_PhoneCall(),
	}}
	adminUpdatesData.AddUpdate(updatePhoneCall.To_Update())
	participantUpdatesData.AddUpdate(updatePhoneCall.To_Update())

	// add message service
	action := &mtproto.TLMessageActionPhoneCall{Data2: &mtproto.MessageAction_Data{
		CallId:   callSession.Id,
		Reason:   request.GetReason(),
		Duration: request.GetDuration(),
	}}
	peer2 := &base.PeerUtil{
		PeerType: base.PEER_USER,
		PeerId:   callSession.ParticipantId,
	}
	message := &mtproto.TLMessageService{Data2: &mtproto.Message_Data{
		Out:    true,
		Date:   int32(time.Now().Unix()),
		FromId: callSession.AdminId,
		ToId:   peer2.ToPeer(),
		Action: action.To_MessageAction(),
	}}
	randomId := base.NextSnowflakeId()
	outbox := message2.CreateMessageOutboxByNew(callSession.AdminId, peer2, randomId, message.To_Message(), func(messageId int32) {
		user.CreateOrUpdateByOutbox(callSession.AdminId, peer2.PeerType, peer2.PeerId, messageId, false, false)
	})
	inboxList, _ := outbox.InsertMessageToInbox(callSession.AdminId, peer2, func(inBoxUserId, messageId int32) {
		user.CreateOrUpdateByInbox(inBoxUserId, base.PEER_USER, peer2.PeerId, messageId, false)
	})

	adminUpdatesData.AddUpdateNewMessage(outbox.Message)
	participantUpdatesData.AddUpdateNewMessage(inboxList[0].Message)

	// 2. add users
	adminUpdatesData.AddUsers(user.GetUsersBySelfAndIDList(callSession.AdminId, []int32{callSession.AdminId, callSession.ParticipantId}))
	participantUpdatesData.AddUsers(user.GetUsersBySelfAndIDList(callSession.ParticipantId, []int32{callSession.AdminId, callSession.ParticipantId}))

	// TODO(@benqi): Add updateReadHistoryInbox
	// 3. sync
	//sync_client.GetSyncClient().PushToUserUpdatesData(toId, updatesData.ToUpdates())
	sync_client.GetSyncClient().PushToUserUpdatesData(callSession.AdminId, adminUpdatesData.ToUpdates())
	sync_client.GetSyncClient().PushToUserUpdatesData(callSession.ParticipantId, participantUpdatesData.ToUpdates())

	/////////////////////////////////////////////////////////////////////////////////
	replyUpdatesData := update2.NewUpdatesLogic(md.UserId)
	replyUpdatesData.AddUpdate(updatePhoneCall.To_Update())
	//
	//if md.UserId == callSession.AdminId {
	//	replyUpdatesData.AddUpdateNewMessage(outbox.Message)
	//} else {
	//	replyUpdatesData.AddUpdateNewMessage(inboxList[0].Message)
	//}
	// 2. add users
	replyUpdatesData.AddUsers(user.GetUsersBySelfAndIDList(md.UserId, []int32{callSession.AdminId, callSession.ParticipantId}))

	glog.Infof("phone.discardCall#78d413a6 - reply {%s}", logger.JsonDebugData(replyUpdatesData))
	return replyUpdatesData.ToUpdates(), nil
}
