/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package main

import (
	"flag"
	"github.com/golang/glog"

	account "github.com/airwide-code/airwide.datacenter/biz_server/account/rpc"
	auth "github.com/airwide-code/airwide.datacenter/biz_server/auth/rpc"
	bots "github.com/airwide-code/airwide.datacenter/biz_server/bots/rpc"
	channels "github.com/airwide-code/airwide.datacenter/biz_server/channels/rpc"
	contacts "github.com/airwide-code/airwide.datacenter/biz_server/contacts/rpc"
	help "github.com/airwide-code/airwide.datacenter/biz_server/help/rpc"
	langpack "github.com/airwide-code/airwide.datacenter/biz_server/langpack/rpc"
	messages "github.com/airwide-code/airwide.datacenter/biz_server/messages/rpc"
	payments "github.com/airwide-code/airwide.datacenter/biz_server/payments/rpc"
	phone "github.com/airwide-code/airwide.datacenter/biz_server/phone/rpc"
	photos "github.com/airwide-code/airwide.datacenter/biz_server/photos/rpc"
	stickers "github.com/airwide-code/airwide.datacenter/biz_server/stickers/rpc"
	updates "github.com/airwide-code/airwide.datacenter/biz_server/updates/rpc"
	users "github.com/airwide-code/airwide.datacenter/biz_server/users/rpc"
	"github.com/airwide-code/airwide.datacenter/mtproto"
	"github.com/airwide-code/airwide.datacenter/baselib/redis_client"
	"github.com/airwide-code/airwide.datacenter/baselib/mysql_client"
	"github.com/BurntSushi/toml"
	"fmt"
	"github.com/airwide-code/airwide.datacenter/biz/dal/dao"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util/service_discovery"
	"google.golang.org/grpc"
	"github.com/airwide-code/airwide.datacenter/biz_server/sync_client"
	"github.com/airwide-code/airwide.datacenter/biz/nbfs_client"
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

type BizServerConfig struct{
	Server 		*RpcServerConfig
	Discovery service_discovery.ServiceDiscoveryServerConfig

	// RpcClient	*RpcClientConfig
	Mysql		[]mysql_client.MySQLConfig
	Redis 		[]redis_client.RedisConfig
	NbfsRpcClient  *service_discovery.ServiceDiscoveryClientConfig
	SyncRpcClient1 *service_discovery.ServiceDiscoveryClientConfig
	SyncRpcClient2 *service_discovery.ServiceDiscoveryClientConfig
}

// 整合各服务，方便开发调试
func main() {
	flag.Parse()

	bizServerConfig := &BizServerConfig{}
	if _, err := toml.DecodeFile("./biz_server.toml", bizServerConfig); err != nil {
		fmt.Errorf("%s\n", err)
		return
	}

	glog.Info(bizServerConfig)

	// 初始化mysql_client、redis_client
	redis_client.InstallRedisClientManager(bizServerConfig.Redis)
	mysql_client.InstallMysqlClientManager(bizServerConfig.Mysql)

	// 初始化redis_dao、mysql_dao
	dao.InstallMysqlDAOManager(mysql_client.GetMysqlClientManager())
	dao.InstallRedisDAOManager(redis_client.GetRedisClientManager())

	nbfs_client.InstallNbfsClient(bizServerConfig.NbfsRpcClient)
	sync_client.InstallSyncClient(bizServerConfig.SyncRpcClient2)

	// InstallNbfsClient

	// Start server
	grpcServer := grpc_util.NewRpcServer(bizServerConfig.Server.Addr, &bizServerConfig.Discovery)
	grpcServer.Serve(func(s *grpc.Server) {
		// AccountServiceImpl
		mtproto.RegisterRPCAccountServer(s, &account.AccountServiceImpl{})

		// AuthServiceImpl
		mtproto.RegisterRPCAuthServer(s, &auth.AuthServiceImpl{})

		mtproto.RegisterRPCBotsServer(s, &bots.BotsServiceImpl{})
		mtproto.RegisterRPCChannelsServer(s, &channels.ChannelsServiceImpl{})

		// ContactsServiceImpl
		mtproto.RegisterRPCContactsServer(s, &contacts.ContactsServiceImpl{})

		mtproto.RegisterRPCHelpServer(s, &help.HelpServiceImpl{})
		mtproto.RegisterRPCLangpackServer(s, &langpack.LangpackServiceImpl{})

		// MessagesServiceImpl
		mtproto.RegisterRPCMessagesServer(s, &messages.MessagesServiceImpl{})

		mtproto.RegisterRPCPaymentsServer(s, &payments.PaymentsServiceImpl{})
		mtproto.RegisterRPCPhoneServer(s, &phone.PhoneServiceImpl{})
		mtproto.RegisterRPCPhotosServer(s, &photos.PhotosServiceImpl{})
		mtproto.RegisterRPCStickersServer(s, &stickers.StickersServiceImpl{})
		mtproto.RegisterRPCUpdatesServer(s, &updates.UpdatesServiceImpl{})

		mtproto.RegisterRPCUsersServer(s, &users.UsersServiceImpl{})
	})
}
