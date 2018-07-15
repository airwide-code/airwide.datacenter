/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package crypto

import (
	"crypto/sha256"
	"crypto/sha1"
	"crypto/rand"
	"encoding/hex"
)

func Sha256Digest(data []byte) []byte {
	r := sha256.Sum256(data)
	return r[:]
}

func Sha1Digest(data []byte) []byte {
	r := sha1.Sum(data)
	return r[:]
}

func GenerateNonce(size int) []byte {
	b := make([]byte, size)
	_, _ = rand.Read(b)
	return b
}

func GenerateStringNonce(size int) string {
	b := make([]byte, size)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
