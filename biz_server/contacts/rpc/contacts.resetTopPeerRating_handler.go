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

// contacts.resetTopPeerRating#1ae373ac category:TopPeerCategory peer:InputPeer = Bool;
func (s *ContactsServiceImpl) ContactsResetTopPeerRating(ctx context.Context, request *mtproto.TLContactsResetTopPeerRating) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ContactsResetTopPeerRating - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	//// TODO(@benqi): Impl ContactsResetTopPeerRating logic
	//_ = helper.FromInputPeer(request.Peer)
	//
	//// TODO(@benqi): 看看客户端代码，什么情况会调用
	//switch request.GetCategory().GetPayload().(type) {
	//case *mtproto.TopPeerCategory_TopPeerCategoryBotsPM:
	//case *mtproto.TopPeerCategory_TopPeerCategoryBotsInline:
	//case *mtproto.TopPeerCategory_TopPeerCategoryCorrespondents:
	//case *mtproto.TopPeerCategory_TopPeerCategoryGroups:
	//case *mtproto.TopPeerCategory_TopPeerCategoryChannels:
	//case *mtproto.TopPeerCategory_TopPeerCategoryPhoneCalls:
	//}

	glog.Infof("ContactsResetTopPeerRating - reply: {true}")
	return mtproto.ToBool(true), nil
}
