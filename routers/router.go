package routers

import (
	"com.miaoyou.server/helper"

	"github.com/gin-gonic/gin"
)

// var zkTracer opentracing.Tracer

//Init 路由初始化
func Init() {

	// 第二步: 初始化 tracer
	// {
	// 	reporter := zkHttp.NewReporter("http://localhost:9411/api/v2/spans")
	// 	defer reporter.Close()
	// 	endpoint, err := zipkin.NewEndpoint("main3", "localhost:80")
	// 	if err != nil {
	// 		log.Fatalf("unable to create local endpoint: %+v\n", err)
	// 	}
	// 	nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	// 	if err != nil {
	// 		log.Fatalf("unable to create tracer: %+v\n", err)
	// 	}
	// 	zkTracer = zkOt.Wrap(nativeTracer)
	// 	opentracing.SetGlobalTracer(zkTracer)
	// }

	router := gin.Default()

	// 添加一个中间键
	// router.Use(func(c *gin.Context) {
	// 	span := zkTracer.StartSpan(c.Request.URL.Path)
	// 	defer span.Finish()
	// 	c.Next()
	// })

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
	// 开发调试用路由
	RegisteDevRoute(router)
	// 中间键
	router.Use(helper.MYMiddle())
	// 注册app路由
	RegisteAppRoute(router)
	// 注册用户路由
	RegisteUsersRoute(router)
	router.Run(":8200") // listen and serve on 0.0.0.0:8080
	// 同步
	// myserver.SyncAppInfo()

}
