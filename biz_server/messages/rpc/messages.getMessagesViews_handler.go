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

// messages.getMessagesViews#c4c8a55d peer:InputPeer id:Vector<int> increment:Bool = Vector<int>;
func (s *MessagesServiceImpl) MessagesGetMessagesViews(ctx context.Context, request *mtproto.TLMessagesGetMessagesViews) (*mtproto.VectorInt, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getMessagesViews#c4c8a55d - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesGetMessagesViews logic
	views := &mtproto.VectorInt{
		Datas: []int32{},
	}

	glog.Infof("messages.getMessagesViews#c4c8a55d - reply: %s", logger.JsonDebugData(views))
	return views, nil
}
