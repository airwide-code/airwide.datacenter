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
	//"fmt"
	"github.com/airwide-code/airwide.datacenter/biz_server/sync_client"
)

// phone.confirmCall#2efe1722 peer:InputPhoneCall g_a:bytes key_fingerprint:long protocol:PhoneCallProtocol = phone.PhoneCall;
func (s *PhoneServiceImpl) PhoneConfirmCall(ctx context.Context, request *mtproto.TLPhoneConfirmCall) (*mtproto.Phone_PhoneCall, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("phone.confirmCall#2efe1722 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	//// TODO(@benqi): check peer
	peer := request.GetPeer().To_InputPhoneCall()

	callSession, err := phone_call.MakePhoneCallLogcByLoad(peer.GetId())
	if err != nil {
		glog.Errorf("invalid peer: {%v}, err: %v", peer, err)
		return nil, err
	}
	// if peer.GetAccessHash() != callSession.AdminAccessHash {
	// 	err = fmt.Errorf("invalid peer: {%v}", peer)
	// 	glog.Errorf("invalid peer: {%v}", peer)
	// 	return nil, err
	// }

	// TODO(@benqi): callSession.SetGA() ???
	callSession.GA = request.GetGA()

	/////////////////////////////////////////////////////////////////////////////////
	updatesData := update2.NewUpdatesLogic(md.UserId)
	// 1. add phoneCallRequested
	updatePhoneCall := &mtproto.TLUpdatePhoneCall{Data2: &mtproto.Update_Data{
		PhoneCall: callSession.ToPhoneCall(callSession.ParticipantId, request.GetKeyFingerprint()).To_PhoneCall(),
	}}
	updatesData.AddUpdate(updatePhoneCall.To_Update())
	// 2. add users
	updatesData.AddUsers(user.GetUsersBySelfAndIDList(callSession.ParticipantId, []int32{md.UserId, callSession.ParticipantId}))
	// 3. sync
	sync_client.GetSyncClient().PushToUserUpdatesData(callSession.ParticipantId, updatesData.ToUpdates())

	/////////////////////////////////////////////////////////////////////////////////
	// 2. reply
	phoneCall := &mtproto.TLPhonePhoneCall{Data2: &mtproto.Phone_PhoneCall_Data{
		PhoneCall: callSession.ToPhoneCall(md.UserId, request.GetKeyFingerprint()).To_PhoneCall(),
		Users:   user.GetUsersBySelfAndIDList(md.UserId, []int32{md.UserId, callSession.ParticipantId}),
	}}

	glog.Infof("phone.confirmCall#2efe1722 - reply: {%v}", phoneCall)
	return phoneCall.To_Phone_PhoneCall(), nil
}
