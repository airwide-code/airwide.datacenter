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

// messages.readMessageContents#36a73f77 id:Vector<int> = messages.AffectedMessages;
func (s *MessagesServiceImpl) MessagesReadMessageContents(ctx context.Context, request *mtproto.TLMessagesReadMessageContents) (*mtproto.Messages_AffectedMessages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesReadMessageContents - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesReadMessageContents logic

	return nil, fmt.Errorf("Not impl MessagesReadMessageContents")
}
