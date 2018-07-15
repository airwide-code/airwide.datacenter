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

// payments.validateRequestedInfo#770a8e74 flags:# save:flags.0?true msg_id:int info:PaymentRequestedInfo = payments.ValidatedRequestedInfo;
func (s *PaymentsServiceImpl) PaymentsValidateRequestedInfo(ctx context.Context, request *mtproto.TLPaymentsValidateRequestedInfo) (*mtproto.Payments_ValidatedRequestedInfo, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("PaymentsValidateRequestedInfo - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl PaymentsValidateRequestedInfo logic

	return nil, fmt.Errorf("Not impl PaymentsValidateRequestedInfo")
}
