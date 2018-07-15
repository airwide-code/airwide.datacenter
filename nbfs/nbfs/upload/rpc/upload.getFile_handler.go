/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package rpc

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/airwide-code/airwide.datacenter/baselib/logger"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util"
	"github.com/airwide-code/airwide.datacenter/mtproto"
	"golang.org/x/net/context"
	photo2 "github.com/airwide-code/airwide.datacenter/nbfs/biz/core/photo"
	"github.com/airwide-code/airwide.datacenter/nbfs/biz/core/document"
)

// upload.getFile#e3a6cfb5 location:InputFileLocation offset:int limit:int = upload.File;
func (s *UploadServiceImpl) UploadGetFile(ctx context.Context, request *mtproto.TLUploadGetFile) (*mtproto.Upload_File, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("upload.getFile#e3a6cfb5 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))
	var (
		uploadFile *mtproto.Upload_File
		err error
	)
	switch request.GetLocation().GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputFileLocation:
		fileLocation := request.GetLocation().To_InputFileLocation()
		uploadFile, err = photo2.GetPhotoFileData(fileLocation.GetVolumeId(),
			fileLocation.GetLocalId(),
			fileLocation.GetSecret(),
			request.GetOffset(),
			request.GetLimit())
	case mtproto.TLConstructor_CRC32_inputEncryptedFileLocation:
	case mtproto.TLConstructor_CRC32_inputDocumentFileLocation:
		fileLocation := request.GetLocation().To_InputDocumentFileLocation()
		uploadFile, err = document.GetDocumentFileData(fileLocation.GetId(),
			fileLocation.GetAccessHash(),
			fileLocation.GetVersion(),
			request.GetOffset(),
			request.GetLimit())
	case mtproto.TLConstructor_CRC32_inputDocumentFileLocationLayer11:
		fileLocation := request.GetLocation().To_InputDocumentFileLocation()
		uploadFile, err = document.GetDocumentFileData(fileLocation.GetId(),
			fileLocation.GetAccessHash(),
			fileLocation.GetVersion(),
			request.GetOffset(),
			request.GetLimit())
	default:
		err = fmt.Errorf("invalid InputFileLocation type: %d", request.GetLocation().GetConstructor())
	}

	if err != nil {
		glog.Error(err)
		return nil, err
	}

	// type:storage.FileType mtime:int bytes:bytes
	glog.Infof("upload.getFile#e3a6cfb5 - reply: {type: %v, mime: %d, len_bytes: %d}",
		uploadFile.GetData2().GetType(),
		uploadFile.GetData2().GetMtime(),
		len(uploadFile.GetData2().GetBytes()))

	return uploadFile, err
}
