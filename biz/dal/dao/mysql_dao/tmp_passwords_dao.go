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

type TmpPasswordsDAO struct {
	db *sqlx.DB
}

func NewTmpPasswordsDAO(db *sqlx.DB) *TmpPasswordsDAO {
	return &TmpPasswordsDAO{db}
}

// insert into devices(auth_id, user_id, password_hash, period, tmp_password, valid_until) values (:auth_id, :user_id, :password_hash, :period, :tmp_password, :valid_until)
// TODO(@benqi): sqlmap
func (dao *TmpPasswordsDAO) Insert(do *dataobject.TmpPasswordsDO) int64 {
	var query = "insert into devices(auth_id, user_id, password_hash, period, tmp_password, valid_until) values (:auth_id, :user_id, :password_hash, :period, :tmp_password, :valid_until)"
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
