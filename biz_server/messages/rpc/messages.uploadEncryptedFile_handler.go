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
    "github.com/airwide-code/airwide.datacenter/mtproto"
    "golang.org/x/net/context"
    "fmt"
    "github.com/airwide-code/airwide.datacenter/baselib/grpc_util"
    "github.com/airwide-code/airwide.datacenter/baselib/logger"
)

// messages.uploadEncryptedFile#5057c497 peer:InputEncryptedChat file:InputEncryptedFile = EncryptedFile;
func (s *MessagesServiceImpl) MessagesUploadEncryptedFile(ctx context.Context, request *mtproto.TLMessagesUploadEncryptedFile) (*mtproto.EncryptedFile, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("MessagesUploadEncryptedFile - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    // TODO(@benqi): Impl MessagesUploadEncryptedFile logic

    return nil, fmt.Errorf("Not impl MessagesUploadEncryptedFile")
}
