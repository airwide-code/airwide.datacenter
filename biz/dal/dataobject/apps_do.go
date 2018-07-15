/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dataobject

type AppsDO struct {
	Id        int32  `db:"id"`
	ApiId     int32  `db:"api_id"`
	ApiHash   string `db:"api_hash"`
	Title     string `db:"title"`
	ShortName string `db:"short_name"`
	CreatedAt string `db:"created_at"`
	DeletedAt string `db:"deleted_at"`
}
