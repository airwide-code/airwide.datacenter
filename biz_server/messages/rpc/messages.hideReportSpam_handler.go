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

// messages.hideReportSpam#a8f1709b peer:InputPeer = Bool;
func (s *MessagesServiceImpl) MessagesHideReportSpam(ctx context.Context, request *mtproto.TLMessagesHideReportSpam) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.hideReportSpam#a8f1709b - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	peer := base.FromInputPeer(request.GetPeer())
	if peer.PeerType == base.PEER_USER || peer.PeerType == base.PEER_CHAT {
		// TODO(@benqi): 入库
	}

	glog.Info("messages.hideReportSpam#a8f1709b - reply: {true}")
	return mtproto.ToBool(true), nil
}
