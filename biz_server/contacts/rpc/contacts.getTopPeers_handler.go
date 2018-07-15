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

// contacts.getTopPeers#d4982db5 flags:# correspondents:flags.0?true bots_pm:flags.1?true bots_inline:flags.2?true phone_calls:flags.3?true groups:flags.10?true channels:flags.15?true offset:int limit:int hash:int = contacts.TopPeers;
func (s *ContactsServiceImpl) ContactsGetTopPeers(ctx context.Context, request *mtproto.TLContactsGetTopPeers) (*mtproto.Contacts_TopPeers, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ContactsGetTopPeers - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ContactsGetTopPeers logic
	topPeers := &mtproto.TLContactsTopPeers{Data2: &mtproto.Contacts_TopPeers_Data{

	}}
	return topPeers.To_Contacts_TopPeers(), nil
	//return nil, fmt.Errorf("Not impl ContactsGetTopPeers")
}
