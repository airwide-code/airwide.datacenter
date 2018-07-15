/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package server

import (
	"testing"
	"fmt"
)

func TestClientSessionManager(t *testing.T) {
	s := newClientSessionManager(100000, []byte{1}, 1)
	s.Start()

	fmt.Println("ready.")
	for i := 0; i < 10; i++ {
		s.onSessionData(&sessionData{ClientConnID{1, 1}, nil, []byte{1}})
	}

	s.Stop()
}
