/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package updates

import (
	"github.com/airwide-code/airwide.datacenter/biz/dal/dao"
)

/*
// TODO(@benqi):
//  使用数据库和REDIS获取sequence
//  redis: sequence
//  暂时不考虑DB等异常处理
func (dao *sequnceModel) NextID(key string) (seq int64) {
	sequenceDAO := dao2.GetSequenceDAO(dao2.CACHE)

	seq, _ = sequenceDAO.Incr(key)
	var do *dataobject.SeqUpdatesNgenDO = nil

	// 使用seq==1做为哨兵减少DB和REDIS操作
	if seq == 1 {
		// seq为1，有两种情况:
		// 1. 没有指定key的seq，第一次生成seq，DB也无纪录
		// 2. redis重新启动，DB里可能已经有值

		SeqUpdatesNgenDAO := dao2.GetSeqUpdatesNgenDAO(dao2.DB_SLAVE)
		do = SeqUpdatesNgenDAO.SelectBySeqName(key)
		if do == nil {
			// DB无值，插入数据
			do = &dataobject.SeqUpdatesNgenDO{
				SeqName:   key,
				Seq:       seq,
				CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			}
		} else {
			// DB有seq
			do.Seq += 1
			sequenceDAO.Set(key, do.Seq)
		}
	} else {
		do = &dataobject.SeqUpdatesNgenDO{
			SeqName: key,
			Seq:     seq,
		}
	}

	// TODO(@benqi): 使用一些策略减少存盘次数
	SeqUpdatesNgenDAO := dao2.GetSeqUpdatesNgenDAO(dao2.DB_MASTER)

	if do.Seq == 1 {
		SeqUpdatesNgenDAO.Insert(do)
	} else {
		SeqUpdatesNgenDAO.UpdateSeqBySeqName(do.Seq, key)
	}

	return
}
*/

func NextSeqId(key string) (seq int64) {
	seq, _ = dao.GetSequenceDAO(dao.CACHE).NextSeqId(key)
	return
}

func CurrentSeqId(key string) (seq int64) {
	var err error
	seq, err = dao.GetSequenceDAO(dao.CACHE).CurrentSeqId(key)
	if err != nil {
		seq = -1
	}
	return
}

func NextPtsId(key string) (seq int64) {
	seq, _ = dao.GetSequenceDAO(dao.CACHE).NextPtsId(key)
	return
}

func CurrentPtsId(key string) (seq int64) {
	seq, _ = dao.GetSequenceDAO(dao.CACHE).CurrentPtsId(key)
	return
}

func NextQtsId(key string) (seq int64) {
	seq, _ = dao.GetSequenceDAO(dao.CACHE).NextQtsId(key)
	return
}

func CurrentQtsId(key string) (seq int64) {
	seq, _ = dao.GetSequenceDAO(dao.CACHE).CurrentQtsId(key)
	return
}

func NextMessageBoxId(key string) (seq int64) {
	seq, _ = dao.GetSequenceDAO(dao.CACHE).NextMessageBoxId(key)
	return
}

func CurrentMessageBoxId(key string) (seq int64) {
	seq, _ = dao.GetSequenceDAO(dao.CACHE).CurrentMessageBoxId(key)
	return
}

func NextChannelPtsId(key string) (seq int64) {
	seq, _ = dao.GetSequenceDAO(dao.CACHE).NextChannelPtsId(key)
	return
}

func CurrentChannelPtsId(key string) (seq int64) {
	seq, _ = dao.GetSequenceDAO(dao.CACHE).CurrentChannelPtsId(key)
	return
}

func NextChannelMessageBoxId(key string) (seq int64) {
	seq, _ = dao.GetSequenceDAO(dao.CACHE).NextChannelMessageBoxId(key)
	return
}

func CurrentChannelMessageBoxId(key string) (seq int64) {
	seq, _ = dao.GetSequenceDAO(dao.CACHE).CurrentChannelMessageBoxId(key)
	return
}
