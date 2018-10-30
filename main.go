package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"resource-backend/pkg/config"
	"resource-backend/routers"
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
