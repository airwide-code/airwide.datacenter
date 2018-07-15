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

// messages.importChatInvite#6c50051c hash:string = Updates;
func (s *MessagesServiceImpl) MessagesImportChatInvite(ctx context.Context, request *mtproto.TLMessagesImportChatInvite) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesImportChatInvite - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesImportChatInvite logic

	return nil, fmt.Errorf("Not impl MessagesImportChatInvite")
}
