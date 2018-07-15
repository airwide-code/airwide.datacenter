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
	"github.com/airwide-code/airwide.datacenter/biz_server/sync_client"
	"github.com/airwide-code/airwide.datacenter/biz/core/user"
)

// account.updateStatus#6628562c offline:Bool = Bool;
func (s *AccountServiceImpl) AccountUpdateStatus(ctx context.Context, request *mtproto.TLAccountUpdateStatus) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.updateStatus#6628562c - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var status *mtproto.UserStatus

	offline := mtproto.FromBool(request.GetOffline())
	if offline {
		// pc端：离开应用程序激活状态（点击其他应用程序）
		statusOffline := &mtproto.TLUserStatusOffline{Data2: &mtproto.UserStatus_Data{
			WasOnline: int32(time.Now().Unix()),
		}}
		status = statusOffline.To_UserStatus()
	} else {
		// pc端：客户端应用程序激活（点击客户端窗口）
		now := time.Now().Unix()
		statusOnline := &mtproto.TLUserStatusOnline{Data2: &mtproto.UserStatus_Data{
			Expires: int32(now + 5*30),
		}}
		status = statusOnline.To_UserStatus()
		user.UpdateUserStatus(md.UserId, now)
	}

	updateUserStatus := &mtproto.TLUpdateUserStatus{Data2: &mtproto.Update_Data{
		UserId: md.UserId,
		Status: status,
	}}
	updates := &mtproto.TLUpdateShort{ Data2: &mtproto.Updates_Data{
		Update: updateUserStatus.To_Update(),
		Date:  int32(time.Now().Unix()),
	}}

	// push to other contacts.
	contactIDList := user.GetContactUserIDList(md.UserId)
	for _, id := range contactIDList {
		sync_client.GetSyncClient().PushToUserUpdatesData(id, updates.To_Updates())
	}

	glog.Infof("account.updateStatus#6628562c - reply: {true}")
	return mtproto.ToBool(true), nil
}
