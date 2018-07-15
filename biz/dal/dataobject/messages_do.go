/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dataobject

type MessagesDO struct {
	Id               int32  `db:"id"`
	UserId           int32  `db:"user_id"`
	UserMessageBoxId int32  `db:"user_message_box_id"`
	DialogMessageId  int64  `db:"dialog_message_id"`
	SenderUserId     int32  `db:"sender_user_id"`
	MessageBoxType   int8   `db:"message_box_type"`
	PeerType         int8   `db:"peer_type"`
	PeerId           int32  `db:"peer_id"`
	RandomId         int64  `db:"random_id"`
	MessageType      int8   `db:"message_type"`
	MessageData      string `db:"message_data"`
	Date2            int32  `db:"date2"`
	Deleted          int8   `db:"deleted"`
	CreatedAt        string `db:"created_at"`
	UpdatedAt        string `db:"updated_at"`
}
