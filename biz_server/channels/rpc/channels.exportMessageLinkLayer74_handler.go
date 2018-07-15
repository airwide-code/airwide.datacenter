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

// Layer74
// channels.exportMessageLink#ceb77163 channel:InputChannel id:int grouped:Bool = ExportedMessageLink;
func (s *ChannelsServiceImpl) ChannelsExportMessageLinkLayer74(ctx context.Context, request *mtproto.TLChannelsExportMessageLinkLayer74) (*mtproto.ExportedMessageLink, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ChannelsExportMessageLink - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ChannelsExportMessageLink logic

	return nil, fmt.Errorf("Not impl ChannelsExportMessageLink")
}
