package main

import (
	"flag"
	"go-telegram-bot/app"
)

func main()  {

	// parse klog/v2 flags
	flag.Parse()

	app.Start()
}
