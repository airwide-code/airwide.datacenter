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
)

// messages.getAllStickers#1c9618b1 hash:int = messages.AllStickers;
func (s *MessagesServiceImpl) MessagesGetAllStickers(ctx context.Context, request *mtproto.TLMessagesGetAllStickers) (*mtproto.Messages_AllStickers, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getAllStickers#1c9618b1 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	stickers := &mtproto.TLMessagesAllStickers{Data2: &mtproto.Messages_AllStickers_Data{
		Hash: 0, // TODO(@benqi): hash规则
		Sets: sticker.GetStickerSetList(request.Hash),
	}}

	glog.Infof("messages.getAllStickers#1c9618b1 - reply: %s", logger.JsonDebugData(stickers))
	return stickers.To_Messages_AllStickers(), nil
}
