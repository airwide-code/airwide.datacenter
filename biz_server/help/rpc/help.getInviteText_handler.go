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
)

// help.getInviteText#4d392343 = help.InviteText;
func (s *HelpServiceImpl) HelpGetInviteText(ctx context.Context, request *mtproto.TLHelpGetInviteText) (*mtproto.Help_InviteText, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("HelpGetInviteText - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	inviteText := &mtproto.TLHelpInviteText{Data2: &mtproto.Help_InviteText_Data{
		Message: "Invited by @benqi",
	}}

	glog.Infof("HelpGetInviteText - reply: %s", logger.JsonDebugData(inviteText))
	return inviteText.To_Help_InviteText(), nil
}
