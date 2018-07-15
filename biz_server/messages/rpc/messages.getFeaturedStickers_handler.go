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

// messages.getFeaturedStickers#2dacca4f hash:int = messages.FeaturedStickers;
func (s *MessagesServiceImpl) MessagesGetFeaturedStickers(ctx context.Context, request *mtproto.TLMessagesGetFeaturedStickers) (*mtproto.Messages_FeaturedStickers, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getFeaturedStickers#2dacca4f - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesGetFeaturedStickers logic
	stickers := &mtproto.TLMessagesFeaturedStickers{Data2: &mtproto.Messages_FeaturedStickers_Data{
		Hash: request.Hash,
		Sets: []*mtproto.StickerSetCovered{},
		Unread: []int64{},
	}}

	glog.Infof("messages.getFeaturedStickers#2dacca4f - reply: %s", logger.JsonDebugData(stickers))
	return stickers.To_Messages_FeaturedStickers(), nil
}
