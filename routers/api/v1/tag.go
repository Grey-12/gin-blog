package v1

import (
	"github.com/Grey-12/gin-blog/models"
	"github.com/Grey-12/gin-blog/pkg/errorCode"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
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
	//data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	//data["total"] = models.GetTagTotal(maps)

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
			log.Println(err.Key, err.Message)
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

}

// DeleteTag 编辑文章标签
func DeleteTag(c *gin.Context) {

}
