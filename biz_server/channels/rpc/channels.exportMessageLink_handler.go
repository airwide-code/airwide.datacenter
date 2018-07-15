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

// channels.exportMessageLink#c846d22d channel:InputChannel id:int = ExportedMessageLink;
func (s *ChannelsServiceImpl) ChannelsExportMessageLink(ctx context.Context, request *mtproto.TLChannelsExportMessageLink) (*mtproto.ExportedMessageLink, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ChannelsExportMessageLink - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ChannelsExportMessageLink logic

	return nil, fmt.Errorf("Not impl ChannelsExportMessageLink")
}
