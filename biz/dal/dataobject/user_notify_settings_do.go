/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dataobject

type UserNotifySettingsDO struct {
	Id           int32  `db:"id"`
	UserId       int32  `db:"user_id"`
	PeerType     int8   `db:"peer_type"`
	PeerId       int32  `db:"peer_id"`
	ShowPreviews int8   `db:"show_previews"`
	Silent       int8   `db:"silent"`
	MuteUntil    int32  `db:"mute_until"`
	Sound        string `db:"sound"`
	IsDeleted    int8   `db:"is_deleted"`
	CreatedAt    string `db:"created_at"`
	UpdatedAt    string `db:"updated_at"`
}
