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
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestReflectTLObject(t *testing.T) {
	mysqlDsn := "root:@/nebulaim?charset=utf8"

	db, err := sqlx.Connect("mysql", mysqlDsn)
	if err != nil {
		glog.Fatalf("Connect mysql %s error: %s", mysqlDsn, err)
		return
	}

	userDialogsDAO := NewUserDialogsDAO(db)

	vals := userDialogsDAO.SelectPinnedDialogs(1)
	fmt.Println(vals)
}
