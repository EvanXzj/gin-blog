package v1

import (
	"net/http"

	"github.com/astaxie/beego/validation"

	"github.com/EvanXzj/gin-blog/models"
	"github.com/EvanXzj/gin-blog/pkg/cerror"
	"github.com/EvanXzj/gin-blog/pkg/logging"
	"github.com/EvanXzj/gin-blog/pkg/setting"
	"github.com/EvanXzj/gin-blog/pkg/util"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

// GetTags Get multiple article tags
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state = -1
	if arg := c.Query("state"); arg != "" {
		state, _ = com.StrTo(arg).Int()
		maps["state"] = state
	}

	code := cerror.SUCCESS

	data["lists"] = models.GetTags(util.GetOffset(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  cerror.GetMsg(code),
		"data": data,
	})
}

// AddTag Add article tag
// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]]
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state, _ := com.StrTo(c.DefaultQuery("state", "0")).Int()
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := cerror.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = cerror.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = cerror.ERROR_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  cerror.GetMsg(code),
		"data": make(map[string]string),
	})
}

// EditTag Update article tag
// @Produce  json
// @Param id path int true "ID"
// @Param name query string true "ID"
// @Param state query int false "State"
// @Param modified_by query string true "ModifiedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags/{id} [put]
func EditTag(c *gin.Context) {
	id, _ := com.StrTo(c.Param("id")).Int()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}
	var state = -1
	if arg := c.Query("state"); arg != "" {
		state, _ = com.StrTo(arg).Int()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := cerror.INVALID_PARAMS
	if !valid.HasErrors() {
		code = cerror.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}

			models.EditTag(id, data)
		} else {
			code = cerror.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  cerror.GetMsg(code),
		"data": make(map[string]string),
	})
}

// DeleteTag Delete tag
func DeleteTag(c *gin.Context) {
	id, _ := com.StrTo(c.Param("id")).Int()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := cerror.INVALID_PARAMS
	if !valid.HasErrors() {
		code = cerror.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = cerror.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  cerror.GetMsg(code),
		"data": make(map[string]string),
	})
}
