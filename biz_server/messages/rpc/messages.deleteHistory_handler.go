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
	"github.com/airwide-code/airwide.datacenter/biz/core/message"
	"github.com/airwide-code/airwide.datacenter/biz/base"
	"github.com/airwide-code/airwide.datacenter/biz_server/sync_client"
)

// messages.deleteHistory#1c015b09 flags:# just_clear:flags.0?true peer:InputPeer max_id:int = messages.AffectedHistory;
func (s *MessagesServiceImpl) MessagesDeleteHistory(ctx context.Context, request *mtproto.TLMessagesDeleteHistory) (*mtproto.Messages_AffectedHistory, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.deleteHistory#1c015b09 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesDeleteHistory logic

	peer := base.FromInputPeer2(md.UserId, request.GetPeer())
	if peer.PeerType == base.PEER_SELF {
		peer.PeerType = base.PEER_USER
	}
	boxIdList := message.GetMessageIdListByDialog(md.UserId, peer)
	if len(boxIdList) > 0 {
		// TOOD(@benqi): delete dialog message.
		message.DeleteByMessageIdList(md.UserId, boxIdList)

		updateDeleteMessages := mtproto.NewTLUpdateDeleteMessages()
		updateDeleteMessages.SetMessages(boxIdList)
		// updateDeleteMessages.SetMaxId(request.MaxId)
		_, err := sync_client.GetSyncClient().SyncOneUpdateData3(md.ServerId, md.AuthId, md.SessionId, md.UserId, md.ClientMsgId, updateDeleteMessages.To_Update())
		if err != nil {
			glog.Error(err)
			return nil, err
		}

		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NOTRETURN_CLIENT)
		glog.Infof("messages.deleteHistory#1c015b09 - reply: {%v}", err)
		return nil, err
	} else {
		affectedHistory := &mtproto.TLMessagesAffectedHistory{Data2: &mtproto.Messages_AffectedHistory_Data{
			Pts:      0,
			PtsCount: 0,
			Offset:   0,
		}}
		glog.Infof("messages.deleteHistory#1c015b09 - reply: {%v}", affectedHistory)
		return affectedHistory.To_Messages_AffectedHistory(), nil
	}
}
