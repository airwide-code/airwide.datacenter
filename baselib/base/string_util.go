/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package base

import "strconv"

func JoinInt32List(s []int32, sep string) string {
	l := len(s)
	if l == 0 {
		return ""
	}

	buf := make([]byte, 0, l*2+len(sep)*l+len(sep)*(l-1))
	for i := 0; i < l; i++ {
		buf = strconv.AppendInt(buf, int64(s[i]), 10)
		// buf = append(buf, sep...)
		if i != l-1 {
			buf = append(buf, sep...)
		}
	}
	return string(buf)
}

func JoinUint32List(s []uint32, sep string) string {
	l := len(s)
	if l == 0 {
		return ""
	}

	buf := make([]byte, 0, l*2+len(sep)*l+len(sep)*(l-1))
	for i := 0; i < l; i++ {
		buf = strconv.AppendUint(buf, uint64(s[i]), 10)
		buf = append(buf, sep...)
		if i != l-1 {
			buf = append(buf, sep...)
		}
	}
	return string(buf)
}

func JoinInt64List(s []int64, sep string) string {
	l := len(s)
	if l == 0 {
		return ""
	}

	buf := make([]byte, 0, l*2+len(sep)*l+len(sep)*(l-1))
	for i := 0; i < l; i++ {
		buf = strconv.AppendInt(buf, s[i], 10)
		buf = append(buf, sep...)
		if i != l-1 {
			buf = append(buf, sep...)
		}
	}
	return string(buf)
}

func JoinUint64List(s []uint64, sep string) string {
	l := len(s)
	if l == 0 {
		return ""
	}

	buf := make([]byte, 0, l*2+len(sep)*l+len(sep)*(l-1))
	for i := 0; i < l; i++ {
		buf = strconv.AppendUint(buf, s[i], 10)
		buf = append(buf, sep...)
		if i != l-1 {
			buf = append(buf, sep...)
		}
	}
	return string(buf)
}

// IsAlNumString returns true if an alpha numeric string consists of characters a-zA-Z0-9
func IsAlNumString(s string) bool {
	c := 0
	for _, r := range s {
		switch {
		case '0' <= r && r <= '9':
			c++
			break
		case 'a' <= r && r <= 'z':
			c++
			break
		case 'A' <= r && r <= 'Z':
			c++
			break
		}
	}
	return len(s) == c
}
