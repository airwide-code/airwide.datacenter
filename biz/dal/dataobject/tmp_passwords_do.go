/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dataobject

type TmpPasswordsDO struct {
	Id           int32  `db:"id"`
	AuthId       int64  `db:"auth_id"`
	UserId       int32  `db:"user_id"`
	PasswordHash string `db:"password_hash"`
	Period       int32  `db:"period"`
	TmpPassword  string `db:"tmp_password"`
	ValidUntil   int32  `db:"valid_until"`
	CreatedAt    string `db:"created_at"`
}
