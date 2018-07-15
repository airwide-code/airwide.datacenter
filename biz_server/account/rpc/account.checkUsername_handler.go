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
	"github.com/airwide-code/airwide.datacenter/baselib/base"
	"github.com/airwide-code/airwide.datacenter/biz/core/account"
)

const (
	kMinimumUserNameLen = 5
)

// account.checkUsername#2714d86c username:string = Bool;
func (s *AccountServiceImpl) AccountCheckUsername(ctx context.Context, request *mtproto.TLAccountCheckUsername) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.checkUsername#2714d86c - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

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
		glog.Error("account.checkUsername#2714d86c - format error: ", err)
		return nil, err
	} else {
		// userId == 0 为username不存在
		userId := account.GetUserIdByUserName(request.Username)
		// username不存在或者不是自身
		if userId > 0 && userId != md.UserId {
			err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USERNAME_OCCUPIED)
			glog.Error("account.checkUsername#2714d86c - exists username: ", err)
			return nil, err
		}
	}

	glog.Infof("account.checkUsername#2714d86c - reply: {true}")
	return mtproto.ToBool(true), nil
}
