/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dataobject

type StickerPacksDO struct {
	Id           int32  `db:"id"`
	StickerSetId int64  `db:"sticker_set_id"`
	Emoticon     string `db:"emoticon"`
	DocumentId   int64  `db:"document_id"`
	CreatedAt    string `db:"created_at"`
}
