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

type AuthKeysDAO struct {
	db *sqlx.DB
}

func NewAuthKeysDAO(db *sqlx.DB) *AuthKeysDAO {
	return &AuthKeysDAO{db}
}

// insert into auth_keys(auth_id, body) values (:auth_id, :body)
// TODO(@benqi): sqlmap
func (dao *AuthKeysDAO) Insert(do *dataobject.AuthKeysDO) int64 {
	var query = "insert into auth_keys(auth_id, body) values (:auth_id, :body)"
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

// select body from auth_keys where auth_id = :auth_id
// TODO(@benqi): sqlmap
func (dao *AuthKeysDAO) SelectByAuthId(auth_id int64) *dataobject.AuthKeysDO {
	var query = "select body from auth_keys where auth_id = ?"
	rows, err := dao.db.Queryx(query, auth_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByAuthId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.AuthKeysDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByAuthId(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByAuthId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}
