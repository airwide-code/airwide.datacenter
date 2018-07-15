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

// messages.clearRecentStickers#8999602d flags:# attached:flags.0?true = Bool;
func (s *MessagesServiceImpl) MessagesClearRecentStickers(ctx context.Context, request *mtproto.TLMessagesClearRecentStickers) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesClearRecentStickers - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesClearRecentStickers logic

	return nil, fmt.Errorf("Not impl MessagesClearRecentStickers")
}
