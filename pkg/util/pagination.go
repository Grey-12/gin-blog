package util

import (
	"github.com/Grey-12/gin-blog/pkg/setting"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := strconv.ParseInt(c.Query("page"), 10, 0)
	if page > 0 {
		result = int(page - 1) * setting.PageSize
	}
	return result
}
