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
)

// messages.getArchivedStickers#57f17692 flags:# masks:flags.0?true offset_id:long limit:int = messages.ArchivedStickers;
func (s *MessagesServiceImpl) MessagesGetArchivedStickers(ctx context.Context, request *mtproto.TLMessagesGetArchivedStickers) (*mtproto.Messages_ArchivedStickers, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getArchivedStickers#57f17692 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesGetArchivedStickers logic
	stickers := &mtproto.TLMessagesArchivedStickers{Data2: &mtproto.Messages_ArchivedStickers_Data{
		Count: 0,
		Sets:  []*mtproto.StickerSetCovered{},
	}}

	glog.Infof("messages.getArchivedStickers#57f17692 - reply: %s", logger.JsonDebugData(stickers))
	return stickers.To_Messages_ArchivedStickers(), nil
}
