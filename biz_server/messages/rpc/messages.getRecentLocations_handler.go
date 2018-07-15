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

// messages.getRecentLocations#249431e2 peer:InputPeer limit:int = messages.Messages;
func (s *MessagesServiceImpl) MessagesGetRecentLocations(ctx context.Context, request *mtproto.TLMessagesGetRecentLocations) (*mtproto.Messages_Messages, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("MessagesGetRecentLocations - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    // TODO(@benqi): Impl MessagesGetRecentLocations logic

    return nil, fmt.Errorf("Not impl MessagesGetRecentLocations")
}
