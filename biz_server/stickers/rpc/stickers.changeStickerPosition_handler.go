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

// stickers.changeStickerPosition#ffb6d4ca sticker:InputDocument position:int = messages.StickerSet;
func (s *StickersServiceImpl) StickersChangeStickerPosition(ctx context.Context, request *mtproto.TLStickersChangeStickerPosition) (*mtproto.Messages_StickerSet, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("StickersChangeStickerPosition - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl StickersChangeStickerPosition logic

	return nil, fmt.Errorf("Not impl StickersChangeStickerPosition")
}
