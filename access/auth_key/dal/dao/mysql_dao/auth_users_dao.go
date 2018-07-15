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
	"github.com/airwide-code/airwide.datacenter/access/auth_key/dal/dataobject"
)

type AuthUsersDAO struct {
	db *sqlx.DB
}

func NewAuthUsersDAO(db *sqlx.DB) *AuthUsersDAO {
	return &AuthUsersDAO{db}
}

// insert into auth_users(auth_id, user_id, hash, device_model, platform, system_version, api_id, app_name, app_version, date_created, date_active, ip, country, region) values (:auth_id, :user_id, :hash, :device_model, :platform, :system_version, :api_id, :app_name, :app_version, :date_created, :date_active, :ip, :country, :region)
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) Insert(do *dataobject.AuthUsersDO) (int64, error) {
	var query = "insert into auth_users(auth_id, user_id, hash, device_model, platform, system_version, api_id, app_name, app_version, date_created, date_active, ip, country, region) values (:auth_id, :user_id, :hash, :device_model, :platform, :system_version, :api_id, :app_name, :app_version, :date_created, :date_active, :ip, :country, :region)"
	r, err := dao.db.NamedExec(query, do)
	if err != nil {
		errDesc := fmt.Sprintf("NamedExec in Insert(%v), error: %v", do, err)
		glog.Error(errDesc)
		// panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		return 0, err
	}

	id, err := r.LastInsertId()
	if err != nil {
		errDesc := fmt.Sprintf("LastInsertId in Insert(%v)_error: %v", do, err)
		glog.Error(errDesc)
		return 0, err
		// panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}
	return id, nil
}

// select auth_id, user_id, hash, device_model, platform, system_version, api_id, app_name, app_version, date_created, date_active, ip, country, region from auth_users where auth_id = :auth_id
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) SelectByAuthId(auth_id int64) (*dataobject.AuthUsersDO, error) {
	var query = "select auth_id, user_id, hash, device_model, platform, system_version, api_id, app_name, app_version, date_created, date_active, ip, country, region from auth_users where auth_id = ?"
	rows, err := dao.db.Queryx(query, auth_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByAuthId(_), error: %v", err)
		glog.Error(errDesc)
		// panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		return nil, err
	}

	defer rows.Close()

	do := &dataobject.AuthUsersDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByAuthId(_), error: %v", err)
			glog.Error(errDesc)
			// panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
			return nil, err
		}
	} else {
		return nil, nil
	}

	err = rows.Err()
	if err != nil {
		errDesc := fmt.Sprintf("rows in SelectByAuthId(_), error: %v", err)
		glog.Error(errDesc)
		return nil, err
		// panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return do, nil
}
