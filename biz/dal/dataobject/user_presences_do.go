/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dataobject

type UserPresencesDO struct {
	Id                int32  `db:"id"`
	UserId            int32  `db:"user_id"`
	LastSeenAt        int64  `db:"last_seen_at"`
	LastSeenAuthKeyId int64  `db:"last_seen_auth_key_id"`
	LastSeenIp        string `db:"last_seen_ip"`
	Version           int64  `db:"version"`
	CreatedAt         string `db:"created_at"`
	UpdatedAt         string `db:"updated_at"`
}
