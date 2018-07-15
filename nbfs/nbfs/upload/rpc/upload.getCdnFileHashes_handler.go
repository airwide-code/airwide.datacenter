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

// upload.getCdnFileHashes#f715c87b file_token:bytes offset:int = Vector<CdnFileHash>;
func (s *UploadServiceImpl) UploadGetCdnFileHashes(ctx context.Context, request *mtproto.TLUploadGetCdnFileHashes) (*mtproto.Vector_CdnFileHash, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("UploadGetCdnFileHashes - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl UploadGetCdnFileHashes logic

	return nil, fmt.Errorf("Not impl UploadGetCdnFileHashes")
}
