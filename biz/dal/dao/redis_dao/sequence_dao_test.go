/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package redis_dao

import (
	"github.com/airwide-code/airwide.datacenter/baselib/redis_client"
	"testing"
)

func TestNextID(t *testing.T) {
	// mysqlDsn := "root:@/nebulaim?charset=utf8"

	//db, err := sqlx.Connect("mysql", mysqlDsn)
	//if err != nil {
	//	glog.Fatalf("Connect mysql %s error: %s", mysqlDsn, err)
	//	return
	//}

	// seqUpdatesNgen := NewSeqUpdatesNgenDAO(db)

	redisConfig := &redis_client.RedisConfig{
		Name:         "test",
		Addr:         "127.0.0.1:6379",
		Idle:         100,
		Active:       100,
		DialTimeout:  1000000,
		ReadTimeout:  1000000,
		WriteTimeout: 1000000,
		IdleTimeout:  15000000,
		DBNum:        "0",
		Password:     "",
	}

	redisPool := redis_client.NewRedisPool(redisConfig)

	_ := NewSequenceDAO(redisPool)
	//seq.NextID("1")
	//seq.NextID("1")
	//seq.NextID("1")
	//seq.NextID("1")
	//seq.NextID("2")
	//seq.NextID("2")
	//seq.NextID("2")
	//seq.NextID("2")
}
