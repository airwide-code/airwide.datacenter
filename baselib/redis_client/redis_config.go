/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package redis_client

import (
	"github.com/airwide-code/airwide.datacenter/baselib/base"
	"fmt"
)

// Redis client config.
type RedisConfig struct {
	Name         	string 			// redis name
	Addr         	string
	Active       	int 			// pool
	Idle         	int 			// pool
	DialTimeout  	base.Duration
	ReadTimeout  	base.Duration
	WriteTimeout 	base.Duration
	IdleTimeout  	base.Duration

	DBNum			string			// db号
	Password 		string			// 密码
}

func (c *RedisConfig) ToRedisCacheConfig() string {
	// config is like {"key":"collection key","conn":"connection info","dbNum":"0"}
	// rc.key = cf["key"]
	// rc.conninfo = cf["conn"]
	// rc.dbNum, _ = strconv.Atoi(cf["dbNum"])
	// rc.password = cf["password"]
	return fmt.Sprintf(`{"conn":"%s", "dbNum":"%s", "password":"%s"}`,
		c.Addr,
		c.DBNum,
		c.Password)
}
