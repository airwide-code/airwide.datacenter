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

// messages.searchGifs#bf9a776b q:string offset:int = messages.FoundGifs;
func (s *MessagesServiceImpl) MessagesSearchGifs(ctx context.Context, request *mtproto.TLMessagesSearchGifs) (*mtproto.Messages_FoundGifs, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesSearchGifs - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesSearchGifs logic

	return nil, fmt.Errorf("Not impl MessagesSearchGifs")
}
