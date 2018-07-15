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

// 客户端未使用

// account.updateDeviceLocked#38df3532 period:int = Bool;
func (s *AccountServiceImpl) AccountUpdateDeviceLocked(ctx context.Context, request *mtproto.TLAccountUpdateDeviceLocked) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AccountUpdateDeviceLocked - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl AccountUpdateDeviceLocked logic

	return nil, fmt.Errorf("Not impl AccountUpdateDeviceLocked")
}
