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

// help.getTermsOfService#350170f3 = help.TermsOfService;
func (s *HelpServiceImpl) HelpGetTermsOfService(ctx context.Context, request *mtproto.TLHelpGetTermsOfService) (*mtproto.Help_TermsOfService, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("HelpGetTermsOfService - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl HelpGetTermsOfService logic

	return nil, fmt.Errorf("Not impl HelpGetTermsOfService")
}
