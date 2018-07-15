/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dataobject

type StickerSetsDO struct {
	Id           int32  `db:"id"`
	StickerSetId int64  `db:"sticker_set_id"`
	AccessHash   int64  `db:"access_hash"`
	Title        string `db:"title"`
	ShortName    string `db:"short_name"`
	Count        int32  `db:"count"`
	Hash         int32  `db:"hash"`
	Official     int8   `db:"official"`
	Mask         int8   `db:"mask"`
	CreatedAt    string `db:"created_at"`
	UpdatedAt    string `db:"updated_at"`
}
