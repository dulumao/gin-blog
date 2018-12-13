package main

import (
	"fmt"
	"gin-blog/routers"
	"gin-blog/setting"
	"log"
	"net/http"
)

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
