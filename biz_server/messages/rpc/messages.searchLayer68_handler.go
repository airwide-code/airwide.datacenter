/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package rpc

import (
	"github.com/golang/glog"
	"github.com/airwide-code/airwide.datacenter/baselib/logger"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util"
	"github.com/airwide-code/airwide.datacenter/mtproto"
	"golang.org/x/net/context"
)

// messages.search#f288a275 flags:# peer:InputPeer q:string from_id:flags.0?InputUser filter:MessagesFilter min_date:int max_date:int offset:int max_id:int limit:int = messages.Messages;
func (s *MessagesServiceImpl) MessagesSearchLayer68(ctx context.Context, request *mtproto.TLMessagesSearchLayer68) (*mtproto.Messages_Messages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesSearch - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Not impl
	messages := mtproto.NewTLMessagesMessages()

	glog.Infof("MessagesGetHistory - reply: %s", messages)
	return messages.To_Messages_Messages(), nil
}
