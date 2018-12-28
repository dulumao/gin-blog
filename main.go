package main

import (
	"fmt"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"log"
	"net/http"
)

// @title 博客 API 文档
// @version 1.0
// @description 描述呢
// @termsOfService http://swagger.io/terms/

// @contact.name zengxianxue
// @contact.url https://github.com/zengxianxue
// @contact.email zengxianxue@gmail.com

// @license.name MIT License
// @license.url https://github.com/zengxianxue/gin-blog/blob/master/LICENSE

// @host localhost:8080
// @BasePath /api/v1

// 密码认证
// @securityDefinitions.basic BasicAuth

// JWT 认证
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name token
func main() {
	router := routers.InitRouter()
	s := http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("start server error: %s", err)
	}
}
