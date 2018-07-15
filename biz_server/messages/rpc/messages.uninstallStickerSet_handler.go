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

// messages.uninstallStickerSet#f96e55de stickerset:InputStickerSet = Bool;
func (s *MessagesServiceImpl) MessagesUninstallStickerSet(ctx context.Context, request *mtproto.TLMessagesUninstallStickerSet) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesUninstallStickerSet - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesUninstallStickerSet logic

	return nil, fmt.Errorf("Not impl MessagesUninstallStickerSet")
}
