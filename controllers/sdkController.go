package controllers

import (
	"com.miaoyou.server/database"
	"com.miaoyou.server/helper"
	"com.miaoyou.server/models"

	"github.com/gin-gonic/gin"
)

//SDKInit SDK初始化
func SDKInit(c *gin.Context) {
	var deviceInfo models.DeviceInfo
	if err := c.ShouldBind(&deviceInfo); err != nil {
		c.JSON(200, helper.Error(err.Error(), nil))
		return
	}
	db := database.Gdb
	var record models.DeviceInfo
	if err := db.Where("pid = ?", deviceInfo.Pid).First(&record).Error; err != nil {
		if err := db.Create(deviceInfo).Error; err != nil {
			c.JSON(200, helper.Error(err.Error(), nil))
			return
		}
	} else {
		if err := db.Model(&deviceInfo).Updates(deviceInfo).Error; err != nil {
			c.JSON(200, helper.Error(err.Error(), nil))
			return
		}
	}
	// token, _ := helper.GetToken()
	// tokenMap := map[string]string{
	// 	"token": token,
	// }
	c.JSON(200, helper.Successful(nil))
}
