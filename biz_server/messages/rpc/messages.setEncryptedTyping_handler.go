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

// messages.setEncryptedTyping#791451ed peer:InputEncryptedChat typing:Bool = Bool;
func (s *MessagesServiceImpl) MessagesSetEncryptedTyping(ctx context.Context, request *mtproto.TLMessagesSetEncryptedTyping) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesSetEncryptedTyping - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesSetEncryptedTyping logic

	return nil, fmt.Errorf("Not impl MessagesSetEncryptedTyping")
}
