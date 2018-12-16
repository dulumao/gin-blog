package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Common() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Next 之前")
		//c.Abort()  不显式调用 abort 将不会阻止调用执行
		c.Next()
		fmt.Println("next 之后")
	}
}
