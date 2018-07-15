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

const (
	// TODO(@benqi): add support user.
	kSupportUserID = int32(2)
)

// help.getSupport#9cdf08cd = help.Support;
func (s *HelpServiceImpl) HelpGetSupport(ctx context.Context, request *mtproto.TLHelpGetSupport) (*mtproto.Help_Support, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("help.getSupport#9cdf08cd - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl HelpGetSupport logic
	reply := &mtproto.TLHelpSupport{ Data2: &mtproto.Help_Support_Data{
		PhoneNumber: "+86 111 1111 1111",
		User:        &mtproto.User{Constructor: mtproto.TLConstructor_CRC32_userEmpty, Data2: &mtproto.User_Data{Id: kSupportUserID}},
	}}

	glog.Infof("help.getSupport#9cdf08cd - reply: {%v}\n", reply)
	return reply.To_Help_Support(), nil
}
