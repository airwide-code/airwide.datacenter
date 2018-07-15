/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package main

import (
	_ "github.com/airwide-code/airwide.datacenter/mtproto"
	"flag"
	"github.com/airwide-code/airwide.datacenter/baselib/app"
	"github.com/airwide-code/airwide.datacenter/push/sync/server"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")

}

func main() {
	flag.Parse()

	instance := server.NewSyncServer("./sync.toml")
	app.DoMainAppInstance(instance)
}

