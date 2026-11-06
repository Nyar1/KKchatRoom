package main

import (
	"KKchatRoom/base"
	"KKchatRoom/base/chat"
)

// main 是程序入口，负责启动聊天室服务。
func main() {
	base.LogSetting()
	chat.Start()
}
