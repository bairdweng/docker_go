package main

import (
	"com.miaoyou.server/database"
	"com.miaoyou.server/myserver"
	"com.miaoyou.server/routers"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// 数据库初始化
	database.InitDataBaseWithDataBase("miaoyou_data")
	// 一定要用协程
	end := make(chan bool, 1)
	go routers.Init()
	go myserver.SyncAppInfo()
	go myserver.GrpcRun()
	<-end

	// select {}
}
