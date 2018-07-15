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

// phone.setCallRating#1c536a34 peer:InputPhoneCall rating:int comment:string = Updates;
func (s *PhoneServiceImpl) PhoneSetCallRating(ctx context.Context, request *mtproto.TLPhoneSetCallRating) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("PhoneSetCallRating - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl PhoneSetCallRating logic

	return nil, fmt.Errorf("Not impl PhoneSetCallRating")
}
