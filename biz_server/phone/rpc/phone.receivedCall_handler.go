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
	"fmt"
	"github.com/airwide-code/airwide.datacenter/biz_server/sync_client"
)

// phone.receivedCall#17d54f61 peer:InputPhoneCall = Bool;
func (s *PhoneServiceImpl) PhoneReceivedCall(ctx context.Context, request *mtproto.TLPhoneReceivedCall) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("phone.receivedCall#17d54f61 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	//// TODO(@benqi): check peer
	peer := request.GetPeer().To_InputPhoneCall()

	callSession, err := phone_call.MakePhoneCallLogcByLoad(peer.GetId())
	if err != nil {
		glog.Errorf("invalid peer: {%v}, err: %v", peer, err)
		return nil, err
	}
	if peer.GetAccessHash() != callSession.ParticipantAccessHash {
		err = fmt.Errorf("invalid peer: {%v}", peer)
		glog.Errorf("invalid peer: {%v}", peer)
		return nil, err
	}

	/////////////////////////////////////////////////////////////////////////////////
	updatesData := update2.NewUpdatesLogic(md.UserId)
	// 1. add phoneCallRequested
	updatePhoneCall := &mtproto.TLUpdatePhoneCall{Data2: &mtproto.Update_Data{
		PhoneCall: callSession.ToPhoneCallWaiting(callSession.AdminId, int32(time.Now().Unix())).To_PhoneCall(),
	}}
	updatesData.AddUpdate(updatePhoneCall.To_Update())
	// 2. add users
	updatesData.AddUsers(user.GetUsersBySelfAndIDList(callSession.AdminId, []int32{md.UserId, callSession.AdminId}))
	// 3. sync
	sync_client.GetSyncClient().PushToUserUpdatesData(callSession.AdminId, updatesData.ToUpdates())

	/////////////////////////////////////////////////////////////////////////////////
	glog.Infof("phone.receivedCall#17d54f61 - reply {true}")
	return mtproto.ToBool(true), nil
}
