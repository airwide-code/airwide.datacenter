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

// messages.getFavedStickers#21ce0b0e hash:int = messages.FavedStickers;
func (s *MessagesServiceImpl) MessagesGetFavedStickers(ctx context.Context, request *mtproto.TLMessagesGetFavedStickers) (*mtproto.Messages_FavedStickers, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getFavedStickers#21ce0b0e - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesGetFavedStickers logic
	stickers := mtproto.TLMessagesFavedStickers{Data2: &mtproto.Messages_FavedStickers_Data{
		Hash: request.Hash,
		Packs: []*mtproto.StickerPack{},
		Stickers: []*mtproto.Document{},
	}}

	glog.Infof("messages.getFavedStickers#21ce0b0e - reply: %s", logger.JsonDebugData(stickers))
	return stickers.To_Messages_FavedStickers(), nil
}
