/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package config

import (
	"github.com/airwide-code/airwide.datacenter/baselib/base"
)

type EtcdConfig struct {
	Name    string
	Root    string
	Addrs   []string
	Timeout base.Duration
}
