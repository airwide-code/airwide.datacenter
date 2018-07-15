/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package load_balancer

import (
	"testing"
	"fmt"
)

func TestKetama(t *testing.T) {
	k := NewKetama(10, nil)
	k.Add("127.0.0.1:10000")
	k.Add("127.0.0.1:10001")
	k.Add("127.0.0.1:10002")
	k.Add("127.0.0.1:10003")

	fmt.Println(k.Get("123"))
	fmt.Println(k.Get("123"))
	fmt.Println(k.Get("234"))
	fmt.Println(k.Get("345"))
}
