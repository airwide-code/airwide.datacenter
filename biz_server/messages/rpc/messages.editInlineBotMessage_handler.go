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

// messages.editInlineBotMessage#b0e08243 flags:# no_webpage:flags.1?true stop_geo_live:flags.12?true id:InputBotInlineMessageID message:flags.11?string reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> geo_point:flags.13?InputGeoPoint = Bool;
func (s *MessagesServiceImpl) MessagesEditInlineBotMessage(ctx context.Context, request *mtproto.TLMessagesEditInlineBotMessage) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.editInlineBotMessage#b0e08243 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesEditInlineBotMessage logic

	return nil, fmt.Errorf("Not impl MessagesEditInlineBotMessage")
}
