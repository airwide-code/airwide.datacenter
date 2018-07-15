/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package mysql_dao

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	"github.com/airwide-code/airwide.datacenter/biz/dal/dataobject"
	"github.com/airwide-code/airwide.datacenter/mtproto"
)

type BannedDAO struct {
	db *sqlx.DB
}

func NewBannedDAO(db *sqlx.DB) *BannedDAO {
	return &BannedDAO{db}
}

// select id from banned where phone = :phone
// TODO(@benqi): sqlmap
func (dao *BannedDAO) CheckBannedByPhone(phone string) *dataobject.BannedDO {
	var query = "select id from banned where phone = ?"
	rows, err := dao.db.Queryx(query, phone)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in CheckBannedByPhone(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.BannedDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in CheckBannedByPhone(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in CheckBannedByPhone(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}
