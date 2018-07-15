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

// messages.search#39e9ea0 flags:# peer:InputPeer q:string from_id:flags.0?InputUser filter:MessagesFilter min_date:int max_date:int offset_id:int add_offset:int limit:int max_id:int min_id:int = messages.Messages;
func (s *MessagesServiceImpl) MessagesSearch(ctx context.Context, request *mtproto.TLMessagesSearch) (*mtproto.Messages_Messages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.search#39e9ea0 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Not impl
	messages := &mtproto.TLMessagesMessages{Data2: &mtproto.Messages_Messages_Data{
		Messages: []*mtproto.Message{},
		Chats:    []*mtproto.Chat{},
		Users:    []*mtproto.User{},
	}}

	glog.Infof("messages.search#39e9ea0 - reply: %s", messages)
	return messages.To_Messages_Messages(), nil
}
