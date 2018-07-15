/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dataobject

type UserPtsUpdatesDO struct {
	Id         int64  `db:"id"`
	UserId     int32  `db:"user_id"`
	Pts        int32  `db:"pts"`
	PtsCount   int32  `db:"pts_count"`
	UpdateType int8   `db:"update_type"`
	UpdateData string `db:"update_data"`
	Date2      int32  `db:"date2"`
	CreatedAt  string `db:"created_at"`
}
