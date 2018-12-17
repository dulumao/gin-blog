package routers

import (
	"gin-blog/middleware"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/utils"
	"gin-blog/routers/api"
	"gin-blog/routers/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func InitRouter() *gin.Engine {

	// 设置运行模式
	gin.SetMode(setting.RunMode)

	// 创建路由
	router := gin.New()

	// 设置表单验证 validator.v9
	binding.Validator = new(utils.ValidatorV9)

	// 使用全局中间件
	//router.Use(middleware.Common())
	router.Use(gin.Logger(), gin.Recovery())

	// 鉴权路由组
	auth := router.Group("/auth")
	{
		auth.POST("/register", api.Register)
		auth.POST("/login", api.Auth)
	}

	// api v1 路由组
	apiV1 := router.Group("/api/v1")
	// 添加认证中间件
	apiV1.Use(middleware.AuthJwtToken())
	{
		//获取标签列表
		//单个中间件
		apiV1.GET("/tags", middleware.Common(), v1.GetTags)
		//新建标签
		apiV1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiV1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiV1.DELETE("/tags/:id", v1.DeleteTag)
	}
	return router
}
