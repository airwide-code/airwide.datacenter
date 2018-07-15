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

// channels.getFullChannel#8736a09 channel:InputChannel = messages.ChatFull;
func (s *ChannelsServiceImpl) ChannelsGetFullChannel(ctx context.Context, request *mtproto.TLChannelsGetFullChannel) (*mtproto.Messages_ChatFull, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("channels.getFullChannel#8736a09 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	if request.Channel.Constructor == mtproto.TLConstructor_CRC32_inputChannelEmpty {
		// TODO(@benqi): chatUser不能是inputUser和inputUserSelf
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error("channels.exportInvite#c7560885 - error: ", err, "; InputPeer invalid")
		return nil, err
	}

	inputChannel := request.GetChannel().To_InputChannel()

	channelLogic, err := channel.NewChannelLogicById(inputChannel.GetChannelId())
	if err != nil {
		glog.Error("channels.getFullChannel#8736a09 - error: ", err)
		return nil, err
	}

	// idList := channelLogic.GetChannelParticipantIdList()
	messagesChatFull := &mtproto.TLMessagesChatFull{Data2: &mtproto.Messages_ChatFull_Data{
		FullChat: 	channel.GetChannelFullBySelfId(md.UserId, channelLogic).To_ChatFull(),
		Chats:      []*mtproto.Chat{channelLogic.ToChannel(md.UserId)},
		Users: 		[]*mtproto.User{},
	}}

	glog.Infof("channels.getFullChannel#8736a09 - reply: %s", logger.JsonDebugData(messagesChatFull))
	return messagesChatFull.To_Messages_ChatFull(), nil

}
