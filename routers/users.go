package routers

import (
	"com.miaoyou.server/controllers"

	"github.com/gin-gonic/gin"
)

// RegisteUsersRoute 注册用户的路由
func RegisteUsersRoute(router *gin.Engine) {
	users := router.Group("/user")
	{
		users.POST("/login", controllers.Login)
		users.POST("/register", controllers.Register)
	}
}
