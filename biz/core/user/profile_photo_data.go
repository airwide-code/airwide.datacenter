/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package user

import (
	"encoding/json"
	"github.com/airwide-code/airwide.datacenter/biz/dal/dao"
)

// type profileData *ProfilePhotoIds

func MakeProfilePhotoData(jsonData string) *ProfilePhotoIds {
	if jsonData == "" {
		return &ProfilePhotoIds{}
	}
	data2 := &ProfilePhotoIds{}
	err := json.Unmarshal([]byte(jsonData), data2)
	if err != nil {
		return &ProfilePhotoIds{}
	}
	return data2
}

func (m *ProfilePhotoIds) AddPhotoId(id int64) {
	idList := make([]int64, 0, len(m.IdList))
	idList = append(idList, id)
	idList = append(idList, m.IdList...)
	m.IdList = idList
	m.Default = id
}

func (m *ProfilePhotoIds) RemovePhotoId(id int64) int64 {
	if len(m.IdList) <= 1 {
		m.IdList = []int64{}
		m.Default = 0
	} else {
		if id == m.Default {
			id = m.IdList[1]
			m.IdList = m.IdList[1:]
		} else {
			for i, j := range m.IdList {
				if j == id {
					m.IdList = append(m.IdList[:i], m.IdList[i+1:]...)
				}
			}
		}
	}
	return m.Default
}

func (m *ProfilePhotoIds) ToJson() string {
	data, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(data)
}

func GetDefaultUserPhotoID(userId int32) int64 {
	do := dao.GetUsersDAO(dao.DB_SLAVE).SelectProfilePhotos(userId)
	if do != nil {
		photoIds := MakeProfilePhotoData(do.Photos)
		return photoIds.Default
	}
	return 0
}

func GetUserPhotoIDList(userId int32) []int64 {
	do := dao.GetUsersDAO(dao.DB_SLAVE).SelectProfilePhotos(userId)
	if do != nil {
		photoIds := MakeProfilePhotoData(do.Photos)
		return photoIds.IdList
	}
	return []int64{}
}

func SetUserPhotoID(userId int32, photoId int64) {
	do := dao.GetUsersDAO(dao.DB_SLAVE).SelectProfilePhotos(userId)
	if do != nil {
		photoIds := MakeProfilePhotoData(do.Photos)
		photoIds.AddPhotoId(photoId)
		dao.GetUsersDAO(dao.DB_MASTER).UpdateProfilePhotos(photoIds.ToJson(), userId)
	}
}

func DeleteUserPhotoID(userId int32, photoId int64) {
	do := dao.GetUsersDAO(dao.DB_SLAVE).SelectProfilePhotos(userId)
	if do != nil {
		photoIds := MakeProfilePhotoData(do.Photos)
		photoIds.RemovePhotoId(photoId)
		dao.GetUsersDAO(dao.DB_MASTER).UpdateProfilePhotos(photoIds.ToJson(), userId)
	}
}
