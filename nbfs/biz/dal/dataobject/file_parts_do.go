/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dataobject

type FilePartsDO struct {
	Id             int64  `db:"id"`
	CreatorId      int64  `db:"creator_id"`
	CreatorUserId  int32  `db:"creator_user_id"`
	FileId         int64  `db:"file_id"`
	FilePartId     int64  `db:"file_part_id"`
	FilePart       int32  `db:"file_part"`
	IsBigFile      int8   `db:"is_big_file"`
	FileTotalParts int32  `db:"file_total_parts"`
	FilePath       string `db:"file_path"`
	FileSize       int64  `db:"file_size"`
	CreatedAt      string `db:"created_at"`
	UpdatedAt      string `db:"updated_at"`
}
