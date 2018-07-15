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
	"github.com/airwide-code/airwide.datacenter/biz/core/account"
	"github.com/airwide-code/airwide.datacenter/biz/nbfs_client"
)

/*
	wallPaper#ccb03657 id:int title:string sizes:Vector<PhotoSize> color:int = WallPaper;
	wallPaperSolid#63117f24 id:int title:string bg_color:int color:int = WallPaper;
 */

// account.getWallPapers#c04cfac2 = Vector<WallPaper>;
func (s *AccountServiceImpl) AccountGetWallPapers(ctx context.Context, request *mtproto.TLAccountGetWallPapers) (*mtproto.Vector_WallPaper, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.getWallPapers#c04cfac2 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))
	//
	wallDataList := account.GetWallPaperList()

	walls := &mtproto.Vector_WallPaper{
		Datas: make([]*mtproto.WallPaper, 0, len(wallDataList)),
	}

	for _, wallData := range wallDataList {
		if wallData.Type == 0 {
			szList, _ := nbfs_client.GetPhotoSizeList(wallData.PhotoId)
			wall := &mtproto.TLWallPaper{Data2: &mtproto.WallPaper_Data{
				Id:    wallData.Id,
				Title: wallData.Title,
				Sizes: szList,
				Color: wallData.Color,
			}}
			walls.Datas = append(walls.Datas, wall.To_WallPaper())
		} else {
			wall := &mtproto.TLWallPaperSolid{Data2: &mtproto.WallPaper_Data{
				Id:      wallData.Id,
				Title:   wallData.Title,
				Color:   wallData.Color,
				BgColor: wallData.BgColor,
			}}
			walls.Datas = append(walls.Datas, wall.To_WallPaper())
		}
	}

	glog.Infof("account.getWallPapers#c04cfac2 - reply: %s", logger.JsonDebugData(walls))
	return walls, nil
}
