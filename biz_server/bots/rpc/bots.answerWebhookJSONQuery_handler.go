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

// bots.answerWebhookJSONQuery#e6213f4d query_id:long data:DataJSON = Bool;
func (s *BotsServiceImpl) BotsAnswerWebhookJSONQuery(ctx context.Context, request *mtproto.TLBotsAnswerWebhookJSONQuery) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("BotsAnswerWebhookJSONQuery - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl BotsAnswerWebhookJSONQuery logic

	return nil, fmt.Errorf("Not impl BotsAnswerWebhookJSONQuery")
}
