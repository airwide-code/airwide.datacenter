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
)

// help.getNearestDc#1fb33026 = NearestDc;
func (s *HelpServiceImpl) HelpGetNearestDc(ctx context.Context, request *mtproto.TLHelpGetNearestDc) (*mtproto.NearestDc, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("HelpGetNearestDc - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	dc := &mtproto.TLNearestDc{Data2: &mtproto.NearestDc_Data{
		Country:   "US",
		ThisDc:    2,
		NearestDc: 2,
	}}
	glog.Infof("HelpGetNearestDc - reply: %s", logger.JsonDebugData(dc))
	return dc.To_NearestDc(), nil
}
