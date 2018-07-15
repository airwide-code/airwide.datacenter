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

// langpack.getStrings#2e1ee318 lang_code:string keys:Vector<string> = Vector<LangPackString>;
func (s *LangpackServiceImpl) LangpackGetStrings(ctx context.Context, request *mtproto.TLLangpackGetStrings) (*mtproto.Vector_LangPackString, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("langpack.getStrings#2e1ee318 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Query from langpack config db
	langpackStrings := &mtproto.Vector_LangPackString{}
	for _, s := range request.Keys {
		s2 := &mtproto.TLLangPackString{Data2: &mtproto.LangPackString_Data{
			Key:   s,
			Value: s, // TODO(@benqi): Query value by key
		}}
		langpackStrings.Datas = append(langpackStrings.Datas, s2.To_LangPackString())
	}

	glog.Infof("langpack.getStrings#2e1ee318 - reply: %s", logger.JsonDebugData(langpackStrings))
	return langpackStrings, nil
}
