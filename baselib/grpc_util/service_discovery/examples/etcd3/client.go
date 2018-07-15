/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
	"github.com/coreos/etcd/clientv3"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util/service_discovery/examples/proto"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util/service_discovery/etcd3"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util/load_balancer"
)

func main() {
	etcdConfg := clientv3.Config{
		Endpoints: []string{"http://127.0.0.1:2379"},
	}
	r := etcd3.NewResolver("/grpc-lb", "test", etcdConfg)
	b := load_balancer.NewBalancer(r, load_balancer.NewRoundRobinSelector())
	c, err := grpc.Dial("", grpc.WithInsecure(),  grpc.WithBalancer(b), grpc.WithTimeout(time.Second*5))
	if err != nil {
		log.Printf("grpc dial: %s", err)
		return
	}
	defer c.Close()

	client := proto.NewEchoServiceClient(c)

	for i := 0; i < 1000; i++ {
		resp, err := client.Echo(context.Background(), &proto.EchoReq{EchoData: "round robin"})
		if err != nil {
			log.Println("aa:", err)
			time.Sleep(time.Second)
			continue
		}
		log.Printf(resp.EchoData)
		time.Sleep(time.Second)
	}
}
