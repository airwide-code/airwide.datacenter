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
)

// 客户端未使用

// contacts.resetSaved#879537f1 = Bool;
func (s *ContactsServiceImpl) ContactsResetSaved(ctx context.Context, request *mtproto.TLContactsResetSaved) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ContactsResetSaved - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): 客户端未调用此请求
	glog.Infof("ContactsResetSaved - reply: {true}")
	return mtproto.ToBool(true), nil
}
