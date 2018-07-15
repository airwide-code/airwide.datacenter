/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package sync_client

import (
	"context"
	"github.com/golang/glog"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util/service_discovery"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util"
	"github.com/airwide-code/airwide.datacenter/mtproto"
)

type syncClient struct {
	client mtproto.RPCSyncClient
}

var (
	syncInstance = &syncClient{}
)

func GetSyncClient() *syncClient {
	return syncInstance
}

func InstallSyncClient(discovery *service_discovery.ServiceDiscoveryClientConfig) {
	conn, err := grpc_util.NewRPCClientByServiceDiscovery(discovery)

	if err != nil {
		glog.Error(err)
		panic(err)
	}

	syncInstance.client = mtproto.NewRPCSyncClient(conn)
}

func (c *syncClient) SyncOneUpdateData2(serverId int32, authKeyId, sessionId int64, pushUserId int32, clientMsgId int64, update *mtproto.Update) (reply *mtproto.ClientUpdatesState, err error) {
	updates := &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
		Updates: []*mtproto.Update{update},
	}}

	m := &mtproto.UpdatesRequest{
		PushType:    mtproto.SyncType_SYNC_TYPE_RPC_RESULT,
		ServerId:    serverId,
		AuthKeyId:   authKeyId,
		SessionId:   sessionId,
		PushUserId:  pushUserId,
		ClientMsgId: clientMsgId,
		Updates:     updates.To_Updates(),
		RpcResult:   &mtproto.RpcResultData{
			AffectedMessages: mtproto.NewTLMessagesAffectedMessages(),
		},
	}
	reply, err = c.client.SyncUpdatesData(context.Background(), m)
	return
}

func (c *syncClient) SyncOneUpdateData3(serverId int32, authKeyId, sessionId int64, pushUserId int32, clientMsgId int64, update *mtproto.Update) (reply *mtproto.ClientUpdatesState, err error) {
	updates := &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
		Updates: []*mtproto.Update{update},
	}}

	m := &mtproto.UpdatesRequest{
		PushType:    mtproto.SyncType_SYNC_TYPE_RPC_RESULT,
		ServerId:    serverId,
		AuthKeyId:   authKeyId,
		SessionId:   sessionId,
		PushUserId:  pushUserId,
		ClientMsgId: clientMsgId,
		Updates:     updates.To_Updates(),
		RpcResult:   &mtproto.RpcResultData{
			AffectedHistory: mtproto.NewTLMessagesAffectedHistory(),
		},
	}
	reply, err = c.client.SyncUpdatesData(context.Background(), m)
	return
}

func (c *syncClient) SyncOneUpdateData(authKeyId, sessionId int64, pushUserId int32, update *mtproto.Update) (reply *mtproto.ClientUpdatesState, err error) {
	updates := &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
		Updates: []*mtproto.Update{update},
	}}

	m := &mtproto.UpdatesRequest{
		PushType:   mtproto.SyncType_SYNC_TYPE_USER_NOTME,
		AuthKeyId:  authKeyId,
		SessionId:  sessionId,
		PushUserId: pushUserId,
		Updates:    updates.To_Updates(),
	}
	reply, err = c.client.SyncUpdatesData(context.Background(), m)
	return
}

func (c *syncClient) PushToUserNotMeOneUpdateData(authKeyId, sessionId int64, pushUserId int32, update *mtproto.Update) (reply *mtproto.VoidRsp, err error) {
	updates := &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
		Updates: []*mtproto.Update{update},
	}}

	m := &mtproto.UpdatesRequest{
		PushType:   mtproto.SyncType_SYNC_TYPE_USER_NOTME,
		AuthKeyId:  authKeyId,
		SessionId:  sessionId,
		PushUserId: pushUserId,
		Updates:    updates.To_Updates(),
	}
	reply, err = c.client.PushUpdatesData(context.Background(), m)
	return
}

func (c *syncClient) PushToUserMeOneUpdateData(authKeyId, sessionId int64, pushUserId int32, update *mtproto.Update) (reply *mtproto.VoidRsp, err error) {
	updates := &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
		Updates: []*mtproto.Update{update},
	}}

	m := &mtproto.UpdatesRequest{
		PushType:   mtproto.SyncType_SYNC_TYPE_USER_ME,
		AuthKeyId:  authKeyId,
		SessionId:  sessionId,
		PushUserId: pushUserId,
		Updates:    updates.To_Updates(),
	}
	reply, err = c.client.PushUpdatesData(context.Background(), m)
	return
}

