/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 */

package app

import (
	"os/signal"
	"syscall"
	"github.com/airwide-code/airwide.datacenter/glog"
	"os"
	"flag"
)

var GAppInstance AppInstance

func init() {
	flag.Parse()
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

type AppInstance interface {
	Initialize() error
	RunLoop()
	Destroy()
}

var ch = make(chan os.Signal, 1)

func DoMainAppInstance(instance AppInstance) {
	if instance == nil {
		// panic("instance is nil!!!!")
		glog.Errorf("instance is nil, will exit.")
		return
	}

	// global
	GAppInstance = instance

	glog.Info("instance initialize...")
	err := instance.Initialize()
	if err != nil {
		glog.Infof("instance initialize error: {%v}", err)
		return
	}

	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)

	glog.Info("instance run_loop...")
	go instance.RunLoop()

	// fmt.Printf("%d", os.Getpid())
	glog.Info("Wait quit...")

	s2 := <-ch
	if i, ok := s2.(syscall.Signal); ok {
		glog.Infof("instance recv os.Exit(%d) signal...", i)
	} else {
		glog.Info("instance exit...", i)
	}

	instance.Destroy()
	glog.Info("instance quited!")
}

func QuitAppInstance() {
	/*notifier := make(chan os.Signal, 1)
	signal.Stop(notifier)*/
	ch <- syscall.SIGQUIT
}
