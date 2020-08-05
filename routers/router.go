package routers

import (
	"MiaoYouGame/controllers"
	"MiaoYouGame/helper"

	"github.com/gin-gonic/gin"
)

//Init 路由初始化
func Init() {

	router := gin.Default()

	v1 := router.Group("/sdk")
	{
		v1.POST("/init", controllers.SDKInit)
	}

	router.GET("/getToken", func(c *gin.Context) {
		token, err := helper.GetToken("test")
		if err == nil {
			c.JSON(200, helper.Successful(token))
		} else {
			c.JSON(200, helper.Error("获取失败", ""))
		}
	})

	router.GET("/validation", func(c *gin.Context) {
		token := c.Query("token")
		deToken, err := helper.ValiteToken(token)
		userName := helper.GetTokenValue("userName", deToken.Claims)
		if err == nil {
			c.JSON(200, helper.Successful(userName))
		} else {
			c.JSON(200, helper.Error("token无效", token))
		}
	})

	RegisteAppRoute(router)

	router.Use(helper.MYMiddle())
	RegisteUsersRoute(router)

	router.Run(":8200") // listen and serve on 0.0.0.0:8080

	helper.Corestart()
}
