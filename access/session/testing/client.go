/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package main

import (
	"github.com/airwide-code/airwide.datacenter/baselib/net2"
	"github.com/golang/glog"
	_ "github.com/airwide-code/airwide.datacenter/mtproto"
	"flag"
	"github.com/airwide-code/airwide.datacenter/mtproto"
	"github.com/airwide-code/airwide.datacenter/baselib/crypto"
	"github.com/gogo/protobuf/proto"
)

type sessionClient struct {
	client *net2.TcpClient
}

func (s *sessionClient) OnNewClient(client *net2.TcpClient) {
	glog.Infof("OnNewConnection")

	req_pq := &mtproto.TLReqPq{
		Nonce: crypto.GenerateNonce(16),
	}

	authKeyMD := &mtproto.AuthKeyMetadata{}
	state := &mtproto.HandshakeState{
		State : mtproto.STATE_pq,
		ResState: mtproto.RES_STATE_NONE,
	}
	state.Ctx, _ = proto.Marshal(authKeyMD)

	smsg := &mtproto.SessionHandshakeMessage{
		State: state,
		MTPMessage: &mtproto.UnencryptedMessage{
			MessageId: 0,
			Object: req_pq,
		},
	}

	zmsg := &mtproto.ZProtoMessage{
		SessionId: 0,
		SeqNum: 0,
		Metadata: &mtproto.ZProtoMetadata{},
		Message: &mtproto.ZProtoRawPayload{
			Payload: smsg.Encode(),
		},
	}

	client.Send(zmsg)
}

func (s *sessionClient) OnClientDataArrived(client *net2.TcpClient, msg interface{}) error {
	glog.Infof("OnDataArrived - recv data: %v", msg)
	return nil
}

func (s *sessionClient) OnClientClosed(client *net2.TcpClient) {
	glog.Infof("OnConnectionClosed")
}

func (s *sessionClient) OnClientTimer(client *net2.TcpClient) {
	glog.Infof("OnTimer")
}

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

func main() {
	flag.Parse()

	client := &sessionClient{}
	client.client = net2.NewTcpClient("session", 1024, "zproto", "127.0.0.1:10000", client)
	client.client.Serve()
	select {}
}

