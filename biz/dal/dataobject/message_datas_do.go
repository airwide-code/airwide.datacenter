/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dataobject

type MessageDatasDO struct {
	Id           int32  `db:"id"`
	DialogId     int64  `db:"dialog_id"`
	MessageId    int64  `db:"message_id"`
	SenderUserId int32  `db:"sender_user_id"`
	PeerType     int8   `db:"peer_type"`
	PeerId       int32  `db:"peer_id"`
	RandomId     int64  `db:"random_id"`
	MessageType  int8   `db:"message_type"`
	MessageData  string `db:"message_data"`
	Date         int32  `db:"date"`
	Deleted      int8   `db:"deleted"`
	CreatedAt    string `db:"created_at"`
	UpdatedAt    string `db:"updated_at"`
}
