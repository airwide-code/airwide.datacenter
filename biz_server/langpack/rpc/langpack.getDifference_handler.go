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

// langpack.getDifference#b2e4d7d from_version:int = LangPackDifference;
func (s *LangpackServiceImpl) LangpackGetDifference(ctx context.Context, request *mtproto.TLLangpackGetDifference) (*mtproto.LangPackDifference, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("LangpackGetDifference - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl LangpackGetDifference logic
	diff := mtproto.NewTLLangPackDifference()
	diff.SetLangCode("en")
	diff.SetVersion(langs.Version)
	diff.SetFromVersion(request.FromVersion)

	if request.FromVersion < langs.Version {
		// TODO(@benqi): 找出不同版本的增量更新数据
	}

	glog.Infof("LangpackGetDifference - reply: %s", logger.JsonDebugData(diff))
	return diff.To_LangPackDifference(), nil
}
