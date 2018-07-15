/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package main

import (
	"time"
	"fmt"
	"github.com/airwide-code/airwide.datacenter/baselib/grpc_util/middleware/examples/zproto"
	"io"
	"math/rand"
	"google.golang.org/grpc"
	"github.com/airwide-code/airwide.datacenter/baselib/base"
	"context"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	conn, err := grpc.Dial("127.0.0.1:22345", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("fail to dial: %v\n", err)
	}
	defer conn.Close()
	client := zproto.NewChatTestClient(conn)
	sess := &zproto.ChatSession{base.Int64ToString(rand.Int63())}
	fmt.Println("sessionId : ", sess.SessionId)

	var message string
	for {
		fmt.Print("> ")
		if n, err := fmt.Scanln(&message); err == io.EOF {
			return
		} else if n > 0 {
			if message == "quit" {
				return
			} else {
				_, err := client.SendChat(context.Background(), &zproto.ChatMessage{SenderSessionId: sess.SessionId, MessageData: message})
				if err != nil {
					fmt.Printf("%v.SendChat(_) = _, %v\n", client, err)
				}
			}
		}
	}
}

