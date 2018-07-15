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
	"time"
	"github.com/airwide-code/airwide.datacenter/biz/base"
	"github.com/airwide-code/airwide.datacenter/biz/core/user"
)

func makeDraftMessageBySaveDraft(request *mtproto.TLMessagesSaveDraft) *mtproto.TLDraftMessage {
	return &mtproto.TLDraftMessage{ Data2: &mtproto.DraftMessage_Data{
		NoWebpage:    request.GetNoWebpage(),
		ReplyToMsgId: request.GetReplyToMsgId(),
		Message:      request.GetMessage(),
		Entities:     request.GetEntities(),
		Date:         int32(time.Now().Unix()),
	}}
}

// messages.saveDraft#bc39e14b flags:# no_webpage:flags.1?true reply_to_msg_id:flags.0?int peer:InputPeer message:string entities:flags.3?Vector<MessageEntity> = Bool;
func (s *MessagesServiceImpl) MessagesSaveDraft(ctx context.Context, request *mtproto.TLMessagesSaveDraft) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.saveDraft#bc39e14b - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		peer *base.PeerUtil
	)

	if request.GetPeer().GetConstructor() ==  mtproto.TLConstructor_CRC32_inputPeerSelf {
		peer = &base.PeerUtil{PeerType: base.PEER_USER, PeerId: md.UserId}
	} else {
		peer = base.FromInputPeer(request.GetPeer())
	}

	draft := makeDraftMessageBySaveDraft(request)

	// TODO(@benqi): 会话未存在如何处理？
	user.SaveDraftMessage(md.UserId, peer.PeerType, peer.PeerId, draft.To_DraftMessage())

	// TODO(@benqi): sync other client

	reply := mtproto.ToBool(true)

	glog.Infof("messages.saveDraft#bc39e14b - reply: {%v}", reply)
	return reply, nil
}
