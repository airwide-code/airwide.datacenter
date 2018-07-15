/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package snowflake

import (
	"testing"
	"log"
	"fmt"
)

func TestID(t *testing.T) {
	id, err := NewIdWorker(0, 0, twepoch)
	if err != nil {
		fmt.Printf("NewIdWorker(0, 0) error(%v)\n", err)
		t.FailNow()
	}
	sid, err := id.NextId()
	if err != nil {
		fmt.Printf("id.NextId() error(%v)\n", err)
		t.FailNow()
	}
	log.Printf("snowflake id: %d\n", sid)
	sids, err := id.NextIds(10)
	if err != nil {
		fmt.Printf("id.NextId() error(%v)\n", err)
		t.FailNow()
	}
	fmt.Printf("snowflake ids: %v\n", sids)
}

func BenchmarkID(b *testing.B) {
	id, err := NewIdWorker(0, 0, twepoch)
	if err != nil {
		fmt.Printf("NewIdWorker(0, 0) error(%v)\n", err)
		b.FailNow()
	}
	for i := 0; i < b.N; i++ {
		if _, err := id.NextId(); err != nil {
			b.FailNow()
		}
	}
}
