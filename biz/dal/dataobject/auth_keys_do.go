/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dataobject

type AuthKeysDO struct {
	Id        int32  `db:"id"`
	AuthId    int64  `db:"auth_id"`
	Body      string `db:"body"`
	CreatedAt string `db:"created_at"`
	DeletedAt string `db:"deleted_at"`
}
