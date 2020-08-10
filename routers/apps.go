package routers

import (
	"com.miaoyou.server/controllers"

	"github.com/gin-gonic/gin"
)

// RegisteAppRoute 注册设备的路由
func RegisteAppRoute(router *gin.Engine) {
	users := router.Group("/app")
	{
		users.POST("/add", controllers.AddApp)
		users.GET("/getAppInfos", controllers.GetAllAppInfos)
		users.GET("/getAppInfo/:bundleID", controllers.GetAppInfoByBundleID)
		users.POST("/addRemark", controllers.AddRemark)
		users.POST("/init", controllers.AppInit)
	}
}
