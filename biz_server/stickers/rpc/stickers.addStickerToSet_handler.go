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

// stickers.addStickerToSet#8653febe stickerset:InputStickerSet sticker:InputStickerSetItem = messages.StickerSet;
func (s *StickersServiceImpl) StickersAddStickerToSet(ctx context.Context, request *mtproto.TLStickersAddStickerToSet) (*mtproto.Messages_StickerSet, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("StickersAddStickerToSet - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl StickersAddStickerToSet logic

	return nil, fmt.Errorf("Not impl StickersAddStickerToSet")
}
