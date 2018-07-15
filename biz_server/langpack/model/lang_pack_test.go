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
	"testing"
)

func TestGetLangPacks(t *testing.T) {
	var langPacks LangPacks

	if _, err := toml.DecodeFile("./lang_pack_en.toml", &langPacks); err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("%v\n", langPacks)
}
