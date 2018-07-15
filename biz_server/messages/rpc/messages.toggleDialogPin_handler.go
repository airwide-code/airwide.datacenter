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

// messages.toggleDialogPin#3289be6a flags:# pinned:flags.0?true peer:InputPeer = Bool;
func (s *MessagesServiceImpl) MessagesToggleDialogPin(ctx context.Context, request *mtproto.TLMessagesToggleDialogPin) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesToggleDialogPin - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesToggleDialogPin logic

	return nil, fmt.Errorf("Not impl MessagesToggleDialogPin")
}
