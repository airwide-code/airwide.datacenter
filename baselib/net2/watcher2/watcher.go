/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package watcher2

import (
	"github.com/coreos/etcd/clientv3"
	// "encoding/json"
	"github.com/golang/glog"
	etcd3 "github.com/coreos/etcd/clientv3"
	"context"
	"fmt"
	"github.com/airwide-code/airwide.datacenter/baselib/net2"
	"encoding/json"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

// see: /baselib/grpc_util/service_discovery/registry.go
type nodeData struct {
	Addr     string
	Metadata map[string]string
}

// TODO(@benqi): grpc_util/serviec_discovery集成
type ClientWatcher struct {
	etcCli      *clientv3.Client
	registryDir string
	serviceName string
	// rootPath    string
	client      *net2.TcpClientGroupManager
	nodes       map[string]*nodeData
}

func NewClientWatcher(registryDir, serviceName string, cfg etcd3.Config, client *net2.TcpClientGroupManager) (watcher *ClientWatcher, err error) {
	var etcdClient *clientv3.Client
	if etcdClient, err = clientv3.New(cfg); err != nil {
		glog.Error("Error: cannot connec to etcd:", err)
		return
	}

	watcher = &ClientWatcher{
		etcCli:      etcdClient,
		registryDir: registryDir,
		serviceName: serviceName,
		client:      client,
		nodes:       map[string]*nodeData{},
	}
	return
}

func (m *ClientWatcher) WatchClients(cb func(etype, addr string)) {
	rootPath := fmt.Sprintf("%s/%s", m.registryDir, m.serviceName)

	resp, err := m.etcCli.Get(context.Background(), rootPath, clientv3.WithPrefix())
	if err != nil {
		glog.Error(err)
	}
	for _, kv := range resp.Kvs {
		m.addClient(kv, cb)
	}

	rch := m.etcCli.Watch(context.Background(), rootPath, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			if ev.Type.String() == "EXPIRE" {
				// TODO(@benqi): 采用何种策略？？
				// n, ok := m.nodes[string(ev.Kv.Key)]
				// if ok {
				//	 delete(m.nodes, string(ev.Kv.Key))
				// }
				// if cb != nil {
				// 	cb("EXPIRE", string(ev.Kv.Key), string(ev.Kv.Value))
				//}
			} else if ev.Type.String() == "PUT" {
				m.addClient(ev.Kv, cb)
			} else if ev.Type.String() == "DELETE" {
				if n, ok := m.nodes[string(ev.Kv.Key)]; ok {
					m.client.RemoveClient(m.serviceName, n.Addr)
					if cb != nil {
						cb("delete", n.Addr)
					}
					delete(m.nodes, string(ev.Kv.Key))
				}
			}
		}
	}
}
func (m *ClientWatcher) addClient(kv *mvccpb.KeyValue, cb func(etype, addr string)) {
	node := &nodeData{}
	err := json.Unmarshal(kv.Value, node)
	if err != nil {
		glog.Error(err)
	}
	if n, ok := m.nodes[string(kv.Key)]; ok {
		if node.Addr != n.Addr {
			m.client.RemoveClient(m.serviceName, n.Addr)
			m.nodes[string(kv.Key)] = node
			if cb != nil {
				cb("delete", n.Addr)
			}
			if cb != nil {
				cb("add", node.Addr)
			}
		}
		m.client.AddClient(m.serviceName, node.Addr)
	} else {
		m.nodes[string(kv.Key)] = node
		m.client.AddClient(m.serviceName, node.Addr)
		if cb != nil {
			cb("add", node.Addr)
		}
	}
}
