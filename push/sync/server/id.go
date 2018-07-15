/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package server

import (
	"github.com/airwide-code/airwide.datacenter/baselib/snowflake"
	"flag"
)

var id *snowflake.IdWorker

// = &snowflake.IdWorker{
//
//}

// TODO(@benqi): 多服务器部署时需要配置驱动workId
const (
	workerId	   	= int64(1)
	dataCenterId	= int64(1)
	twepoch        	= int64(1288834974657)
)

func init()  {
	flag.Parse()
	id, _ = snowflake.NewIdWorker(workerId, dataCenterId, twepoch)
}

func NextId() (int64) {
	r, _ := id.NextId()
	return r
}
