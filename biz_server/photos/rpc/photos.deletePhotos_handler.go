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

// photos.deletePhotos#87cf7f2f id:Vector<InputPhoto> = Vector<long>;
func (s *PhotosServiceImpl) PhotosDeletePhotos(ctx context.Context, request *mtproto.TLPhotosDeletePhotos) (*mtproto.VectorLong, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("PhotosDeletePhotos - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl PhotosDeletePhotos logic

	return nil, fmt.Errorf("Not impl PhotosDeletePhotos")
}
