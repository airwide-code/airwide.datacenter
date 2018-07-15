/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package logger

import (
	"encoding/json"
)

var (
	// empty
	emptyBytes = []byte("{}")
)

func JsonDebugData(message interface{}) ([]byte) {
	if data, err := json.Marshal(message); err == nil {
		return data
	}
	return emptyBytes
}
