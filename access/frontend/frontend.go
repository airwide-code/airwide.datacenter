/*
 *  Copyright (c) 2018
 *  http://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 */

package main

import (
   "github.com/airwide-code/airwide.datacenter/baselib/app"
   "github.com/airwide-code/airwide.datacenter/access/frontend/server"
   "flag"
)

func init() {
   flag.Set("alsologtostderr", "true")
   flag.Set("log_dir", "false")
}

func main() {
   flag.Parse()

   instance := server.NewFrontendServer("./frontend.toml")
   app.DoMainAppInstance(instance)
}

