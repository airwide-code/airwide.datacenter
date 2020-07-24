/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package account

import (
	"github.com/airwide-code/airwide.datacenter/biz/dal/dao"
	"github.com/airwide-code/airwide.datacenter/biz/dal/dataobject"
	"github.com/airwide-code/airwide.datacenter/biz/base"
	"github.com/airwide-code/airwide.datacenter/mtproto"
	base2 "github.com/airwide-code/airwide.datacenter/baselib/base"
)

func GetNotifySettings(userId int32, peer *base.PeerUtil) *mtproto.PeerNotifySettings {
	do := dao.GetUserNotifySettingsDAO(dao.DB_SLAVE).SelectByPeer(userId, int8(peer.PeerType), peer.PeerId)

	// var mute_until int32 = 0
	if do == nil {
		settings := &mtproto.TLPeerNotifySettings{Data2: &mtproto.PeerNotifySettings_Data{
			ShowPreviews: true,
			Silent:       false,
			MuteUntil:    0,
			Sound:        "default",
		}}
		return settings.To_PeerNotifySettings()
	} else {
		settings := &mtproto.TLPeerNotifySettings{Data2: &mtproto.PeerNotifySettings_Data{
			ShowPreviews: do.ShowPreviews == 1,
			Silent:       do.Silent == 1,
			MuteUntil:    do.MuteUntil,
			Sound:        do.Sound,
		}}
		return settings.To_PeerNotifySettings()
	}
}

func SetNotifySettings(userId int32, peer *base.PeerUtil, settings *mtproto.TLInputPeerNotifySettings) {
	subordinate := dao.GetUserNotifySettingsDAO(dao.DB_SLAVE)
	main := dao.GetUserNotifySettingsDAO(dao.DB_MASTER)

	var (
		showPreviews = base2.BoolToInt8(settings.GetShowPreviews())
		silent = base2.BoolToInt8(settings.GetSilent())
	)

	do := subordinate.SelectByPeer(userId, int8(peer.PeerType), peer.PeerId)
	if do == nil {
		do = &dataobject.UserNotifySettingsDO{
			UserId:       userId,
			PeerType:     int8(peer.PeerType),
			PeerId:       peer.PeerId,
			ShowPreviews: showPreviews,
			Silent:       silent,
			MuteUntil:    settings.GetMuteUntil(),
			Sound:        settings.GetSound(),
		}
		main.Insert(do)
	} else {
		main.UpdateByPeer(showPreviews, silent, settings.GetMuteUntil(), settings.GetSound(), 0, userId, int8(peer.PeerType), peer.PeerId)
	}
}

func ResetNotifySettings(userId int32) {
	subordinate := dao.GetUserNotifySettingsDAO(dao.DB_SLAVE)
	main := dao.GetUserNotifySettingsDAO(dao.DB_MASTER)

	main.DeleteNotAll(userId)
	do := subordinate.SelectByAll(userId)
	if do == nil {
		do = &dataobject.UserNotifySettingsDO{}
		do.UserId = userId
		do.PeerType = base.PEER_ALL
		do.PeerId = 0
		do.ShowPreviews = 1
		do.Silent = 0
		do.MuteUntil = 0
		main.Insert(do)
	} else {
		main.UpdateByPeer(1, 0, 0, "default", 0, userId, base.PEER_ALL, 0)
	}
}
