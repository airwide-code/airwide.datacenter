/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dataobject

type ReportsDO struct {
	Id        int64  `db:"id"`
	UserId    int32  `db:"user_id"`
	PeerType  int32  `db:"peer_type"`
	PeerId    int32  `db:"peer_id"`
	Reason    int8   `db:"reason"`
	Content   string `db:"content"`
	CreatedAt string `db:"created_at"`
}
