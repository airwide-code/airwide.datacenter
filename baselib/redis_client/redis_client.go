/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package redis_client

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/golang/glog"
	"time"
)

type redisConn struct {
	pool   *RedisPool
	conn   redis.Conn
}

type RedisPool struct {
	*redis.Pool
	env string
}

func NewRedisPool(c *RedisConfig) (pool *RedisPool) {
	pool = &RedisPool{env: fmt.Sprintf("[%s]tcp@%s", c.Name, c.Addr)}
	cnop := redis.DialConnectTimeout(time.Duration(c.DialTimeout))
	rdop := redis.DialReadTimeout(time.Duration(c.ReadTimeout))
	wrop := redis.DialWriteTimeout(time.Duration(c.WriteTimeout))

	dialFunc := func() (rconn redis.Conn, err error) {
		rconn, err = redis.Dial("tcp", c.Addr, cnop, rdop, wrop)
		if err != nil {
			glog.Errorf("Redis connect %s error: %s", pool.env, err)
			return
		}

		if c.Password != "" {
			if _, err = rconn.Do("AUTH", c.Password); err != nil {
				glog.Errorf("Redis %s AUTH(password: %s) error: %s", pool.env, c.Password, err)
				rconn.Close()
				rconn = nil
				return
			}
		}

		// TODO(@benqi):  检查c.DBNum，必须是数字
		_, err = rconn.Do("SELECT", c.DBNum)
		if err != nil {
			glog.Errorf("Redis %s SELECT %s error: %s", pool.env, c.DBNum, err)
			rconn.Close()
			rconn = nil
		}
		return
	}

	pool.Pool = &redis.Pool{
		MaxActive: 	 c.Active,
		MaxIdle:     c.Idle,
		IdleTimeout: time.Duration(c.IdleTimeout),
		Dial:        dialFunc,
	}
	return
}

func (p *RedisPool) Get() redis.Conn {
	return &redisConn{
		pool: p,
		conn: p.Pool.Get(),
	}
}

func (p *RedisPool) Close() error {
	return p.Pool.Close()
}

func (c *redisConn) Err() error {
	return c.conn.Err()
}

func (c *redisConn) Close() error {
	return c.conn.Close()
}

func (c *redisConn) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	if err = c.Err(); err != nil {
		return
	}
	return c.conn.Do(commandName, args...)
}

// NOTE not goroutine safe
func (c *redisConn) Send(commandName string, args ...interface{}) (err error) {
	if err = c.Err(); err != nil {
		return
	}
	return c.conn.Send(commandName, args...)
}

func (c *redisConn) Flush() error {
	return c.conn.Flush()
}

// NOTE not goroutine safe
func (c *redisConn) Receive() (reply interface{}, err error) {
	if err = c.Err(); err != nil {
		return
	}
	return c.conn.Receive()
}

