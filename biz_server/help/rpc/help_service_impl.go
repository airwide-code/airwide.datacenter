/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package rpc

import (
	"github.com/BurntSushi/toml"
	model2 "github.com/airwide-code/airwide.datacenter/biz_server/help/model"
)

const (
	CONFIG_FILE = "./config.toml"

	// date = 1509066502,    2017/10/27 09:08:22
	// expires = 1509070295, 2017/10/27 10:11:35
	EXPIRES_TIMEOUT = 3600 // 超时时间设置为3600秒

	// support user: @benqi
	SUPPORT_USER_ID = 2
)

var config model2.Config

func init() {
	if _, err := toml.DecodeFile(CONFIG_FILE, &config); err != nil {
		panic(err)
	}
}

type HelpServiceImpl struct {
}
