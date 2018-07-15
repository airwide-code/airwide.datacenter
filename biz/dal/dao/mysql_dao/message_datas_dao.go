/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package mysql_dao

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	"github.com/airwide-code/airwide.datacenter/biz/dal/dataobject"
	"github.com/airwide-code/airwide.datacenter/mtproto"
)

type MessageDatasDAO struct {
	db *sqlx.DB
}

func NewMessageDatasDAO(db *sqlx.DB) *MessageDatasDAO {
	return &MessageDatasDAO{db}
}

// insert into message_datas(dialog_id, message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, `date`) values (:dialog_id, :message_id, :sender_user_id, :peer_type, :peer_id, :random_id, :message_type, :message_data, :date)
// TODO(@benqi): sqlmap
func (dao *MessageDatasDAO) Insert(do *dataobject.MessageDatasDO) int64 {
	var query = "insert into message_datas(dialog_id, message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, `date`) values (:dialog_id, :message_id, :sender_user_id, :peer_type, :peer_id, :random_id, :message_type, :message_data, :date)"
	r, err := dao.db.NamedExec(query, do)
	if err != nil {
		errDesc := fmt.Sprintf("NamedExec in Insert(%v), error: %v", do, err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	id, err := r.LastInsertId()
	if err != nil {
		errDesc := fmt.Sprintf("LastInsertId in Insert(%v)_error: %v", do, err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}
	return id
}

// select dialog_id, message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, `date` from message_datas where deleted = 0 and message_id in (:idList) order by id desc
// TODO(@benqi): sqlmap
func (dao *MessageDatasDAO) SelectByMessageIdList(idList []int64) []dataobject.MessageDatasDO {
	var q = "select dialog_id, message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, `date` from message_datas where deleted = 0 and message_id in (?) order by id desc"
	query, a, err := sqlx.In(q, idList)
	rows, err := dao.db.Queryx(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByMessageIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.MessageDatasDO
	for rows.Next() {
		v := dataobject.MessageDatasDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByMessageIdList(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByMessageIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return values
}

// select dialog_id, message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, `date` from message_datas where message_id = :message_id and deleted = 0 limit 1
// TODO(@benqi): sqlmap
func (dao *MessageDatasDAO) SelectByMessageId(message_id int64) *dataobject.MessageDatasDO {
	var query = "select dialog_id, message_id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, `date` from message_datas where message_id = ? and deleted = 0 limit 1"
	rows, err := dao.db.Queryx(query, message_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByMessageId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.MessageDatasDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByMessageId(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByMessageId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do
}
