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
	"github.com/airwide-code/airwide.datacenter/biz/core/channel"
	"github.com/airwide-code/airwide.datacenter/biz/dal/dao"
	"github.com/airwide-code/airwide.datacenter/biz/core/user"
)

// channels.getParticipant#546dd7a6 channel:InputChannel user_id:InputUser = channels.ChannelParticipant;
func (s *ChannelsServiceImpl) ChannelsGetParticipant(ctx context.Context, request *mtproto.TLChannelsGetParticipant) (*mtproto.Channels_ChannelParticipant, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("channels.getParticipant#546dd7a6 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	if request.Channel.Constructor == mtproto.TLConstructor_CRC32_inputChannelEmpty {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error("channels.exportInvite#c7560885 - error: ", err, "; InputPeer invalid")
		return nil, err
	}

	var userId = md.UserId
	if request.UserId.GetConstructor() == mtproto.TLConstructor_CRC32_inputUserEmpty {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error("channels.exportInvite#c7560885 - error: ", err, "; InputPeer invalid")
		return nil, err
	} else if request.UserId.GetConstructor() == mtproto.TLConstructor_CRC32_inputUser {
		userId = request.UserId.GetData2().GetUserId()
	}

	// GetData2().GetUserId()
	inputChannel := request.GetChannel().To_InputChannel()
	do := dao.GetChannelParticipantsDAO(dao.DB_SLAVE).SelectByUserId(inputChannel.GetChannelId(), userId)
	if do == nil {
		err := fmt.Errorf("not find userId in (%v, %d)", inputChannel, userId)
		glog.Error(err)
		return nil, err
	}

	channelParticipant := &mtproto.TLChannelsChannelParticipant{Data2: &mtproto.Channels_ChannelParticipant_Data{
		Participant: channel.MakeChannelParticipant2ByDO(md.UserId, do),
		Users: user.GetUsersBySelfAndIDList(md.UserId, []int32{do.UserId, do.InviterUserId}),
	}}

	glog.Infof("channels.getParticipant#546dd7a6 - reply: {%v}", channelParticipant)
	return channelParticipant.To_Channels_ChannelParticipant(), nil
}
