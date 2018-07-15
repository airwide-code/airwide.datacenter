/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dataobject

type UsersDO struct {
	Id             int32  `db:"id"`
	AccessHash     int64  `db:"access_hash"`
	FirstName      string `db:"first_name"`
	LastName       string `db:"last_name"`
	Username       string `db:"username"`
	Phone          string `db:"phone"`
	CountryCode    string `db:"country_code"`
	Bio            string `db:"bio"`
	About          string `db:"about"`
	State          int32  `db:"state"`
	IsBot          int8   `db:"is_bot"`
	Banned         int64  `db:"banned"`
	BannedReason   string `db:"banned_reason"`
	AccountDaysTtl int32  `db:"account_days_ttl"`
	Photos         string `db:"photos"`
	Deleted        int8   `db:"deleted"`
	DeletedReason  string `db:"deleted_reason"`
	CreatedAt      string `db:"created_at"`
	UpdatedAt      string `db:"updated_at"`
	BannedAt       string `db:"banned_at"`
	DeletedAt      string `db:"deleted_at"`
}
