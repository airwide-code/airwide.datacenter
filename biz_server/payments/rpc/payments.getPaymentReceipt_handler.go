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

// payments.getPaymentReceipt#a092a980 msg_id:int = payments.PaymentReceipt;
func (s *PaymentsServiceImpl) PaymentsGetPaymentReceipt(ctx context.Context, request *mtproto.TLPaymentsGetPaymentReceipt) (*mtproto.Payments_PaymentReceipt, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("PaymentsGetPaymentReceipt - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl PaymentsGetPaymentReceipt logic

	return nil, fmt.Errorf("Not impl PaymentsGetPaymentReceipt")
}
