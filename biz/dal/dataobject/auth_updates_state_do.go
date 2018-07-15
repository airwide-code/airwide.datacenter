/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dataobject

type AuthUpdatesStateDO struct {
	Id        int32  `db:"id"`
	AuthKeyId int64  `db:"auth_key_id"`
	UserId    int32  `db:"user_id"`
	Pts       int32  `db:"pts"`
	Pts2      int32  `db:"pts2"`
	Qts       int32  `db:"qts"`
	Qts2      int32  `db:"qts2"`
	Seq       int32  `db:"seq"`
	Seq2      int32  `db:"seq2"`
	Date      int32  `db:"date"`
	Date2     int32  `db:"date2"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
	DeletedAt string `db:"deleted_at"`
}
