/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package mysql_dao

import (
	"fmt"
	"strings"
	"testing"
)

func TestCheckExists(t *testing.T) {
	params := make(map[string]interface{})
	params["username"] = "n1"
	params["username2"] = "n2"

	names := make([]string, 0, len(params))
	fmt.Println(len(names))
	for k, v := range params {
		names = append(names, k+" = :"+k)
		fmt.Println("k: ", k, ", v: ", v)
	}
	sql := fmt.Sprintf("SELECT 1 FROM users WHERE %s LIMIT 1", strings.Join(names, " AND "))
	fmt.Println("checkExists - sql: {", sql, "}, params: ", params)
}
