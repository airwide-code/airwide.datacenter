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

// account.confirmPhone#5f2178c3 phone_code_hash:string phone_code:string = Bool;
func (s *AccountServiceImpl) AccountConfirmPhone(ctx context.Context, request *mtproto.TLAccountConfirmPhone) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AccountConfirmPhone - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl AccountConfirmPhone logic

	return nil, fmt.Errorf("Not impl AccountConfirmPhone")
}
