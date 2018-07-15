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
	model2 "github.com/airwide-code/airwide.datacenter/biz_server/langpack/model"
)

const (
	LANG_PACK_EN_FILE = "./lang_pack_en.toml"
)

var langs model2.LangPacks

func init() {
	if _, err := toml.DecodeFile(LANG_PACK_EN_FILE, &langs); err != nil {
		panic(err)
	}
}

type LangpackServiceImpl struct {
}
