/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package dao

import (
	"github.com/airwide-code/airwide.datacenter/nbfs/biz/dal/dao/mysql_dao"
	"github.com/jmoiron/sqlx"
	"github.com/golang/glog"
	"sync"
)

const (
	DB_MASTER 		= "immain"
	DB_SLAVE 		= "imsubordinate"
)

type MysqlDAOList struct {
	FilePartsDAO             *mysql_dao.FilePartsDAO
	FilesDAO                 *mysql_dao.FilesDAO
	PhotoDatasDAO            *mysql_dao.PhotoDatasDAO
	DocumentsDAO			 *mysql_dao.DocumentsDAO
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

		daoList.FilePartsDAO = mysql_dao.NewFilePartsDAO(v)
		daoList.FilesDAO = mysql_dao.NewFilesDAO(v)
		daoList.PhotoDatasDAO = mysql_dao.NewPhotoDatasDAO(v)
		daoList.DocumentsDAO = mysql_dao.NewDocumentsDAO(v)

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

func GetFilePartsDAO(dbName string) (dao *mysql_dao.FilePartsDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.FilePartsDAO
	}
	return
}

func GetFilesDAO(dbName string) (dao *mysql_dao.FilesDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.FilesDAO
	}
	return
}

func GetPhotoDatasDAO(dbName string) (dao *mysql_dao.PhotoDatasDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.PhotoDatasDAO
	}
	return
}

func GetDocumentsDAO(dbName string) (dao *mysql_dao.DocumentsDAO) {
	daoList := GetMysqlDAOList(dbName)
	// err := mysqlDAOManager.daoListMap[dbName]
	if daoList != nil {
		dao = daoList.DocumentsDAO
	}
	return
}
