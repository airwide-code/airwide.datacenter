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

// messages.exportChatInvite#7d885289 chat_id:int = ExportedChatInvite;
func (s *MessagesServiceImpl) MessagesExportChatInvite(ctx context.Context, request *mtproto.TLMessagesExportChatInvite) (*mtproto.ExportedChatInvite, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesExportChatInvite - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesExportChatInvite logic

	return nil, fmt.Errorf("Not impl MessagesExportChatInvite")
}
