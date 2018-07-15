/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dataobject

type FilesDO struct {
	Id            int64  `db:"id"`
	FileId        int64  `db:"file_id"`
	AccessHash    int64  `db:"access_hash"`
	CreatorId     int64  `db:"creator_id"`
	CreatorUserId int32  `db:"creator_user_id"`
	FilePartId    int64  `db:"file_part_id"`
	FileParts     int32  `db:"file_parts"`
	FileSize      int64  `db:"file_size"`
	FilePath      string `db:"file_path"`
	Ext           string `db:"ext"`
	IsBigFile     int8   `db:"is_big_file"`
	Md5Checksum   string `db:"md5_checksum"`
	UploadName    string `db:"upload_name"`
	CreatedAt     string `db:"created_at"`
}
