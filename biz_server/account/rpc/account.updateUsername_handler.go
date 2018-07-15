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
	user2 "github.com/airwide-code/airwide.datacenter/biz/core/user"
	"github.com/airwide-code/airwide.datacenter/biz/core/account"
	"github.com/airwide-code/airwide.datacenter/baselib/base"
	"github.com/airwide-code/airwide.datacenter/biz_server/sync_client"
)

// account.updateUsername#3e0bdd7c username:string = User;
func (s *AccountServiceImpl) AccountUpdateUsername(ctx context.Context, request *mtproto.TLAccountUpdateUsername) (*mtproto.User, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.updateUsername#3e0bdd7c - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): wrapper checkUserName func

	// Check username format
	// You can choose a username on Telegram.
	// If you do, other people will be able to find
	// you by this username and contact you
	// without knowing your phone number.
	//
	// You can use a-z, 0-9 and underscores.
	// Minimum length is 5 characters.";
	//
	if len(request.Username) < kMinimumUserNameLen || !base.IsAlNumString(request.Username) {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USERNAME_INVALID)
		glog.Error("account.updateUsername#3e0bdd7c - format error: ", err)
		return nil, err
	} else {
		// userId == 0 为username不存在
		userId := account.GetUserIdByUserName(request.Username)
		// username不存在或者不是自身
		if userId > 0 && userId != md.UserId {
			err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USERNAME_OCCUPIED)
			glog.Error("account.updateUsername#3e0bdd7c - exists username: ", err)
			return nil, err
		}
	}

	// affected
	account.ChangeUserNameByUserId(md.UserId, request.Username)

	user := user2.GetUserById(md.UserId, md.UserId)
	// 要考虑到数据库主从同步问题
	user.SetUsername(request.Username)

	// sync to other sessions
	// updateUserName#a7332b73 user_id:int first_name:string last_name:string username:string = Update;
	updateUserName := &mtproto.TLUpdateUserName{Data2: &mtproto.Update_Data{
		UserId:    md.UserId,
		FirstName: user.GetFirstName(),
		LastName:  user.GetLastName(),
		Username:  request.Username,
	}}
	sync_client.GetSyncClient().PushToUserUpdateShortData(md.UserId, updateUserName.To_Update())

	// TODO(@benqi): push to other contacts

	glog.Infof("account.updateUsername#3e0bdd7c - reply: %s", logger.JsonDebugData(user))
	return user.To_User(), nil
}
