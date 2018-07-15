/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package base

import (
	"strconv"
)

func StringToInt32(s string) (int32, error) {
	i, err := strconv.Atoi(s)
	return int32(i), err
}

func StringToUint32(s string) (uint32, error) {
	i, err := strconv.Atoi(s)
	return uint32(i), err
}

func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func StringToUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Int32ToString(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

func BoolToInt8(b bool) int8 {
	if b {
		return 1
	} else {
		return 0
	}
}
