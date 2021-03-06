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
)

// GetArticle 获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := errorCode.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistArticleById(id) {
			data = models.GetArticle(id)
			code = errorCode.SUCCESS
		} else {
			code = errorCode.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			logging.SugarLogger.Errorf("Err key = %v, Err msg=%v", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": errorCode.GetMsg(code),
		"data": data,
	})
}

// GetArticles 获取多个文章
func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(c.Query("state")).MustInt()
		maps["state"] = state

		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(c.Query("tag_id")).MustInt()
		maps["tag_id"] = tagId

		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}

	code := errorCode.INVALID_PARAMS
	if !valid.HasErrors() {
		code = errorCode.SUCCESS
		data["lists"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)

	} else {
		for _, err := range valid.Errors {
			logging.SugarLogger.Errorf("Err key = %v, Err msg=%v", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": errorCode.GetMsg(code),
		"data": data,
	})

}

// AddArticle 新增文章
func AddArticle(c *gin.Context) {
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := errorCode.INVALID_PARAMS
	if !valid.HasErrors() {
		ok, _ := models.ExistTagByID(tagId)
		if  ok{
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state

			models.AddArticle(data)
			code = errorCode.SUCCESS
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
		"data": make(map[string]interface{}),
	})
}

// EditArticle 修改文章
func EditArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(c.Query("state")).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	code := errorCode.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleById(id) {
			ok, _ := models.ExistTagByID(id)
			if  ok{
				data := make(map[string]interface{})
				if tagId > 0 {
					data["tag_id"] = tagId
				}
				if title != "" {
					data["title"] = title
				}
				if desc != "" {
					data["desc"] = desc
				}
				if content != "" {
					data["content"] = content
				}
				data["modified_by"] = modifiedBy
				models.EditArticle(id, data)
			} else {
				code = errorCode.ERROR_NOT_EXIST_ARTICLE
			}
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

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := errorCode.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleById(id) {
			models.DeleteArticle(id)
			code = errorCode.SUCCESS
		} else {
			code = errorCode.ERROR_NOT_EXIST_ARTICLE
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

