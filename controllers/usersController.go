package controllers

import (
	"fmt"
	"time"

	"com.miaoyou.server/database"
	"com.miaoyou.server/helper"
	"com.miaoyou.server/models"

	"github.com/gin-gonic/gin"
)

// Login 登录的具体实现
func Login(c *gin.Context) {
	userName := c.PostForm("userName")
	if err := helper.VerylyParamsByUserName(userName); err != nil {
		c.JSON(200, helper.Error(err.Error(), nil))
		return
	}
	passWord := c.PostForm("passWord")
	if err := helper.VerylyParamsByPassWord(passWord); err != nil {
		c.JSON(200, helper.Error(err.Error(), nil))
		return
	}
	db := database.Gdb
	var user models.User
	if err := db.Where("user_name = ?", userName).First(&user).Error; err != nil {
		c.JSON(200, helper.Error("用户不存在", nil))
	} else {
		if helper.ValationPassWord(passWord, user.Password) == true {
			userID := fmt.Sprintf("%d", user.ID)
			token, err := helper.GetToken(userID)
			if err != nil {
				c.JSON(200, helper.Error("获取token失败", nil))
			} else {
				result := map[string]string{
					"token": token,
				}
				c.JSON(200, helper.Successful(result))
			}
		} else {
			c.JSON(200, helper.Error("密码错误", nil))
		}
	}
}

// Register 注册
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(200, helper.Error(err.Error(), nil))
		return
	}

	userName := c.PostForm("userName")
	if err := helper.VerylyParamsByUserName(userName); err != nil {
		c.JSON(200, helper.Error(err.Error(), nil))
		return
	}
	passWord := c.PostForm("passWord")
	if err := helper.VerylyParamsByPassWord(passWord); err != nil {
		c.JSON(200, helper.Error(err.Error(), nil))
		return
	}
	db := database.Gdb
	if err := db.Where("user_name = ?", userName).First(&user).Error; err != nil {
		enp, err := helper.EnPassword(user.Password)
		if err != nil {
			c.JSON(200, helper.Error(err.Error(), nil))
			return
		}
		user.Time = time.Now().Unix()
		user.Password = enp
		if err := db.Create(&user).Error; err != nil {
			c.JSON(200, helper.Error(err.Error(), nil))
			return
		}
		c.JSON(200, helper.Successful("注册成功"))
	} else {
		c.JSON(200, helper.Error("用户已存在", nil))
	}
}
