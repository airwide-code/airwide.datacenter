/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package dao

import (
	"github.com/airwide-code/airwide.datacenter/access/auth_key/dal/dao/mysql_dao"
	"github.com/jmoiron/sqlx"
	"github.com/golang/glog"
	"sync"
)

const (
	DB_MASTER 		= "immaster"
	DB_SLAVE 		= "imslave"
)

type MysqlDAOList struct {
	// auth_key
	AuthKeysDAO  *mysql_dao.AuthKeysDAO
	AuthUsersDAO *mysql_dao.AuthUsersDAO
}

// TODO(@benqi): 一主多从
type MysqlDAOManager struct {
	daoListMap map[string]*MysqlDAOList
}

var mysqlDAOManager = &MysqlDAOManager{make(map[string]*MysqlDAOList)}

func InstallMysqlDAOManager(clients sync.Map/*map[string]*sqlx.DB*/) {
	clients.Range(func(key, value interface{}) bool {
		k, _ := key.(string)
		v, _ := value.(*sqlx.DB)

		daoList := &MysqlDAOList{}
		// auth_key
		daoList.AuthKeysDAO = mysql_dao.NewAuthKeysDAO(v)
		daoList.AuthUsersDAO = mysql_dao.NewAuthUsersDAO(v)

		mysqlDAOManager.daoListMap[k] = daoList
		return true
	})
}

func  GetMysqlDAOListMap() map[string]*MysqlDAOList {
	return mysqlDAOManager.daoListMap
}

func  GetMysqlDAOList(dbName string) (daoList *MysqlDAOList) {
	daoList, ok := mysqlDAOManager.daoListMap[dbName]
	if !ok {
		glog.Errorf("GetMysqlDAOList - Not found daoList: %s", dbName)
	}
	return
}

func GetAuthKeysDAO(dbName string) (dao *mysql_dao.AuthKeysDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.AuthKeysDAO
	}
	return
}

func GetAuthUsersDAO(dbName string) (dao *mysql_dao.AuthUsersDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.AuthUsersDAO
	}
	return
}
