package main

import (
	"KKchatRoom/base/dbmodel"
	"KKchatRoom/configResources"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log/slog"
)

func main() {
	// 从环境变量读取连接信息，避免把密码写进代码。
	config, err := configResources.LoadConfig("?")
	if err != nil {
		slog.Error("Start err:", "tools.LoadConfig(", err.Error())
		return
	}
	dsn := config.Database.Dsn
	if dsn == "" {
		slog.Error("请设置 MYSQL_DSN")
		return
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("连接 MySQL 失败", "err", err)
		return
	}

	if err := db.AutoMigrate(&dbmodel.User{}); err != nil {
		slog.Error("创建消息表失败", "err", err)
		return
	}

	slog.Info("消息表创建成功")
}
