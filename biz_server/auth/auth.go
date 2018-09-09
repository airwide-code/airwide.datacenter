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

	auth "github.com/airwide-code/airwide.datacenter/biz_server/auth/rpc"
	"github.com/airwide-code/airwide.datacenter/mtproto"
	"github.com/airwide-code/airwide.datacenter/baselib/redis_client"
	"github.com/airwide-code/airwide.datacenter/baselib/mysql_client"
	"github.com/BurntSushi/toml"
	"fmt"
	"github.com/airwide-code/airwide.datacenter/biz/dal/dao"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util/service_discovery"
	"google.golang.org/grpc"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

type RpcServerConfig struct {
	Addr string
}

//type RpcClientConfig struct {
//	ServiceName string
//	Addr string
//}

type authServerConfig struct{
	Server 		*RpcServerConfig
	Discovery service_discovery.ServiceDiscoveryServerConfig

	// RpcClient	*RpcClientConfig
	Mysql		[]mysql_client.MySQLConfig
	Redis 		[]redis_client.RedisConfig
}

// Integrate services to facilitate development and debugging
func main() {
	flag.Parse()

	config := &authServerConfig{}
	if _, err := toml.DecodeFile("./auth.toml", config); err != nil {
		fmt.Errorf("%s\n", err)
		return
	}

	glog.Info(config)

	// initialization mysql_client、redis_client
	redis_client.InstallRedisClientManager(config.Redis)
	mysql_client.InstallMysqlClientManager(config.Mysql)

	// initialization redis_dao、mysql_dao
	dao.InstallMysqlDAOManager(mysql_client.GetMysqlClientManager())
	dao.InstallRedisDAOManager(redis_client.GetRedisClientManager())

	// Start server
	grpcServer := grpc_util.NewRpcServer(config.Server.Addr, &config.Discovery)
	grpcServer.Serve(func(s *grpc.Server) {
		// AuthServiceImpl
		mtproto.RegisterRPCAuthServer(s, &auth.AuthServiceImpl{})
	})
}
