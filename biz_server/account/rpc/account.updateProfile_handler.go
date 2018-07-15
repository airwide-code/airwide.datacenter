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
	"github.com/airwide-code/airwide.datacenter/biz_server/sync_client"
)

// account.updateProfile#78515775 flags:# first_name:flags.0?string last_name:flags.1?string about:flags.2?string = User;
func (s *AccountServiceImpl) AccountUpdateProfile(ctx context.Context, request *mtproto.TLAccountUpdateProfile) (*mtproto.User, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.updateProfile#78515775 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Check first_name and last_name invalid. has err: FIRSTNAME_INVALID, LASTNAME_INVALID

	// Check format
	// about长度<70并且可以为emtpy
	// first_name必须有值
	if len(request.FirstName) > 0 && len(request.About) > 0 {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_FIRSTNAME_INVALID)
		glog.Error(err)
		return nil, err
	}

	user := user2.GetUserById(md.UserId, md.UserId)

	if len(request.FirstName) > 0 {
		account.UpdateFirstAndLastName(md.UserId, request.FirstName, request.LastName)

		// return new first_name and last_name.
		user.SetFirstName(request.FirstName)
		user.SetLastName(request.LastName)
	} else {
		account.UpdateAbout(md.UserId, request.About)
	}

	// sync to other sessions
	// updateUserName#a7332b73 user_id:int first_name:string last_name:string username:string = Update;
	updateUserName := &mtproto.TLUpdateUserName{Data2: &mtproto.Update_Data{
		UserId:    md.UserId,
		FirstName: user.GetFirstName(),
		LastName:  user.GetLastName(),
		Username:  user.GetUsername(),
	}}

	sync_client.GetSyncClient().PushToUserUpdateShortData(md.UserId, updateUserName.To_Update())
	// TODO(@benqi): push to other contacts

	glog.Infof("account.updateProfile#78515775 - reply: {%v}", user)
	return user.To_User(), nil
}
