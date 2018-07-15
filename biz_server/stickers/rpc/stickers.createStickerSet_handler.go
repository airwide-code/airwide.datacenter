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

// stickers.createStickerSet#9bd86e6a flags:# masks:flags.0?true user_id:InputUser title:string short_name:string stickers:Vector<InputStickerSetItem> = messages.StickerSet;
func (s *StickersServiceImpl) StickersCreateStickerSet(ctx context.Context, request *mtproto.TLStickersCreateStickerSet) (*mtproto.Messages_StickerSet, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("StickersCreateStickerSet - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl StickersCreateStickerSet logic

	return nil, fmt.Errorf("Not impl StickersCreateStickerSet")
}
