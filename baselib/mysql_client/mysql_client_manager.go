/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package mysql_client

import (
	"github.com/jmoiron/sqlx"
	"github.com/golang/glog"
	"fmt"
	"sync"
)

type MysqlClientManager struct {
	mysqlClients sync.Map
}

var mysqlClients = &MysqlClientManager{}

func InstallMysqlClientManager(configs []MySQLConfig) {
	for _, config := range configs {
		client := NewSqlxDB(&config)
		if client == nil {
			err := fmt.Errorf("InstallMysqlClientManager - NewSqlxDB {%v} error!", config)
			panic(err)
		}

		if config.Name == "" {
			err := fmt.Errorf("InstallMysqlClientManager - config error: config.Name is empty")
			panic(err)
		}
		if val, ok := mysqlClients.mysqlClients.Load(config.Name); ok {
			err := fmt.Errorf("InstallMysqlClientManager - config error: dublicated config.Name {%v}", val)
			panic(err)
		}
		mysqlClients.mysqlClients.Store(config.Name, client)
	}
}

func GetMysqlClient(dbName string) (client *sqlx.DB) {
	if val, ok := mysqlClients.mysqlClients.Load(dbName); ok {
		if client, ok = val.(*sqlx.DB); ok {
			return
		}
	}

	glog.Errorf("GetMysqlClient - Not found client: %s", dbName)
	return
}

func GetMysqlClientManager() sync.Map {
	return mysqlClients.mysqlClients
}
