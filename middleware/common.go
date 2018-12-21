package middleware

import (
	"errors"
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

func Common() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Next 之前")
		//c.Abort()  不显式调用 abort 将不会阻止调用执行
		c.Next()
		fmt.Println("next 之后")
	}
}

func SentryRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			flags := map[string]string{
				"endpoint": c.Request.RequestURI,
			}
			if err := recover(); err != nil {
				debug.PrintStack()
				rvalStr := fmt.Sprint(err)
				packet := raven.NewPacket(rvalStr, raven.NewException(errors.New(rvalStr), raven.NewStacktrace(2, 3, nil)))
				eventID, err := raven.Capture(packet, flags)
				fmt.Println(eventID, err)
				c.Writer.WriteHeader(http.StatusInternalServerError)

			}
		}()
		c.Next()
	}
}
