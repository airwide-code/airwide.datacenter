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

type wallPaperDataList []dataobject.WallPapersDO

func GetWallPaperList() wallPaperDataList {
	return dao.GetWallPapersDAO(dao.DB_SLAVE).SelectAll()
}
