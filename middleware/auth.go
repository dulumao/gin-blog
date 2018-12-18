package middleware

import (
	"gin-blog/pkg/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func AuthJwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		//cookie, err := c.Request.Cookie("token")
		token := c.Query("token")
		result := gin.H{}
		claims, err := utils.ParseToken(token)
		if err == nil && claims.ExpiresAt > time.Now().Unix() {
			c.Next()
			return
		}
		result["message"] = "token 无效"
		c.JSON(200, result)
		c.Abort() // 结束调用链
	}
}
