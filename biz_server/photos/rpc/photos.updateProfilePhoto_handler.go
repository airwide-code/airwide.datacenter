/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
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
	photo2 "github.com/airwide-code/airwide.datacenter/biz/core/photo"
	"github.com/airwide-code/airwide.datacenter/biz/nbfs_client"
)

// photos.updateProfilePhoto#f0bb5152 id:InputPhoto = UserProfilePhoto;
func (s *PhotosServiceImpl) PhotosUpdateProfilePhoto(ctx context.Context, request *mtproto.TLPhotosUpdateProfilePhoto) (*mtproto.UserProfilePhoto, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("photos.updateProfilePhoto#f0bb5152 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		photo *mtproto.UserProfilePhoto
	)

	if request.GetId().GetConstructor() == mtproto.TLConstructor_CRC32_inputPhotoEmpty {
		photo = mtproto.NewTLUserProfilePhotoEmpty().To_UserProfilePhoto()
	} else {
		id := request.GetId().To_InputPhoto()
		// TODO(@benqi): check inputPhoto.access_hash

		sizes, _ := nbfs_client.GetPhotoSizeList(id.GetId())
		photo = photo2.MakeUserProfilePhoto(id.GetId(), sizes)
	}

	// TODO(@benqi): sync update.

	glog.Infof("photos.uploadProfilePhoto#4f32c098 - reply: %s", logger.JsonDebugData(photo))
	return photo, nil
}
