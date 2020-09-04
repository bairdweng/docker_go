package routers

import (
	"sync"

	"com.miaoyou.server/controllers"
	"com.miaoyou.server/coroutine"
	"com.miaoyou.server/database"
	"com.miaoyou.server/models"
	"github.com/robfig/cron"
)

// 声明一个等待组
var wg sync.WaitGroup

// SyncAppInfo 同步app信息
func SyncAppInfo() {
	c := cron.New()
	// 定时5秒，每5秒执行print5
	c.AddFunc("*/20 * * * * *", runAppInfo)
	c.Start()
	// 让main函数永不退出
	// select {}
	// defer c.Stop()
}

func runAppInfo() {
	// println("runAppInfo")
	// coroutine.LockExample()
	db := database.Gdb
	var infos []models.AppInfo
	db.Find(&infos)
	// print("syncAppInfo")
	if len(infos) > 0 {
		// 限制并发为5个吧。
		ch := make(chan string, 5)
		for _, info := range infos {
			// 使用协程去执行
			// 任务加1啦
			wg.Add(1)
			go getInfo(ch, info.AppID)

		}
	}
	wg.Wait()
	println("所有任务完成了")
	coroutine.CreateChanel()
}
func getInfo(ch chan string, appID string) {
	controllers.GetAppInfo(appID, ch, CallBack)
}

// CallBack 回调
func CallBack(code int, appID string, ch chan string) {
	wg.Done()
	ch <- "hello"
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
	<-ch

}
