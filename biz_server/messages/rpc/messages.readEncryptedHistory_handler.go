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

// messages.readEncryptedHistory#7f4b690a peer:InputEncryptedChat max_date:int = Bool;
func (s *MessagesServiceImpl) MessagesReadEncryptedHistory(ctx context.Context, request *mtproto.TLMessagesReadEncryptedHistory) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesReadEncryptedHistory - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesReadEncryptedHistory logic

	return nil, fmt.Errorf("Not impl MessagesReadEncryptedHistory")
}
