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

// help.getAppChangelog#9010ef6f prev_app_version:string = Updates;
func (s *HelpServiceImpl) HelpGetAppChangelog(ctx context.Context, request *mtproto.TLHelpGetAppChangelog) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("HelpGetAppChangelog - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl HelpGetAppChangelog logic

	return nil, fmt.Errorf("Not impl HelpGetAppChangelog")
}
