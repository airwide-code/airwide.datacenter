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

// messages.getDocumentByHash#338e2464 sha256:bytes size:int mime_type:string = Document;
func (s *MessagesServiceImpl) MessagesGetDocumentByHash(ctx context.Context, request *mtproto.TLMessagesGetDocumentByHash) (*mtproto.Document, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesGetDocumentByHash - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesGetDocumentByHash logic

	return nil, fmt.Errorf("Not impl MessagesGetDocumentByHash")
}
