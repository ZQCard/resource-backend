package routers

import (
	"github.com/gin-gonic/gin"
	"resource-backend/controllers"
	"resource-backend/middleware/cors"
	"resource-backend/middleware/jwt"
)

func InitRouter() *gin.Engine {
	// 初始化router
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	// 使用跨域中间件
	router.Use(cors.Cors())

	// 获取token
	router.POST("/auth", controllers.Login)

	// 视频列表
	api := router.Group("/")
	api.Use(jwt.JWT())
	{
		api.GET("/videos", controllers.Videos)
		// 添加视频
		api.POST("/video", controllers.VideoAdd)
		// 视频详情
		api.GET("/video", controllers.VideoView)
		// 更新视频
		api.PUT("/video", controllers.VideoUpdate)
		// 删除视频
		api.DELETE("/video", controllers.VideoDelete)
	}
	return router
}