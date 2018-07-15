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
	"github.com/airwide-code/airwide.datacenter/biz/dal/dao"
	"github.com/airwide-code/airwide.datacenter/biz_server/sync_client"
)

// messages.deleteMessages#e58e95d2 flags:# revoke:flags.0?true id:Vector<int> = messages.AffectedMessages;
func (s *MessagesServiceImpl) MessagesDeleteMessages(ctx context.Context, request *mtproto.TLMessagesDeleteMessages) (*mtproto.Messages_AffectedMessages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.deleteMessages#e58e95d2 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		deleteIdList = request.GetId()
	)

	deleteMessages := &mtproto.TLUpdateDeleteMessages{Data2: &mtproto.Update_Data{
		Messages: deleteIdList,
	}}

	state, err := sync_client.GetSyncClient().SyncOneUpdateData(md.AuthId, md.SessionId, md.UserId, deleteMessages.To_Update())
	if err != nil {
		return nil, err
	}

	affectedMessages := &mtproto.TLMessagesAffectedMessages{Data2: &mtproto.Messages_AffectedMessages_Data{
		Pts:      state.Pts,
		PtsCount: state.PtsCount,
	}}

	// 1. delete messages
	// 2. updateTopMessage
	if request.GetRevoke() {
		//  消息撤回
		doList := dao.GetMessagesDAO(dao.DB_SLAVE).SelectPeerDialogMessageIdList(md.UserId, request.GetId())
		deleteIdListMap := make(map[int32][]int32)
		for _, do := range doList {
			if messageIdList, ok := deleteIdListMap[do.UserId]; !ok {
				deleteIdListMap[do.UserId] = []int32{do.UserMessageBoxId}
			} else {
				deleteIdListMap[do.UserId] = append(messageIdList, do.UserMessageBoxId)
			}
		}

		glog.Info("messages.deleteMessages#e58e95d2 - deleteIdListMap: ", deleteIdListMap)
		for k, v := range deleteIdListMap {
			deleteMessages := &mtproto.TLUpdateDeleteMessages{Data2: &mtproto.Update_Data{
				Messages: v,
			}}
			sync_client.GetSyncClient().PushToUserOneUpdateData(k, deleteMessages.To_Update())
			dao.GetMessagesDAO(dao.DB_MASTER).DeleteMessagesByMessageIdList(k, v)
		}
		dao.GetMessagesDAO(dao.DB_MASTER).DeleteMessagesByMessageIdList(md.UserId, deleteIdList)
		// TODO(@benqi): 更新dialog的TopMessage
	} else {
		// 删除消息
		dao.GetMessagesDAO(dao.DB_MASTER).DeleteMessagesByMessageIdList(md.UserId, deleteIdList)

		// TODO(@benqi): 更新dialog的TopMessage
	}

	glog.Infof("messages.deleteMessages#e58e95d2 - reply: %s", logger.JsonDebugData(affectedMessages))
	return affectedMessages.To_Messages_AffectedMessages(), nil
}
