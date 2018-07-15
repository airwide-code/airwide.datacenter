/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package account

import (
	"github.com/airwide-code/airwide.datacenter/biz/dal/dataobject"
	"github.com/airwide-code/airwide.datacenter/biz/dal/dao"
)

func InsertReportData(userId, peerType, peerId, reason int32, text string) bool {
	do := &dataobject.ReportsDO{
		UserId: userId,
		PeerType: peerType,
		PeerId: peerId,
		Reason: int8(reason),
		Content: text,
	}
	do.Id = dao.GetReportsDAO(dao.DB_MASTER).Insert(do)
	return do.Id > 0
}
