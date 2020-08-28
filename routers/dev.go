package routers

import (
	"com.miaoyou.server/controllers"
	"github.com/gin-gonic/gin"
)

// RegisteDevRoute 注册设备的路由
func RegisteDevRoute(router *gin.Engine) {
	users := router.Group("/dev")
	{
		users.GET("/go", controllers.Start)
	}
}
