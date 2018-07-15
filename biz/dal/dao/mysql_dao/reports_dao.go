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

type ReportsDAO struct {
	db *sqlx.DB
}

func NewReportsDAO(db *sqlx.DB) *ReportsDAO {
	return &ReportsDAO{db}
}

// insert into reports(user_id, peer_type, peer_id, reason, content) values (:user_id, :peer_type, :peer_id, :reason, :content)
// TODO(@benqi): sqlmap
func (dao *ReportsDAO) Insert(do *dataobject.ReportsDO) int64 {
	var query = "insert into reports(user_id, peer_type, peer_id, reason, content) values (:user_id, :peer_type, :peer_id, :reason, :content)"
	r, err := dao.db.NamedExec(query, do)
	if err != nil {
		errDesc := fmt.Sprintf("NamedExec in Insert(%v), error: %v", do, err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	id, err := r.LastInsertId()
	if err != nil {
		errDesc := fmt.Sprintf("LastInsertId in Insert(%v)_error: %v", do, err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}
	return id
}
