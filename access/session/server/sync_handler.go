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
	"github.com/gogo/protobuf/proto"
	"github.com/golang/glog"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util"
	"github.com/airwide-code/airwide.datacenter/baselib/net2"
	"github.com/airwide-code/airwide.datacenter/mtproto"
)

func init() {
	proto.RegisterType((*mtproto.ConnectToSessionServerReq)(nil), "mtproto.ConnectToSessionServerReq")
	proto.RegisterType((*mtproto.SessionServerConnectedRsp)(nil), "mtproto.SessionServerConnectedRsp")
	proto.RegisterType((*mtproto.PushUpdatesData)(nil), "mtproto.PushUpdatesData")
	proto.RegisterType((*mtproto.VoidRsp)(nil), "mtproto.VoidRsp")
}

type syncHandler struct {
	smgr *sessionManager
}

func newSyncHandler(smgr *sessionManager) *syncHandler {
	s := &syncHandler{
		smgr: smgr,
	}
	return s
}

func protoToRawPayload(m proto.Message) (*mtproto.ZProtoRawPayload, error) {
	x := mtproto.NewEncodeBuf(128)
	x.UInt(mtproto.SYNC_DATA)
	n := proto.MessageName(m)
	x.Int(int32(len(n)))
	x.Bytes([]byte(n))
	b, err := proto.Marshal(m)
	x.Bytes(b)
	return &mtproto.ZProtoRawPayload{Payload: x.GetBuf()}, err
}

func (s *syncHandler) onSyncData(conn *net2.TcpConnection, buf []byte) (*mtproto.ZProtoRawPayload, error) {
	dbuf := mtproto.NewDecodeBuf(buf)
	len2 := int(dbuf.Int())
	messageName := string(dbuf.Bytes(len2))
	message, err := grpc_util.NewMessageByName(messageName)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	err = proto.Unmarshal(buf[4+len2:], message)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	switch message.(type) {
	case *mtproto.ConnectToSessionServerReq:
		glog.Infof("onSyncData - request(ConnectToSessionServerReq): {%v}", message)
		return protoToRawPayload(&mtproto.SessionServerConnectedRsp{
			ServerId:   1,
			ServerName: "session",
		})
	case *mtproto.PushUpdatesData:
		glog.Infof("onSyncData - request(PushUpdatesData): {%v}", message)

		// TODO(@benqi): dispatch to session_client
		pushData, _ := message.(*mtproto.PushUpdatesData)
		dbuf := mtproto.NewDecodeBuf(pushData.GetUpdatesData())
		mdata := &messageData{
			confirmFlag:  true,
			compressFlag: false,
			obj:          dbuf.Object(),
		}
		if mdata.obj == nil {
			glog.Errorf("onSyncData - recv invalid pushData: {%v}", message)
		} else {
			md := &mtproto.ZProtoMetadata{}
			// push
			// s.smgr.pushToSessionData(pushData.GetAuthKeyId(), pushData.GetSessionId(), md, mdata)
			s.smgr.onSyncData2(pushData.GetAuthKeyId(), pushData.GetSessionId(), md, mdata)
		}

		return protoToRawPayload(&mtproto.VoidRsp{})
	default:
		err := fmt.Errorf("invalid register proto type: {%v}", message)
		glog.Error(err)
		return nil, err
	}
}
