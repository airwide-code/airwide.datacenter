/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package rpc

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/airwide-code/airwide.datacenter/baselib/logger"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util"
	"github.com/airwide-code/airwide.datacenter/mtproto"
	"golang.org/x/net/context"
)

// account.changePhone#70c32edb phone_number:string phone_code_hash:string phone_code:string = User;
func (s *AccountServiceImpl) AccountChangePhone(ctx context.Context, request *mtproto.TLAccountChangePhone) (*mtproto.User, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AccountChangePhone - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl AccountChangePhone logic

	return nil, fmt.Errorf("Not impl AccountChangePhone")
}
