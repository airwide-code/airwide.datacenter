/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dataobject

type AuthUsersDO struct {
	Id            int32  `db:"id"`
	AuthId        int64  `db:"auth_id"`
	UserId        int32  `db:"user_id"`
	Hash          int64  `db:"hash"`
	DeviceModel   string `db:"device_model"`
	Platform      string `db:"platform"`
	SystemVersion string `db:"system_version"`
	ApiId         int32  `db:"api_id"`
	AppName       string `db:"app_name"`
	AppVersion    string `db:"app_version"`
	DateCreated   int32  `db:"date_created"`
	DateActive    int32  `db:"date_active"`
	Ip            string `db:"ip"`
	Country       string `db:"country"`
	Region        string `db:"region"`
	DeletedAt     int64  `db:"deleted_at"`
	CreatedAt     string `db:"created_at"`
	UpdatedAt     string `db:"updated_at"`
}
