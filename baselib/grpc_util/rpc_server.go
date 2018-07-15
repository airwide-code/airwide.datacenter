/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package grpc_util

import (
	"net"
	"time"
	"google.golang.org/grpc"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util/middleware/recovery2"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util/service_discovery"
	"github.com/golang/glog"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util/service_discovery/etcd3"
	"github.com/coreos/etcd/clientv3"
	"os/signal"
	"syscall"
	"os"
)

type RPCServer struct {
	addr string
	registry *etcd3.EtcdReigistry
	s        *grpc.Server
}

func NewRpcServer(addr string, discovery *service_discovery.ServiceDiscoveryServerConfig) *RPCServer {
	etcdConfg := clientv3.Config{
		Endpoints: discovery.EtcdAddrs,
	}

	registry, err := etcd3.NewRegistry(
		etcd3.Option{
			EtcdConfig:  etcdConfg,
			RegistryDir: "/nebulaim",
			ServiceName: discovery.ServiceName,
			NodeID:      discovery.NodeID,
			NData: etcd3.NodeData{
				Addr: discovery.RPCAddr,
				//Metadata: map[string]string{"weight": "1"},
			},
			Ttl: time.Duration(discovery.TTL), // * time.Second,
		})
	if err != nil {
		glog.Fatal(err)
		// return nil
	}

	s := grpc_recovery2.NewRecoveryServer2(BizUnaryRecoveryHandler, BizUnaryRecoveryHandler2, BizStreamRecoveryHandler)
	rs := &RPCServer{
		addr:     addr,
		registry: registry,
		s:        s,
	}
	return rs
}

// type func RegisterRPCServerHandler(s *grpc.Server)
type RegisterRPCServerFunc func(s *grpc.Server)

func (s *RPCServer) Serve(regFunc RegisterRPCServerFunc) {
	// defer s.GracefulStop()
	listener, err := net.Listen("tcp", s.addr)

	if err != nil {
		glog.Error("failed to listen: %v", err)
		return
	}
	glog.Infof("rpc listening on:%s", s.addr)

	if regFunc != nil {
		regFunc(s.s)
	}

	defer s.s.GracefulStop()
	go s.registry.Register()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s2 := <-ch
		glog.Infof("exit...")
		s.registry.Deregister()
		if i, ok := s2.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}

	}()

	if err := s.s.Serve(listener); err != nil {
		glog.Fatalf("failed to serve: %s", err)
	}
}

func (s *RPCServer) Stop() {
	s.s.GracefulStop()
}
