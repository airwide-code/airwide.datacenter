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

// account.sendConfirmPhoneCode#1516d7bd flags:# allow_flashcall:flags.0?true hash:string current_number:flags.0?Bool = auth.SentCode;
func (s *AccountServiceImpl) AccountSendConfirmPhoneCode(ctx context.Context, request *mtproto.TLAccountSendConfirmPhoneCode) (*mtproto.Auth_SentCode, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AccountSendConfirmPhoneCode - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl AccountSendConfirmPhoneCode logic

	return nil, fmt.Errorf("Not impl AccountSendConfirmPhoneCode")
}
