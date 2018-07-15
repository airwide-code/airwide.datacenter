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
	"github.com/airwide-code/airwide.datacenter/biz/core/account"
)

/*
	selfUser: hash = 0, flag = 1
	other:  hash and flag load from db
 */

// account.getAuthorizations#e320c158 = account.Authorizations;
func (s *AccountServiceImpl) AccountGetAuthorizations(ctx context.Context, request *mtproto.TLAccountGetAuthorizations) (*mtproto.Account_Authorizations, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.getAuthorizations#e320c158 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	sessionList := account.GetAuthorizationList(md.AuthId, md.UserId)
	authorizations := &mtproto.TLAccountAuthorizations{Data2: &mtproto.Account_Authorizations_Data{
		Authorizations: sessionList,
	}}

	glog.Infof("account.getAuthorizations#e320c158 - reply: {%s}", logger.JsonDebugData(authorizations))
	return authorizations.To_Account_Authorizations(), nil
}
