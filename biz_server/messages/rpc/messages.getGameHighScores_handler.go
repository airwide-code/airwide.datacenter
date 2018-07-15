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

// messages.getGameHighScores#e822649d peer:InputPeer id:int user_id:InputUser = messages.HighScores;
func (s *MessagesServiceImpl) MessagesGetGameHighScores(ctx context.Context, request *mtproto.TLMessagesGetGameHighScores) (*mtproto.Messages_HighScores, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesGetGameHighScores - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesGetGameHighScores logic

	return nil, fmt.Errorf("Not impl MessagesGetGameHighScores")
}