func (c *syncClient) PushToUserOneUpdateData(pushUserId int32, update *mtproto.Update) (reply *mtproto.VoidRsp, err error) {
	updates := &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
		Updates: []*mtproto.Update{update},
	}}

	m := &mtproto.UpdatesRequest{
		PushType:   mtproto.SyncType_SYNC_TYPE_USER,
		// AuthKeyId:  authKeyId,
		// SessionId:  sessionId,
		PushUserId: pushUserId,
		Updates:    updates.To_Updates(),
	}
	reply, err = c.client.PushUpdatesData(context.Background(), m)
	return
}

func (c *syncClient) PushToUserNotMeUpdateShortData(authKeyId, sessionId int64, pushUserId int32, update *mtproto.Update) (reply *mtproto.VoidRsp, err error) {
	updates := &mtproto.TLUpdateShort{Data2: &mtproto.Updates_Data{
		Update: update,
	}}

	m := &mtproto.UpdatesRequest{
		PushType:   mtproto.SyncType_SYNC_TYPE_USER_NOTME,
		AuthKeyId:  authKeyId,
		SessionId:  sessionId,
		PushUserId: pushUserId,
		Updates:    updates.To_Updates(),
	}
	reply, err = c.client.PushUpdatesData(context.Background(), m)
	return
}

func (c *syncClient) PushToUserMeUpdateShortData(authKeyId, sessionId int64, pushUserId int32, update *mtproto.Update) (reply *mtproto.VoidRsp, err error) {
	updates := &mtproto.TLUpdateShort{Data2: &mtproto.Updates_Data{
		Update: update,
	}}

	m := &mtproto.UpdatesRequest{
		PushType:   mtproto.SyncType_SYNC_TYPE_USER_ME,
		AuthKeyId:  authKeyId,
		SessionId:  sessionId,
		PushUserId: pushUserId,
		Updates:    updates.To_Updates(),
	}
	reply, err = c.client.PushUpdatesData(context.Background(), m)
	return
}

func (c *syncClient) PushToUserUpdateShortData(pushUserId int32, update *mtproto.Update) (reply *mtproto.VoidRsp, err error) {
	updates := &mtproto.TLUpdateShort{Data2: &mtproto.Updates_Data{
		Update: update,
	}}

	m := &mtproto.UpdatesRequest{
		PushType:   mtproto.SyncType_SYNC_TYPE_USER,
		// AuthKeyId:  authKeyId,
		// SessionId:  sessionId,
		PushUserId: pushUserId,
		Updates:    updates.To_Updates(),
	}
	reply, err = c.client.PushUpdatesData(context.Background(), m)
	return
}

func (c *syncClient) SyncUpdatesData(authKeyId, sessionId int64, pushUserId int32, updates *mtproto.Updates) (reply *mtproto.ClientUpdatesState, err error) {
	m := &mtproto.UpdatesRequest{
		PushType:   mtproto.SyncType_SYNC_TYPE_USER_NOTME,
		AuthKeyId:  authKeyId,
		SessionId:  sessionId,
		PushUserId: pushUserId,
		Updates:    updates,
	}
	reply, err = c.client.SyncUpdatesData(context.Background(), m)
	return
}

func (c *syncClient) PushToUserNotMeUpdatesData(authKeyId, sessionId int64, pushUserId int32, updates *mtproto.Updates) (reply *mtproto.VoidRsp, err error) {
	m := &mtproto.UpdatesRequest{
		PushType:   mtproto.SyncType_SYNC_TYPE_USER_NOTME,
		AuthKeyId:  authKeyId,
		SessionId:  sessionId,
		PushUserId: pushUserId,
		Updates:    updates,
	}
	reply, err = c.client.PushUpdatesData(context.Background(), m)
	return
}

func (c *syncClient) PushToUserMeUpdatesData(authKeyId, sessionId int64, pushUserId int32, updates *mtproto.Updates) (reply *mtproto.VoidRsp, err error) {
	m := &mtproto.UpdatesRequest{
		PushType:   mtproto.SyncType_SYNC_TYPE_USER_ME,
		AuthKeyId:  authKeyId,
		SessionId:  sessionId,
		PushUserId: pushUserId,
		Updates:    updates,
	}
	reply, err = c.client.PushUpdatesData(context.Background(), m)
	return
}

func (c *syncClient) PushToUserUpdatesData(pushUserId int32, updates *mtproto.Updates) (reply *mtproto.VoidRsp, err error) {
	m := &mtproto.UpdatesRequest{
		PushType:   mtproto.SyncType_SYNC_TYPE_USER,
		// AuthKeyId:  authKeyId,
		// SessionId:  sessionId,
		PushUserId: pushUserId,
		Updates:    updates,
	}
	reply, err = c.client.PushUpdatesData(context.Background(), m)
	return
}
