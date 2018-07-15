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
    "github.com/airwide-code/airwide.datacenter/mtproto"
    "golang.org/x/net/context"
    "fmt"
    "github.com/airwide-code/airwide.datacenter/baselib/grpc_util"
    "github.com/airwide-code/airwide.datacenter/baselib/logger"
)

// messages.readMentions#f0189d3 peer:InputPeer = messages.AffectedHistory;
func (s *MessagesServiceImpl) MessagesReadMentions(ctx context.Context, request *mtproto.TLMessagesReadMentions) (*mtproto.Messages_AffectedHistory, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("MessagesReadMentions - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    // TODO(@benqi): Impl MessagesReadMentions logic

    return nil, fmt.Errorf("Not impl MessagesReadMentions")
}
