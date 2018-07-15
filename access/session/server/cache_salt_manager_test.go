/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package server

import (
	"math/rand"
	"testing"
)

// GetOrInsertSaltList
func TestGetOrInsertSaltList(t *testing.T) {
	id := rand.Int63()
	salts, err := GetOrInsertSaltList(id, 32)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(salts)
	}

	var salt int64 = 0
	salt, err = GetOrInsertSalt(id)
	t.Log(salt)

	if CheckBySalt(id, salt) {
		t.Logf("CheckBySalt(%d, %d) = true", id, salt)
	}

	if !CheckBySalt(id, 123) {
		t.Logf("CheckBySalt(%d, 123) = false", id)
	}
}

func BenchmarkGetSalt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val, _ := GetOrInsertSalt(rand.Int63())
		_ = val
	}
}
