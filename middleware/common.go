package middleware

import (
	"gin-blog/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 通用的 panic 处理
func CommonRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := recover(); err != nil {
			if err := recover(); err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": e.ERROR,
					"msg":  e.GetMsg(e.ERROR),
					"data": nil,
				})
			}
		}
	}
}
