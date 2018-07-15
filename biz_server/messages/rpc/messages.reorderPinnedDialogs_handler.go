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

// messages.reorderPinnedDialogs#959ff644 flags:# force:flags.0?true order:Vector<InputPeer> = Bool;
func (s *MessagesServiceImpl) MessagesReorderPinnedDialogs(ctx context.Context, request *mtproto.TLMessagesReorderPinnedDialogs) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesReorderPinnedDialogs - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesReorderPinnedDialogs logic

	return nil, fmt.Errorf("Not impl MessagesReorderPinnedDialogs")
}
