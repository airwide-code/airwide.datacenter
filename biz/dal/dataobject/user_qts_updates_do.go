/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dataobject

type UserQtsUpdatesDO struct {
	Id         int64  `db:"id"`
	UserId     int32  `db:"user_id"`
	Qts        int32  `db:"qts"`
	UpdateType int32  `db:"update_type"`
	UpdateData []byte `db:"update_data"`
	Date2      int32  `db:"date2"`
	CreatedAt  string `db:"created_at"`
}
