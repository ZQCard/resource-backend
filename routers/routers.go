package routers

import (
	"resource-backend/controllers"
	"resource-backend/middleware/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 初始化router
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	// 使用跨域中间件
	router.Use(cors.Cors())

	// 视频列表
	router.GET("/videos", controllers.Videos)
	// 添加视频
	router.POST("/video", controllers.VideoAdd)
	return router
}