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
	"github.com/airwide-code/airwide.datacenter/biz/core/account"
)

// account.getNotifySettings#12b3ad31 peer:InputNotifyPeer = PeerNotifySettings;
func (s *AccountServiceImpl) AccountGetNotifySettings(ctx context.Context, request *mtproto.TLAccountGetNotifySettings) (*mtproto.PeerNotifySettings, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.getNotifySettings#12b3ad31 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		settings *mtproto.PeerNotifySettings
	)

	switch request.GetPeer().GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputNotifyPeer:
		peer := base.FromInputNotifyPeer(request.GetPeer())
		settings = account.GetNotifySettings(md.UserId, peer)
	case mtproto.TLConstructor_CRC32_inputNotifyUsers,
		mtproto.TLConstructor_CRC32_inputNotifyChats,
		mtproto.TLConstructor_CRC32_inputNotifyAll:

		peerSettings := &mtproto.TLPeerNotifySettings{Data2: &mtproto.PeerNotifySettings_Data{
			ShowPreviews: true,
			Silent:       false,
			MuteUntil:    0,
			Sound:        "default",
		}}
		settings = peerSettings.To_PeerNotifySettings()
	}

	glog.Infof("account.getNotifySettings#12b3ad31 - reply: %s", logger.JsonDebugData(settings))
	return settings, nil
}
