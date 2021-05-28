package main

import (
	"fmt"
	"github.com/Grey-12/gin-blog/pkg/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Test",
		})
	})

	s := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.HTTPPort),
		Handler: router,
		ReadTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTImeOut,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}
