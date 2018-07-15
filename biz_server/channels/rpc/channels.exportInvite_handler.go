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

// channels.exportInvite#c7560885 channel:InputChannel = ExportedChatInvite;
func (s *ChannelsServiceImpl) ChannelsExportInvite(ctx context.Context, request *mtproto.TLChannelsExportInvite) (*mtproto.ExportedChatInvite, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("channels.exportInvite#c7560885 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	if request.Channel.Constructor == mtproto.TLConstructor_CRC32_inputChannelEmpty {
		// TODO(@benqi): chatUser不能是inputUser和inputUserSelf
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error("channels.exportInvite#c7560885 - error: ", err, "; InputPeer invalid")
		return nil, err
	}

	channelLogic, err := channel.NewChannelLogicById(request.GetChannel().GetData2().GetChannelId())
	if err != nil {

	}

	exportedChatInvite := &mtproto.TLChatInviteExported{Data2: &mtproto.ExportedChatInvite_Data{
		Link: channelLogic.ExportedChatInvite(),
	}}

	glog.Infof("channels.exportInvite#c7560885 - reply: {%v}", exportedChatInvite)
	return exportedChatInvite.To_ExportedChatInvite(), nil
}
