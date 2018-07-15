/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package rpc

import (
	"github.com/golang/glog"
	"github.com/airwide-code/airwide.datacenter/baselib/logger"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util"
	"github.com/airwide-code/airwide.datacenter/mtproto"
	"golang.org/x/net/context"
	"github.com/airwide-code/airwide.datacenter/nbfs/biz/core/file"
	"io/ioutil"
	"fmt"
)

// upload.saveFilePart#b304a621 file_id:long file_part:int bytes:bytes = Bool;
func (s *UploadServiceImpl) UploadSaveFilePart(ctx context.Context, request *mtproto.TLUploadSaveFilePart) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("upload.saveFilePart#b304a621 - metadata: %s, request: {file_id: %d, file_part: %d, bytes_len: %d}",
		logger.JsonDebugData(md),
		request.FileId,
		request.FilePart,
		len(request.Bytes))

	// TODO(@benqi): Check file_part <= bigFileSize/kMaxFileSize, Check len(bytes) <= kMaxFilePartSize

	filePartLogic, err := file.MakeFilePartData(md.AuthId, request.FileId, request.FilePart == 0, false)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	ioutil.WriteFile(fmt.Sprintf("/tmp/uploads/%d_%d.tmp", request.GetFileId(), request.GetFilePart()), request.GetBytes(), 0644)
	err = filePartLogic.SaveFilePart(request.FilePart, request.Bytes)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	glog.Infof("upload.saveFilePart#b304a621 - reply: {true}")
	return mtproto.ToBool(true), nil
}
