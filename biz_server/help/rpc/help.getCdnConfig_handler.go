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

// help.getCdnConfig#52029342 = CdnConfig;
func (s *HelpServiceImpl) HelpGetCdnConfig(ctx context.Context, request *mtproto.TLHelpGetCdnConfig) (*mtproto.CdnConfig, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("HelpGetCdnConfig - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl HelpGetCdnConfig logic

	return nil, fmt.Errorf("Not impl HelpGetCdnConfig")
}
