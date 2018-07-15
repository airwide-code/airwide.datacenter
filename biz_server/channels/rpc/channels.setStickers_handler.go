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

// channels.setStickers#ea8ca4f9 channel:InputChannel stickerset:InputStickerSet = Bool;
func (s *ChannelsServiceImpl) ChannelsSetStickers(ctx context.Context, request *mtproto.TLChannelsSetStickers) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ChannelsSetStickers - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ChannelsSetStickers logic

	return nil, fmt.Errorf("Not impl ChannelsSetStickers")
}
