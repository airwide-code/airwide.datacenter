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

// payments.getPaymentForm#99f09745 msg_id:int = payments.PaymentForm;
func (s *PaymentsServiceImpl) PaymentsGetPaymentForm(ctx context.Context, request *mtproto.TLPaymentsGetPaymentForm) (*mtproto.Payments_PaymentForm, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("PaymentsGetPaymentForm - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl PaymentsGetPaymentForm logic

	return nil, fmt.Errorf("Not impl PaymentsGetPaymentForm")
}
