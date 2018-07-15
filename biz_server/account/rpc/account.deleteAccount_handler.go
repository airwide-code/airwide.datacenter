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
	"github.com/airwide-code/airwide.datacenter/biz/dal/dao"
	base2 "github.com/airwide-code/airwide.datacenter/baselib/base"
)

/*
  reset my account?
  delete成功，转到注册页面，失败处理:

	if (error.text.equals("2FA_RECENT_CONFIRM")) {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("ResetAccountCancelledAlert", R.string.ResetAccountCancelledAlert));
	} else if (error.text.startsWith("2FA_CONFIRM_WAIT_")) {
		Bundle params = new Bundle();
		params.putString("phoneFormated", requestPhone);
		params.putString("phoneHash", phoneHash);
		params.putString("code", phoneCode);
		params.putInt("startTime", ConnectionsManager.getInstance().getCurrentTime());
		params.putInt("waitTime", Utilities.parseInt(error.text.replace("2FA_CONFIRM_WAIT_", "")));
		setPage(8, true, params, false);
	} else {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), error.text);
	}
 */

// account.deleteAccount#418d4e0b reason:string = Bool;
func (s *AccountServiceImpl) AccountDeleteAccount(ctx context.Context, request *mtproto.TLAccountDeleteAccount) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AccountDeleteAccount - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl AccountDeleteAccount logic
	affected := dao.GetUsersDAO(dao.DB_MASTER).Delete(
		request.GetReason(),
		base2.NowFormatYMDHMS(),
		md.UserId)

	deletedOk := affected == 1
	// TODO(@benqi): 1. Clear account data 2. Kickoff other client

	glog.Infof("AccountDeleteAccount - reply: {%v}", deletedOk)
	return mtproto.ToBool(deletedOk), nil
}
