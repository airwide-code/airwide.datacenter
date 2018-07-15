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
	"github.com/airwide-code/airwide.datacenter/biz/core/phone_call"
	"fmt"
)

// phone.saveCallDebug#277add7e peer:InputPhoneCall debug:DataJSON = Bool;
func (s *PhoneServiceImpl) PhoneSaveCallDebug(ctx context.Context, request *mtproto.TLPhoneSaveCallDebug) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("phone.saveCallDebug#277add7e - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	//// TODO(@benqi): check peer
	peer := request.GetPeer().To_InputPhoneCall()

	callSession, err := phone_call.MakePhoneCallLogcByLoad(peer.GetId())
	if err != nil {
		glog.Errorf("invalid peer: {%v}, err: %v", peer, err)
		return nil, err
	}

	if md.UserId == callSession.AdminId {
		if peer.GetAccessHash() != callSession.AdminAccessHash {
			err = fmt.Errorf("invalid peer: {%v}", peer)
			glog.Errorf("invalid peer: {%v}", peer)
			return nil, err
		}

		callSession.SetAdminDebugData(request.GetDebug().GetData2().GetData())
	} else {
		if peer.GetAccessHash() != callSession.ParticipantAccessHash {
			err = fmt.Errorf("invalid peer: {%v}", peer)
			glog.Errorf("invalid peer: {%v}", peer)
			return nil, err
		}

		callSession.SetParticipantDebugData(request.GetDebug().GetData2().GetData())
	}

	glog.Infof("phone.saveCallDebug#277add7e - reply: {true}")
	return mtproto.ToBool(true), nil
}
