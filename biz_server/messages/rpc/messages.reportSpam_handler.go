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
)

// messages.reportSpam#cf1592db peer:InputPeer = Bool;
func (s *MessagesServiceImpl) MessagesReportSpam(ctx context.Context, request *mtproto.TLMessagesReportSpam) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.reportSpam#cf1592db - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	peer := base.FromInputPeer(request.GetPeer())
	if peer.PeerType == base.PEER_USER || peer.PeerType == base.PEER_CHAT {
		// TODO(@benqi): 入库
	}

	glog.Info("messages.reportSpam#cf1592db - reply: {true}")
	return mtproto.ToBool(true), nil
}
