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

// auth.sendInvites#771c1d97 phone_numbers:Vector<string> message:string = Bool;
func (s *AuthServiceImpl) AuthSendInvites(ctx context.Context, request *mtproto.TLAuthSendInvites) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AuthSendInvites - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl AuthSendInvites logic

	return nil, fmt.Errorf("Not impl AuthSendInvites")
}
