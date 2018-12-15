package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Common() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Next 之前")
		c.Abort()
		c.Next()
		fmt.Println("next 之后")
	}
}
