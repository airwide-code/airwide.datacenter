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
	update2 "github.com/airwide-code/airwide.datacenter/biz/core/update"
)

// 执行getState后，获取最新的pts, qts and seq
// updates.getState#edd4882a = updates.State;
func (s *UpdatesServiceImpl) UpdatesGetState(ctx context.Context, request *mtproto.TLUpdatesGetState) (*mtproto.Updates_State, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("updates.getState#edd4882a  - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	state := update2.GetServerUpdatesState(md.AuthId, md.UserId)
	glog.Infof("updates.getState#edd4882a  - reply: %s", logger.JsonDebugData(state))
	return state.To_Updates_State(), nil
}
