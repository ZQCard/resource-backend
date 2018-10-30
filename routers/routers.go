package routers

import (
	"gin-crud/controllers"
	"gin-crud/middleware/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 初始化router
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	// 使用跨域中间件
	router.Use(cors.Cors())

	router.GET("/videos", controllers.Videos)

	return router
}