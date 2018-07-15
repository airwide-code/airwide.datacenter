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
	"github.com/airwide-code/airwide.datacenter/mtproto"
	"fmt"
	"github.com/gogo/protobuf/proto"
)

func TestReflect(t *testing.T) {
	req := &mtproto.ConnectToSessionServerReq{}
	fmt.Println(proto.MessageName(req))
	// m, _ = protoToRawPayload(req)
}
