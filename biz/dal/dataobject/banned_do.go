/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dataobject

type BannedDO struct {
	Id           int32  `db:"id"`
	Phone        string `db:"phone"`
	BannedTime   int64  `db:"banned_time"`
	Expires      int64  `db:"expires"`
	BannedReason string `db:"banned_reason"`
	Log          string `db:"log"`
	State        int8   `db:"state"`
	CreatedAt    int32  `db:"created_at"`
	UpdatedAt    int32  `db:"updated_at"`
}
