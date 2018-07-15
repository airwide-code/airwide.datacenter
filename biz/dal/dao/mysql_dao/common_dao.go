/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
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
	"strings"
)

type CommonDAO struct {
	db *sqlx.DB
}

func NewCommonDAO(db *sqlx.DB) *CommonDAO {
	return &CommonDAO{db}
}

// 检查是否存在
// TODO(@benqi): SELECT count(id) 是否会快一点？
func (dao *CommonDAO) CheckExists(table string, params map[string]interface{}) bool {
	if len(params) == 0 {
		glog.Errorf("CheckExists - [%s] error: params empty!", table)
		return false
	}

	names := make([]string, 0, len(params))
	for k, v := range params {
		names = append(names, k+" = :"+k)
		glog.Info("k: ", k, ", v: ", v)
	}
	sql := fmt.Sprintf("SELECT 1 FROM %s WHERE %s LIMIT 1", table, strings.Join(names, " AND "))
	glog.Info("checkExists - sql: ", sql, ", params: ", params)
	rows, err := dao.db.NamedQuery(sql, params)
	if err != nil {
		glog.Errorf("CheckExists - [%s] error: %s", table, err)
		return false
	}

	defer rows.Close()

	if rows.Next() {
		return true
	} else {
		return false
	}
}

//
//func (dao *CommonDAO) InsertOrUpdate(table string, params map[string]interface{}) bool {
//	return true
//}
//
//func (dao *CommonDAO) GetOrInsert(table string, params map[string]interface{}) bool {
//	return true
//}
