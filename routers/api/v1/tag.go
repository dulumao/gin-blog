package v1

import (
	"gin-blog/models"
	"gin-blog/pkg/e"
	"gin-blog/pkg/forms"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary 获取标签
// @Description 获取标签
// @Tags tag
// @Accept json
// @Produce json
// @Param name query string false "指定 name 查询"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /tags [get]
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

// @Summary 添加标签
// @Description 添加标签
// @Tags tag
// @Accept json
// @Produce json
// @Param name query string true "标签名字"
// @Param created_by query string true "创建人"
// @Param status query string true "状态"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /tags [post]
func AddTag(c *gin.Context) {
	var addForm forms.AddTagForm
	err := c.ShouldBind(&addForm)
	result := gin.H{}
	if err != nil {
		result["code"] = e.InvalidParams
		result["message"] = e.GetMsg(e.InvalidParams)
		c.JSON(200, result)
		return
	}
	exist := models.TagIsExistByName(addForm.Name)
	if exist {
		result["code"] = e.ErrorExistTag
		result["message"] = e.GetMsg(e.ErrorExistTag)
		c.JSON(200, result)
		return
	}

	_, err = models.CreateTag(addForm.Name, addForm.CreatedBy, addForm.Status)
	if err != nil {
		result["code"] = e.ERROR
		result["message"] = e.GetMsg(e.ERROR)
		c.JSON(200, result)
		return
	}
	result["code"] = e.SUCCESS
	result["message"] = e.GetMsg(e.SUCCESS)
	c.JSON(200, result)
}

//修改文章标签
func EditTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var tagForm forms.EditTagForm
	err := c.ShouldBind(&tagForm)
	result := gin.H{}
	if err != nil || id <= 0 {
		result["code"] = e.InvalidParams
		result["message"] = e.GetMsg(e.InvalidParams)
		c.JSON(200, result)
		return
	}

	size := models.UpdateTagById(id, tagForm)
	if size < 1 {
		result["code"] = e.ERROR
		result["message"] = "更新失败"
		c.JSON(200, result)
		return
	}
	result["code"] = e.SUCCESS
	result["message"] = e.GetMsg(e.SUCCESS)
	c.JSON(200, result)

}

//删除文章标签
func DeleteTag(c *gin.Context) {
	result := gin.H{}
	id, _ := strconv.Atoi(c.Param("id"))
	if id <= 0 {
		result["code"] = e.InvalidParams
		result["message"] = e.GetMsg(e.InvalidParams)
		c.JSON(200, result)
		return
	}
	size := models.DeleteTagById(id)
	if size < 1 {
		result["code"] = e.ERROR
		result["message"] = "删除失败"
		c.JSON(200, result)
		return
	}
	result["code"] = e.SUCCESS
	result["message"] = e.GetMsg(e.SUCCESS)
	c.JSON(200, result)
}
