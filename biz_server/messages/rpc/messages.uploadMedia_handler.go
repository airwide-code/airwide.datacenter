/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package rpc

import (
	// "fmt"
	"github.com/golang/glog"
	"github.com/airwide-code/airwide.datacenter/baselib/logger"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util"
	"github.com/airwide-code/airwide.datacenter/mtproto"
	"golang.org/x/net/context"
)

// messages.uploadMedia#519bc2b1 peer:InputPeer media:InputMedia = MessageMedia;
func (s *MessagesServiceImpl) MessagesUploadMedia(ctx context.Context, request *mtproto.TLMessagesUploadMedia) (*mtproto.MessageMedia, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.uploadMedia#519bc2b1 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	messageMedia := makeMediaByInputMedia(md.AuthId, request.GetMedia())

	//// TODO(@benqi): Impl MessagesUploadMedia logic
	//return nil, fmt.Errorf("Not impl MessagesUploadMedia")

	glog.Infof("messages.uploadMedia#519bc2b1 - reply: %s", logger.JsonDebugData(messageMedia))
	return messageMedia, nil
}
