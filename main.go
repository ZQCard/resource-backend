package main

import (
	"fmt"
	"gin-crud/pkg/config"
	"gin-crud/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.SetMode(config.ServerSettings.RunMode)

	router := routers.InitRouter()

	s := &http.Server{
		Addr:fmt.Sprintf(":%d", config.ServerSettings.HTTPPort),
		Handler:router,
		ReadHeaderTimeout:config.ServerSettings.ReadTimeOut,
		WriteTimeout:config.ServerSettings.WriteTimeOut,
	}
	s.ListenAndServe()
}
