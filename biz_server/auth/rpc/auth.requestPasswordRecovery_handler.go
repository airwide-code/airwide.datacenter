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
	"github.com/airwide-code/airwide.datacenter/biz/core/account"
)

// Forgot password?

// auth.requestPasswordRecovery#d897bc66 = auth.PasswordRecovery;
func (s *AuthServiceImpl) AuthRequestPasswordRecovery(ctx context.Context, request *mtproto.TLAuthRequestPasswordRecovery) (*mtproto.Auth_PasswordRecovery, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("auth.requestPasswordRecovery#d897bc66 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	passwordLogic, err := account.MakePasswordData(md.UserId)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	passwordRecovery, err := passwordLogic.RequestPasswordRecovery()
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	glog.Infof("auth.requestPasswordRecovery#d897bc66 - reply: %s\n", logger.JsonDebugData(passwordRecovery))
	return passwordRecovery, nil
}
