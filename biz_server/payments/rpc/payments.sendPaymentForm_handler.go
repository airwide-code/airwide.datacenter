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

// payments.sendPaymentForm#2b8879b3 flags:# msg_id:int requested_info_id:flags.0?string shipping_option_id:flags.1?string credentials:InputPaymentCredentials = payments.PaymentResult;
func (s *PaymentsServiceImpl) PaymentsSendPaymentForm(ctx context.Context, request *mtproto.TLPaymentsSendPaymentForm) (*mtproto.Payments_PaymentResult, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("PaymentsSendPaymentForm - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl PaymentsSendPaymentForm logic

	return nil, fmt.Errorf("Not impl PaymentsSendPaymentForm")
}
