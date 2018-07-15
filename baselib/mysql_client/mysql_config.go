/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package mysql_client

import (
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/jmoiron/sqlx"
	"github.com/golang/glog"
)

type MySQLConfig struct {
	Name   string // for trace
	DSN    string // data source name
	Active int    // pool
	Idle   int    // pool
}

func NewSqlxDB(c* MySQLConfig) (db *sqlx.DB) {
	db, err := sqlx.Connect("mysql", c.DSN)
	if err != nil {
		glog.Errorf("Connect db error: %s", err)
	}

	db.SetMaxOpenConns(c.Active)
	db.SetMaxIdleConns(c.Idle)
	return
}
