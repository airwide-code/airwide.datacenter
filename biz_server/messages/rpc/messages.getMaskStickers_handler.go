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

// messages.getMaskStickers#65b8c79f hash:int = messages.AllStickers;
func (s *MessagesServiceImpl) MessagesGetMaskStickers(ctx context.Context, request *mtproto.TLMessagesGetMaskStickers) (*mtproto.Messages_AllStickers, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getMaskStickers#65b8c79f - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesGetMaskStickers logic
	stickers := &mtproto.TLMessagesAllStickers{Data2: &mtproto.Messages_AllStickers_Data{
		Hash: request.Hash,
		Sets: []*mtproto.StickerSet{},
	}}

	glog.Infof("messages.getMaskStickers#65b8c79f - reply: %s", logger.JsonDebugData(stickers))
	return stickers.To_Messages_AllStickers(), nil
}
