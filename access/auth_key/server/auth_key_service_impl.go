/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package server

import (
	"context"
	"github.com/airwide-code/airwide.datacenter/mtproto"
	"github.com/golang/glog"
	"github.com/airwide-code/airwide.datacenter/baselib/logger"
	"github.com/airwide-code/airwide.datacenter/access/auth_key/dal/dao"
	"encoding/base64"
)

type AuthKeyServiceImpl struct {
}

// rpc QueryAuthKey(AuthKeyRequest) returns (AuthKeyData);
func (s *AuthKeyServiceImpl) QueryAuthKey(ctx context.Context, request *mtproto.AuthKeyRequest) (*mtproto.AuthKeyData, error) {
	glog.Infof("auth_key.queryAuthKey - request: %s", logger.JsonDebugData(request))

	authKeyData := &mtproto.AuthKeyData{
		Result:    0,
		AuthKeyId: request.AuthKeyId,
	}

	// Check auth_key_id
	if request.AuthKeyId == 0 {
		authKeyData.Result = 1000
	}

	// TODO(@benqi): cache auth_key
	do, err := dao.GetAuthKeysDAO(dao.DB_MASTER).SelectByAuthId(request.AuthKeyId)
	if err != nil {
		glog.Error(err)
		authKeyData.Result = 1001
	} else {
		if do == nil {
			glog.Errorf("read keyData error: not find keyId = %d", request.AuthKeyId)
			authKeyData.Result = 1002
		} else {
			authKeyData.AuthKey, err = base64.RawStdEncoding.DecodeString(do.Body)
			if err != nil {
				glog.Errorf("read keyData error - keyId = %d, %v", request.AuthKeyId, err)
				authKeyData.Result = 1003
			}
		}
	}

	glog.Infof("queryAuthKey {auth_key_id: %d} ok.", request.AuthKeyId)
	return authKeyData, nil
}

// rpc QueryUserId(AuthKeyIdRequest) returns (UserIdResponse);
func (s *AuthKeyServiceImpl) QueryUserId(ctx context.Context, request *mtproto.AuthKeyIdRequest) (*mtproto.UserIdResponse, error) {
	glog.Infof("auth_key.queryUserId - request: %s", logger.JsonDebugData(request))

	userId := &mtproto.UserIdResponse{
		Result:    0,
		AuthKeyId: request.AuthKeyId,
	}

	// Check auth_key_id
	if request.AuthKeyId == 0 {
		userId.Result = 1000
	}

	// TODO(@benqi): cache auth_key
	do, err := dao.GetAuthUsersDAO(dao.DB_MASTER).SelectByAuthId(request.AuthKeyId)
	if err != nil {
		glog.Error(err)
		userId.Result = 1001
	} else {
		if do == nil {
			glog.Errorf("getUserId error: not find keyId = %d", request.AuthKeyId)
			userId.Result = 1002
		} else {
			userId.UserId = do.UserId
		}
	}

	glog.Infof("queryUserId {auth_key_id: %d} ok.", request.AuthKeyId)
	return userId, nil
}
