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
	"github.com/airwide-code/airwide.datacenter/biz/core/user"
	"time"
	"github.com/airwide-code/airwide.datacenter/biz/nbfs_client"
)

/*
 rpc_requst:
	body: { photos_getUserPhotos
	  user_id: { inputUserSelf },
	  offset: 1 [INT],
	  max_id: 0 [LONG],
	  limit: 5 [INT],
	},

 rpc_result:
  body: { rpc_result
    req_msg_id: 6537205080566771468 [LONG],
    result: { photos_photosSlice
      count: 1 [INT],
      photos: [ vector<0x0> ],
      users: [ vector<0x0> ],
    },
  },
 */

// photos.getUserPhotos#91cd32a8 user_id:InputUser offset:int max_id:long limit:int = photos.Photos;
func (s *PhotosServiceImpl) PhotosGetUserPhotos(ctx context.Context, request *mtproto.TLPhotosGetUserPhotos) (*mtproto.Photos_Photos, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("photos.getUserPhotos#91cd32a8 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl PhotosGetUserPhotos logic
	var userId int32 = 0
	switch request.GetUserId().GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputUserSelf:
		userId = md.UserId
	case mtproto.TLConstructor_CRC32_inputUser:
		userId = request.GetUserId().GetData2().GetUserId()
	default:
		// TODO(@benqi): bad request
	}

	photos := mtproto.NewTLPhotosPhotos()
	photoIdList := user.GetUserPhotoIDList(userId)
	// idList := []int32{}
	for _, photoId := range photoIdList {
		sizes, _ := nbfs_client.GetPhotoSizeList(photoId)
		// photo2 := photo2.MakeUserProfilePhoto(photoId, sizes)
		photo := &mtproto.TLPhoto{ Data2: &mtproto.Photo_Data{
			Id:          photoId,
			HasStickers: false,
			AccessHash:  photoId, // photo2.GetFileAccessHash(file.GetData2().GetId(), file.GetData2().GetParts()),
			Date:        int32(time.Now().Unix()),
			Sizes:       sizes,
		}}
		photos.Data2.Photos = append(photos.Data2.Photos, photo.To_Photo())
	}
	// if idList

	glog.Infof("photos.getUserPhotos#91cd32a8 - reply: %s", logger.JsonDebugData(photos))
	return photos.To_Photos_Photos(), nil
}
