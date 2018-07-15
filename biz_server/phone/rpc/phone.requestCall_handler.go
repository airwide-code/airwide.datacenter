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
	update2 "github.com/airwide-code/airwide.datacenter/biz/core/update"
	"github.com/airwide-code/airwide.datacenter/biz/core/user"
	"github.com/airwide-code/airwide.datacenter/biz/core/phone_call"
	"github.com/airwide-code/airwide.datacenter/biz_server/sync_client"
)

// phone.requestCall#5b95b3d4 user_id:InputUser random_id:int g_a_hash:bytes protocol:PhoneCallProtocol = phone.PhoneCall;
func (s *PhoneServiceImpl) PhoneRequestCall(ctx context.Context, request *mtproto.TLPhoneRequestCall) (*mtproto.Phone_PhoneCall, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("phone.requestCall#5b95b3d4 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		err error
		participantId int32
	)

	switch request.GetUserId().GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputUser:
		// TODO(@benqi): Check access_hash
		participantId = request.GetUserId().GetData2().GetUserId()
	default:
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error("inputUser is empty or self, err: ", err)
		return nil, err
	}

	callSession := phone_call.NewPhoneCallLogic(md.UserId, participantId, request.GetGAHash(), request.GetProtocol().To_PhoneCallProtocol())

	/////////////////////////////////////////////////////////////////////////////////
	updatesData := update2.NewUpdatesLogic(md.UserId)
	// 1. add updateUserStatus
	//var status *mtproto.UserStatus
	statusOnline := &mtproto.TLUserStatusOnline{Data2: &mtproto.UserStatus_Data{
		Expires: int32(time.Now().Unix() + 5*30),
	}}
	// status = statusOnline.To_UserStatus()
	updateUserStatus := &mtproto.TLUpdateUserStatus{Data2: &mtproto.Update_Data{
		UserId: md.UserId,
		Status: statusOnline.To_UserStatus(),
	}}
	updatesData.AddUpdate(updateUserStatus.To_Update())
	// 2. add phoneCallRequested
	updatePhoneCall := &mtproto.TLUpdatePhoneCall{Data2: &mtproto.Update_Data{
		PhoneCall: callSession.ToPhoneCallRequested().To_PhoneCall(),
	}}
	updatesData.AddUpdate(updatePhoneCall.To_Update())
	// 3. add users
	updatesData.AddUsers(user.GetUsersBySelfAndIDList(participantId, []int32{md.UserId, participantId}))
	sync_client.GetSyncClient().PushToUserUpdatesData(participantId, updatesData.ToUpdates())

	/////////////////////////////////////////////////////////////////////////////////
	// 2. reply
	phoneCall := &mtproto.TLPhonePhoneCall{Data2: &mtproto.Phone_PhoneCall_Data{
		PhoneCall: callSession.ToPhoneCallWaiting(md.UserId, 0).To_PhoneCall(),
		Users:   user.GetUsersBySelfAndIDList(md.UserId, []int32{md.UserId, participantId}),
	}}

	glog.Infof("phone.requestCall#5b95b3d4 - reply: {%v}", phoneCall)
	return phoneCall.To_Phone_PhoneCall(), nil
}
