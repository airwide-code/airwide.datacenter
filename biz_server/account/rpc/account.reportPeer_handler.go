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
	"github.com/airwide-code/airwide.datacenter/biz/base"
	"github.com/airwide-code/airwide.datacenter/biz/core/account"
)

/*
 Android client source code:
	if (ChatObject.isChannel(currentChat) && !currentChat.creator && (!currentChat.megagroup || currentChat.username != null && currentChat.username.length() > 0)) {
		headerItem.addSubItem(report, LocaleController.getString("ReportChat", R.string.ReportChat));
	}
 */
// account.reportPeer#ae189d5f peer:InputPeer reason:ReportReason = Bool;
func (s *AccountServiceImpl) AccountReportPeer(ctx context.Context, request *mtproto.TLAccountReportPeer) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.reportPeer#ae189d5f - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Check peer invalid
	peer := request.Peer
	// TODO(@benqi): Check peer access_hash
	if peer.GetConstructor() != mtproto.TLConstructor_CRC32_inputPeerChannel {
		// TODO(@benqi): Add INPUT_PEER_INVALID code
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error(err)
		return nil, err
	} else {
		// TODO(@benqi): !currentChat.creator && (!currentChat.megagroup || currentChat.username != null && currentChat.username.length() > 0)
	}

	// peer := helper.FromInputPeer(request.GetPeer())
	reason := base.FromReportReason(request.GetReason())

	text := ""
	if reason == base.REASON_OTHER {
		text = request.GetReason().GetData2().GetText()
	}

	account.InsertReportData(md.UserId, base.PEER_CHANNEL, peer.GetData2().GetChannelId(), int32(reason), text)

	glog.Infof("account.reportPeer#ae189d5f - reply: {true}",)
	return mtproto.ToBool(true), nil
}
