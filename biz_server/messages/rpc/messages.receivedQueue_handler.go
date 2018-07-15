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

// messages.receivedQueue#55a5bb66 max_qts:int = Vector<long>;
func (s *MessagesServiceImpl) MessagesReceivedQueue(ctx context.Context, request *mtproto.TLMessagesReceivedQueue) (*mtproto.VectorLong, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesReceivedQueue - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesReceivedQueue logic

	return nil, fmt.Errorf("Not impl MessagesReceivedQueue")
}
