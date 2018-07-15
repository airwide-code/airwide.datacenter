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

// messages.searchGlobal#9e3cacb0 q:string offset_date:int offset_peer:InputPeer offset_id:int limit:int = messages.Messages;
func (s *MessagesServiceImpl) MessagesSearchGlobal(ctx context.Context, request *mtproto.TLMessagesSearchGlobal) (*mtproto.Messages_Messages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesSearchGlobal - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesSearchGlobal logic

	return nil, fmt.Errorf("Not impl MessagesSearchGlobal")
}
