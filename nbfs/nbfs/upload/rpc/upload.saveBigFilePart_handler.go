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
)

// upload.saveBigFilePart#de7b673d file_id:long file_part:int file_total_parts:int bytes:bytes = Bool;
func (s *UploadServiceImpl) UploadSaveBigFilePart(ctx context.Context, request *mtproto.TLUploadSaveBigFilePart) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("upload.saveBigFilePart#de7b673d - metadata: %s, request: {file_id: %d, file_part: %d, bytes_len: %d}",
		logger.JsonDebugData(md),
		request.FileId,
		request.FilePart,
		len(request.Bytes))

	// TODO(@benqi): Check file_total_parts >= bigFileSize/kMaxFileSize, Check len(bytes) <= kMaxFilePartSize, Check file_part <= file_total_parts

	filePartLogic, err := file.MakeFilePartData(md.AuthId, request.FileId, request.FilePart == 0, true)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	err = filePartLogic.SaveFilePart(request.FilePart, request.Bytes)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	glog.Infof("upload.saveBigFilePart#de7b673d - reply: {true}")
	return mtproto.ToBool(true), nil

}
