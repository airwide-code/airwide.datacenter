/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package rpc

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/airwide-code/airwide.datacenter/baselib/logger"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util"
	"github.com/airwide-code/airwide.datacenter/mtproto"
	"golang.org/x/net/context"
)

// stickers.removeStickerFromSet#f7760f51 sticker:InputDocument = messages.StickerSet;
func (s *StickersServiceImpl) StickersRemoveStickerFromSet(ctx context.Context, request *mtproto.TLStickersRemoveStickerFromSet) (*mtproto.Messages_StickerSet, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("StickersRemoveStickerFromSet - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl StickersRemoveStickerFromSet logic

	return nil, fmt.Errorf("Not impl StickersRemoveStickerFromSet")
}
