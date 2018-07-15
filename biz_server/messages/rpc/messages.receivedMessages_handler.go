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

// messages.receivedMessages#5a954c0 max_id:int = Vector<ReceivedNotifyMessage>;
func (s *MessagesServiceImpl) MessagesReceivedMessages(ctx context.Context, request *mtproto.TLMessagesReceivedMessages) (*mtproto.Vector_ReceivedNotifyMessage, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.receivedMessages#5a954c0 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// @benai: android client and tdesktop 未使用
	return &mtproto.Vector_ReceivedNotifyMessage{Datas: []*mtproto.ReceivedNotifyMessage{}}, nil
}
