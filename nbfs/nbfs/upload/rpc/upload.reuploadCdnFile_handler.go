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

// upload.reuploadCdnFile#1af91c09 file_token:bytes request_token:bytes = Vector<CdnFileHash>;
func (s *UploadServiceImpl) UploadReuploadCdnFile(ctx context.Context, request *mtproto.TLUploadReuploadCdnFile) (*mtproto.Vector_CdnFileHash, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("UploadReuploadCdnFile - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl UploadReuploadCdnFile logic

	return nil, fmt.Errorf("Not impl UploadReuploadCdnFile")
}
