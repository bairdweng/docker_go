package middle

import (
	"com.miaoyou.server/helper"
	"com.miaoyou.server/redis"
	"github.com/gin-gonic/gin"
)

// MYMiddle 自定义的中间键
func MYMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.PostForm("token")
		// 兼容get获取
		if len(token) == 0 {
			token, _ = c.GetQuery("token")
		}
		// 兼容header获取
		if len(token) == 0 {
			token = c.Request.Header.Get("token")
		}
		path := c.Request.URL.Path
		if path == "/user/login" || path == "/user/register" {
			c.Next()
			return
		}
		if len(token) == 0 {
			c.Abort()
			c.JSON(200, helper.Error("token不能为空", nil))
			return
		}

		_, error := helper.ValiteToken(token)
		if error != nil {
			c.Abort()
			c.JSON(200, helper.Error("token无效", nil))
			return
		}
		// 根据token寻找用户id
		uid := helper.GetUserIDByToken(token)
		if uid == 0 {
			c.Abort()
			c.JSON(200, helper.Error("token未匹配用户", nil))
			return
		}
		if rtoken, err := redis.GetToken(uid); err == nil {
			if rtoken != token {
				c.Abort()
				c.JSON(200, helper.Error("token无效，用户已经重新登录", map[string]uint{
					"code": 2300,
				}))
				return
			}
		} else {
			c.Abort()
			c.JSON(200, helper.Error("token无效", nil))
			return
		}
		c.Next()
	}
}
