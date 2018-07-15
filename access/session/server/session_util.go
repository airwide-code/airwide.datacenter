/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package server

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/airwide-code/airwide.datacenter/baselib/app"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util"
	"github.com/airwide-code/airwide.datacenter/baselib/net2"
	"github.com/airwide-code/airwide.datacenter/mtproto"
)

func sendDataByConnection(conn *net2.TcpConnection, sessionID uint64, md *mtproto.ZProtoMetadata, buf []byte) error {
	smsg := &mtproto.ZProtoSessionData{
		MTPMessage: &mtproto.MTPRawMessage{
			Payload: buf,
		},
	}
	zmsg := &mtproto.ZProtoMessage{
		SessionId: sessionID,
		Metadata:  md,
		SeqNum:    2,
		Message: &mtproto.ZProtoRawPayload{
			Payload: smsg.Encode(),
		},
	}
	return conn.Send(zmsg)
}

func sendDataByConnID(connID, sessionID uint64, md *mtproto.ZProtoMetadata, buf []byte) error {
	sessionServer, ok := app.GAppInstance.(*SessionServer)
	if !ok {
		err := fmt.Errorf("not use app instance framework")
		glog.Error(err)
		return err
	}
	return sessionServer.SendToClientData(connID, sessionID, md, buf)
}

func getBizRPCClient() (*grpc_util.RPCClient, error) {
	sessionServer, ok := app.GAppInstance.(*SessionServer)
	if !ok {
		err := fmt.Errorf("not use app instance framework")
		glog.Error(err)
		return nil, err
	}
	return sessionServer.bizRpcClient, nil
}

func getNbfsRPCClient() (*grpc_util.RPCClient, error) {
	sessionServer, ok := app.GAppInstance.(*SessionServer)
	if !ok {
		err := fmt.Errorf("not use app instance framework")
		glog.Error(err)
		return nil, err
	}
	return sessionServer.nbfsRpcClient, nil
}

func getSyncRPCClient() (mtproto.RPCSyncClient, error) {
	sessionServer, ok := app.GAppInstance.(*SessionServer)
	if !ok {
		err := fmt.Errorf("not use app instance framework")
		glog.Error(err)
		return nil, err
	}
	return sessionServer.syncRpcClient, nil
}

func deleteClientSessionManager(authKeyID int64) {
	if sessionServer, ok := app.GAppInstance.(*SessionServer); !ok {
		err := fmt.Errorf("not use app instance framework")
		glog.Error(err)
	} else {
		sessionServer.sessionManager.onCloseSessionClientManager(authKeyID)
	}
}
