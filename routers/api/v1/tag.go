package v1

import (
	"github.com/Grey-12/gin-blog/models"
	"github.com/Grey-12/gin-blog/pkg/errorCode"
	"github.com/Grey-12/gin-blog/pkg/setting"
	"github.com/Grey-12/gin-blog/pkg/util"
	"github.com/gin-gonic/gin"
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
	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": errorCode.GetMsg(code),
		"data": data,
	})

}

// AddTag 新增文章标签
func AddTag(c *gin.Context) {

}

// EditTag 编辑文章标签
func EditTag(c *gin.Context) {

}

// DeleteTag 编辑文章标签
func DeleteTag(c *gin.Context) {

}
