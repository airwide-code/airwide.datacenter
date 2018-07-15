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

// contacts.exportCard#84e53737 = Vector<int>;
func (s *ContactsServiceImpl) ContactsExportCard(ctx context.Context, request *mtproto.TLContactsExportCard) (*mtproto.VectorInt, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ContactsExportCard - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ContactsExportCard logic

	return nil, fmt.Errorf("Not impl ContactsExportCard")
}
