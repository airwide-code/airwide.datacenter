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
    "github.com/airwide-code/airwide.datacenter/mtproto"
    "golang.org/x/net/context"
    "fmt"
    "github.com/airwide-code/airwide.datacenter/baselib/grpc_util"
    "github.com/airwide-code/airwide.datacenter/baselib/logger"
)

// help.getRecentMeUrls#3dc0f114 referer:string = help.RecentMeUrls;
func (s *HelpServiceImpl) HelpGetRecentMeUrls(ctx context.Context, request *mtproto.TLHelpGetRecentMeUrls) (*mtproto.Help_RecentMeUrls, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("HelpGetRecentMeUrls - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    // TODO(@benqi): Impl HelpGetRecentMeUrls logic

    return nil, fmt.Errorf("Not impl HelpGetRecentMeUrls")
}
