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

// messages.getRecentStickers#5ea192c9 flags:# attached:flags.0?true hash:int = messages.RecentStickers;
func (s *MessagesServiceImpl) MessagesGetRecentStickers(ctx context.Context, request *mtproto.TLMessagesGetRecentStickers) (*mtproto.Messages_RecentStickers, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getRecentStickers#5ea192c9 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesGetRecentStickers logic
	stickers := &mtproto.TLMessagesRecentStickers{Data2: &mtproto.Messages_RecentStickers_Data{
		Hash: request.Hash,
		Stickers: []*mtproto.Document{},
	}}

	glog.Infof("messages.getRecentStickers#5ea192c9 - reply: %s", logger.JsonDebugData(stickers))
	return stickers.To_Messages_RecentStickers(), nil
}
