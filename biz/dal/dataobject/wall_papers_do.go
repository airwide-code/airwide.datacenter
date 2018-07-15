/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dataobject

type WallPapersDO struct {
	Id        int32  `db:"id"`
	Type      int8   `db:"type"`
	Title     string `db:"title"`
	Color     int32  `db:"color"`
	BgColor   int32  `db:"bg_color"`
	PhotoId   int64  `db:"photo_id"`
	CreatedAt string `db:"created_at"`
	DeletedAt int64  `db:"deleted_at"`
}
