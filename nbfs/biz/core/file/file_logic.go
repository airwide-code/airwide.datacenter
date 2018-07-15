/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package file

import (
	"github.com/airwide-code/airwide.datacenter/nbfs/biz/dal/dataobject"
	"fmt"
	"github.com/airwide-code/airwide.datacenter/nbfs/biz/base"
	"math/rand"
	"github.com/airwide-code/airwide.datacenter/nbfs/biz/dal/dao"
	"path"
	"os"
	"github.com/airwide-code/airwide.datacenter/nbfs/biz/core"
	"strings"
	// "github.com/golang/glog"
)


// inputFile#f52ff27f id:long parts:int name:string md5_checksum:string = InputFile;
// inputFileBig#fa4f0bb5 id:long parts:int name:string = InputFile;
type fileData struct {
	*dataobject.FilesDO
}

// TODO(@benqi): 是否要加去重字段？？
func NewFileData(filePartId int64, filePath, uploadName string, fileSize int64, md5Checksum string) (*fileData, error) {
	var fileId = base.NextSnowflakeId()
	var ext = path.Ext(uploadName)
	ext = strings.ToLower(ext)
	data2 := &dataobject.FilesDO{
		FileId:      fileId,
		AccessHash:  int64(rand.Uint64()),
		FilePartId:  filePartId,
		FileSize:    fileSize,
		FilePath:    fmt.Sprintf("/0/%d%s", fileId, ext),
		Ext:         ext,
		Md5Checksum: md5Checksum,
		UploadName:  uploadName,
	}

	//var oldpath = core.NBFS_DATA_PATH + filePath
	//var onewpath = core.NBFS_DATA_PATH + data2.FilePath
	//
	//f, err := os.Create(onewpath)
	//if err != nil {
	//	glog.Error(err)
	//	return nil, err
	//}
	//defer f.Close()
	//
	//for i := 0; i <
	//_, err = f.Write(bytes)
	//if err != nil {
	//	glog.Error(err)
	//	return err
	//}
	//f.Sync()

	// var oldpath = core.NBFS_DATA_PATH + filePath
	// var newpath =
	// os.Rename(core.NBFS_DATA_PATH + filePath)
	data2.Id = dao.GetFilesDAO(dao.DB_MASTER).Insert(data2)

	err := os.Rename(core.NBFS_DATA_PATH + filePath, core.NBFS_DATA_PATH + data2.FilePath)
	if err != nil {
		return nil, err
	}
	return &fileData{FilesDO: data2}, nil
}

func MakeFileDataByLoad(fileId, accessHash int64) (*fileData, error) {
	data2 := dao.GetFilesDAO(dao.DB_SLAVE).Select(fileId)
	if data2 == nil {
		return nil, fmt.Errorf("not found file_id: %d", fileId)
	}

	if data2.AccessHash != accessHash {
		return nil, fmt.Errorf("invalid access_hash: %d", accessHash)
	}
	return &fileData{FilesDO: data2}, nil
}
