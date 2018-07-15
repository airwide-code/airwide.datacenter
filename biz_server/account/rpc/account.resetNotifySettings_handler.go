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
	"github.com/airwide-code/airwide.datacenter/biz/core/account"
	// "github.com/airwide-code/airwide.datacenter/biz_server/sync_client"
	// peer2 "github.com/airwide-code/airwide.datacenter/biz/core/peer"
)

// account.resetNotifySettings#db7e1747 = Bool;
func (s *AccountServiceImpl) AccountResetNotifySettings(ctx context.Context, request *mtproto.TLAccountResetNotifySettings) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.resetNotifySettings#db7e1747 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	account.ResetNotifySettings(md.UserId)

	// TODO(@benqi): update notify setting
	/*
	 Android client source:
		} else if (update instanceof TLRPC.TL_updateNotifySettings) {
			TLRPC.TL_updateNotifySettings updateNotifySettings = (TLRPC.TL_updateNotifySettings) update;
			if (update.notify_settings instanceof TLRPC.TL_peerNotifySettings && updateNotifySettings.peer instanceof TLRPC.TL_notifyPeer) {
	           ......
	        }
	    }
	 */

	//peer := &peer2.PeerUtil{}
	//peer.PeerType = peer2.PEER_ALL
	//update := mtproto.NewTLUpdateNotifySettings()
	//update.SetPeer(peer.ToNotifyPeer())
	//updateSettings := mtproto.NewTLPeerNotifySettings()
	//updateSettings.SetShowPreviews(true)
	//updateSettings.SetSilent(false)
	//updateSettings.SetMuteUntil(0)
	//updateSettings.SetSound("default")
	//update.SetNotifySettings(updateSettings.To_PeerNotifySettings())
	//
	//sync_client.GetSyncClient().PushToUserMeOneUpdateData(md.AuthId, md.SessionId, md.UserId, update.To_Update())

	glog.Infof("account.resetNotifySettings#db7e1747 - reply: {true}")
	return mtproto.ToBool(true), nil
}
