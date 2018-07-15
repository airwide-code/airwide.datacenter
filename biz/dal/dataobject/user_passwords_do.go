/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dataobject

type UserPasswordsDO struct {
	Id          int64  `db:"id"`
	UserId      int32  `db:"user_id"`
	ServerSalt  string `db:"server_salt"`
	Hash        string `db:"hash"`
	Salt        string `db:"salt"`
	Hint        string `db:"hint"`
	Email       string `db:"email"`
	HasRecovery int8   `db:"has_recovery"`
	Code        string `db:"code"`
	CodeExpired int32  `db:"code_expired"`
	Attempts    int32  `db:"attempts"`
	State       int8   `db:"state"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}
