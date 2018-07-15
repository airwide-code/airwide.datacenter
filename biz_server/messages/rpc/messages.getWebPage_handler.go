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

// messages.getWebPage#32ca8f91 url:string hash:int = WebPage;
func (s *MessagesServiceImpl) MessagesGetWebPage(ctx context.Context, request *mtproto.TLMessagesGetWebPage) (*mtproto.WebPage, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesGetWebPage - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesGetWebPage logic

	return nil, fmt.Errorf("Not impl MessagesGetWebPage")
}
