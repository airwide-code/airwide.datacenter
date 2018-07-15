/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package photo

import (
	//"testing"
	//"fmt"
	//"github.com/disintegration/imaging"
	//"bytes"
	//"github.com/airwide-code/airwide.datacenter/baselib/mysql_client"
	"github.com/airwide-code/airwide.datacenter/nbfs/biz/dal/dao"
	//"image"
	"github.com/airwide-code/airwide.datacenter/baselib/mysql_client"
	"testing"
	"github.com/airwide-code/airwide.datacenter/nbfs/biz/core/file"
	"fmt"
)

func init()  {
	// rand.Seed(time.Now().UnixNano())
	mysqlConfig1 := mysql_client.MySQLConfig{
		Name:   "immaster",
		DSN:    "root:@/nebulaim?charset=utf8",
		Active: 5,
		Idle:   2,
	}
	mysqlConfig2 := mysql_client.MySQLConfig{
		Name:   "imslave",
		DSN:    "root:@/nebulaim?charset=utf8",
		Active: 5,
		Idle:   2,
	}
	mysql_client.InstallMysqlClientManager([]mysql_client.MySQLConfig{mysqlConfig1, mysqlConfig2})
	dao.InstallMysqlDAOManager(mysql_client.GetMysqlClientManager())
}

// go test -v -run=TestUploadPhotoFile
func TestUploadPhotoFile(t *testing.T) {
	var (
		fileId = int64(986511829842923520)
		accessHash = int64(2540815227215546042)
	)

	fileData, err := file.MakeFileDataByLoad(fileId, accessHash)
	if err != nil {
		fmt.Errorf("not found <%d, %d>", fileId, accessHash)
		return
	}

	UploadPhotoFile(fileData.FileId, 1, fileData.FilePath, fileData.Ext, false)
	UploadPhotoFile(fileData.FileId, 2, fileData.FilePath, fileData.Ext, true)

	//func MakeFileDataByLoad(fileId, accessHash int64) (*fileData, error) {
	//
	//}
}
