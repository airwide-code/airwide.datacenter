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
	"github.com/airwide-code/airwide.datacenter/biz/core/channel"
)

// channels.checkUsername#10e6bd2c channel:InputChannel username:string = Bool;
func (s *ChannelsServiceImpl) ChannelsCheckUsername(ctx context.Context, request *mtproto.TLChannelsCheckUsername) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("channels.checkUsername#10e6bd2c - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var checked bool
	if request.GetChannel().GetConstructor() == mtproto.TLConstructor_CRC32_inputChannelEmpty {
		checked = false
	} else {
		checked = channel.CheckChannelUserName(request.GetChannel().GetData2().GetChannelId(), request.GetUsername())
	}

	glog.Infof("channels.checkUsername#10e6bd2c - reply: {%v}", checked)
	return mtproto.ToBool(checked), nil
}
