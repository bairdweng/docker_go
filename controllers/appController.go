package controllers

import (
	"fmt"
	"net/http"

	"com.miaoyou.server/database"
	"com.miaoyou.server/helper"
	"com.miaoyou.server/models"

	"github.com/gin-gonic/gin"
)

// AddApp 添加App
func AddApp(c *gin.Context) {
	var info models.AppInfo
	if err := c.ShouldBind(&info); err != nil {
		c.JSON(200, helper.Error(err.Error(), nil))
		return
	}
	db := database.Gdb
	if err := db.Where("app_id = ?", info.AppID).First(&info).Error; err == nil {
		c.JSON(200, helper.Error(info.AppName+"已存在啦哦", nil))
		return
	}
	if err := db.Create(&info).Error; err != nil {

		c.JSON(200, helper.Error(err.Error(), nil))
		return
	}
	c.JSON(200, helper.Successful("添加成功"))
}

// GetAppInfo 获取app信息
func GetAppInfo(appID string, callback func(int, string)) {
	resp, err := http.Get("https://apps.apple.com/cn/app/id" + appID)
	if err != nil {
		fmt.Println(err)
		return
	}
	callback(resp.StatusCode, appID)
}
