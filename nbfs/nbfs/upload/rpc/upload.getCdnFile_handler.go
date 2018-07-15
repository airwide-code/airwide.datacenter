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

// upload.getCdnFile#2000bcc3 file_token:bytes offset:int limit:int = upload.CdnFile;
func (s *UploadServiceImpl) UploadGetCdnFile(ctx context.Context, request *mtproto.TLUploadGetCdnFile) (*mtproto.Upload_CdnFile, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("UploadGetCdnFile - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl UploadGetCdnFile logic

	return nil, fmt.Errorf("Not impl UploadGetCdnFile")
}
