/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dataobject

type DevicesDO struct {
	Id        int64  `db:"id"`
	AuthKeyId int64  `db:"auth_key_id"`
	UserId    int32  `db:"user_id"`
	TokenType int8   `db:"token_type"`
	Token     string `db:"token"`
	State     int8   `db:"state"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
