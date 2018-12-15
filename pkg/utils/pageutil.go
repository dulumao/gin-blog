package utils

import (
	"gin-blog/pkg/setting"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPage(c *gin.Context) int {
	result := 0
	pageNo := c.DefaultQuery("page_no", "0")
	page, _ := strconv.Atoi(pageNo)
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result
}
