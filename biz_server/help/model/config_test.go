/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package model

import (
	// "fmt"
	// "github.com/airwide-code/airwide.datacenter/helper/orm"
	// _ "github.com/go-sql-driver/mysql" // import your used driver
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/airwide-code/airwide.datacenter/baselib/logger"
	"testing"
)

func TestGetAuthKey(t *testing.T) {
	var config Config

	if _, err := toml.DecodeFile("./config_test.toml", &config); err != nil {
		fmt.Errorf("%s\n", err)
		return
	}

	fmt.Printf("%s\n", logger.JsonDebugData(config))
}
