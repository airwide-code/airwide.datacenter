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
	"github.com/airwide-code/airwide.datacenter/biz/core/user"
)

// account.getPrivacy#dadbc950 key:InputPrivacyKey = account.PrivacyRules;
func (s *AccountServiceImpl) AccountGetPrivacy(ctx context.Context, request *mtproto.TLAccountGetPrivacy) (*mtproto.Account_PrivacyRules, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.getPrivacy#dadbc950 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	privacyLogic := account.MakePrivacyLogic(md.UserId)
	rulesData := privacyLogic.GetPrivacy(account.FromInputPrivacyKey(request.Key))

	var rules *mtproto.TLAccountPrivacyRules
	if rulesData == nil {
		// TODO(@benqi): return nil or empty
		// rules = mtproto.NewTLAccountPrivacyRules()
		rules = &mtproto.TLAccountPrivacyRules{ Data2: &mtproto.Account_PrivacyRules_Data{
			Rules: []*mtproto.PrivacyRule{mtproto.NewTLPrivacyValueAllowAll().To_PrivacyRule()},
		}}
	} else {
		idList := rulesData.PickAllUserIdList()
		if len(idList) == 0 {
			rules = &mtproto.TLAccountPrivacyRules{ Data2: &mtproto.Account_PrivacyRules_Data{
				Rules: rulesData.ToPrivacyRuleList(),
			}}
		} else {
			rules = &mtproto.TLAccountPrivacyRules{ Data2: &mtproto.Account_PrivacyRules_Data{
				Rules: rulesData.ToPrivacyRuleList(),
				Users: user.GetUsersBySelfAndIDList(md.UserId, idList),
			}}
		}
	}

	glog.Infof("account.getPrivacy#dadbc950 - reply: %s", logger.JsonDebugData(rules))
	return rules.To_Account_PrivacyRules(), nil
}
