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
	webpage2 "github.com/airwide-code/airwide.datacenter/biz/core/webpage"
)

// messages.getWebPagePreview#25223e24 message:string = MessageMedia;
func (s *MessagesServiceImpl) MessagesGetWebPagePreview(ctx context.Context, request *mtproto.TLMessagesGetWebPagePreview) (*mtproto.MessageMedia, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getWebPagePreview#25223e24 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	webpage := webpage2.GetWebPagePreview(request.Message)
	media := &mtproto.TLMessageMediaWebPage{Data2: &mtproto.MessageMedia_Data{
		Webpage: webpage,
	}}

	glog.Infof("messages.getWebPagePreview#25223e24 - reply: %s\n", logger.JsonDebugData(media))
	return media.To_MessageMedia(), nil
}
