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

// messages.discardEncryption#edd923c5 chat_id:int = Bool;
func (s *MessagesServiceImpl) MessagesDiscardEncryption(ctx context.Context, request *mtproto.TLMessagesDiscardEncryption) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesDiscardEncryption - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesDiscardEncryption logic

	return nil, fmt.Errorf("Not impl MessagesDiscardEncryption")
}
