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

// messages.getSavedGifs#83bf3d52 hash:int = messages.SavedGifs;
func (s *MessagesServiceImpl) MessagesGetSavedGifs(ctx context.Context, request *mtproto.TLMessagesGetSavedGifs) (*mtproto.Messages_SavedGifs, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesGetSavedGifs - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesGetSavedGifs logic
	stickers := mtproto.TLMessagesSavedGifs{Data2: &mtproto.Messages_SavedGifs_Data{
		Hash: request.Hash,
		Gifs: []*mtproto.Document{},
	}}

	glog.Infof("MessagesGetSavedGifs - reply: %s\n", logger.JsonDebugData(stickers))
	return stickers.To_Messages_SavedGifs(), nil
}
