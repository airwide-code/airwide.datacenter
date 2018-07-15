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

// messages.reorderStickerSets#78337739 flags:# masks:flags.0?true order:Vector<long> = Bool;
func (s *MessagesServiceImpl) MessagesReorderStickerSets(ctx context.Context, request *mtproto.TLMessagesReorderStickerSets) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesReorderStickerSets - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesReorderStickerSets logic

	return nil, fmt.Errorf("Not impl MessagesReorderStickerSets")
}
