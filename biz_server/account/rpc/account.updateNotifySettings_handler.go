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
	"github.com/airwide-code/airwide.datacenter/biz_server/sync_client"
	"github.com/airwide-code/airwide.datacenter/biz/core/account"
	updates2 "github.com/airwide-code/airwide.datacenter/biz/core/update"
	"github.com/airwide-code/airwide.datacenter/biz/core/user"
)

/*
 Android client's account.updateNotifySettings#84be5b93 source code:

	TLRPC.TL_account_updateNotifySettings req = new TLRPC.TL_account_updateNotifySettings();
	req.settings = new TLRPC.TL_inputPeerNotifySettings();
	req.settings.sound = "default";
	int mute_type = preferences.getInt("notify2_" + dialog_id, 0);
	if (mute_type == 3) {
		req.settings.mute_until = preferences.getInt("notifyuntil_" + dialog_id, 0);
	} else {
		req.settings.mute_until = mute_type != 2 ? 0 : Integer.MAX_VALUE;
	}
	req.settings.show_previews = preferences.getBoolean("preview_" + dialog_id, true);
	req.settings.silent = preferences.getBoolean("silent_" + dialog_id, false);
	req.peer = new TLRPC.TL_inputNotifyPeer();
	((TLRPC.TL_inputNotifyPeer) req.peer).peer = MessagesController.getInputPeer((int) dialog_id);
 */

// account.updateNotifySettings#84be5b93 peer:InputNotifyPeer settings:InputPeerNotifySettings = Bool;
func (s *AccountServiceImpl) AccountUpdateNotifySettings(ctx context.Context, request *mtproto.TLAccountUpdateNotifySettings) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.updateNotifySettings#84be5b93 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): by android client source code, we only impl inputNotifyPeer - (inputNotifyPeer#b8bc5b0c peer:InputPeer = InputNotifyPeer)

	// peer := request.GetPeer()
	if request.GetPeer().GetConstructor() != mtproto.TLConstructor_CRC32_inputNotifyPeer {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error(err, ": peer only is inputNotifyPeer")
		return nil, err
	}

	peer := base.FromInputPeer(request.GetPeer().GetData2().GetPeer())
	settings := request.GetSettings().To_InputPeerNotifySettings()

	account.SetNotifySettings(md.UserId, peer, settings)

	// sync
	updateNotifySettings := &mtproto.TLUpdateNotifySettings{Data2: &mtproto.Update_Data{
		Peer_28:        peer.ToNotifyPeer(),
		NotifySettings: &mtproto.PeerNotifySettings{
			Constructor: mtproto.TLConstructor_CRC32_peerNotifySettings,
			Data2: &mtproto.PeerNotifySettings_Data{
				ShowPreviews: settings.GetShowPreviews(),
				Silent:       settings.GetSilent(),
				MuteUntil:    settings.GetMuteUntil(),
				Sound:        settings.GetSound(),
			},
		},
	}}
	notifySettingUpdates := updates2.NewUpdatesLogic(md.UserId)
	notifySettingUpdates.AddUpdate(updateNotifySettings.To_Update())

	switch peer.PeerType {
	case base.PEER_SELF:
		// TODO(@benqi): PeerUtil - PEER_SELF
		notifySettingUpdates.AddUser(user.GetUserById(md.UserId, peer.PeerId).To_User())
	case base.PEER_USER:
		notifySettingUpdates.AddUser(user.GetUserById(md.UserId, peer.PeerId).To_User())
	case base.PEER_CHAT:
		// TODO(@benqi): impl
	case base.PEER_CHANNEL:
		// TODO(@benqi): impl
	}
	sync_client.GetSyncClient().SyncUpdatesData(md.AuthId, md.SessionId, md.UserId, notifySettingUpdates.ToUpdates())

	glog.Infof("account.updateNotifySettings#84be5b93 - reply: {trur}")
	return mtproto.ToBool(true), nil
}
