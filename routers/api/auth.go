package api

import (
	"gin-blog/models"
	"gin-blog/pkg/e"
	"gin-blog/pkg/forms"
	"gin-blog/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 注册用户
func Register(c *gin.Context) {
	var regUser forms.RegisterUserForm
	err := c.ShouldBind(&regUser)
	result := gin.H{}
	if err != nil {
		result["code"] = e.InvalidParams
		result["message"] = e.GetMsg(e.InvalidParams)
		c.JSON(200, result)
		return
	}
	exist := models.UserIsExistByUsername(regUser.Username)
	if exist {
		result["code"] = e.ERROR
		result["message"] = "用户已存在"
		c.JSON(200, result)
		return
	}
	_, err = models.CreateUser(regUser.Username, regUser.Password, regUser.Name, regUser.Phone)
	if err != nil {
		result["code"] = e.ERROR
		result["message"] = e.GetMsg(e.ERROR)
		c.JSON(200, result)
		return
	}
	result["code"] = e.SUCCESS
	result["message"] = e.GetMsg(e.SUCCESS)
	c.JSON(200, result)
	return
}

// 获取token
func Auth(c *gin.Context) {
	result := gin.H{}
	var authForm forms.AuthForm
	err := c.ShouldBind(&authForm)
	if err != nil {
		result["code"] = e.InvalidParams
		result["message"] = e.GetMsg(e.InvalidParams)
		c.JSON(200, result)
		return
	}
	user := models.CheckPassword(authForm.Username, authForm.Password)
	if user != nil {
		token, err := utils.GenerateToken(user.Username)
		if err == nil {
			c.SetCookie("jwt_token", token, 60*60, "/", "", false, false)
			result["code"] = e.SUCCESS
			result["message"] = e.GetMsg(e.SUCCESS)
			result["token"] = token
			c.JSON(http.StatusOK, result)
			return
		}
	}
	result["code"] = e.ErrorAuth
	result["message"] = e.GetMsg(e.ErrorAuth)
	c.JSON(http.StatusOK, result)
}
