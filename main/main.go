package main

import (
	"com.miaoyou.server/database"
	"com.miaoyou.server/redis"
	"com.miaoyou.server/routers"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 链接redis
	redis.ConnectRedis()
	// 链接数据库
	database.InitDataBaseWithDataBase("miaoyou_data")
	// 路由初始化
	routers.Init()
}
