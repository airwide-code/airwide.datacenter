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

// payments.getSavedInfo#227d824b = payments.SavedInfo;
func (s *PaymentsServiceImpl) PaymentsGetSavedInfo(ctx context.Context, request *mtproto.TLPaymentsGetSavedInfo) (*mtproto.Payments_SavedInfo, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("PaymentsGetSavedInfo - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl PaymentsGetSavedInfo logic

	return nil, fmt.Errorf("Not impl PaymentsGetSavedInfo")
}
