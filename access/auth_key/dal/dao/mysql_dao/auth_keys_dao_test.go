/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package mysql_dao

import (
	"testing"
	"github.com/airwide-code/airwide.datacenter/baselib/mysql_client"
	"github.com/airwide-code/airwide.datacenter/access/auth_key/dal/dataobject"
	"fmt"
)

func init()  {
	mysqlConfig := mysql_client.MySQLConfig{
		Name:   "immaster",
		DSN:    "root:@/nebulaim?charset=utf8",
		Active: 5,
		Idle:   2,
	}
	mysql_client.InstallMysqlClientManager([]mysql_client.MySQLConfig{mysqlConfig})
	// InstallMysqlDAOManager(mysql_client.GetMysqlClientManager())
}

func TestCheckExists(t *testing.T) {
	authKeysDAO := NewAuthKeysDAO(mysql_client.GetMysqlClient("immaster"))
	do := &dataobject.AuthKeysDO{
		AuthId: 2,
		Body:   "123",
	}

	fmt.Println(authKeysDAO.Insert(do))
}
