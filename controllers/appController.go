package controllers

import (
	"fmt"
	"net/http"

	"com.miaoyou.server/database"
	"com.miaoyou.server/helper"
	"com.miaoyou.server/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// AddApp 添加App
func AddApp(c *gin.Context) {
	var info models.AppInfo
	token := c.PostForm("token")
	userID := helper.GetUserIDByToken(token)
	info.CreateID = userID
	if err := c.ShouldBind(&info); err != nil {
		c.JSON(200, helper.Error(err.Error(), nil))
		return
	}
	db := database.Gdb
	if err := db.Where("app_id = ?", info.AppID).First(&info).Error; err == nil {
		c.JSON(200, helper.Error(info.AppName+"已存在", nil))
		return
	}
	if err := db.Where("bundle = ?", info.BundleID).First(&info).Error; err == nil {
		c.JSON(200, helper.Error(info.BundleID+"已存在", nil))
		return
	}

	if err := db.Create(&info).Error; err != nil {

		c.JSON(200, helper.Error(err.Error(), nil))
		return
	}
	c.JSON(200, helper.Successful("添加成功"))
}

// GetAllAppInfos 获取所有的app信息
func GetAllAppInfos(c *gin.Context) {
	var infos []models.AppInfoResult
	db := database.Gdb
	if err := db.Table("app_info").Find(&infos).Error; err != nil {
		c.JSON(200, helper.Error(err.Error(), nil))
		return
	}
	var newInfos []models.AppInfoResult
	for _, info := range infos {
		var remark models.AppRemarkInfoResult
		db.Table("app_remark_info").Where("id=?", info.ID).Find(&remark)
		info.Remark = remark
		newInfos = append(newInfos, info)
	}
	c.JSON(200, helper.Successful(newInfos))
}

// GetAppInfoByBundleID 根据包ID获取app信息
func GetAppInfoByBundleID(c *gin.Context) {
	db := database.Gdb
	bundleID := c.Param("bundleID")
	var info models.AppInfoResult
	if err := db.Table("app_info").Where("bundle_id = ?", bundleID).First(&info).Error; err != nil {
		c.JSON(200, helper.Error(err.Error(), nil))
		return
	}
	c.JSON(200, helper.Successful(info))
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

// AddRemark app添加备注
func AddRemark(c *gin.Context) {
	var remarkInfo models.AppRemarkInfo
	db := database.Gdb
	var appInfo models.AppInfo
	if err := c.ShouldBind(&remarkInfo); err != nil {
		c.JSON(200, helper.Error(err.Error(), nil))
		return
	}
	if err := db.Where("id = ?", remarkInfo.ID).First(&appInfo).Error; err != nil {
		c.JSON(200, helper.Error("app不存在", nil))
		return
	}
	// 判断是否存在，存在更新，否则创建
	var firstRemark models.AppRemarkInfo
	if err := db.Where("id = ?", remarkInfo.ID).First(&firstRemark).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			db.Create(&remarkInfo)
		}
	} else {
		if err := db.Model(&remarkInfo).Updates(remarkInfo).Error; err != nil {
			c.JSON(200, helper.Error(err.Error(), nil))
			return
		}
	}
	c.JSON(200, helper.Successful(nil))
}
