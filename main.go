package main

import (
	"KKchatRoom/base"
)

// main 是程序入口，负责启动聊天室服务。
func main() {
	base.LogSetting()
	base.Start()
}
