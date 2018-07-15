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

// bots.sendCustomRequest#aa2769ed custom_method:string params:DataJSON = DataJSON;
func (s *BotsServiceImpl) BotsSendCustomRequest(ctx context.Context, request *mtproto.TLBotsSendCustomRequest) (*mtproto.DataJSON, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("BotsSendCustomRequest - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl BotsSendCustomRequest logic

	return nil, fmt.Errorf("Not impl BotsSendCustomRequest")
}
