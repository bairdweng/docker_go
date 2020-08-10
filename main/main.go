package main

import (
	"log"

	"com.miaoyou.server/controllers"
	"com.miaoyou.server/database"
	"com.miaoyou.server/models"
	"com.miaoyou.server/routers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/cron"
	// log "github.com/sirupsen/logrus"
)

func main() {
	log.Println("Starting...")

	// 定义一个cron运行器
	c := cron.New()
	// 定时5秒，每5秒执行print5
	c.AddFunc("*/300 * * * * *", syncAppInfo)

	// 开始
	c.Start()
	defer c.Stop()
	// 数据库初始化
	database.InitDataBaseWithDataBase("miaoyou_data")
	db := database.Gdb
	// 自动同步
	db.AutoMigrate(&models.AppInfo{}, &models.User{}, &models.AppRemarkInfo{})
	routers.Init()
}

// syncAppInfo
func syncAppInfo() {
	log.Println("Run 5s cron")
	db := database.Gdb
	var infos []models.AppInfo
	db.Find(&infos)
	if len(infos) > 0 {
		for _, info := range infos {
			controllers.GetAppInfo(info.AppID, CallBack)
		}
	}
}

// CallBack 回调
func CallBack(code int, appID string) {
	db := database.Gdb
	var info models.AppInfo
	if err := db.Where("app_id = ?", appID).First(&info).Error; err != nil {
		print("app信息没有找到")
		return
	}
	if code == 404 {
		info.Status = "下架"
	} else {
		info.Status = "在线"
	}

	if err := db.Model(&info).Updates(&info).Error; err != nil {
		print("app信息更新失败")
		return
	}
	print("状态更新成功啦：appId: ", appID)
}
