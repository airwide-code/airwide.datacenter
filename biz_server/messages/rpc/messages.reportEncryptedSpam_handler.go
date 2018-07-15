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
)

// messages.reportEncryptedSpam#4b0c8c0f peer:InputEncryptedChat = Bool;
func (s *MessagesServiceImpl) MessagesReportEncryptedSpam(ctx context.Context, request *mtproto.TLMessagesReportEncryptedSpam) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesReportEncryptedSpam - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// peer := helper.FromInputPeer(request.GetPeer())
	//
	// if peer.PeerType == helper.PEER_USER || peer.PeerType == helper.PEER_CHAT {
	//	// TODO(@benqi): 入库
	// }

	glog.Info("MessagesReportEncryptedSpam - reply: {true}")
	return mtproto.ToBool(true), nil
}
