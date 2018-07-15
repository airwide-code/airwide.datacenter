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

type WallPapersDAO struct {
	db *sqlx.DB
}

func NewWallPapersDAO(db *sqlx.DB) *WallPapersDAO {
	return &WallPapersDAO{db}
}

// insert into wall_papers(type, title, color, bg_color, photo_id) values (:type, :title, :color, :bg_color, :photo_id)
// TODO(@benqi): sqlmap
func (dao *WallPapersDAO) Insert(do *dataobject.WallPapersDO) int64 {
	var query = "insert into wall_papers(type, title, color, bg_color, photo_id) values (:type, :title, :color, :bg_color, :photo_id)"
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

// select id, type, title, color, bg_color, photo_id from wall_papers where deleted_at = 0
// TODO(@benqi): sqlmap
func (dao *WallPapersDAO) SelectAll() []dataobject.WallPapersDO {
	var query = "select id, type, title, color, bg_color, photo_id from wall_papers where deleted_at = 0"
	rows, err := dao.db.Queryx(query)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectAll(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.WallPapersDO
	for rows.Next() {
		v := dataobject.WallPapersDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectAll(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectAll(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}
