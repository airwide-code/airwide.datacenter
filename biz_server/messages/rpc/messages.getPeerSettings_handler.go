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

// messages.getPeerSettings#3672e09c peer:InputPeer = PeerSettings;
func (s *MessagesServiceImpl) MessagesGetPeerSettings(ctx context.Context, request *mtproto.TLMessagesGetPeerSettings) (*mtproto.PeerSettings, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getPeerSettings#3672e09c - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	peerSettings := &mtproto.TLPeerSettings{Data2: &mtproto.PeerSettings_Data{
		ReportSpam: false,
	}}

	glog.Infof("messages.getPeerSettings#3672e09c - reply: %s", logger.JsonDebugData(peerSettings))
	return peerSettings.To_PeerSettings(), nil
}
