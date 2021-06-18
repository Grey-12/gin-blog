package main

import (
	"fmt"
	"github.com/Grey-12/gin-blog/pkg/setting"
	"github.com/Grey-12/gin-blog/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.HTTPPort),
		Handler: router,
		ReadTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTImeOut,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}
