/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dataobject

type UserPrivacysDO struct {
	Id        int32  `db:"id"`
	UserId    int32  `db:"user_id"`
	KeyType   int8   `db:"key_type"`
	Rules     string `db:"rules"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
