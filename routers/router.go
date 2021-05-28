package routers

import (
	"github.com/Grey-12/gin-blog/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Test",
		})
	})
	return r
}