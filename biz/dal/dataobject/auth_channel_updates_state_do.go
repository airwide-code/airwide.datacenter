/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dataobject

type AuthChannelUpdatesStateDO struct {
	Id        int32  `db:"id"`
	AuthKeyId int64  `db:"auth_key_id"`
	UserId    int32  `db:"user_id"`
	ChannelId int32  `db:"channel_id"`
	Pts       int32  `db:"pts"`
	Pts2      int32  `db:"pts2"`
	Date      int32  `db:"date"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
	DeletedAt string `db:"deleted_at"`
}
