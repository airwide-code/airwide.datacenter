/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package redis_client

import (
	"github.com/golang/glog"
	"fmt"
)

type redisClientManager struct{
	// TODO(@benqi): 使用sync.Map，动态添加和卸载数据库
	redisClients map[string]*RedisPool
}

var redisClients = &redisClientManager{make(map[string]*RedisPool)}

func  InstallRedisClientManager(configs []RedisConfig) {
	for _, config := range configs {
		client := NewRedisPool(&config)
		if client == nil {
			err := fmt.Errorf("InstallRedisClient - NewRedisPool {%v} error!", config)
			panic(err)
			// continue
		}

		// TODO(@benqi): 检查config数据合法性
		redisClients.redisClients[config.Name] = client
	}
}

func  GetRedisClient(redisName string) (client *RedisPool) {
	client, ok := redisClients.redisClients[redisName]
	if !ok {
		glog.Errorf("GetRedisClient - Not found client: %s", redisName)
	}
	return
}

func  GetRedisClientManager() map[string]*RedisPool {
	return redisClients.redisClients
}
