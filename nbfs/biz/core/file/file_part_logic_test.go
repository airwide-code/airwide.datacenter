/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package file

import (
	"testing"
	"github.com/airwide-code/airwide.datacenter/baselib/mysql_client"
	"github.com/airwide-code/airwide.datacenter/nbfs/biz/dal/dao"
	"io/ioutil"
	"math/rand"
	"fmt"
	"time"
	"github.com/airwide-code/airwide.datacenter/nbfs/biz/core"
)

func init()  {
	rand.Seed(time.Now().UnixNano())
	mysqlConfig1 := mysql_client.MySQLConfig{
		Name:   "immain",
		DSN:    "root:@/nebulaim?charset=utf8",
		Active: 5,
		Idle:   2,
	}
	mysqlConfig2 := mysql_client.MySQLConfig{
		Name:   "imsubordinate",
		DSN:    "root:@/nebulaim?charset=utf8",
		Active: 5,
		Idle:   2,
	}
	mysql_client.InstallMysqlClientManager([]mysql_client.MySQLConfig{mysqlConfig1, mysqlConfig2})
	dao.InstallMysqlDAOManager(mysql_client.GetMysqlClientManager())
}

// go test -v -run=TestSaveFilePart
func TestSaveFilePart(t *testing.T) {
	var uploadName = "./test002.jpeg"
	buf, err := ioutil.ReadFile(uploadName)
	if err != nil {
		panic(err)
	}

	sz := len(buf) / kMaxFilePartSize

	var creatorId= rand.Int63()
	var filePartId= rand.Int63()

	var blockSize= kMaxFilePartSize
	// = sz % kMaxFilePartSize
	var uploadFileName string

	var filePart *filePartData
	for i := 0; i <= sz; i++ {
		var isNew = i == 0
		filePart, err = MakeFilePartData(creatorId, filePartId, isNew, false)
		if i == sz {
			blockSize = len(buf) % kMaxFilePartSize
			uploadFileName = filePart.FilePath
		}

		filePart.SaveFilePart(int32(i), buf[i*kMaxFilePartSize:i*kMaxFilePartSize+blockSize])
		fmt.Println(*filePart.FilePartsDO, ", file_part: ", i, ", block_size: ", blockSize)
	}

	md5, _ := core.CalcMd5File(uploadName)
	fileLogic, _ := NewFileData(filePartId, uploadFileName, uploadName, int64(len(buf)), md5)
	_ = fileLogic
}
