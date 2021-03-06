package v1

import (
	"github.com/Grey-12/gin-blog/models"
	"github.com/Grey-12/gin-blog/pkg/errorCode"
	"github.com/Grey-12/gin-blog/pkg/logging"
	"github.com/Grey-12/gin-blog/pkg/setting"
	"github.com/Grey-12/gin-blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"strconv"
)

// GetTags 获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	// var state int64 = -1
	if arg := c.Query("state"); arg != "" {
		state, err := strconv.ParseInt(arg, 10, 0)
		if err != nil {
			//TODO
		}
		maps["state"] = int(state)
	}
	code := errorCode.SUCCESS
	listTag, err := models.GetTags(util.GetPage(c), setting.PageSize, maps)
	if err != nil {

	}
	data["lists"] = listTag
	count, err := models.GetTagTotal(maps)
	if err != nil {
		count = 0
	}
	data["total"] = count

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": errorCode.GetMsg(code),
		"data": data,
	})
}

// AddTag 新增文章标签
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createBy := c.Query("create_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createBy, "create_by").Message("创建人不能为空")
	valid.MaxSize(createBy, 100, "create_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := errorCode.INVALID_PARAMS

	if !valid.HasErrors() {
		ok, _ := models.ExistTagByName(name)
		if !ok{
			code = errorCode.SUCCESS
			models.AddTag(name, state, createBy)
		} else {
			code = errorCode.ERROR_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logging.SugarLogger.Errorf("Err key = %v, Err msg=%v", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": errorCode.GetMsg(code),
		"data": make(map[string]string),
	})
}

// EditTag 编辑文章标签
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := errorCode.INVALID_PARAMS
	if !valid.HasErrors() {
		code = errorCode.SUCCESS
		ok, _ := models.ExistTagByID(id)
		if ok{
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
			for _, err := range valid.Errors {
				logging.SugarLogger.Errorf("Err key = %v, Err msg=%v", err.Key, err.Message)
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": errorCode.GetMsg(code),
		"data": make(map[string]string),
	})
}

// DeleteTag 编辑文章标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := errorCode.INVALID_PARAMS
	if !valid.HasErrors() {
		code = errorCode.SUCCESS
		ok, _ := models.ExistTagByID(id)
		if ok {
			models.DeleteTag(id)
		} else {
			code = errorCode.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logging.SugarLogger.Errorf("Err key = %v, Err msg=%v", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": errorCode.GetMsg(code),
		"data": make(map[string]string),
	})
}
