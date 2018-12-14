package v1

import (
	"gin-blog/models"
	"gin-blog/models/forms"
	"gin-blog/pkg/e"
	"gin-blog/pkg/utils"
	"gin-blog/setting"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	// 获取查询参数
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	code := e.SUCCESS
	data["lists"] = models.GetTags(utils.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

//新增文章标签
func AddTag(c *gin.Context) {
	var addForm forms.TagAddForm
	err := c.ShouldBind(&addForm)
	result := gin.H{}
	if err != nil {
		result["code"] = e.INVALID_PARAMS
		result["message"] = e.GetMsg(e.INVALID_PARAMS)
		c.JSON(200, result)
		log.Print(err)
		return
	}
	exist := models.TagIsExist(addForm.Name)
	if exist {
		result["code"] = e.ERROR_EXIST_TAG
		result["message"] = e.GetMsg(e.ERROR_EXIST_TAG)
		c.JSON(200, result)
		return
	}

	models.CreateTag(addForm.Name, addForm.CreatedBy, int(addForm.Status))
	result["code"] = e.SUCCESS
	result["message"] = e.GetMsg(e.SUCCESS)
	c.JSON(200, result)
}

//修改文章标签
func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}
