/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package main

import (
	"github.com/airwide-code/airwide.datacenter/baselib/app"
	"flag"
	"github.com/airwide-code/airwide.datacenter/access/auth_key/server"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

func main() {
	flag.Parse()

	instance := server.NewAuthKeyServer("./auth_key.toml")
	app.DoMainAppInstance(instance)
}

