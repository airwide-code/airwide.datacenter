/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package photo

import (
	"github.com/airwide-code/airwide.datacenter/mtproto"
)

func MakeUserProfilePhoto(photoId int64, sizes []*mtproto.PhotoSize) *mtproto.UserProfilePhoto {
	if len(sizes) == 0 {
		return mtproto.NewTLUserProfilePhotoEmpty().To_UserProfilePhoto()
	}

	// TODO(@benqi): check PhotoSize is photoSizeEmpty
	photo := &mtproto.TLUserProfilePhoto{Data2: &mtproto.UserProfilePhoto_Data{
		PhotoId: photoId,
		PhotoSmall: sizes[0].GetData2().GetLocation(),
		PhotoBig: sizes[len(sizes)-1].GetData2().GetLocation(),
	}}

	return photo.To_UserProfilePhoto()
}

func MakeChatPhoto(sizes []*mtproto.PhotoSize) *mtproto.ChatPhoto {
	if len(sizes) == 0 {
		return mtproto.NewTLChatPhotoEmpty().To_ChatPhoto()
	}

	// TODO(@benqi): check PhotoSize is photoSizeEmpty
	photo := &mtproto.TLChatPhoto{Data2: &mtproto.ChatPhoto_Data{
		PhotoSmall: sizes[0].GetData2().GetLocation(),
		PhotoBig: sizes[len(sizes)-1].GetData2().GetLocation(),
	}}

	return photo.To_ChatPhoto()
}

