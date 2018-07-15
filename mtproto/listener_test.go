/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package mtproto

import (
	"log"
	"net"
	"testing"
	// "fmt"
	// net2 "github.com/airwide-code/airwide.datacenter/baselib/net2"
	// "github.com/airwide-code/airwide.datacenter/baselib/net2/codec"
)

func TTestListener(t *testing.T) {
	//lengthBasedFrame := codec.NewLengthBasedFrame(1024)
	//
	//// server, err := net2.Listen("tcp", "0.0.0.0:12345",
	////	lengthBasedFrame, 0 /* sync send */,
	////	net2.HandlerFunc(serverSessionLoop)
	////)
	//
	//lsn := listen("server", "0.0.0.0:12345")
	//
	//server := net2.NewServer(lsn, lengthBasedFrame, 1024, net2.HandlerFunc(serverSessionLoop))
	//
	//server.Listener().Addr().String()
	//server.Serve()
}

//func serverSessionLoop(session *net2.Session) {
//	//// log.Println("OnNewSession: ")
//	//for {
//	//	line, err := session.Receive()
//	//	if err != nil {
//	//		return
//	//	}
//	//
//	//	fmt.Print(line)
//	//	err = session.Send(line)
//	//	if err != nil {
//	//		return
//	//	}
//	//}
//}

func listen(who, addr string) net.Listener {
	var lsn net.Listener
	var err error

	lsn, err = net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("setup %s listener at %s failed - %s", who, addr, err)
	}

	lsn, _ = Listen(func() (net.Listener, error) {
		return lsn, nil
	})

	log.Printf("setup %s listener at - %s", who, lsn.Addr())
	return lsn
}
