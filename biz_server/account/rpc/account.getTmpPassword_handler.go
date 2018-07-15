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
	"time"
)

// account.getTmpPassword#4a82327e password_hash:bytes period:int = account.TmpPassword;
func (s *AccountServiceImpl) AccountGetTmpPassword(ctx context.Context, request *mtproto.TLAccountGetTmpPassword) (*mtproto.Account_TmpPassword, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AccountGetTmpPassword - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Check password_hash invalid, android source code
	// byte[] hash = new byte[currentPassword.current_salt.length * 2 + passwordBytes.length];
	// System.arraycopy(currentPassword.current_salt, 0, hash, 0, currentPassword.current_salt.length);
	// System.arraycopy(passwordBytes, 0, hash, currentPassword.current_salt.length, passwordBytes.length);
	// System.arraycopy(currentPassword.current_salt, 0, hash, hash.length - currentPassword.current_salt.length, currentPassword.current_salt.length);

	// account.tmpPassword#db64fd34 tmp_password:bytes valid_until:int = account.TmpPassword;
	tmpPassword := mtproto.NewTLAccountTmpPassword()
	tmpPassword.SetTmpPassword([]byte("01234567899876543210"))
	tmpPassword.SetValidUntil(int32(time.Now().Unix()) + request.Period)

	glog.Infof("AccountServiceImpl - reply: %s", logger.JsonDebugData(tmpPassword))
	return tmpPassword.To_Account_TmpPassword(), nil
}
