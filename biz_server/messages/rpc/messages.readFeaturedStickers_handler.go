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

// messages.readFeaturedStickers#5b118126 id:Vector<long> = Bool;
func (s *MessagesServiceImpl) MessagesReadFeaturedStickers(ctx context.Context, request *mtproto.TLMessagesReadFeaturedStickers) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesReadFeaturedStickers - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesReadFeaturedStickers logic

	return nil, fmt.Errorf("Not impl MessagesReadFeaturedStickers")
}
