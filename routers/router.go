package routers

import (
	"com.miaoyou.server/helper"

	"github.com/gin-gonic/gin"
)

//Init 路由初始化
func Init() {

	router := gin.Default()

	router.GET("/getToken", func(c *gin.Context) {
		token, err := helper.GetToken(1)
		if err == nil {
			c.JSON(200, helper.Successful(token))
		} else {
			c.JSON(200, helper.Error("获取失败", ""))
		}
	})

	router.GET("/validation", func(c *gin.Context) {
		token := c.Query("token")
		deToken, err := helper.ValiteToken(token)
		userName := helper.GetTokenValue("userToken", deToken.Claims)
		if err == nil {
			c.JSON(200, helper.Successful(userName))
		} else {
			c.JSON(200, helper.Error("token无效", token))
		}
	})

	router.Use(helper.MYMiddle())
	// 注册app路由
	RegisteAppRoute(router)
	// 注册用户路由
	RegisteUsersRoute(router)
	router.Run(":8200") // listen and serve on 0.0.0.0:8080
}
