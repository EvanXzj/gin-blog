package main

import (
	"fmt"
	"net/http"

	"github.com/EvanXzj/gin-blog/pkg/setting"
	"github.com/EvanXzj/gin-blog/routers"
)

func main() {
	r := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        r,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
