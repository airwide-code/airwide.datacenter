package main

import (
	"github.com/airwide-code/airwide.datacenter/baselib/app"
	"github.com/airwide-code/airwide.datacenter/baselib/net2/examples/multi_proxy/server"
)

func main() {
	instance := &server.MultiProtoInsance{}
	// app.AppInstance(instance)
	app.DoMainAppInstance(instance)
}
