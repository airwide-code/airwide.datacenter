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

/*
 // account.noPassword#96dabc18 new_salt:bytes email_unconfirmed_pattern:string = account.Password;
 // account.password#7c18141c current_salt:bytes new_salt:bytes hint:string has_recovery:Bool email_unconfirmed_pattern:string = account.Password;
 */

// account.getPassword#548a30f5 = account.Password;
func (s *AccountServiceImpl) AccountGetPassword(ctx context.Context, request *mtproto.TLAccountGetPassword) (*mtproto.Account_Password, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.getPassword#548a30f5 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	passwordLogic, err := account.MakePasswordData(md.UserId)
	if err != nil {
		glog.Error("account.getPassword#548a30f5 - error: ", err)
		return nil, err
	}

	// DoGetPassword
	password := passwordLogic.GetPassword()

	glog.Infof("account.getPassword#548a30f5 - reply: %s", logger.JsonDebugData(password))
	return password, nil
}
