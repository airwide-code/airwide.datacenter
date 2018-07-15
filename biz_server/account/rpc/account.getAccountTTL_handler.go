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

// account.getAccountTTL#8fc711d = AccountDaysTTL;
func (s *AccountServiceImpl) AccountGetAccountTTL(ctx context.Context, request *mtproto.TLAccountGetAccountTTL) (*mtproto.AccountDaysTTL, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.getAccountTTL#8fc711d - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	days := account.GetAccountDaysTTL(md.UserId)
	ttl := &mtproto.TLAccountDaysTTL{ Data2: &mtproto.AccountDaysTTL_Data{
		Days: days,
	}}

	glog.Infof("account.getAccountTTL#8fc711d - reply: %s\n", logger.JsonDebugData(ttl))
	return ttl.To_AccountDaysTTL(), nil
}
