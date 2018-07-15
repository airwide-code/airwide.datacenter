/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package service_discovery

import (
	"github.com/airwide-code/airwide.datacenter/baselib/base"
)

type ServiceDiscoveryServerConfig struct {
	ServiceName string
	NodeID		string
	RPCAddr     string
	EtcdAddrs   []string
	Interval    base.Duration
	TTL         base.Duration
}

type ServiceDiscoveryClientConfig struct {
	ServiceName string
	EtcdAddrs   []string
	Balancer    string
}

