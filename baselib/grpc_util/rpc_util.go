/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package grpc_util

import (
	"fmt"
	"reflect"
	"github.com/gogo/protobuf/proto"
)

func NewMessageByName(mname string) (proto.Message, error) {
	mt := proto.MessageType(mname)
	if mt == nil {
		return nil, fmt.Errorf("unknown message type %q", mname)
	}

	return reflect.New(mt.Elem()).Interface().(proto.Message), nil
}

