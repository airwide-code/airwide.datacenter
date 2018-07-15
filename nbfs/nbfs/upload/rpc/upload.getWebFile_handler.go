/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
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
)

// upload.getWebFile#24e6818d location:InputWebFileLocation offset:int limit:int = upload.WebFile;
func (s *UploadServiceImpl) UploadGetWebFile(ctx context.Context, request *mtproto.TLUploadGetWebFile) (*mtproto.Upload_WebFile, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("UploadGetWebFile - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl UploadGetWebFile logic

	return nil, fmt.Errorf("Not impl UploadGetWebFile")
}
