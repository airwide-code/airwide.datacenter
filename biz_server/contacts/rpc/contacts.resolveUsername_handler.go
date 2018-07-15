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
	"github.com/airwide-code/airwide.datacenter/biz/core/user"
)

// contacts.resolveUsername#f93ccba3 username:string = contacts.ResolvedPeer;
func (s *ContactsServiceImpl) ContactsResolveUsername(ctx context.Context, request *mtproto.TLContactsResolveUsername) (*mtproto.Contacts_ResolvedPeer, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("contacts.resolveUsername#f93ccba3 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ContactsResolveUsername logic
	do := dao.GetUsersDAO(dao.DB_SLAVE).SelectByUsername(request.GetUsername())
	if do == nil {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USERNAME_INVALID)
		glog.Error(err)
		return nil, err
	}

	peer := &mtproto.TLPeerUser{Data2: &mtproto.Peer_Data{
		UserId: do.Id,
	}}
	resolvedPeer := &mtproto.TLContactsResolvedPeer{Data2: &mtproto.Contacts_ResolvedPeer_Data{
		Peer:  peer.To_Peer(),
		Chats: []*mtproto.Chat{},
		Users: []*mtproto.User{user.GetUserById(md.UserId, do.Id).To_User()},
	}}

	glog.Infof("contacts.resolveUsername#f93ccba3 - reply: {%v}", resolvedPeer)
	return resolvedPeer.To_Contacts_ResolvedPeer(), nil
}
