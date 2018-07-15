/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package sticker

import (
	"github.com/airwide-code/airwide.datacenter/mtproto"
	"github.com/airwide-code/airwide.datacenter/biz/dal/dao"
	"github.com/airwide-code/airwide.datacenter/biz/dal/dataobject"
	"github.com/golang/glog"
)

//public static final int TYPE_IMAGE = 0;
//public static final int TYPE_MASK = 1;
//public static final int TYPE_FAVE = 2;

//type Logic struct {
//
//}

// stickerSet#cd303b41 flags:# installed:flags.0?true archived:flags.1?true official:flags.2?true masks:flags.3?true id:long access_hash:long title:string short_name:string count:int hash:int = StickerSet;

func makeStickerSet(do *dataobject.StickerSetsDO) *mtproto.StickerSet {
	sitckers := &mtproto.TLStickerSet{Data2: &mtproto.StickerSet_Data{
		Installed:  true,
		Id:         do.StickerSetId,
		AccessHash: do.AccessHash,
		Title:      do.Title,
		ShortName:  do.ShortName,
		Hash:       do.Hash,
	}}
	return sitckers.To_StickerSet()
}

func GetStickerSetList(hash int32) []*mtproto.StickerSet {
	//
	doList := dao.GetStickerSetsDAO(dao.DB_SLAVE).SelectAll()
	stickers := make([]*mtproto.StickerSet, len(doList))
	for i := 0; i < len(doList); i++ {
		stickers[i] = makeStickerSet(&doList[i])
	}
	return stickers
}

func GetStickerSet(stickerset *mtproto.InputStickerSet) *mtproto.StickerSet {
	var (
		inputSet = stickerset.GetData2()
		set *mtproto.StickerSet
	)

	switch stickerset.GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputStickerSetID:
		do := dao.GetStickerSetsDAO(dao.DB_SLAVE).SelectByID(inputSet.GetId(), inputSet.GetAccessHash())
		if do != nil {
			set = makeStickerSet(do)
		}
	case mtproto.TLConstructor_CRC32_inputStickerSetShortName:
		do := dao.GetStickerSetsDAO(dao.DB_SLAVE).SelectByShortName(inputSet.GetShortName())
		if do != nil {
			set = makeStickerSet(do)
		}
	case mtproto.TLConstructor_CRC32_inputStickerSetEmpty:
		glog.Error("stickerset is inputStickerSetEmpty")
	}

	return set
}

func GetStickerPackList(setId int64) ([]*mtproto.StickerPack, []int64) {
	doList := dao.GetStickerPacksDAO(dao.DB_SLAVE).SelectBySetID(setId)
	packs := make([]*mtproto.StickerPack, len(doList))
	idList := make([]int64, len(doList))
	for i := 0; i < len(doList); i++ {
		packs[i] = &mtproto.StickerPack{
			Constructor: mtproto.TLConstructor_CRC32_stickerPack,
			Data2: 		 &mtproto.StickerPack_Data{
				Emoticon:  doList[i].Emoticon,
				Documents: []int64{doList[i].DocumentId},
			},
		}
		idList[i] = doList[i].DocumentId
	}
	return packs, idList
}

//func GetStickerDocumentList(idList []int64) []*mtproto.Document {
//	return nil
//}
