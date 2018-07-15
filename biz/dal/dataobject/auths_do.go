/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dataobject

type AuthsDO struct {
	Id             int32  `db:"id"`
	AuthId         int64  `db:"auth_id"`
	ApiId          int32  `db:"api_id"`
	DeviceModel    string `db:"device_model"`
	SystemVersion  string `db:"system_version"`
	AppVersion     string `db:"app_version"`
	SystemLangCode string `db:"system_lang_code"`
	LangPack       string `db:"lang_pack"`
	LangCode       string `db:"lang_code"`
	ConnectionHash int64  `db:"connection_hash"`
	CreatedAt      string `db:"created_at"`
	UpdatedAt      string `db:"updated_at"`
	DeletedAt      string `db:"deleted_at"`
}
