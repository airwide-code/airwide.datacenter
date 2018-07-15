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

// account.setAccountTTL#2442485e ttl:AccountDaysTTL = Bool;
func (s *AccountServiceImpl) AccountSetAccountTTL(ctx context.Context, request *mtproto.TLAccountSetAccountTTL) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AccountSetAccountTTL - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Check ttl
	ttl := request.GetTtl().GetData2().GetDays()
	if ttl <= 0 || ttl > 365 {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error("ttl_days error: ", err)
		return nil, err
	}
	account.SetAccountDaysTTL(md.UserId, ttl)

	glog.Infof("account.setAccountTTL#2442485e - reply: {true}")
	return mtproto.ToBool(true), nil
}
