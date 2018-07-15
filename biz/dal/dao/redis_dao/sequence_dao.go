/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

// TODO(@benqi): 可以使用如下方法来生成seq
// - 如果运维能保证redis数据可靠性，可移除数据库seq_updates_ngen的存储
// - 可使用[seqsvr](https://github.com/nebula-in/seqsvr)服务来生成seq
// - 可调研艺龙的序列号生成器
// - 直接使用etcd或zk
package redis_dao

import (
	"github.com/garyburd/redigo/redis"
	"github.com/airwide-code/airwide.datacenter/baselib/redis_client"
	"github.com/golang/glog"
)

const (
	// TODO(@benqi): 使用更紧凑的前缀
	seqUpdatesNgenId = "seq_updates_ngen_"
	ptsUpdatesNgenId = "pts_updates_ngen_"
	qtsUpdatesNgenId = "qts_updates_ngen_"
	boxUpdatesNgenId = "message_box_ngen_"
	channelPtsUpdatesNgenId = "channel_pts_updates_ngen_"
	channelBoxUpdatesNgenId = "channel_message_box_ngen_"
)

type SequenceDAO struct {
	redis *redis_client.RedisPool
	// ngen  *SeqUpdatesNgenDAO
}

func NewSequenceDAO(redis *redis_client.RedisPool) *SequenceDAO {
	return &SequenceDAO{
		redis: redis,
		// ngen:  ngen,
	}
}

func (dao *SequenceDAO) NextSeqId(key string) (seq int64, err error) {
	return dao.FetchNextSequence(seqUpdatesNgenId + key)
}

func (dao *SequenceDAO) CurrentSeqId(key string) (seq int64, err error) {
	return dao.GetCurrentSequence(seqUpdatesNgenId + key)
}

func (dao *SequenceDAO) NextPtsId(key string) (seq int64, err error) {
	return dao.FetchNextSequence(ptsUpdatesNgenId + key)
}

func (dao *SequenceDAO) CurrentPtsId(key string) (seq int64, err error) {
	return dao.GetCurrentSequence(ptsUpdatesNgenId + key)
}

func (dao *SequenceDAO) NextQtsId(key string) (seq int64, err error) {
	return dao.FetchNextSequence(qtsUpdatesNgenId + key)
}

func (dao *SequenceDAO) CurrentQtsId(key string) (seq int64, err error) {
	return dao.GetCurrentSequence(qtsUpdatesNgenId + key)
}

func (dao *SequenceDAO) NextMessageBoxId(key string) (seq int64, err error) {
	return dao.FetchNextSequence(boxUpdatesNgenId + key)
}

func (dao *SequenceDAO) CurrentMessageBoxId(key string) (seq int64, err error) {
	return dao.GetCurrentSequence(boxUpdatesNgenId + key)
}

func (dao *SequenceDAO) NextChannelPtsId(key string) (seq int64, err error) {
	return dao.FetchNextSequence(channelPtsUpdatesNgenId + key)
}

func (dao *SequenceDAO) CurrentChannelPtsId(key string) (seq int64, err error) {
	return dao.GetCurrentSequence(channelPtsUpdatesNgenId + key)
}

func (dao *SequenceDAO) NextChannelMessageBoxId(key string) (seq int64, err error) {
	return dao.FetchNextSequence(channelBoxUpdatesNgenId + key)
}

func (dao *SequenceDAO) CurrentChannelMessageBoxId(key string) (seq int64, err error) {
	return dao.GetCurrentSequence(channelBoxUpdatesNgenId + key)
}

func (dao *SequenceDAO) FetchNextSequence(key string) (seq int64, err error) {
	conn := dao.redis.Get()
	defer conn.Close()

	// 设置键
	seq, err = redis.Int64(conn.Do("INCR", key))
	if err != nil {
		glog.Errorf("FetchNextSequence - INCR {%s}, error: {%v}", key, err)
	}

	return
}

func (dao *SequenceDAO) GetCurrentSequence(key string) (seq int64, err error) {
	conn := dao.redis.Get()
	defer conn.Close()

	seq, err = redis.Int64(conn.Do("GET", key))
	if err != nil {
		glog.Errorf("GetCurrentSequence - GET {%s}, error: {%v}", key, err)
	}

	return
}
