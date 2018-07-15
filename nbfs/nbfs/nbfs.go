/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package main

import (
	"flag"
	"github.com/golang/glog"

	upload "github.com/airwide-code/airwide.datacenter/nbfs/nbfs/upload/rpc"
	photo "github.com/airwide-code/airwide.datacenter/nbfs/nbfs/photo/rpc"
	"github.com/airwide-code/airwide.datacenter/mtproto"
	"github.com/airwide-code/airwide.datacenter/baselib/mysql_client"
	"github.com/BurntSushi/toml"
	"fmt"
	"github.com/airwide-code/airwide.datacenter/nbfs/biz/dal/dao"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util/service_discovery"
	"google.golang.org/grpc"
	"github.com/airwide-code/airwide.datacenter/nbfs/biz/core"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

type RpcServerConfig struct {
	Addr string
}

type NbfsConfig struct {
	DataPath string
}

type uploadServerConfig struct {
	Nbfs      *NbfsConfig
	Server    *RpcServerConfig
	Discovery service_discovery.ServiceDiscoveryServerConfig
	Mysql     []mysql_client.MySQLConfig
}

// 整合各服务，方便开发调试
func main() {
	flag.Parse()

	config := &uploadServerConfig{}
	if _, err := toml.DecodeFile("./nbfs.toml", config); err != nil {
		fmt.Errorf("%s\n", err)
		return
	}

	glog.Info(config.Nbfs, ", ", config.Server, ", ", config.Discovery, ", ", config.Mysql)

	// Init
	core.InitNbfsDataPath(config.Nbfs.DataPath)

	// 初始化mysql_client、redis_client
	// redis_client.InstallRedisClientManager(config.Redis)
	mysql_client.InstallMysqlClientManager(config.Mysql)

	// 初始化redis_dao、mysql_dao
	dao.InstallMysqlDAOManager(mysql_client.GetMysqlClientManager())
	// dao.InstallRedisDAOManager(redis_client.GetRedisClientManager())

	// Start server
	grpcServer := grpc_util.NewRpcServer(config.Server.Addr, &config.Discovery)
	grpcServer.Serve(func(s *grpc.Server) {
		mtproto.RegisterRPCUploadServer(s, &upload.UploadServiceImpl{DataPath: config.Nbfs.DataPath})
		mtproto.RegisterRPCNbfsServer(s, &photo.PhotoServiceImpl{})
	})
}
