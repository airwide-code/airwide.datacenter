/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package core

import (
	"github.com/golang/glog"
	"crypto/md5"
	"fmt"
	"os"
	"io"
	"github.com/airwide-code/airwide.datacenter/mtproto"
)

//const (
//	NBFS_DATA_PATH = "/opt/nbfs"
//	// kMaxFilePartSize = 32768
//)

var NBFS_DATA_PATH = "/opt/nbfs"

func InitNbfsDataPath(path string) {
	NBFS_DATA_PATH = path
}

func MakeStorageFileType(extName string) *mtproto.Storage_FileType {
	fileType := &mtproto.Storage_FileType{Data2: &mtproto.Storage_FileType_Data{}}

	switch extName {
	case ".partial":
		fileType.Constructor = mtproto.TLConstructor_CRC32_storage_filePartial
	case ".jpeg", ".jpg":
		fileType.Constructor = mtproto.TLConstructor_CRC32_storage_fileJpeg
	case ".gif":
		fileType.Constructor = mtproto.TLConstructor_CRC32_storage_fileGif
	case ".png":
		fileType.Constructor = mtproto.TLConstructor_CRC32_storage_filePng
	case ".pdf":
		fileType.Constructor = mtproto.TLConstructor_CRC32_storage_filePdf
	case ".mp3":
		fileType.Constructor = mtproto.TLConstructor_CRC32_storage_fileMp3
	case ".mov":
		fileType.Constructor = mtproto.TLConstructor_CRC32_storage_fileMov
	case ".mp4":
		fileType.Constructor = mtproto.TLConstructor_CRC32_storage_fileMp4
	case ".webp":
		fileType.Constructor = mtproto.TLConstructor_CRC32_storage_fileWebp
	default:
		// fileType.Constructor = mtproto.TLConstructor_CRC32_storage_fileUnknown
		fileType.Constructor = mtproto.TLConstructor_CRC32_storage_filePartial
	}

	return fileType
}

// TODO(@benqi): remove to baselib
func CalcMd5File(filename string) (string, error) {
	// fileName := core.NBFS_DATA_PATH + m.data.FilePath
	f, err := os.Open(filename)
	if err != nil {
		glog.Error(err)
		return "", err
	}

	defer f.Close()

	md5Hash := md5.New()
	if _, err := io.Copy(md5Hash, f); err != nil {
		// fmt.Println("Copy", err)
		glog.Error("Copy - ", err)
		return "", err
	}

	return fmt.Sprintf("%x", md5Hash.Sum(nil)), nil
}
