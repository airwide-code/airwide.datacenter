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

// contacts.importCard#4fe196fe export_card:Vector<int> = User;
func (s *ContactsServiceImpl) ContactsImportCard(ctx context.Context, request *mtproto.TLContactsImportCard) (*mtproto.User, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ContactsImportCard - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ContactsImportCard logic

	return nil, fmt.Errorf("Not impl ContactsImportCard")
}
