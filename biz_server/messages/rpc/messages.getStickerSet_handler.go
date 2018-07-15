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
	"github.com/airwide-code/airwide.datacenter/biz/core/sticker"
	"github.com/airwide-code/airwide.datacenter/biz/nbfs_client"
)

// messages.getStickerSet#2619a90e stickerset:InputStickerSet = messages.StickerSet;
func (s *MessagesServiceImpl) MessagesGetStickerSet(ctx context.Context, request *mtproto.TLMessagesGetStickerSet) (*mtproto.Messages_StickerSet, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getStickerSet#2619a90e - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): check inputStickerSetEmpty
	set := sticker.GetStickerSet(request.GetStickerset())
	packs, idList := sticker.GetStickerPackList(set.GetData2().GetId())
	var (
		documents []*mtproto.Document
		err error
	)

	if len(idList) == 0 {
		documents = []*mtproto.Document{}
	} else {
		documents, err = nbfs_client.GetDocumentByIdList(idList)
		if err != nil {
			glog.Error(err)
			documents = []*mtproto.Document{}
		}
	}

	reply := &mtproto.TLMessagesStickerSet{Data2: &mtproto.Messages_StickerSet_Data{
		Set:       set,
		Packs:     packs,
		Documents: documents,
	}}

	glog.Infof("messages.getStickerSet#2619a90e - reply: %s", logger.JsonDebugData(reply))
	return reply.To_Messages_StickerSet(), nil
}
