package app

import (
	"go-telegram-bot/telegram"
	"os"
	"os/signal"
	"syscall"
)

// Start 服务启动
func Start() {
	// 机器人启动
	go func() {
		telegram.BotStart()
	}()

	//创建监听退出chan
	signalChan := make(chan os.Signal, 1)
	//监听指定信号 ctrl+c kill
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
}
