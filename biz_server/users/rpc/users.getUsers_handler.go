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
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util"
	"github.com/airwide-code/airwide.datacenter/baselib/logger"
	"github.com/airwide-code/airwide.datacenter/biz/core/user"
	"github.com/airwide-code/airwide.datacenter/mtproto"
	"golang.org/x/net/context"
)

// users.getUsers#d91a548 id:Vector<InputUser> = Vector<User>;
func (s *UsersServiceImpl) UsersGetUsers(ctx context.Context, request *mtproto.TLUsersGetUsers) (*mtproto.Vector_User, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("UsersGetUsers - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl UsersGetUsers logic
	userList := &mtproto.Vector_User{
		Datas: make([]*mtproto.User, 0, len(request.Id)),
	}

	for _, inputUser := range request.Id {
		switch inputUser.GetConstructor() {
		case mtproto.TLConstructor_CRC32_inputUserSelf:
			userData := user.GetUserById(md.GetUserId(), md.GetUserId())
			userList.Datas = append(userList.Datas, userData.To_User())
		case mtproto.TLConstructor_CRC32_inputUser:
			userData := user.GetUserById(md.GetUserId(), inputUser.GetData2().GetUserId())
			userList.Datas = append(userList.Datas, userData.To_User())
		case mtproto.TLConstructor_CRC32_inputUserEmpty:
		}
	}

	glog.Infof("users.getUsers#d91a548 - reply: ", logger.JsonDebugData(userList))
	return userList, nil
}
