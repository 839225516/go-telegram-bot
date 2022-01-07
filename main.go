package main

import (
	"flag"
	"go-telegram-bot/app"
	"k8s.io/klog/v2"
)

func main()  {

	klog.InitFlags(nil)

	// parse klog/v2 flags
	flag.Parse()

	// make sure we flush before exiting
	defer klog.Flush()

	app.Start()
}
