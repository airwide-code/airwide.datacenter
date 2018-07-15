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
	"github.com/airwide-code/airwide.datacenter/biz/core/contact"
	"github.com/airwide-code/airwide.datacenter/biz/core/user"
)

// contacts.search#11f812d8 q:string limit:int = contacts.Found;
func (s *ContactsServiceImpl) ContactsSearch(ctx context.Context, request *mtproto.TLContactsSearch) (*mtproto.Contacts_Found, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("contacts.search#11f812d8 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// Check query string and limit
	if len(request.Q) < 5 || request.Limit < 1 {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error(err, ": query or limit invalid")
		return nil, err
	}

	contactLogic := contact.MakeContactLogic(md.UserId)
	idList := contactLogic.SearchContacts(request.Q, request.Limit)

	// results
	results := make([]*mtproto.Peer, 0, len(idList))
	for _, id := range idList {
		peer := &mtproto.TLPeerUser{Data2: &mtproto.Peer_Data{
			UserId: id,
		}}
		results = append(results, peer.To_Peer())
	}

	// users
	users := user.GetUsersBySelfAndIDList(md.UserId, idList)

	found := &mtproto.TLContactsFound{Data2: &mtproto.Contacts_Found_Data{
		Results: results,
		Users:   users,
	}}

	glog.Infof("contacts.search#11f812d8 - reply: ", logger.JsonDebugData(found))
	return found.To_Contacts_Found(), nil
}
