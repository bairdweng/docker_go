package helper

import (
	"github.com/gin-gonic/gin"
)

// MYMiddle 自定义的中间键
func MYMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.PostForm("token")

		if len(token) == 0 {
			c.Abort()
			c.JSON(200, Error("token不能为空", nil))
			return
		}
		_, error := ValiteToken(token)
		if error != nil {
			c.Abort()
			c.JSON(200, Error("token无效", nil))
			return
		}
		c.Next()
	}
}
